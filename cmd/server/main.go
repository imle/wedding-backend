package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"

	"wedding/ent"
	"wedding/ent/backroomuser"
	"wedding/pkg/server"
	"wedding/pkg/util"
)

var (
	pgHost     string
	pgPort     int
	pgUsername string
	pgDatabase string
	pgPassword string
	pgSSLMode  string
)

var (
	authZone   string
	authSecret string
)

func main() {
	app := cli.App{
		Name:      "wedding",
		ArgsUsage: "[[hostname]:port]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "pg-host",
				Destination: &pgHost,
				EnvVars:     []string{"PG_HOST"},
				Value:       "localhost",
			},
			&cli.IntFlag{
				Name:        "pg-port",
				Destination: &pgPort,
				EnvVars:     []string{"PG_PORT"},
				Value:       5432,
			},
			&cli.StringFlag{
				Name:        "pg-username",
				Destination: &pgUsername,
				EnvVars:     []string{"PG_USERNAME"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-database",
				Destination: &pgDatabase,
				EnvVars:     []string{"PG_DATABASE"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-password",
				Destination: &pgPassword,
				EnvVars:     []string{"PG_PASSWORD"},
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "pg-sslmode",
				Destination: &pgSSLMode,
				EnvVars:     []string{"PG_SSL_MODE"},
				Value:       "disable",
			},
			&cli.BoolFlag{
				Name:  "generate",
				Value: false,
			},
			&cli.StringFlag{
				Name:        "auth-zone",
				Destination: &authZone,
				EnvVars:     []string{"AUTH_ZONE"},
				Value:       "default",
			},
			&cli.StringFlag{
				Name:        "auth-secret",
				Destination: &authSecret,
				EnvVars:     []string{"AUTH_SECRET"},
				Value:       "default",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "add-backroom-user",
				Action: func(ctx *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)

					fmt.Print("Enter Username: ")
					username, err := reader.ReadString('\n')
					if err != nil {
						return err
					}

					fmt.Print("Enter Password: ")
					bytePassword, err := terminal.ReadPassword(syscall.Stdin)
					if err != nil {
						return err
					}
					fmt.Println()

					username = strings.TrimSpace(username)
					password := strings.TrimSpace(string(bytePassword))

					hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
					if err != nil {
						return err
					}

					// Connect to db.
					client, err := ent.Open("postgres", getPgConnectionString(), ent.Log(log.Println))
					if err != nil {
						return err
					}
					defer client.Close()

					save, err := client.BackroomUser.Create().
						SetUsername(username).
						SetPassword(string(hash)).
						Save(ctx.Context)
					if err != nil {
						return err
					}

					log.Printf("user created: (%s)\n", save.String())

					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			// Connect to db.
			client, err := ent.Open("postgres", getPgConnectionString(), ent.Log(log.Println))
			if err != nil {
				return err
			}
			defer client.Close()

			// TODO: Remove auto migrate.
			if err := client.Schema.Create(ctx.Context); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}

			// Gen some fake data if asked for.
			if ctx.Bool("generate") {
				generateFakeData(ctx.Context, client)
			}

			// For non-release mode we want to see the queries.
			if gin.Mode() != gin.ReleaseMode {
				client = client.Debug()
			}

			// Default port and listen on all.
			serverAddress := ctx.Args().First()
			if serverAddress == "" {
				serverAddress = ":8080"
			}

			// Setup gin.
			var router = gin.Default()

			// Setup CORS.
			corsConfig := cors.DefaultConfig()
			corsConfig.AllowOrigins = []string{"*"}
			router.Use(cors.New(corsConfig))

			// Setup JWT Auth.
			authMiddleware, err := getJwtAuthMiddleware(ctx.Context, client, router)
			if err != nil {
				return err
			}

			// Register auth handlers.
			router.POST("/api/login", authMiddleware.LoginHandler)
			auth := router.Group("/api/auth")
			// Refresh time can be longer than token timeout
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)

			// Register data handlers.
			server.RegisterAdminAPIv1(client, router.Group("/api/admin/v1/", authMiddleware.MiddlewareFunc()))
			server.RegisterAPIv1(client, router.Group("/api/v1/invitees"))

			// Create server.
			srv := &http.Server{
				Addr:    serverAddress,
				Handler: router,
			}

			// Initializing the server in a goroutine to enable graceful handling of shutdown.
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")

			shutdownCtx, cancel := context.WithTimeout(ctx.Context, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exiting")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPgConnectionString() string {
	connData := map[string]string{
		"host":     pgHost,
		"port":     strconv.Itoa(pgPort),
		"user":     pgUsername,
		"dbname":   pgDatabase,
		"password": pgPassword,
		"sslmode":  pgSSLMode,
	}
	var connectionString string
	for key, value := range connData {
		if value == "" {
			continue
		}
		connectionString += key + "=" + value + " "
	}

	return connectionString
}

func generateFakeData(ctx context.Context, client *ent.Client) {
	imleParty := client.InviteeParty.Create().
		SetName("Steven Immediate Family").
		SetCode(util.RandomString(10)).
		SaveX(ctx)

	client.Invitee.Create().
		SetName("Susan Hixson").
		SetParty(imleParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Todd Hixson").
		SetParty(imleParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Ryan Hixson").
		SetParty(imleParty).
		SaveX(ctx)

	smithParty := client.InviteeParty.Create().
		SetName("Savannah Immediate Family").
		SetCode(util.RandomString(10)).
		SaveX(ctx)

	client.Invitee.Create().
		SetName("Harold Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Kimberly Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Joseph Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Chandler Smith").
		SetParty(smithParty).
		SaveX(ctx)
}

const identityKey = "username"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func getJwtAuthMiddleware(ctx context.Context, client *ent.Client, engine *gin.Engine) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       authZone,
		Key:         []byte(authSecret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*ent.BackroomUser); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &ent.BackroomUser{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var credentials login
			if err := c.ShouldBindBodyWith(&credentials, binding.JSON); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			user, err := client.BackroomUser.Query().Where(backroomuser.Username(credentials.Username)).Only(ctx)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*ent.BackroomUser); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			var credentials login

			// Already did this once, would have errored earlier.
			_ = c.ShouldBindBodyWith(&credentials, binding.JSON)
			user, _ := client.BackroomUser.Query().Where(backroomuser.Username(credentials.Username)).Only(ctx)

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   token,
				"expire":  t.Format(time.RFC3339),
				"message": "login successfully",
				"user":    user,
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

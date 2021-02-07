package main

import (
	"bufio"
	"os"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"

	"wedding/ent"
)

var AddBackroomUser = cli.Command{
	Name: "add-backroom-user",
	Action: func(ctx *cli.Context) error {
		reader := bufio.NewReader(os.Stdin)

		log.Print("Enter Username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		log.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			return err
		}
		log.Println()

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
}

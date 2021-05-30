package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"wedding/ent"
	"wedding/pkg/util"
)

var (
	csvFile         string
	shouldDeleteAll bool
)

var ImportGuestList = cli.Command{
	Name:      "import-guest-list",
	ArgsUsage: "[< stdin]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "csv",
			Destination: &csvFile,
			TakesFile:   true,
		},
		&cli.BoolFlag{
			Name:        "clean",
			Destination: &shouldDeleteAll,
		},
	},
	Action: func(ctx *cli.Context) error {
		if shouldDeleteAll {
			shouldContinue := askForConfirmation("Are you sure you want to clear the db?", false)

			if !shouldContinue {
				log.Exit(2)
			}
		}

		config, err := ProvideEntConfig()
		if err != nil {
			return err
		}

		client, err := ent.Open("postgres", config.ConnectionString, ent.Log(log.Println))
		if err != nil {
			return err
		}
		defer client.Close()

		if shouldDeleteAll {
			client.Invitee.Delete().ExecX(ctx.Context)
			client.InviteeParty.Delete().ExecX(ctx.Context)
		}

		var reader io.ReadCloser
		if csvFile != "" {
			reader, err = os.OpenFile(csvFile, os.O_RDONLY, 0666)
			if err != nil {
				return err
			}
		} else {
			reader = os.Stdin
		}

		defer reader.Close()
		csvReader := csv.NewReader(reader)

		_, err = csvReader.Read()
		if err != nil {
			return err
		}

		var invitees []*ent.Invitee
		var parties = map[string]*ent.InviteeParty{}

		tx, err := client.Tx(ctx.Context)
		if err != nil {
			return err
		}

		defer func() {
			_ = tx.Rollback()
		}()

		for {
			read, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}

				return err
			}

			partyName := strings.TrimSpace(read[1])
			if partyName == "" {
				partyName = strings.TrimSpace(read[0])
			}

			if _, ok := parties[partyName]; !ok {
				party, err := tx.InviteeParty.Create().
					SetName(partyName).
					SetCode(util.RandomString(10)).
					Save(ctx.Context)
				if err != nil {
					return err
				}

				parties[partyName] = party
			}

			query := tx.Invitee.Create().
				SetName(read[0]).
				SetParty(parties[partyName])

			if hasPlusOne := strings.ToLower(read[2]) == "true"; hasPlusOne {
				query.SetHasPlusOne(hasPlusOne)

				if name := strings.ToLower(read[3]); name != "" {
					query.SetPlusOneName(name)
				}
			}

			if isBridalParty := strings.ToLower(read[4]) == "true"; isBridalParty {
				query.SetIsBridesmaid(isBridalParty)
			}

			if isGroomsman := strings.ToLower(read[5]) == "true"; isGroomsman {
				query.SetIsGroomsman(isGroomsman)
			}

			if isChild := strings.ToLower(read[6]) == "true"; isChild {
				query.SetIsChild(isChild)
			}

			if phone := strings.TrimSpace(read[7]); phone != "" {
				query.SetPhone(phone)
			}

			//if address := strings.TrimSpace(read[8]); address != "" {
			//	query.SetAddress(address)
			//}

			if email := strings.TrimSpace(read[9]); email != "" {
				query.SetEmail(email)
			}

			invitee, err := query.Save(ctx.Context)
			if err != nil {
				return err
			}

			invitees = append(invitees, invitee)
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		return nil
	},
}

func askForConfirmation(text string, defaultValue bool) bool {
	var response string

	fmt.Printf("%s ", text)
	if defaultValue {
		fmt.Print("[Y/n]: ")
	} else {
		fmt.Print("[y/N]: ")
	}

	_, err := fmt.Scanln(&response)
	if err != nil {
		return defaultValue
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return defaultValue
	}
}

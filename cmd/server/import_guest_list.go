package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"wedding/ent"
	"wedding/pkg/util"
)

var (
	csvFile string
)

var ImportGuestList = cli.Command{
	Name: "import-guest-list",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "csv",
			Destination: &csvFile,
			Value:       "./import.csv",
		},
	},
	Action: func(ctx *cli.Context) error {
		client, err := ent.Open("postgres", getPgConnectionString(), ent.Log(log.Println))
		if err != nil {
			return err
		}
		defer client.Close()

		file, err := os.OpenFile(csvFile, os.O_RDONLY, 0666)
		if err != nil {
			return err
		}

		reader := csv.NewReader(file)

		_, err = reader.Read()
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
			read, err := reader.Read()
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

			if _, ok := parties[read[1]]; !ok {
				party, err := tx.InviteeParty.Create().
					SetName(partyName).
					SetCode(util.RandomString(10)).
					Save(ctx.Context)
				if err != nil {
					return err
				}

				parties[read[1]] = party
			}

			query := tx.Invitee.Create().
				SetName(read[0]).
				SetParty(parties[read[1]])

			if hasPlusOne := strings.ToLower(read[2]) == "true"; hasPlusOne {
				query.SetHasPlusOne(hasPlusOne)

				if name := strings.ToLower(read[3]); name != "" {
					query.SetPlusOneName(name)
				}
			}

			//if isBridalParty := strings.ToLower(read[4]) == "true"; isBridalParty {
			//	query.SetIsBridalParty(isBridalParty)
			//}
			//
			//if isGroomsman := strings.ToLower(read[5]) == "true"; isGroomsman {
			//	query.SetIsGroomsman(isGroomsman)
			//}

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

package database

import (
	"context"
	"database/sql"
	"encoding/json"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/semconv"

	"wedding/pkg/health"

	"wedding/ent"
)

type EntConfig struct {
	ConnectionString string
	Debug            bool
}

type dbHealthCheck struct {
	db *sql.DB
}

func (c *dbHealthCheck) IsHealthy(ctx context.Context) bool {
	return c.db.PingContext(ctx) == nil
}

func ProvideEntClient(cfg *EntConfig, hm *health.Monitor) (*ent.Client, func(), error) {
	driverName := "postgres"
	driverName, err := otelsql.Register(driverName, semconv.DBSystemPostgres.Value.AsString())
	if err != nil {
		return nil, nil, err
	}

	drv, err := entsql.Open(driverName, cfg.ConnectionString)
	if err != nil {
		return nil, nil, err
	}

	hm.RegisterService("database", &dbHealthCheck{
		db: drv.DB(),
	})

	var driver dialect.Driver = drv
	if cfg.Debug {
		driver = dialect.DebugWithContext(drv, func(ctx context.Context, entry dialect.LogEntry) {
			fields := log.Fields{
				"entAction": entry.Action,
			}
			if entry.Query != "" {
				fields["query"] = entry.Query

				args, err := json.Marshal(entry.Args)
				if err != nil {
					log.WithContext(ctx).Error(err)
					fields["queryArgs"] = entry.Args
				} else {
					fields["queryArgs"] = string(args)
				}
			}
			if entry.TxID != "" {
				fields["queryTxID"] = entry.TxID
			}
			if entry.TxOpt != nil {
				fields["queryTxOptions"] = entry.TxOpt
			}

			log.WithContext(ctx).WithFields(fields).Debug(entry)
		})
	}

	client := ent.NewClient(ent.Driver(driver))

	cleanup := func() {
		err = client.Close()
		if err != nil {
			log.Error(err)
		}
	}

	return client, cleanup, err
}

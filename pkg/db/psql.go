package db

import (
	"app/config"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDb(pgCfg config.PgConfig) (*pgx.Conn, error) {

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)

	fmt.Println(dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database %v\n", err)
		return nil, err
	}
	return conn, nil
}

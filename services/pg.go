package services

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"os"
)


func NewDbConnect() (db *pg.DB) {
	dbconfig := Config{
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port_db"),
		Username: viper.GetString("username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db_name"),
		Timezone: viper.GetString("timezone"),
	}

	// connect to Postgres
	pgDB, err := dial(dbconfig)
	fmt.Println(err)

	return pgDB
}


// Dial creates new database connection to postgres
func dial(cfg Config) (*pg.DB, error) {

	dbpg := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.DBName,
		OnConnect: func(ctx context.Context, db *pg.Conn) error {
			_, err := db.Exec("SET TIME ZONE ?", cfg.Timezone)
			return err
		},
	})

	//dbpg.AddQueryHook(dbLogger{})
	err := dbpg.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbpg, nil
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Timezone string
}


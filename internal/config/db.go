package config

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabasePoolConfig struct {
	minCon            int
	maxCon            int
	maxConnsLifeTime  time.Duration
	keepAliveInterval time.Duration
}

type DatabaseConfig struct {
	host     string
	port     int
	user     string
	password string
	schema   string
	DatabasePoolConfig
}

func (dbc *DatabaseConfig) createDatabaseUrl() string {

	u := url.URL{
		Host:   fmt.Sprintf("%s:%d", dbc.host, dbc.port),
		User:   url.UserPassword(dbc.user, dbc.password),
		Scheme: "postgres",
		Path:   dbc.schema,
	}

	val := url.Values{}
	val.Add("sslmode", "disable")
	u.RawQuery = val.Encode()

	return u.String()
}

func (dbc *DatabaseConfig) GetDatabaseConnPool() *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbUrl := dbc.createDatabaseUrl()
	log.Println(dbUrl)

	pgxConf, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("fail to create url database err{%v}", err)
	}

	pgxConf.MinConns = int32(dbc.minCon)
	pgxConf.MaxConns = int32(dbc.maxCon)
	pgxConf.MaxConnLifetime = dbc.maxConnsLifeTime
	pgxConf.HealthCheckPeriod = dbc.keepAliveInterval

	connPool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		log.Fatalf("fail to set cofig database err{%v}", err)
	}

	return connPool
}

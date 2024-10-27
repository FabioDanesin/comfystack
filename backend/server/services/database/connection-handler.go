package utils

import (
	envvars "comfystack/services/env-vars"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"comfystack/types"
	"comfystack/utils/functional"
	"fmt"
	"strings"
	"sync"
)

const postgresDbRoot string = "postgresql"

var dbStrings sync.Map

func buildUserSpecs(userspec types.Userspec) string {
	return fmt.Sprintf("%s:%s", userspec.User, userspec.Password)
}

func buildHostSpecs(host types.HostSpec) string {
	return fmt.Sprintf("%s:%d", host.Host, host.Port)
}

func buildPgParams(params []types.PgParam) string {
	mappedValues := functional.Map(params, func(v types.PgParam) string {
		return fmt.Sprintf("%s=%s", v.Name, v.Value)
	})

	return strings.Join(mappedValues, "&")
}

func buildPostgresqlConnectionString(c types.PostgresqlConnString) string {
	hostSpec := buildHostSpecs(c.Hostspec)
	userSpecs := buildUserSpecs(c.Userspec)
	pgParams := buildPgParams(c.Paramspecs)

	baseConnString := fmt.Sprintf("%s://%s@%s/%s", postgresDbRoot, userSpecs, hostSpec, c.Dbname)
	if len(pgParams) == 0 {
		return baseConnString
	} else {
		return fmt.Sprintf(baseConnString+"?%s", pgParams)
	}
}

var postgresDbConnection *pgxpool.Pool = nil

func tryBuildPgsqlConnection(conn string) (*pgxpool.Pool, bool) {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, true
	} else {
		return pool, false
	}
}

func GetConnectionString() *pgxpool.Pool {
	if postgresDbConnection == nil {
		conn := buildPostgresqlConnectionString(envvars.Instance.Dbconn)
		if len(conn) == 0 {
			return nil
		} else if pool, err := tryBuildPgsqlConnection(conn); err {
			return nil
		} else {
			postgresDbConnection = pool
		}
	}
	return postgresDbConnection
}

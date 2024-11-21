package utils

import (
	envvars "comfystack/services/env-vars"
	"database/sql"

	"comfystack/types"
	"comfystack/utils/functional"
	"fmt"
	"strings"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

func tryBuildPgsqlConnection(conn string) (*bun.DB, bool) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, db == nil
}

func GetConnectionString() *bun.DB {
	conn := buildPostgresqlConnectionString(envvars.Instance.Dbconn)
	pool, err := tryBuildPgsqlConnection(conn)
	if len(conn) == 0 || err {
		return nil
	} else {
		return pool
	}
}

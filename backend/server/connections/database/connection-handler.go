package utils

import (
	"comfystack/connections/utils/functional"
	"comfystack/types"
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
	userSpecs := buildUserSpecs(c.Userspec)
	hostSpec := buildHostSpecs(c.Hostspec)
	pgParams := buildPgParams(c.Paramspecs)

	baseConnString := fmt.Sprintf("%s://%s@%s/%s", postgresDbRoot, userSpecs, hostSpec, c.Dbname)
	if len(pgParams) == 0 {
		return baseConnString
	} else {
		return fmt.Sprintf(baseConnString+"?%s", pgParams)
	}
}

func AddNewPostgresqlConnection(c types.PostgresqlConnString) {
	connStr := buildPostgresqlConnectionString(c)
	name := c.Dbname
	dbStrings.Store(name, connStr)
}

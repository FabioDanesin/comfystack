package utils

import (
	"comfystack/database/utils/functional"
	"fmt"
	"strings"
	"sync"
)

const postgresDbRoot string = "postgresql"

var dbStrings sync.Map

type Userspec struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type HostSpec struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type PgParam struct {
	Name  string
	Value string
}

type PostgresqlConnString struct {
	Userspec   Userspec  `json:"userspec"`
	Hostspec   HostSpec  `json:"connection"`
	Dbname     string    `json:"name"`
	Paramspecs []PgParam `json:"params"`
}

func buildUserSpecs(userspec Userspec) string {
	return fmt.Sprintf("%s:%s", userspec.User, userspec.Password)
}

func buildHostSpecs(host HostSpec) string {
	return fmt.Sprintf("%s:%d", host.Host, host.Port)
}

func buildPgParams(params []PgParam) string {
	mappedValues := functional.Map(params, func(v PgParam) string {
		return fmt.Sprintf("%s=%s", v.Name, v.Value)
	})

	return strings.Join(mappedValues, "&")
}

func buildPostgresqlConnectionString(c PostgresqlConnString) string {
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

func AddNewPostgresqlConnection(c PostgresqlConnString) {
	connStr := buildPostgresqlConnectionString(c)
	name := c.Dbname
	dbStrings.Store(name, connStr)
}

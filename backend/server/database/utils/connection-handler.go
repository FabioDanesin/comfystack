package utils

type Userspec struct {
	user     string
	password string
}

type HostSpec struct {
	host string
	port uint16
}

type ParamSpecs struct {
	user     string
	password string
}

type PostgresqlConnString struct {
	userspec   string
	hostspec   string
	dbname     string
	paramspecs string
}

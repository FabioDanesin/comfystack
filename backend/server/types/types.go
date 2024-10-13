package types

type SiteConnectionType struct {
	Root  string `json:"root"`
	Port  int    `json:"port"`
	Https bool   `json:"https"`
}

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

type EnvironmentVariablesFileStructure struct {
	Dbconn      []PostgresqlConnString `json:"dbconn"`
	SiteOptions SiteConnectionType     `json:"siteopts"`
}

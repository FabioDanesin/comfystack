package models

import "github.com/uptrace/bun"

type Utente struct {
	bun.BaseModel `bun:"table:Utenti,alias:u"`
	Id            int32  `bun:"id,pk,autoincrement"`
	Nome          string `bun:"nome,notnull"`
	Nickname      string `bun:"nickname,notnull,unique"`
	Password      string `bun:"password,notnull"`
}

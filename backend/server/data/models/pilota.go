package models

type Pilota struct {
	id       int    `bun:"id,pk,autoincrement"`
	nome     string `bun:",notnull,table:piloti"`
	nickname string `bun:",notnull,table:piloti"`
}

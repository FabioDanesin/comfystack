package models

import "github.com/uptrace/bun"

type Token struct {
	bun.BaseModel `bun:"table:TokenAutenticazione,alias:T"`
	Token         string  `bun:"type:uuid,default:gen_random_uuid(),pk"`
	UserId        int32   `bun:"userid"`
	User          *Utente `bun:"rel:has-one,join:userid=id"`
}

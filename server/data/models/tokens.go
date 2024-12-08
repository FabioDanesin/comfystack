package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Token struct {
	bun.BaseModel `bun:"table:TokenAutenticazione,alias:T"`
	Token         string    `bun:"token,type:uuid,default:gen_random_uuid(),pk"`
	UserId        int32     `bun:"userid"`
	CreatedDate   time.Time `bun:"createdDate,nullzero,notnull,default:current_timestamp"`
	ValidityTime  time.Time `bun:"validityTime,nullzero,notnull"` // Il tempo di validità del token
	User          *Utente   `bun:"rel:has-one,join:userid=id"`
}

func (t *Token) OnInsertHook(ctx context.Context, query bun.Query) error {

	switch query.(type) {
	case *bun.InsertQuery:
		t.ValidityTime = t.CreatedDate.Add(time.Hour * 8)
	}

	return nil
}

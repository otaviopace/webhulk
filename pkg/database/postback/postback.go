package postback

import "github.com/grvcoelho/webhulk/pkg/database/model"

type Postback struct {
	model.Model
	Status  string `db:"status"`
	Payload string `db:"payload"`
}

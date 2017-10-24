package database

type Postback struct {
	Status  string `db:"status"`
	Payload string `db:"payload"`
}

func NewPostback(status, payload string) (*Postback, error) {
	return &Postback{
		Status:  status,
		Payload: payload,
	}, nil
}

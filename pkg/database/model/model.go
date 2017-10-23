import "time"

type Model struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

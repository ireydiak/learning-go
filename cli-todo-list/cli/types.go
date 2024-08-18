package cli

import "time"

const StatusPending = "pending"
const StatusDone = "done"

type TodoItem struct {
	Id        int       `csv:"id"`
	Title     string    `csv:"title"`
	CreatedAt time.Time `csv:"created_at"`
	Status    string    `csv:"status"`
}

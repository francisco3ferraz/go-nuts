package orders

import "time"

type Order struct {
	ID          string
	Customer    string
	AmountCents int64
	CreatedAt   time.Time
}

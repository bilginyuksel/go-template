package expense

import "time"

type Expense struct {
	Title       string
	Description string
	Price       float32
	At          time.Time
}

package subscription

import "time"

type Subscription struct {
	ID string

	Company     string
	Service     string
	Price       float32
	Description string
	Start       time.Time
	End         time.Time

	MonthlyPayday int

	Settings Settings

	NoticeAt time.Time
}

type Settings struct {
	Notify     bool
	BeforeDays int
}

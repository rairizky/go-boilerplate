package dummy

import "time"

type Dummy struct {
	ID         int
	Name       string
	Email      string
	Token      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

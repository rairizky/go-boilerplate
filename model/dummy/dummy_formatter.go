package dummy

import "time"

type DummyFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

func FormatDummy(dummy Dummy, token string) DummyFormatter {
	formatter := DummyFormatter{
		ID:         dummy.ID,
		Name:       dummy.Name,
		Email:      dummy.Email,
		Token:      token,
		Created_at: dummy.Created_at,
		Updated_at: dummy.Updated_at,
		Deleted_at: dummy.Deleted_at,
	}

	return formatter
}

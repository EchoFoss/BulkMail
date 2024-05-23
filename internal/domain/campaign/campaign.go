package campaign

import (
	"github.com/rs/xid"
	"time"
)

type Campaign struct {
	ID        string    `json:"_"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}

func NewCampain(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, nameIsRequiredErr
	}

	if content == "" {
		return nil, contentIsRequiredErr
	}

	if len(emails) == 0 {
		return nil, contactIsRequiredErr
	}

	contacts := make([]Contact, len(emails))
	for idx, email := range emails {
		contacts[idx].Email = email
	}
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil
}

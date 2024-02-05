package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}

// construtor é o que inicializa uma classe quando é passado por parametro
func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	} else if content == "" {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("emails is required")
	}

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}

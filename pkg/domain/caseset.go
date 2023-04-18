package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Caseset struct {
	id    id.SnowFlakeID
	name  string
	point int
}

// NewCaseset 不変値: id
func NewCaseset(id id.SnowFlakeID) *Caseset {
	return &Caseset{id: id}
}

func (c *Caseset) GetID() id.SnowFlakeID {
	return c.id
}

func (c *Caseset) GetName() string {
	return c.name
}

func (c *Caseset) GetPoint() int {
	return c.point
}

func (c *Caseset) SetName(name string) error {
	if utf8.RuneCountInString(name) > 32 {
		return errors.New("NameLengthError")
	}
	c.name = name
	return nil
}

func (c *Caseset) SetPoint(point int) error {
	if point < 0 || point > 5000 || point%100 != 0 {
		return errors.New("InvalidPointError")
	}
	c.point = point
	return nil
}

package schema

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

func errCannotBeBlank(field string) error {
	return fmt.Errorf("%s cannot be blank", field)

}

func (o *Opening) Validate() error {
	if o.ID <= 0 {
		return fmt.Errorf("invalid Id")
	}
	if o.Role == "" {
		return errCannotBeBlank("Role")
	}
	if o.Company == "" {
		return errCannotBeBlank("Company")
	}
	if o.Location == "" {
		return errCannotBeBlank("Location")
	}

	if o.Link == "" {
		return errCannotBeBlank("Link")
	}

	if o.Salary <= 0 {
		return fmt.Errorf("Salary cannot be zero or under")
	}
	return nil
}

type OpeningResponse struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CretedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

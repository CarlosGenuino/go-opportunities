package handler

import "fmt"

func logErrorParamRequired(field string, prop string) error {
	return fmt.Errorf("Field %s type %s is required", field, prop)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" && c.Company == "" && c.Location == "" && c.Link == "" && c.Remote == nil && c.Salary <= 0 {
		return fmt.Errorf("Request Body Empty or Malformated")
	}
	if c.Role == "" {
		return logErrorParamRequired("role", "string")
	}
	if c.Company == "" {
		return logErrorParamRequired("company", "string")
	}
	if c.Location == "" {
		return logErrorParamRequired("location", "string")
	}

	if c.Link == "" {
		return logErrorParamRequired("link", "string")
	}

	if c.Remote == nil {
		return logErrorParamRequired("remote", "bool")
	}

	if c.Salary <= 0 {
		return logErrorParamRequired("salary", "int64")
	}
	return nil
}

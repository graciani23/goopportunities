package handler

import (
	"fmt"
)

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpeningHandler
type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf(("request body is empty or malformed"))
	}
	if r.Role == "" {
		return errParamsIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamsIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamsIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errParamsIsRequired("link", "string")
	}

	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role == "" || r.Company == "" || r.Location == "" || r.Remote == nil || r.Link == "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}

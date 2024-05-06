package handlers

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamMissing(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

func (req *CreateOpeningRequest) validate() error {
	if req.Role == "" {
		return errParamMissing("role", "string")
	}

	if req.Company == "" {
		return errParamMissing("company", "string")
	}

	if req.Location == "" {
		return errParamMissing("location", "string")
	}

	if req.Link == "" {
		return errParamMissing("link", "string")
	}

	if req.Salary <= 0 {
		return errParamMissing("salary", "int64")
	}

	if req.Remote == nil {
		return errParamMissing("remote", "bool")
	}

	return nil
}

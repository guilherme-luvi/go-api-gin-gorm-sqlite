package handlers

import "fmt"

// Auxiliar function to return an error when a required parameter is missing
func errParamMissing(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

// CreateOpeningRequest struct represents the request to create an opening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
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

// UpdateOpeningRequest struct represents the request to update an opening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *UpdateOpeningRequest) validate() error {
	// if any of the fields is provided, validate is truthy
	if req.Role != "" || req.Company != "" || req.Location != "" || req.Remote != nil || req.Link != "" || req.Salary > 0 {
		return nil
	}

	// if none of the fields is provided, return error
	return fmt.Errorf("at least one field must be provided")
}

// CreateUserRequest struct represents the request to create a user
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) validate() error {
	if req.Name == "" {
		return errParamMissing("name", "string")
	}

	if req.Email == "" {
		return errParamMissing("email", "string")
	}

	if req.Password == "" {
		return errParamMissing("password", "string")
	}

	return nil
}

// UpdateUserRequest struct represents the request to update a user
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *UpdateUserRequest) validate() error {
	// if any of the fields is provided, validate is truthy
	if req.Name != "" || req.Email != "" || req.Password != "" {
		return nil
	}

	// if none of the fields is provided, return error
	return fmt.Errorf("at least one field must be provided")
}

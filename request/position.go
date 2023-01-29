package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	PositionRequest struct {
		Name   string `json:"name" validate:"required"`
		Salary int    `json:"salary" validate:"required"`
	}
)

func (req PositionRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Salary, validation.Required),
	)
}

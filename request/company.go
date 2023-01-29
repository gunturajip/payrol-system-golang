package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	CompanyRequest struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
		Address string `json:"address"`
	}

	TopupCompanyBalance struct {
		Balance int `json:"balance" validate:"required"`
	}
)

func (comp CompanyRequest) Validate() error {
	return validation.ValidateStruct(
		&comp,
		validation.Field(&comp.Name, validation.Required),
		validation.Field(&comp.Balance, validation.Required),
		validation.Field(&comp.Address, validation.Required),
	)
}

func (req TopupCompanyBalance) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Balance, validation.Required),
	)
}

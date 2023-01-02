package form

import "github.com/gookit/validate"

type Validator interface {
	Validate()
}

func Validate(f Validator) validate.Errors {
	v := validate.Struct(f)

	if v.Validate() {
		return nil
	}

	v.BindSafeData(f)

	return v.Errors
}

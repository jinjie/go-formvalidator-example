package form

import (
	"github.com/gookit/validate"
)

type Form struct {
	Validator
}

// Example of a custom validator
func (f Form) MustBeJinJie(val string) bool {
	return val == "Jin Jie"
}

// Custom messages
func (f Form) Messages() map[string]string {
	return validate.MS{
		"mustBeJinJie": "Only Accepts Jin Jie",
	}
}

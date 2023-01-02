package form

type UserForm struct {
	Validator
	Form
	FirstName string `json:"firstName" validate:"required|mustBeJinJie" filter:"trim"`
	LastName  string `json:"lastName" filter:"trim"`
	Email     string `json:"email" validate:"required|email"`
}

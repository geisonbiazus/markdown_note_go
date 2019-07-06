package commands

type ValidationResult struct {
	Valid  bool
	Errors []Error
}

type Error struct {
	Field string
	Type  string
}

type Validator struct {
	Errors []Error
}

func NewValidator() *Validator {
	return &Validator{
		Errors: []Error{},
	}
}

func (v *Validator) ValidateRequired(field, value string) {
	if value == "" {
		v.Errors = append(v.Errors, Error{Field: field, Type: "REQUIRED"})
	}
}

func (v *Validator) Result() ValidationResult {
	return ValidationResult{
		Valid:  len(v.Errors) == 0,
		Errors: v.Errors,
	}
}

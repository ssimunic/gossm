package validate

// Validator is used to validate something
type Validator interface {
	// Validate returns <true, nil> if valid, <false, error> if invalid
	Validate() (bool, error)
}

// ValidateAll validates all validators
func ValidateAll(validators ...Validator) (bool, error) {
	for _, validator := range validators {
		if ok, err := validator.Validate(); !ok {
			return false, err
		}
	}
	return true, nil
}

package validate

// Validator is used to validate something
type Validator interface {
	// Validate returns nil if valid, error if invalid
	Validate() error
}

// ValidateAll validates all validators
func ValidateAll(validators ...Validator) error {
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

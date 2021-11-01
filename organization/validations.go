package organization

func (org *Organization) Validate() []map[string]string {
	var err map[string]string
	var ValidationErrors []map[string]string

	if (*org).Name == "" {
		err = map[string]string{"firstname": "name of organization can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*org).Address == "" {
		err = map[string]string{"lastname": "laname of organization can not be blank"}
		ValidationErrors = append(ValidationErrors, err)

	}

	if len(ValidationErrors) > 0 {
		return ValidationErrors
	}

	return nil
}

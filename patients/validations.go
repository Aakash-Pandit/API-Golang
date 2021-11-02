package patients

func (patient *Patient) Validate() []map[string]string {
	var err map[string]string
	var ValidationErrors []map[string]string

	if (*patient).FirstName == "" {
		err = map[string]string{"firstname": "first name of patient can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*patient).Email == "" {
		err = map[string]string{"email": "Address of patient can not be blank"}
		ValidationErrors = append(ValidationErrors, err)

	}

	if len(ValidationErrors) > 0 {
		return ValidationErrors
	}

	return nil
}

func (medicine *Medicine) Validate() []map[string]string {
	var err map[string]string
	var ValidationErrors []map[string]string

	if (*medicine).Name == "" {
		err = map[string]string{"name": "name of medicine can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*medicine).Cost == 0 {
		err = map[string]string{"cost": "cost of medicine can not be zero"}
		ValidationErrors = append(ValidationErrors, err)

	}

	if len(ValidationErrors) > 0 {
		return ValidationErrors
	}

	return nil
}

package users

func (user *User) Validate() []map[string]string {
	var err map[string]string
	var ValidationErrors []map[string]string

	if (*user).FirstName == "" {
		err = map[string]string{"firstname": "firstname of user can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*user).Email == "" {
		err = map[string]string{"email": "email of user can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*user).Password == "" {
		err = map[string]string{"password": "password of user can not be blank"}
		ValidationErrors = append(ValidationErrors, err)

	}

	if len(ValidationErrors) > 0 {
		return ValidationErrors
	}

	return nil
}

func (userAuthentication *UserAuthentication) Validate() []map[string]string {
	var err map[string]string
	var ValidationErrors []map[string]string

	if (*userAuthentication).Email == "" {
		err = map[string]string{"email": "email of user can not be blank"}
		ValidationErrors = append(ValidationErrors, err)
	}

	if (*userAuthentication).Password == "" {
		err = map[string]string{"password": "password of user can not be blank"}
		ValidationErrors = append(ValidationErrors, err)

	}

	if len(ValidationErrors) > 0 {
		return ValidationErrors
	}

	return nil
}

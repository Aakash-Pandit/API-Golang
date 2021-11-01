package patients

import "errors"

func (patient *Patient) Validate() error {
	if (*patient).FirstName == "" {
		return errors.New("first name of Patient can not be blank")
	}

	return nil
}

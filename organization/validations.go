package organization

import "errors"

func (org *Organization) Validate() error {
	if (*org).Name == "" {
		return errors.New("name of organization can not be blank")
	}

	return nil
}

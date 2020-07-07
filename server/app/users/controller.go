package users

import (
	usermodels "stolencarsproject/server/app/users/models"
)

func GetUserDetails(email string) (usermodels.User, error) {
	data, err := usermodels.GetUserByEmailPswrd(email)
	if err != nil {
		return usermodels.User{}, err
	}
	return data, nil
}

func CreateUser(userDetails *usermodels.User) error {
	err := usermodels.CreateUser(userDetails)
	if err != nil {
		return err
	}
	return nil
}

func GetAvailablePO() (usermodels.User, error) {

	data, err := usermodels.GetAvailablePO()
	if err != nil {
		return usermodels.User{}, err
	}

	return data, nil
}

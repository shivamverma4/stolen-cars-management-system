package users

import (
	"fantasytipster/server/app/utils"
	"fmt"
	"net/http"
	usermodels "stolencarsproject/server/app/users/models"

	"github.com/labstack/echo"
)

func HandleGetUser(c echo.Context) (err error) {

	email := fmt.Sprint(c.Param("email"))

	userDetails, err := GetUserDetails(email)
	if fmt.Sprint(err) == "mongo: no documents in result" {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = "No User Found"
		return c.JSON(http.StatusNotFound, resp)
	} else if err != nil {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = userDetails
	resp.Message = "User retrieved"
	return c.JSON(http.StatusOK, resp)
}

func HandleCreateUser(c echo.Context) (err error) {

	user := fmt.Sprint(c.Param("userType"))

	u := new(usermodels.User)
	if err = c.Bind(u); err != nil {
		return
	}
	if _err := ValidateUserDetails(u); (_err != utils.CustomHTTPError{}) {
		return c.JSON(http.StatusBadRequest, _err)
	}

	if user == "owner" {
		u.UserType = 1
	} else if user == "police" {
		u.UserType = 9
	}

	u.Availability = 1
	err = CreateUser(u)
	if err != nil {
		_error := utils.GenerateError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusOK, _error)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = map[string]bool{
		"created": true,
	}
	resp.Message = "User created"
	return c.JSON(http.StatusOK, resp)
}

func ValidateUserDetails(user *usermodels.User) (_error utils.CustomHTTPError) {

	if utils.ValidateEmail(user.Email) == false {
		_error := utils.GenerateError(http.StatusBadRequest, "Please enter valid email")
		return _error
	}

	if CheckUserExists(user.Email) == true {
		_error := utils.GenerateError(http.StatusBadRequest, "Email provided already exists")
		return _error
	}
	return utils.CustomHTTPError{}
}

func CheckUserExists(email string) bool {

	_, err := GetUserDetails(email)
	if fmt.Sprint(err) == "mongo: no documents in result" {
		return false
	}

	return true
}

func HandleGetAvailablePO(c echo.Context) (err error) {

	userDetails, err := GetAvailablePO()
	if fmt.Sprint(err) == "mongo: no documents in result" {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = "No User Found"
		return c.JSON(http.StatusNotFound, resp)
	} else if err != nil {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = userDetails
	resp.Message = "User retrieved"
	return c.JSON(http.StatusOK, resp)
}

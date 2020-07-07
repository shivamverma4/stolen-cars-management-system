package cars

import (
	"fantasytipster/server/app/utils"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"net/http"
	carsmodels "stolencarsproject/server/app/cars/models"
	"stolencarsproject/server/app/users"

	"github.com/labstack/echo"
)

type StolenCar struct {
	RegNumber   string `json:"regnum"`
	Color       string `json:"color"`
	OwnerID     string `json:"ownerID"`
	POID        string `json:"poID"`
	Description string `json:"description"`
	Status      uint   `json:"status"` // {"0": unassigned,"1": assigned, "5": not found. "9": found}
}

type StolenCarStatus struct {
	OID    string `json:"oID"`
	POID   string `json:"poID"`
	Status uint   `json:"status"`
}

func HandleGetStolenCarDetails(c echo.Context) (err error) {

	regID := fmt.Sprint(c.Param("regID"))
	userType := fmt.Sprint(c.Param("userType"))
	uID := fmt.Sprint(c.Param("uID"))

	carDetails, err := GetCarDetails(regID, userType, uID)
	if fmt.Sprint(err) == "mongo: no documents in result" {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = "No Car Found"
		return c.JSON(http.StatusNotFound, resp)
	} else if err != nil {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = carDetails
	resp.Message = "Car Details retrieved"
	return c.JSON(http.StatusOK, resp)
}

func HandleCreateStolenCarRecord(c echo.Context) (err error) {

	sc := new(StolenCar)
	if err = c.Bind(sc); err != nil {
		return
	}
	if _err := ValidateStolenCarsDetails(sc.RegNumber); (_err != utils.CustomHTTPError{}) {
		return c.JSON(http.StatusBadRequest, _err)
	}

	var oid primitive.ObjectID
	if len(sc.OwnerID) > 0 {
		oid, _ = primitive.ObjectIDFromHex(sc.OwnerID)
	}

	carStolenStatus := uint(0)

	availablePO, err := users.GetAvailablePO()
	if fmt.Sprint(err) == "mongo: no documents in result" {
		carStolenStatus = uint(0)
		fmt.Println("no available PO found")
	} else if err != nil {
		_error := utils.GenerateError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusBadRequest, _error)
	}

	if availablePO.UserType != 9 || availablePO.Availability != 1 {
		carStolenStatus = uint(0)
		fmt.Println("Wrong PO fetched")
	}

	if len(availablePO.Email) > 0 {
		carStolenStatus = 1
	}

	sc_cars := carsmodels.Car{
		RegNumber:   sc.RegNumber,
		Color:       sc.Color,
		OwnerID:     oid,
		POID:        availablePO.ID,
		Description: sc.Description,
		Status:      carStolenStatus,
	}

	err = UpdateAvailabilityOfPO(availablePO.ID, 0)
	if err != nil {
		fmt.Println("Error while updating POs availability ")
		_error := utils.GenerateError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusBadRequest, _error)
	}

	err = CreateCarRecord(sc_cars)
	if err != nil {
		_error := utils.GenerateError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusOK, _error)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = map[string]bool{"created": true}
	resp.Message = "Stolen Car Record added successfully"
	return c.JSON(http.StatusOK, resp)
}

func ValidateStolenCarsDetails(regNumber string) (_error utils.CustomHTTPError) {

	if CheckStolenCarExists(regNumber) == true {
		_error := utils.GenerateError(http.StatusBadRequest, "Registration Number of the Stolen Car already exists")
		return _error
	}
	fmt.Println("Stolen Car validated successful")
	return utils.CustomHTTPError{}
}

func CheckStolenCarExists(regNumber string) bool {

	_, err := GetCarDetails(regNumber, "", "")
	if fmt.Sprint(err) == "mongo: no documents in result" {
		return false
	}
	return true
}

func HandleChangeStolenCarStatus(c echo.Context) (err error) {

	scStatus := new(StolenCarStatus)
	if err = c.Bind(scStatus); err != nil {
		return
	}

	oid, _ := primitive.ObjectIDFromHex(scStatus.OID)
	poID, _ := primitive.ObjectIDFromHex(scStatus.POID)

	err = UpdateStolenCarStatus(oid, scStatus.Status)
	if err != nil {
		_error := utils.GenerateError(http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusOK, _error)
	}

	// assign Any UnAssigned Stolen Cars To This Police Officer
	unAssignedCar, err := GetUnAssignedStolenCars()
	if err != nil {

		err = UpdateAvailabilityOfPO(poID, 1)
		if err != nil {
			_error := utils.GenerateError(http.StatusBadRequest, err.Error())
			return c.JSON(http.StatusOK, _error)
		}

		resp := utils.CustomHTTPResponse{}
		resp.Data = map[string]bool{"created": true}
		resp.Message = "Stolen Car Status changed successfully and now the police officer is available for assigning any stolen car"
		return c.JSON(http.StatusOK, resp)
	}

	if len(unAssignedCar.Color) > 0 {
		err = AssignPOToStolenCar(unAssignedCar.ID, poID)
		if err != nil {
			_error := utils.GenerateError(http.StatusBadRequest, err.Error())
			return c.JSON(http.StatusOK, _error)
		}
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = map[string]bool{"created": true}
	resp.Message = "Stolen Car Status changed successfully and PO is assigned successfully to another stolen car"
	return c.JSON(http.StatusOK, resp)
}

func HandleGetAllStolenCarsForUser(c echo.Context) (err error) {

	uID := fmt.Sprint(c.Param("uID"))
	userType := fmt.Sprint(c.Param("userType"))

	if len(uID) < 24 {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = "Invalid User ID"
		return c.JSON(http.StatusBadRequest, resp)
	}

	stolenCars, err := GetStolenCarsForUser(uID, userType)
	if err != nil {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = stolenCars
	resp.Message = "All Stolen Cars retrieved"
	return c.JSON(http.StatusOK, resp)
}

func HandleAssignUnassignedStolenCar(c echo.Context) (err error) {

	uID := fmt.Sprint(c.Param("uID"))
	oid, _ := primitive.ObjectIDFromHex(uID)

	err = AssignUnassignedStolenCar(oid)
	if fmt.Sprint(err) == "mongo: no documents in result" {
		resp := utils.CustomHTTPResponse{}
		resp.Data = map[string]bool{"created": false}
		resp.Message = "All stolen cars Assigned"
		return c.JSON(http.StatusOK, resp)
	} else if err != nil {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = map[string]bool{"created": true}
	resp.Message = "Unassigned stolen car assigned"
	return c.JSON(http.StatusOK, resp)
}

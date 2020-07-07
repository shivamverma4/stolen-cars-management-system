package cars

import (
	"encoding/json"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	carsmodels "stolencarsproject/server/app/cars/models"
	usermodels "stolencarsproject/server/app/users/models"
)

type SCar struct {
	ID          string `json:"_id"`
	RegNumber   string `json:"regnum"`
	Color       string `json:"color"`
	OwnerID     string `json:"ownerID"`
	POID        string `json:"poID"`
	Description string `json:"description"`
	Status      uint   `json:"status"` // {"0": unassigned,"1": assigned, "5": not found. "9": found}
}

type defaultMongoContainer struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

func GetCarDetails(regID string, userType string, uID string) (carsmodels.Car, error) {
	data, err := carsmodels.GetCarDetails(regID, userType, uID)
	if err != nil {
		return carsmodels.Car{}, err
	}
	return data, nil
}

func CreateCarRecord(carDetails carsmodels.Car) error {

	err := carsmodels.CreateCarRecord(carDetails)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStolenCarStatus(oid primitive.ObjectID, status uint) error {
	err := carsmodels.UpdateStolenCarStatus(oid, status)
	if err != nil {
		return err
	}
	return nil
}

func GetUnAssignedStolenCars() (carsmodels.Car, error) {
	data, err := carsmodels.GetUnAssignedStolenCars()
	if err != nil {
		return carsmodels.Car{}, err
	}
	return data, nil
}

func AssignPOToStolenCar(oID primitive.ObjectID, poID primitive.ObjectID) (err error) {
	err = carsmodels.AssignPOToStolenCar(oID, poID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAvailabilityOfPO(poID primitive.ObjectID, availability uint) error {
	err := usermodels.UpdateAvailabilityOfPO(poID, availability)
	if err != nil {
		return err
	}
	return nil
}

func GetStolenCarsForUser(uID string, userType string) (stolenCars []SCar, err error) {
	cars, err := carsmodels.GetStolenCarsForUser(uID, userType)
	if err != nil {
		return nil, err
	}

	for _, res := range cars {
		var carDetails []defaultMongoContainer
		marshalledOneStolenCar, _ := json.Marshal(res)
		err = json.Unmarshal(marshalledOneStolenCar, &carDetails)
		if err != nil {
			fmt.Println("error while unmarshalling stolen cars list")
			return
		}

		oneStolenCar := make(map[string]interface{})
		for _, details := range carDetails {
			oneStolenCar[details.Key] = details.Value
		}

		var oneCar SCar
		marshalledStolenCar, _ := json.Marshal(oneStolenCar)
		err = json.Unmarshal(marshalledStolenCar, &oneCar)
		if err != nil {
			fmt.Println("error while unmarshalling stolen cars list")
			return
		}

		stolenCars = append(stolenCars, oneCar)

	}

	return
}

func AssignUnassignedStolenCar(poID primitive.ObjectID) error {

	unassignedStolenCars, err := GetUnAssignedStolenCars()
	if err != nil {
		return err
	}

	err = AssignPOToStolenCar(unassignedStolenCars.ID, poID)
	if err != nil {
		return err
	}

	err = UpdateAvailabilityOfPO(poID, 0)
	if err != nil {
		return err
	}

	return nil
}

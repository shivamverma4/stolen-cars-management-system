package carsmodels

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	mongoController "stolencarsproject/server/common/services/mongo"
)

const carsCollection = "stolencars"

type Car struct {
	ID          primitive.ObjectID `bson:"_id"`
	RegNumber   string             `json:"regnum", bson:"regnum"`
	Color       string             `json:"color", bson:"color"`
	OwnerID     primitive.ObjectID `bson:"ownerID"`
	POID        primitive.ObjectID `bson:"poID"`
	Description string             `json:"description", bson:"description"`
	Status      uint               `json:"status", bson:"status"` // {"0": unassigned,"1": assigned, "5": not found. "9": found}
}

func GetCarDetails(regID string, userType string, uID string) (carData Car, err error) {

	var searchElements bson.D
	if len(regID) > 0 {
		searchElements = append(searchElements, bson.E{"regnum", regID})
	}

	oid, _ := primitive.ObjectIDFromHex(uID)
	if userType == "owner" {
		searchElements = append(searchElements, bson.E{"ownerID", oid})
	} else if userType == "police" {
		searchElements = append(searchElements, bson.E{"poID", oid})
	}
	idDoc := searchElements

	result, err := mongoController.FindOne(carsCollection, idDoc)
	if err != nil {
		fmt.Println("mongo get stolen car by regID and uID failure")
		return
	}

	err = result.Decode(&carData)
	if err != nil {
		fmt.Println("mongo data : ", carData, ", error : ", err)
		return Car{}, err
	}

	return
}

func CreateCarRecord(carDetails Car) (err error) {

	idDoc := bson.M{"regnum": carDetails.RegNumber, "color": carDetails.Color, "ownerID": carDetails.OwnerID, "poID": carDetails.POID, "status": carDetails.Status, "description": carDetails.Description}

	err = mongoController.InsertOne(carsCollection, idDoc)
	if err != nil {
		fmt.Println("mongo getUserByEmail failure")
		return
	}
	return nil
}

func UpdateStolenCarStatus(oid primitive.ObjectID, status uint) (err error) {

	filter := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}
	update := bson.M{"$set": bson.M{"status": status}}

	err = mongoController.UpdateOne(carsCollection, filter, update)
	if err != nil {
		fmt.Println("mongo updateStolenCarStatus failure")
		return
	}
	return nil
}

func GetUnAssignedStolenCars() (carData Car, err error) {

	idDoc := bson.D{{"status", 0}}

	result, err := mongoController.FindOne(carsCollection, idDoc)
	if err != nil {
		fmt.Println("mongo get unassigned stolen car failure")
		return
	}

	err = result.Decode(&carData)
	if err != nil {
		fmt.Println("mongo data : ", carData, ", error : ", err)
		return Car{}, err
	}

	return

}

func AssignPOToStolenCar(oID primitive.ObjectID, poID primitive.ObjectID) (err error) {

	filter := bson.M{
		"_id": bson.M{
			"$eq": oID,
		},
	}
	update := bson.M{"$set": bson.M{"poID": poID, "status": 1}}

	err = mongoController.UpdateOne(carsCollection, filter, update)
	if err != nil {
		fmt.Println("mongo assignPOToStolenCar failure")
		return
	}
	return nil
}

func GetStolenCarsForUser(uID string, userType string) (stolenCars []interface{}, err error) {

	oID, _ := primitive.ObjectIDFromHex(uID)

	var searchElements bson.D
	if userType == "owner" {
		searchElements = append(searchElements, bson.E{"ownerID", oID})
	} else if userType == "police" {
		searchElements = append(searchElements, bson.E{"poID", oID})
	}

	idDoc := searchElements

	stolenCars, err = mongoController.FindAll(carsCollection, idDoc)
	if err != nil {
		fmt.Println("mongo get stolen cars for added by or assigned to user failure")
		return
	}

	return
}

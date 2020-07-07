package usermodels

import (
	"errors"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	mongoController "stolencarsproject/server/common/services/mongo"
	"stolencarsproject/server/config"
)

const userCollection = "sc_users"

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `json:"name", bson:"name"`
	Email        string             `json:"email", bson:"email"`
	UserType     uint               `json:"usertype", bson:"usertype"`         // {"1": car_owner, "9": police_officer}
	Availability uint               `json:"availability", bson:"availability"` // {"0": not available,"1": available}
}

func GetUserByEmailPswrd(email string) (userData User, err error) {

	idDoc := bson.D{{"email", email}}

	result, err := mongoController.FindOne(userCollection, idDoc)
	if err != nil {
		fmt.Println("mongo getUserByEmail failure")
		return
	}

	err = result.Decode(&userData)
	if err != nil {
		fmt.Println("mongo data : ", userData, ", error : ", err)
		return User{}, err
	}

	return
}

func CreateUser(userDetails *User) (err error) {
	fmt.Println("creating user : ", userDetails)
	idDoc := bson.M{"name": userDetails.Name, "email": userDetails.Email, "usertype": userDetails.UserType, "availability": userDetails.Availability}

	err = mongoController.InsertOne(userCollection, idDoc)
	if err != nil {
		fmt.Println("mongo getUserByEmail failure")
		return
	}
	return nil
}

func GetAvailablePO() (userData User, err error) {

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "stolencars"}, {"localField", "_id"}, {"foreignField", "poID"}, {"as", "found_profiles"}}}}
	matchStage := bson.D{{"$match", bson.D{{"availability", 1}, {"usertype", 9}}}}

	mongoResult := mongoController.GetAggregate(userCollection, matchStage, lookupStage, config.MongoDbName)
	if mongoResult == nil || mongoResult.Current == nil {
		fmt.Println("mongoResult or mongoResult.Current empty : ", mongoResult)
		err = errors.New("mongo: no documents in result")
		return
	}

	err = mongoResult.Decode(&userData)
	if err != nil {
		fmt.Println("mongo data : ", userData, ", error : ", err)
		return User{}, err
	}

	return
}

func UpdateAvailabilityOfPO(poId primitive.ObjectID, availability uint) (err error) {

	filter := bson.M{
		"_id": bson.M{
			"$eq": poId,
		},
	}
	update := bson.M{"$set": bson.M{"availability": availability}}

	err = mongoController.UpdateOne(userCollection, filter, update)
	if err != nil {
		fmt.Println("mongo updateStolenCarStatus failure")
		return
	}
	return nil
}

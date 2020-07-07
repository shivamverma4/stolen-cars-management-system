package utils

import (
	// "fmt"
	"math/rand"
	"regexp"
	// "stolencarsproject/server/config"
	"strconv"
	"time"
)

type ErrorType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomHTTPError struct {
	Error ErrorType `json:"error"`
}

type CustomHTTPResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return
}

func GenerateError(errorCode int, msg string) (_error CustomHTTPError) {
	_error = CustomHTTPError{
		Error: ErrorType{
			Code:    errorCode,
			Message: msg,
		},
	}
	return
}

func ConvertToUint(str string) uint {
	uintNumber, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(uintNumber)

}

func ConvertIntegerListToSet(listValues []uint) map[uint]bool {
	set := make(map[uint]bool)
	for _, listValue := range listValues {
		set[listValue] = true
	}
	return set
}

func ConvertStringListToSet(listValues []string) map[string]bool {
	set := make(map[string]bool)
	for _, listValue := range listValues {
		set[listValue] = true
	}
	return set
}

func RandomFloatBetweenTwoNumber(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

package parklots
/***************
Test Car Object
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestCarObject(t *testing.T) {
	
	//Creates a Car Object
	testCar :=  CreateCar("KA-01-HH-1234","White");

	//Test if Registration Number is captured correctly
	assert.Equal(t, testCar.registration_number, "KA-01-HH-1234", "Registration Number should be equal.")

	//Test if Color
	assert.Equal(t, testCar.color, "White", "Color should be same.")

	//Test if MatchRegistrationNumber returns correct values
	assert.Equal(t, MatchRegistrationNumber(testCar,"KA-01-HH-1234"), true, "It should return true.")
	assert.Equal(t, MatchRegistrationNumber(testCar,"KZ-01-HH-1234"), false, "It should return false.")

	//Test if MatchCarColor returns returns correct values
	assert.Equal(t, MatchCarColor(testCar,"White"), true, "It should return true.")
	assert.Equal(t, MatchCarColor(testCar,"Black"), false, "It should return false.")

	//Test if the Car Information in Text is correct
	assert.Equal(t, RetrieveCarStatusText(testCar), "KA-01-HH-1234 White", "There is car information in a sentence")

	//Test Empty Car
	testEmptyCar :=  CreateCar("","");
	
	assert.Equal(t, RetrieveCarStatusText(testEmptyCar), "No Car", "No Car is correct")
}
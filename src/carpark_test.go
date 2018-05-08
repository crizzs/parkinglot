package parklots
/***************
Test Carpark Object
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	
)

func TestCarparkObject(t *testing.T) {
	
	//Creates a Carpark Object
	testCarpark :=  CreateCarpark(6);

	//Test if the amount of car park slots is the same as created
	assert.Equal(t, len(testCarpark.allSlot), 6)

	//Creates a Car Object
	testCar := CreateCar("KA-01-HH-1234","White");

	//Adds to a slot not available
	updatedCarpark, availabilityCode := AddCarToSlot(testCarpark,testCar,7)

	//Signified error
	assert.Equal(t, availabilityCode, 1, "It should throw 1 which is an error.")

	//Adds a Car to Carpark Object
	updatedCarpark, availabilityCode = AddCarToSlot(testCarpark,testCar,2)

	//Availability Code shows not occupied and being claimed
	assert.Equal(t, availabilityCode, 2, "Available and being claimed.")

	//Checks if the availability checking function works ([1] in array is 2nd item)
	//Should return 3 where car is found
	assert.Equal(t, CheckSlotAvailability(updatedCarpark,2),3, "Being claimed and occupied.")
	
	//This is beyond the array size and is being caught
	assert.Equal(t, CheckSlotAvailability(updatedCarpark,7),1, "It should throw 1 which is an error.")

	//The car in slot should match with the car we get
	//This tests both adding and retrieving
	retrievedCar,retrievedCarAvailabilityCode := GetCarInSlot(updatedCarpark,2)

	assert.Equal(t, retrievedCarAvailabilityCode,3, "Being claimed and occupied.")

	//Assert this as the same car
	assert.Equal(t, testCar,retrievedCar, "Should be the same vehicle")

	//Retrieve car that does not exist
	retrievedCar,retrievedCarAvailabilityCode = GetCarInSlot(updatedCarpark,7)

	//Check for error
	assert.Equal(t, retrievedCarAvailabilityCode,1, "It should throw 1 which is an error.")

	//Removes car at no 2 slot
	updatedCarpark = RemoveCarFromSlot(updatedCarpark,2)

	//The space should be returned
	assert.Equal(t, CheckSlotAvailability(updatedCarpark,2),2, "It is free again!!.")

}
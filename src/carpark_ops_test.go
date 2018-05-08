package parklots
/***************
Test Business Logic
On Carpark Operations
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCarparkMethod(t *testing.T) {
	//Test Create Carpark with Parking slots
	carparkResultStr := Create_Parking_Lot(6)
	assert.Equal(t, carparkResultStr, "Created a parking lot with 6 slots")

	carparkResultStr = Create_Parking_Lot(-1)
	assert.Equal(t, carparkResultStr, "The amount of slots is incorrect")
}

func TestParkCarMethod(t *testing.T) {
	//Test Create Carpark with Parking slots
	Create_Parking_Lot(6)

	//Test Park Seven Cars
	Park_Car("KA-01-HH-1234","White")
	Park_Car("KA-01-HH-9999","White")
	Park_Car("KA-01-BB-0001","Black")
	Park_Car("KA-01-HH-7777","Red")
	Park_Car("KA-01-HH-2701","Blue")
	str := Park_Car("KA-01-HH-3141","Black")
	//Check output is correct when all 6 slots successfully occupied
	assert.Equal(t, str, "Allocated slot number: 6")

	carparkFullStr := Park_Car("DL-12-AA-9999","White")
	//Will show carpark is full
	assert.Equal(t, carparkFullStr, "Sorry, parking lot is full")

	//Create carpark with no slot
	Create_Parking_Lot(0)

	carparkFullStr = Park_Car("KA-01-HH-3141","Black")
	//Will show carpark is full
	assert.Equal(t, carparkFullStr, "Sorry, parking lot is full")
}

func TestLeaveCarparkMethod(t *testing.T) {
	//Initialise 6 slots and fill them
	Create_Parking_Lot(6)

	Park_Car("KA-01-HH-1234","White")
	Park_Car("KA-01-HH-9999","White")
	Park_Car("KA-01-BB-0001","Black")
	Park_Car("KA-01-HH-7777","Red")
	Park_Car("KA-01-HH-2701","Blue")
	Park_Car("KA-01-HH-3141","Black")

	//Check car has been removed
	str := Leave_Carpark(4)
	//Assert string return and slot availability
	assert.Equal(t, str, "Slot number 4 is free")

	assert.Equal(t, CheckSlotAvailability(multistorey_carpark,4),2, "Car has left and space is opened.")
	//Check space could be filled again
	parkEmptySlotStr := Park_Car("KA-01-P-333","White")
	//Correctly, it should be allocated slot 4
	assert.Equal(t, parkEmptySlotStr, "Allocated slot number: 4")

}

func TestStatusMethod(t *testing.T) {
	Park_Car("KA-01-HH-1234","White")
	Park_Car("KA-01-HH-9999","White")
	Park_Car("KA-01-BB-0001","Black")
	Park_Car("KA-01-HH-7777","Red")
	Park_Car("KA-01-HH-2701","Blue")
	Park_Car("KA-01-HH-3141","Black")

	Leave_Carpark(4)
	//Check Status Information of Carpark
	str := Status()

	//Status should print correctly
	assert.Equal(t, str, "Slot No. Registration No Colour\n1 KA-01-HH-1234 White\n2 KA-01-HH-9999 White\n3 KA-01-BB-0001 Black\n5 KA-01-HH-2701 Blue\n6 KA-01-HH-3141 Black")
	
}

func TestFindAllCarsOfSameColor(t *testing.T) {
	
	Park_Car("KA-01-P-333","White")
	//Check the color white (Find all Reg No)
	assert.Equal(t, Find_All_Cars_Of_Same_Color("White","reg"), "KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333")
	assert.Equal(t, Find_All_Cars_Of_Same_Color("Blue","reg"), "KA-01-HH-2701")

	//Check non-existence color (Find all Reg No)
	assert.Equal(t, Find_All_Cars_Of_Same_Color("Purple","reg"), "Not found")

	//Check the color white (Find all Slot Numbers)
	assert.Equal(t, Find_All_Cars_Of_Same_Color("White","slot"), "1, 2, 4")
	assert.Equal(t, Find_All_Cars_Of_Same_Color("Black","slot"), "3, 6")

	//Check non-existence color (Find all Reg No)
	assert.Equal(t, Find_All_Cars_Of_Same_Color("Purple","slot"), "Not found")


}

func TestFindParticularRegisteredNumber(t *testing.T) {
	//It should be able to find the registered number
	assert.Equal(t, Find_Slot_Num_For_Registered_Car("KA-01-HH-3141"), "6")
	assert.Equal(t, Find_Slot_Num_For_Registered_Car("KA-01-BB-0001"), "3")
	//Cannot be found example
	assert.Equal(t, Find_Slot_Num_For_Registered_Car("MH-04-AY-1111"), "Not found")
}
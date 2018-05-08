package parklots

import(
	"fmt"
	"strings"
)
/*
Check Availability of Slot
*******Legends*******
1 -> SlotNumber is bigger than carpark
2 -> Not occupied 
3 -> Car is found
*/
//Initialized and Store Carpark Object for manuveuring
var multistorey_carpark = CreateCarpark(0)
var lengthOfCarpark = len(multistorey_carpark.allSlot)
var listOfReplies = map[int]string{
									1: "Sorry, parking lot is full",
									2: "Allocated slot number: ",
									3: "The car does not exist in slot",
									4: "Slot number slotNum is free",
									5: "Created a parking lot with slotNum slots",
									6: "Not found"}
/*
This creates the Parking Lot of a given size
*/
func Create_Parking_Lot(totalAmtOfSlots int) string{
	multistorey_carpark = CreateCarpark(totalAmtOfSlots)
	
	if totalAmtOfSlots < 0 {
		return "The amount of slots is incorrect"
	}
	//Reset the length of carpark
	lengthOfCarpark = len(multistorey_carpark.allSlot)
	return strings.Replace(getRepliesText(5), "slotNum", fmt.Sprintf("%d",totalAmtOfSlots), -1)
}
/*
Parking of car into carpark
*/
func Park_Car(reg_no string,color string) string{
	//function variables for parking operation 
	incomingCar := CreateCar(reg_no,color);
	
	currentSlotNumber := 1 

	for lengthOfCarpark >= currentSlotNumber{

		availabilityCode := CheckSlotAvailability(multistorey_carpark,currentSlotNumber)

		if availabilityCode == 1 {
			return getRepliesText(1)
		}else if availabilityCode == 2{
			carparkUpdated, _ := AddCarToSlot(multistorey_carpark,incomingCar,currentSlotNumber)
			multistorey_carpark = carparkUpdated
			return getRepliesText(2) + fmt.Sprintf("%d",currentSlotNumber)
		}else{
			currentSlotNumber++
		}
	}
	//When the for loop runs out, there is no slot
	return getRepliesText(1) 
}
/*
Function to remove car from carpark when it leaves
*/
func Leave_Carpark(slotNum int) string{
	availNum := CheckSlotAvailability(multistorey_carpark,slotNum)

	if slotNum <= 0 ||  availNum == 1  || availNum == 2 {
		return getRepliesText(3)
	}
	//Otherwise we will remove the car from slot	
	carparkUpdated  := RemoveCarFromSlot(multistorey_carpark,slotNum)
	multistorey_carpark = carparkUpdated

	return strings.Replace(getRepliesText(4), "slotNum", fmt.Sprintf("%d",slotNum), -1);
}
/*
Function is to show the status of the entire carpark
*/
func Status() string{
	currentSlotNumber := 1
	
	statusStr := "Slot No. Registration No Colour"

	for lengthOfCarpark >= currentSlotNumber{
		retrievedCar,availNum := GetCarInSlot(multistorey_carpark,currentSlotNumber)
		//Got car
		if availNum == 3 {
			statusStr += "\n"+ fmt.Sprintf("%d",currentSlotNumber) + " " + RetrieveCarStatusText(retrievedCar)
		}
		currentSlotNumber++
	}
	return statusStr
}
/*
Find all cars that are of the same color
Usable for both Slot Number or Registration Number finding
slot -> Slot Number
reg -> Registration Number

(Refactored to cater for both scenarios)
*/
func Find_All_Cars_Of_Same_Color(color string,findSlotNoOrRegNo string) string{
	currentSlotNumber := 1
	foundCarsOfSameColor := ""

	for lengthOfCarpark >= currentSlotNumber{
		availNum := CheckSlotAvailability(multistorey_carpark,currentSlotNumber)
		if availNum == 3 {
			//If so, retrieve car to examine its color
			retrievedCar,_ := GetCarInSlot(multistorey_carpark,currentSlotNumber)
			checkColor := MatchCarColor(retrievedCar,color)
			
			//For text formatting
			if checkColor && foundCarsOfSameColor == "" && findSlotNoOrRegNo == "reg" {
				foundCarsOfSameColor += retrievedCar.registration_number 
			}else if checkColor && foundCarsOfSameColor != ""  && findSlotNoOrRegNo == "reg" {
				foundCarsOfSameColor += ", "+retrievedCar.registration_number 
			}

			if checkColor && foundCarsOfSameColor == "" && findSlotNoOrRegNo == "slot" {
				foundCarsOfSameColor += fmt.Sprintf("%d",currentSlotNumber)
			}else if checkColor && foundCarsOfSameColor != ""  && findSlotNoOrRegNo == "slot" {
				foundCarsOfSameColor += ", "+fmt.Sprintf("%d",currentSlotNumber)
			}
		}
		currentSlotNumber++
	}
	if(foundCarsOfSameColor == ""){
		return getRepliesText(6)
	}
	return foundCarsOfSameColor
}
/*
Find the slot number for a registered car number
*/
func Find_Slot_Num_For_Registered_Car(registeredNum string) string{
	currentSlotNumber := 1

	//This is being brought for integer manipulation
	var numOfCarpark = len(multistorey_carpark.allSlot)
	/*
	To search for slot from backward and forward towards middle point with Big O Notation (n/2)
	*/
	for numOfCarpark >= currentSlotNumber {
		retrievedForwardCar,availForwardNum := GetCarInSlot(multistorey_carpark,currentSlotNumber)
		retrievedBackwardCar,availBackwardNum := GetCarInSlot(multistorey_carpark,numOfCarpark)

		//Forward
		if availForwardNum == 3 {
			checkMatchForwardRegNum := MatchRegistrationNumber(retrievedForwardCar,registeredNum)
			if checkMatchForwardRegNum {
				return fmt.Sprintf("%d",currentSlotNumber)
			}
		}
		//Backward
		if availBackwardNum == 3 {
			checkMatchBackwardRegNum := MatchRegistrationNumber(retrievedBackwardCar,registeredNum)
			if checkMatchBackwardRegNum {
				return fmt.Sprintf("%d",numOfCarpark)
			}
		}
		currentSlotNumber++
		numOfCarpark--	
	}
	return getRepliesText(6)
}
/*
This will generate a reply based on number code
*/
func getRepliesText(txtNum int) string{
	return listOfReplies[txtNum];
}
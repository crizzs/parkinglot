package parklots

/***************
Created Model 
class for carpark
slots
****************/
type carpark struct {  
	allSlot []car
}
//Creates a Carpark Object
func CreateCarpark(totalAmountOfSlots int) carpark{
	if totalAmountOfSlots < 0 {
		//By default is nothing
		created_carpark := carpark{make([]car, 0)}
		return created_carpark
	}

	created_carpark := carpark{make([]car, totalAmountOfSlots)}
	return created_carpark
}
//Adds a Car Object to a Carpark slot
func AddCarToSlot(c carpark,carObject car,slotNumber int) (carpark,int){
	 
	//Checks if slot existed for vehicle (Different Scenarios)
	availabilityResult := CheckSlotAvailability(c,slotNumber)
	//Process the adding of car
	if availabilityResult == 2 {
		//Declaration of Slot to a car
		c.allSlot[slotNumber-1] = carObject
		return c, availabilityResult
	}else{
		return c,availabilityResult
	}	
}
//Removes a Car Object from a slot
func RemoveCarFromSlot(c carpark,slotNumber int) carpark{
	c.allSlot[slotNumber-1] = car{}
	return c
}
/*
Returns Car in Slot (SlotNumber Minus 1) -Array starts from 0
*/
func GetCarInSlot(c carpark,slotNumber int) (car,int){
	
	availabilityResult := CheckSlotAvailability(c,slotNumber)

	//Checks if the slot number existed in reality
	if len(c.allSlot) == 0 || len(c.allSlot) < slotNumber {
		return car{},availabilityResult
	}else{
		return c.allSlot[slotNumber-1],availabilityResult
	}
}
/*
Check Availability of Slot
*******Legends*******
1 -> SlotNumber is bigger than carpark
2 -> Not occupied 
3 -> Car is found
*/
func CheckSlotAvailability(c carpark,slotNumber int) int{
	var actualArrayPosition = slotNumber-1
	if len(c.allSlot) == 0 || len(c.allSlot) < slotNumber {
		return 1
	}else if c.allSlot[actualArrayPosition].registration_number == "" && c.allSlot[actualArrayPosition].color == ""{
		return 2
	}else{
		return 3
	}
}
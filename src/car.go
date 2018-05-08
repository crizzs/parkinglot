package parklots

/***************
Created Model 
class for car
****************/
type car struct {  
	registration_number string
	color string
}
//Creates a Car Object
func CreateCar(registration_number string,color string) car{
	created_car := car{registration_number,color}
	return created_car
}
//Match the Registration Number
func MatchRegistrationNumber(c car,input_registration_number string)  bool{
	if c.registration_number == input_registration_number{ 
		return true
	}else{
		return false
	}
}
//Match the Car Color
func MatchCarColor(c car,input_color string) bool{
	if c.color == input_color{ 
		return true
	}else{
		return false
	}
}
//Retrieve Car Information in Text Format
func RetrieveCarStatusText(c car) string{
	if c.registration_number == ""{
		return "No Car"
	}
	return c.registration_number + " " + c.color
}
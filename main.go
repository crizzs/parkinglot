/*****************
Contains the main method
of Parking Lot Application
*****************/
package main

import (
	"log"
	"flag"
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	parklots "./src"
)

func main(){
	//Parse the flag for either file input or cli inputs
	flag.Parse()
	/*********
	There should only be one flag which is the file path
	**********/
	if len(flag.Args()) == 1 {
		outputStr := FileProcessor(flag.Args()[0])
		//Displays Result to Users
		fmt.Print(outputStr)
		return 
	}
	//Activate Interactive Command Line Terminal
	ActivateInteractiveCommandTerminal()
}
/*
This reads the file and process parking commands
*/
func FileProcessor(path string) string{
	//Check if path existed
	var _, err = os.Stat(path)
	
	//If file does not exist, activate interactive command terminal
	if os.IsNotExist(err) {
		ActivateInteractiveCommandTerminal()
	}

	//Once file is checked to exist, we shall proceed to open the file
	var file, openFileErr = os.Open(path)
	checkError(openFileErr)

	defer file.Close()
	//Scan the file line by line
	scanner := bufio.NewScanner(file)
	displayStr := ""
	lineCounter := 1

	for scanner.Scan() {
		//Remove Carriage Return
		var eachCommand = trimRN(scanner.Text());

		if lineCounter == 1 {
			//Split by Spaces
			var parseFirstCommand = strings.Split(eachCommand," ")
			//Check for valid carpark slots creation
			if parseFirstCommand[0] != "create_parking_lot" {
				return "You must first create parking slots for your carpark."
			}
			if _, err := strconv.Atoi(parseFirstCommand[1]); err != nil {
				return "The parking slots creation needs an integer."
			}
			strReturn,checkValidity := ProcessCommand(eachCommand)
			if checkValidity{
				displayStr += strReturn
			}
		}else{
			//Concat for final outputs
			strReturn,checkValidity := ProcessCommand(eachCommand)
			if checkValidity{
				displayStr += "\n"+strReturn
			}
		}
		
		lineCounter++
	}
	//Wrong case is an empty string display
    return displayStr
}
/*
This is the interactive command terminal to interact with the program
*/
func ActivateInteractiveCommandTerminal(){

	commandReader := bufio.NewScanner(os.Stdin)
	trackLineCounter := 1
	trackCarparkCreated := false
	fmt.Println("Input:")
	for commandReader.Scan() {

		var commandRetrieved = commandReader.Text()
		if commandRetrieved == "bye"{
			break
		}
		if trackLineCounter == 1 {
			//Split by Spaces
			var parseFirstCommand = strings.Split(commandRetrieved," ")
			//Check for valid carpark slots creation
			if parseFirstCommand[0] != "create_parking_lot" {
				fmt.Println("Note: You must first create parking slots for your carpark.")
			}
			if len(parseFirstCommand) == 2{
				slotNum := convertToInt(parseFirstCommand[1])
				if slotNum == -1 ||  slotNum == 0{
					fmt.Println("Note: The parking slots creation needs an integer bigger than 0.")
				}
				strReturn,checkValidity := ProcessCommand(commandRetrieved)
				if checkValidity && slotNum > 0{
					fmt.Println("\nOutput:\n"+strReturn+"\n\nInput:")
					trackCarparkCreated = true
					trackLineCounter++
				}
			}else{
				fmt.Println("\nInput:")
			}
		}else if trackLineCounter> 1 && trackCarparkCreated == true{
			strReturn,checkValidity := ProcessCommand(commandRetrieved)
			if checkValidity{
					fmt.Println("\nOutput:\n"+strReturn+"\n\nInput:")
					trackLineCounter++
			}
		}
	}
}
//Parse and Process the commands in accordance to requirement 
//Boolean is to ensure there is good return for concatenation
func ProcessCommand(commandStr string) (string,bool){
	var splitCommand = strings.Split(commandStr," ")

	var actionInstigated = splitCommand[0]

	//All the scenarios of a carpark operations (Ensured it is safe for split array range)
	if actionInstigated == "create_parking_lot"  && len(splitCommand)==2{
		slotNum := convertToInt(splitCommand[1]) 
		if slotNum == -1 {
				return "Note: An number is required to activate this feature.",true
		}
		return parklots.Create_Parking_Lot(slotNum),true

	}else if actionInstigated == "park" && len(splitCommand)==3{
		return parklots.Park_Car(splitCommand[1],splitCommand[2]),true
	}else if actionInstigated == "leave"  && len(splitCommand)==2{
		slotNum := convertToInt(splitCommand[1]) 
		if slotNum == -1 {
				return "Note: An number is required to activate this feature.",true
		}
		return parklots.Leave_Carpark(slotNum),true
	}else if actionInstigated == "status" {
		return parklots.Status(),true
	}else if actionInstigated == "registration_numbers_for_cars_with_colour"  && len(splitCommand)==2{
		return parklots.Find_All_Cars_Of_Same_Color(splitCommand[1],"reg"),true
	}else if actionInstigated == "slot_numbers_for_cars_with_colour"  && len(splitCommand)==2{
		return parklots.Find_All_Cars_Of_Same_Color(splitCommand[1],"slot"),true
	}else if actionInstigated == "slot_number_for_registration_number"  && len(splitCommand)==2{
		return parklots.Find_Slot_Num_For_Registered_Car(splitCommand[1]),true
	}else{
		return "" , false
	}
	return "",false
}
/*
Function to convert string to integer
*/
func convertToInt(str string) int{
	slotNum, err := strconv.Atoi(str); 
		if err != nil {
			return -1
	}
	return slotNum
}
/*
Carriage Return stripper
*/
func trimRN(str string) string{
	return strings.TrimRight(str,"\r\n");
}
/*
Checks for errors and log it
*/
func checkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}
package main
/***************
Test Main Package
****************/
import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestFileInput(t *testing.T) {
	assert.Equal(t, FileProcessor("sample_file/file_inputs.txt"), "Created a parking lot with 6 slots\nAllocated slot number: 1\nAllocated slot number: 2\nAllocated slot number: 3\nAllocated slot number: 4\nAllocated slot number: 5\nAllocated slot number: 6\nSlot number 4 is free\nSlot No. Registration No Colour\n1 KA-01-HH-1234 White\n2 KA-01-HH-9999 White\n3 KA-01-BB-0001 Black\n5 KA-01-HH-2701 Blue\n6 KA-01-HH-3141 Black\nAllocated slot number: 4\nSorry, parking lot is full\nKA-01-HH-1234, KA-01-HH-9999, KA-01-P-333\n1, 2, 4\n6\nNot found")
	//This will produce a output required on the cmd screen
	assert.Equal(t, FileProcessor("sample_file/errors_input.txt"), "The parking slots creation needs an integer.")
}

func TestInteractiveTerminal(t *testing.T) {
	ActivateInteractiveCommandTerminal()
}

func TestProcessCommandFunction(t *testing.T) {
	//Test all seven commands (File Input)
	outputOne,_ := ProcessCommand("create_parking_lot 6")
	outputTwo,_ := ProcessCommand("park KA-01-HH-1234 White")
	outputThree,_ := ProcessCommand("leave 4")
	outputFour,_ := ProcessCommand("status")
	outputFive,_ := ProcessCommand("registration_numbers_for_cars_with_colour White")
	outputSix,_ := ProcessCommand("slot_numbers_for_cars_with_colour White")
	outputSeven,_ := ProcessCommand("slot_number_for_registration_number KA-01-HH-3141")

	//Asserts all Values for File Inputs on Test Process
	assert.Equal(t, outputOne , "Created a parking lot with 6 slots")
	assert.Equal(t, outputTwo , "Allocated slot number: 1")
	assert.Equal(t, outputThree , "The car does not exist in slot")
	assert.Equal(t, outputFour , "Slot No. Registration No Colour\n1 KA-01-HH-1234 White")
	assert.Equal(t, outputFive , "KA-01-HH-1234")
	assert.Equal(t, outputSix , "1")
	assert.Equal(t, outputSeven , "Not found")
}	
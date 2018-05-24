# Parking Lot Application

This application can be easily cross-compiled into other OS distribution using Golang cross-compiler.

<font color="#0099ff" size=12 >Folder Structure</font>
-----------------
|--parkingapp （Directory）    
|&emsp;&emsp;|-------src（Contains the parklots package and unit testing scripts)  
|&emsp;&emsp;|-------sample_file  (The file to test output)   
|&emsp;&emsp;|-------main （Executable for Mac OS）   
|&emsp;&emsp;|-------main.go （Entry point）
|&emsp;&emsp;|-------parking_lot.sh （A shell script to run GO test and starts application）  
|-------|-Testify Package for TDD @ https://github.com/stretchr/testify  


How to start?
-----------------
Use Makefile to build
### 1. `go get github.com/crizzs/parkinglot`
### 2. Type `make` - This will run the test scripts for UNIT and Main Package testing 
### 3A. `./main` - This will start the interactive command terminal
### 3B. `./main sample_file/file_inputs.txt` - This will run the file input (You can change the file directory)


-----------------------------------------


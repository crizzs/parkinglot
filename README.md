# Parking Lot Application

Please place the application into your `$GOPATH/src`. You don't require any added dependencies except for testify.

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

### 1. `Place the folder into your GO workspace`
### 2. `./parking_lot.sh` - This will run the test scripts for UNIT and Main Package testing 
### 3A. `go run main.go` - This will start the interactive command terminal
### 3B. `go run main.go sample_file/file_inputs.txt` - This will run the file input (You can change the file directory)


-----------------------------------------


# learning_golang
Learning Go Language
A new Journey to Go...

- Static typed language

- first.go
  - It is just a hello world program
- second.go
  - It is just how we create a function and import packages
 - third.go
  - It Shows how we can define a data type to a given variable, paramters or return type
  - Three different syntax to look up on
  	 - Syntax 1 
  	 	We define the varaible in two different lines
  	 - Syntax 2 
  	  	We define in one line and if the data type is same we assign it at the end
  	  - Syntax 3 
  	    We let go define the varaible but the data type cannot be change once it is compiled. 
  	    Basically it does not support dynamic data type.
     E.g.
     ``` go
     //Syntax 1:
		var num1 float64 = 5.7
		var num2 float64 = 9.7
	 //Syntax 2:
		var num1, num2, float64 = 5.7, 9.7
	 //Syntax 3:
	 	num1, num2 := 5.7, 9.7
     ```
 - fourth.go 
   Example of pointers in Go it is pretty similar to the pointer in C
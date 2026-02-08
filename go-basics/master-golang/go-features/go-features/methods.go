// //Methods are define inside struct here and Function is like normal function.package practice

// package main 

// import(
// 	"fmt"
// )


// type Employee struct{
// 	Name string 
// } 

// //Mehthods (Expanding the struct)
// func (e *Employee) updateName(newName string){
// 	e.Name = newName
// }

// func (e *Employee) printName(){
// 	fmt.Println(e.Name)
// }


// func main(){
// var emp Employee
// emp.Name = "Alok Tripathi"

// emp.printName()

// emp.updateName("Mudit Kapoor")

// emp.printName()

// }



package main

import(
	"fmt"
)


type Employee struct{
	Name string 
	Age int16
}


func (e *Employee) updateName(newName string){
	e.Name=newName
}

func (e *Employee) updateAge(newAge int16){
	e.Age = newAge
}


func (e *Employee) logEmopoyee(){
	fmt.Printf("Employee Name = %s  Age = %v \n", e.Name, e.Age)
}



func main(){
	var emp Employee

	emp.Name = "Alok"
	emp.Age = 21

	emp.logEmopoyee()


	emp.Name = "Mudit"
	emp.Age = 22

	emp.logEmopoyee()

	emp.Name = "Debs"
	emp.Age = 20

	emp.logEmopoyee()
}
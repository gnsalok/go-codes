package main 


import(
	"fmt"
	"math"
)

func main(){
rect := Rectangle{20,50}
cir := Circle{2}


//This is the way to get the resutlt via receiver function or methods.
fmt.Println("Area of Rectangle = ", rect.area())
fmt.Println("Area of Cirle is = ",cir.area())

 
//This is Polymorphism   (we got the same result) 
fmt.Println("Area of Rectangle =",getArea(rect))
fmt.Println("Area of Circle =",getArea(cir))

}

type Shape interface{
	 area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

func(r Rectangle) area() float64{
	return r.height * r.width
}

func(c Circle) area() float64{
	return math.Pi * math.Pow(c.radius,2)
}

func getArea(shape Shape) float64{
	return shape.area()
}
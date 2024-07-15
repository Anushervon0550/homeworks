package main

import (
	
	"fmt"
)

func main() {
//Zadanie 1
	PrintGeerting()
	DisplaySeparator()
	NotifyStart()
//Zadanie 2
	GetWelcomeMessage()
	fmt.Println(GetWelcomeMessage())

	GetPiValue()
	fmt.Println(GetPiValue())

	IsServerActive()
	fmt.Println(IsServerActive())
//Zadanie 3
	
	PrintNumber(4)
	GreetUser("Anushervon")
	ToggleLight(true)
//Zadanie 4
 	first,second:=4 ,5

	Add(first,second )
	fmt.Println(Add(first,second))

	C:=Concat("Salon ","barodar!")
	fmt.Println(C)

	fmt.Println(IsEven(60))
	
//Zadanie 5
	adder:=func(a, b int)int{
		 c:=a + b 
		return c
}
	fmt.Println(adder(4,6))

	concatenator:=func(a,b string)string{
		c:=a+b
		return c
	}
	fmt.Println(concatenator("Hello","what time is it?"))
	
	isPositive:=func(a int)bool{
		if a >0 {
			return true
		} else {
			return false
		}
}

	fmt.Println(isPositive(5))
//Zadanie 6
fmt.Println(Calculate(45, 65,addc))
}
//Zadanie 1
func PrintGeerting() {
	fmt.Println("Hello, World!")
}
func DisplaySeparator(){
	fmt.Println("Hello my name is Anush")
}
func NotifyStart(){
	 fmt.Println("Program started")
}
//Zadanie 2
func GetWelcomeMessage()string{
	return "Welcome!"
}
func GetPiValue()float64{
	pi:=3.14
	return pi
}
func IsServerActive()bool{
	return true
}
//Zadanie 3
func PrintNumber(number int){
	fmt.Println(number)
}
func  GreetUser(name string){
	fmt.Printf("Hello %s\n",name)
}
func  ToggleLight(toggle bool){
	if toggle{
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
}
}
//Zadanie 4
func Add(first,second int)int{
	sum:=first+second
	return sum
}
func Concat(a1,a2 string)string{
	return a1+a2
}
func IsEven(b int)bool{
 if b%2 ==0 {
	return true
 } else {
	return false
 }
}
func addc(a, b int) int {
	return a + b
}

func Calculate(a, b int, f func(a, b int) int) int {
	ravno := f(a, b)
	return ravno
}

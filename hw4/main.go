package main

import "fmt"

import (
	"fmt"
)

func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	GetGrade()
	GetDiscount()
	GetTemperatureDescription()
	CheckNumber()
	CheckAge()
	CheckPassword()
	println(Add(5, 343))
	CompareStrings()
	Max()
	operation:=func(a, b int)int  {
		return a-b
	}
	a,b:=3,6
	result:=operation(a,b)
	fmt.Println(a,b,result)
	a1:="что небудь"
	b1:="!!!!"
	concat:=func (a1, b1 string) string {
		
		c:=a1+b1
		return c
		
	}
	fmt.Println(concat(a1, b1))
		n1:=114
		n2:=45
	multiply:=func (n1, n2 int) int {
		return n1+n2
	}
		
		result1:=multiply(n1,n2)
	fmt.Println(result1)
		
	}

}

func PrintGreeting() {
var dayType string
fmt.Scan(&dayType)
switch dayType{
case "утро":
	println("Доброе утро!")
case "день":
	println("Добрый день!")
case "вечер":
	println("Добрый вечер!")
}

}

func PrintWeather(){
var weatherType string
fmt.Scan(&weatherType)
switch weatherType{
case "солнично":
	println("Солнично")
case "облочно":
	println("Облочно")
case "дождливо":
	println("Дождливо")
}
 }

func PrintTrafficLight(){
var lightColor string
fmt.Scan(&lightColor)
switch lightColor{
case "красный":

	println("Стоп")

case "желтый":

	println("Внимание")

 case "зеленый":

		println("Идите")
	}

 }

	func GetGrade(){
		var score uint8

fmt.Scan(&score)
switch score{
case 5:

	println("A")

case 4:

	println("B")

case 3:

		println("C")
	}

}
func GetDiscount(){
var amount float64
fmt.Scan(&amount)

 	if amount >1000 {
 		println("скидка 10%")
	}else if amount>500{

		println("скидка 5%")
	}else   {

		println("скидка 0%")
	}

 }

	func GetTemperatureDescription(){
		var temperature int
		fmt.Scan(&temperature)

	if temperature<=10{
		println("Холодно")
	}else if temperature>10 && temperature<=25{

		println("Тепло")
	}else {

		println("Жарко")
	}

}
func CheckNumber(){
var a int
fmt.Scan(&a)

	if  a>0 {
		println("Положительное")
	}else if a<0 {

		println("Отрицательное")
	}else{

		println("Ноль")
	}

}
func  CheckAge(){
var b uint
fmt.Scan(&b)

	if b>17 {
		println("Совершеннолетний")
	}else  {

		println("Несовершеннолетний")
	}

}
func CheckPassword() {
	var a int
	fmt.Scan(&a)
	if a == 1234 {
		println("Пароль верный")
	} else {
		println("Пароль неверный")
	}
}

func Add(a, b int)int{
 s:=a+b
 if s<0{
	s=-s
 }
 return s

}
func CompareStrings() {
	a := 65
	b :=564
if a==b {
	println("равны")
}else{
	println("не равны")
}
 }
func Max(){
	var a int=34534
	var b int=345
if a>b {
	fmt.Println(a)
}else if b>a {
	fmt.Println(b)
}else {
	println("равны")
}

}

package main

import "fmt"

func main() {
	println(Temp(4))
	println(vozr(18))
	println(scor(1))
	println(sco(5))
	println(bmi(91))
	println(kw(6))
	println(udw(5))
	println(check(6))
	println(check1(5))
	cou(56)
	print(battary(99))
	println(wes(60))
	println(uspi(60))
	// println(wess())
	p1:= Product{Name: "Nike", Price: 155}
	fmt.Println(p1)
	fmt.Printf("%+v\n",p1)
	// a,b:=12,15

	person := Person{
		FirstName: "Иван",
		LastName:  "Петров",
		Age:       30,
	}
	per(person)
}

type Temperatur float64

func Temp(tem Temperatur) string {
	if tem > 0 {
		return "выше 0"
	} else if tem < 0 {
		return "ниже 0"
	} else {
		return "равно 0"
	}
}

type Age int

func vozr(age Age) string {
	if age <= 19 || age >= 13 {
		return "подросток"
	} else {
		return "не подросток"
	}
}

type Speed float64

func scor(speed Speed) string {
	if speed <= 120 || speed < 0 {
		return "Допустимая скорость"
	} else {
		return "не допустимая скорость"
	}
}

type Score int

func sco(num Score) string {
	if num > 0 {
		return "положительный"
	} else if num < 0 {
		return "отрицательнай"
	} else {
		return "нулевое"
	}
}

type BMI float64

func bmi(num BMI) string {
	if num >= 0 && num <= 40 {
		return "Underweight"
	} else if num >= 41 && num <= 75 {
		return "Normal weight"
	} else if num >= 80 && num < 100 {
		return "Overweight"
	} else {
		return "Obesity"
	}
}

type UnaryOperation int

func kw(a UnaryOperation) int {
	c := a * a
	return int(c)
}
func udw(num UnaryOperation) int {
	c := num * 2
	return int(c)
}

type Check int

func check(num Check) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
func check1(num Check) bool {
	if num > 0 {
		return true
	} else {
		return false
	}
}

// type Operation int

//	func sum(a,b Operation)int  {
//		return a-b
//	}
//
//	func  min(a,b Operation )int  {
//		return a-b
//	}
//
//	func  umn(a,b Operation)int  {
//		return a*b
//	}
type Count = int

func cou(num Count) {
	for num > 0 {
		fmt.Println(num)
		num--
	}
}
type BatteryLevel=int
func battary(num BatteryLevel)string{
	if num>=0 && num<=40{
	return "низкий уровень заряда"
}else if num>40 && num<70 {
	return "средний уровень заряда"
}else if num>=70 && num<=100{
	return "высокий уровень заряда"
}else{
	return "error"
}

}
type Weight float64
func wes(num Weight )string{
	if num>20 && num<40{
	return "Light"
}else if num>40 && num<65{
	return "Medium"
}else{
	return "Heavy"
}
}
type Grade int
func uspi(num Grade)string{
	if num>=50{
		return "удовлетворительной"
	} else{
		return "не удовлетворительной"
	}
}
// type BloodPressure float64
// func wess(num1 BMI, num2 float64 )string{
// 	if num1>=45; num2>=60 && num1<=65;  num2<=80 {
// 		 return "Healthy"
// 	}else if num1>=65; num2>=80 && num1<=100 ; num2<=100{
// 		return  "At risk" 
// 	}else {
// 		return "Unhealthy"
// 	}
// }
type Product struct {
	Name   string 
	Price  float64
}
type Person struct{
	FirstName string
	LastName string 
	Age int
}
func per(p Person){
	fmt.Printf("Полное имя: %s %s\n", p.FirstName, p.LastName)
	fmt.Printf("Возраст: %d лет\n", p.Age)
}
type Product struct{
	Name string
	Price int
}
func prod(p Product)
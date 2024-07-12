package main
import (
	"fmt"
	"math"
)

func main() {
	Zadanie1()
	Zadanie2()
	Zadanie3()
	Zadanie4()
	Zadanie5()
	Zadanie6()
	Zadanie7()
	Zadanie8()
	Zadanie9()
	Zadanie10()
	Zadanie11()
	Zadanie12()
	Zadanie13()
	Zadanie14()
	Zadanie15()
}
func Zadanie1() {
	//1) Дана сторона квадрата a. Найти его периметр P = 4·a.

	var num1 string = "Задача 1 (P)="

	var a = 5

	P := 4 * a

	fmt.Println(num1, P)
}
func Zadanie2(){
	
	//2) Дана сторона квадрата a. Найти его площадь S = a

	var num string="Задача 2 (S)="

	var a float64=8
	
	S:=a*a

	fmt.Println(num ,S)
}
func Zadanie3(){
		//3)Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b).

		var a float64=8
	
		var b float64=9
	
		var num1 string="Задача 3 (S)="

		S:=a*b
		var num2 string="Задача 3 (P) ="
		P:=2*(a+b)
		fmt.Println(num1, S, num2, P)
}
func Zadanie4(){
	//4) Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14.

	var num string="Задача 4 (L)="
	 var  pi=3.14

	var d float64=6
	
	L:=pi*d

	fmt.Println(num ,L)

}
func Zadanie5(){
//5) Дана длина ребра куба a. Найти объем куба V = a3 и площадь его поверхности S = 6·a

var a float64=8

var num string="Задача 5 (V)="

V:=a*a*a

var num1 string="Задача 5 (S)="

S:=6*a

fmt.Println(num , V, num1, S)
}
func Zadanie6(){
//6) Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).

var a = 34
var b = 45
var c = 24

var num string="Задача 6 (V)="

V:=a*b*c

var num1 string="Задача 6 (S)="

S:=2*(a*b)+(b*c)+(a*c)

fmt.Println(num, V, num1, S)

}
func Zadanie7(){
	// 7) Найти длину окружности L и площадь круга S заданного радиуса R: L = 2·π·R, S = π·R

	var pi= 3.14

	var R float64= 5

	var num string="Задача 7 (S)="

	S:=pi*R

	var num1 string="Задача 7 (L)="

	L:=2*pi*R

	fmt.Println(num, S ,num1, L)
}

func Zadanie8 (){
//8) Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.

	var a =65
	
	var b =68

	var num string="Задача 8 (Sr)="

	Sr:=(a*b)/2
	fmt.Println(num, Sr)
}
func Zadanie9(){
//9) Даны два неотрицательных числа a и b. Найти их среднее геометрическое, то есть квадратный корень из их произведения: √a·b.

	var a float64=8

	var b float64=9

	var num string="Задача 9 (c)="

	c:=math.Sqrt(a*b)

	fmt.Println(num,c)
}
func Zadanie10(){
//10) Даны два ненулевых числа. Найти сумму, разность, произведение и частное их квадратов
	
var num string="Задача 10 " 
	
		var  a uint=65
	
	var  b uint=95
	
	var num1 string="sum ="
	
	sum:=a+b
	
	var num2 string="r ="
	
	r:=a-b
	
	var num3 string="p ="
	
	p:=(a+b)^2
	
	var num4 string="ch ="
	
	ch:=a^2+b^2

	fmt.Println(num, num1, sum, num2, r ,num3, p, num4, ch)
}
func Zadanie11(){
//11) Дано расстояние L в сантиметрах. Используя операцию деления нацело, найти количество полных метров в нем (1 метр = 100 см).
	var num1 string="Задача 11 "
	
	var L float64=150

 	Metr:=L/100 

 	var num string="Metr ="

 	fmt.Println(num1,num,Metr)
}
func Zadanie12(){
//12) Дана масса M в килограммах. Используя операцию деления нацело, найти количество полных тонн в ней (1 тонна = 1000 кг).
	var num string="Задача 12 "

	var kg float64=456

	tonna:=kg/1000

	var num1 string="Tonna ="

	fmt.Println(num,num1,tonna)
}
func Zadanie13(){
//13) Дан размер файла в байтах. Используя операцию деления нацело, найти количество полных килобайтов, которые занимает данный файл (1 килобайт = 1024 байта).
	var num string="Задача 13 "

	var b float64=57567

	KB:=b/1024

	var num1 string="KB ="

	fmt.Println(num,num1,KB)
}
func Zadanie14(){
//14) Даны целые положительные числа A и B (A > B). На отрезке длины A размещено максимально возможное количество отрезков длины B (без наложений). Используя операцию деления нацело, найти количество отрезков B, размещенных на отрезке A.
	var num string="Задача 14 "

	var a float64=6657

	var b float64=45
	
	var num1 string="c ="

	c:=a/b

	fmt.Println(num,num1,c)
}
func Zadanie15(){
//15) Даны целые положительные числа A и B (A > B). На отрезке длины A размещено максимально возможное количество отрезков длины B (без наложений). Используя операцию взятия остатка от деления нацело, найти длину незанятой части отрезка A.

var num string="Задача 15 "

var a =45645

var b =5676

	var num1 string="c ="

	c:=a/b
	
	var num2 string="d ="
	
	d:=a%b

	fmt.Println(num,num1,c,num2,d)
}

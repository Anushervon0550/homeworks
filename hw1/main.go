package main
import "fmt"

func main() {
// int-ы
	var num int
	num=6
	fmt.Println(num)

	var num2 int8
	num2=120
	fmt.Println(num2)

	var num3 int16
	num3=-6343
	fmt.Println(num3)

	var num4 int32
	num4=-533453
	fmt.Println(num4)

	var num5 int64
	num5=64564574574
	fmt.Println(num5)

// float-ы

	var num6 float32=678.8867
	fmt.Println(num6) 
	
	var num7 float64=8789.7897897
	fmt.Println(num7)

//unit-ы

	var num8 uint=657
	fmt.Println(num8)

	var num9 uint8=254
	fmt.Println(num9)
	
	var num10 uint16=344
	fmt.Println(num10)
	
	var num11 uint32=534646
	fmt.Println(num11)
	
	var num12 uint64=623536
	fmt.Println(num12)
	
//bool
	
	var B bool=true
	fmt.Println(B)

	var a bool=true
	fmt.Println(a)
	
//rune

	var c rune=56
	fmt.Println(c)

//byte

	var d byte=1
	fmt.Println(d)

//string

	var name string="первый проект"
	fmt.Println(name)
//complex

var name1 complex64=4+3i
fmt.Println(name1)

var name2 complex128=1+2i
fmt.Println(name2)
//Arifmetica

	var num14 int=8
	
	var num15 int8=45
	
	num16:=int8(num14)+num15	

	num17:=int8(num14)-num15

	num18:=int8(num14)/num15

	num19:=int8(num14)*num15

	fmt.Println(num16,num17,num18,num19)


	var num20 int=46
	
	var num21 int16=455
	
	num22:=int16(num20)+num21	

	num23:=int16(num20)-num21

	num24:=int16(num20)/num21

	num25:=int16(num20)*num21

	fmt.Println(num22,num23,num24,num25)
 


}
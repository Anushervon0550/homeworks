package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 43, 53, 34, 455, 34}
	max :=ma(a)
	fmt.Printf("Максимальный элемент в массиве: %d\n",max)

	
	b := []int{8,4,5,7,8,90,5,3,2,}
	mi := fi(b)
	fmt.Printf("Минимальный элемент в массиве: %d\n", mi)

		
	c := []int{-1, 1, 2, 6, 3, -12, 11,2, -3, 4}
	t := cp(c)
	fmt.Printf("Количество положительных чисел в массиве: %d\n", t)
}
func ma(a []int)int{
	m:=  a[0]
	for _, b := range a {
		if b>m {
			m=b
		}
	}
	return m
}
func fi(b []int) int {
	m := b[0]
	for _, a := range b {
		if a < m {
			m = a
		}
	}
	return m
}
func cp(a []int) int {
	b := 0
	for _, c := range a {
		if c > 0 {
			b++
		}
	}
	return b
}
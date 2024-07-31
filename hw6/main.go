package main

import "fmt"

func main() {
	sav(100000, 50000)

	cred(3000000, 300000)

	tran(1, 1)

	contribution(1)}
	func sav(a, b int) {
		a = 100000
		println("Начальный баланс равен ", a)
		b = 50000
		println("ежемесечное пополнение равно ", b)
		for i := 1; ; i++ {
			c := a + b*i
			println("за ", i, " месяц обший баланс равен ", c)
			if c >= 500000 {
				println("Достигнут лимит накоплений")
				break
			}
		}
	
	}
	func cred(a, b int) {
		a = 3000000
		println("ваш кредитный счет равен ", a)
		b = 300000
		println(" фиксированная ставка в 10 % ежемесечных платежей равно ", b)
		for i := 1; ; i++ {
			c := a - b*i
			println("остаток вашего кридитного счета равен ", c)
			if c < 500000 {
				println("кредит почти погашен ")
				break
			}
		}
	
	}
	
func tran(a, b int) {
	a = 500000
	println("Начальный баланс счета равен ", a)
	fmt.Scan(&b)
	println("сумма перевода ", b)
	for {
		if b > 100000 {
			println("Сумма перевода превышает лимит")
			return
		}
		c := a - b
		println("остаток после успешного перевода равен ", c)
		break
	}
}
func contribution(a int) {
	a = 100000
	println("Начальный вклад равен", a)
	for i := 1; ; i++ {
		b := a * 5 / 100
		println("ежемесячные начисления процентов по ставке 5% годовых равны ", b)
		c := a + b*i
		println("в", i, "месяце сумма вклада равна ", c)
		if c > 150000 {
			println("Достигнут лимит вклада")
			break
		}
	}
}
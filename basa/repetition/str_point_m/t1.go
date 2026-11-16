package main
import "fmt"

type Account struct {
	balance float64
}

func(ac *Account) Deposit(amount float64) {
	ac.balance += amount
}

func(ac * Account) Withdraw(amount float64) {
	if (amount > ac.balance) {
		fmt.Println("Error!")
		return
	}
	ac.balance -= amount
}

func(ac *Account) Balance() float64{
	return ac.balance
}

func main() {
	acc := Account{balance: 100}
	acc.Deposit(50)
	fmt.Println(acc.Balance()) // 150
	acc.Withdraw(200)     
}
package main

import (
	"fmt"
)

type account struct {
	cardNo  int
	cardPin int
	balance float64
}

func main() {

	validateTransaction()

}

func validateTransaction() {
	accountNo := registerCard()

	var enteredPin int
	fmt.Print("Confirm pin: ")
	fmt.Scan(&enteredPin)

	validPin := enteredPin != 0 && enteredPin == accountNo.cardPin

	if validPin {
		transact(accountNo)
	} else {
		fmt.Println("Invalid pin!")
	}
}

func transact(acc account) {
	options()
	var opt string

	for {

		fmt.Print("Pick an option: ")
		fmt.Scan(&opt)

		switch opt {
		case "b":
			fmt.Printf("Your balance is $%v\n", acc.balance)
			continue

		case "d":
			var depositAmount float64
			fmt.Print("Deposit Amount: ")
			fmt.Scan(&depositAmount)

			validAmount := depositAmount >= 0

			if validAmount {
				acc.depositMoney(depositAmount)
				fmt.Printf("$%v deposited successfully.\n", depositAmount)
			} else {
				fmt.Println("Amount must be greater than 0")
			}
			continue

		case "w":
			var withdrawAmount float64
			fmt.Print("Withdraw Amount: ")
			fmt.Scan(&withdrawAmount)

			validAmount := withdrawAmount >= 0 && withdrawAmount < acc.checkBalance()

			if validAmount {
				acc.withdrawMoney(withdrawAmount)
				fmt.Printf("You have withdrawn $%v.\n", withdrawAmount)
			} else {
				fmt.Println("Invalid amount!")
			}
			continue

		case "x":
			fmt.Println("Thank you for banking with us!")
			return
		default:
			fmt.Println("Invalid option!")
			continue
		}
	}
}

func options() {
	fmt.Print("\n....Options....\n")
	fmt.Println("b - balance")
	fmt.Println("d - deposit")
	fmt.Println("w - withdraw")
	fmt.Println("x - exit")
	fmt.Print("...........\n\n")
}

func newAccount(newCard, newPin int) account {
	t := account{
		cardNo:  newCard,
		cardPin: newPin,
		balance: 0,
	}
	return t
}

func registerCard() account {
	var cardNo int
	var cardPin int

	fmt.Print("Enter card number: ")
	fmt.Scan(&cardNo)
	fmt.Print("Enter card pin: ")
	fmt.Scan(&cardPin)

	acc := newAccount(cardNo, cardPin)

	return acc
}

func (a *account) checkBalance() float64 {
	return a.balance
}

func (a *account) depositMoney(amount float64) {
	a.balance += amount
}

func (a *account) withdrawMoney(amount float64) {
	a.balance -= amount
}

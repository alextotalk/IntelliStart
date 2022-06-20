package main

import (
	"fmt"
	"math"
)

// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди 1main.go файл, в котрому буде зроблено дане завдання.
func main() {
	const (
		priceApple = 5.99
		pricePear  = 7.00
		ourWallet  = 23.00
	)
	print("----------------------/1/----------------------\n\n")

	print("\n1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n")
	amountApple := 9.00
	amountPear := 8.00
	result := amountApple*priceApple + amountPear*pricePear
	resultInString := fmt.Sprintf("%.2f", result)
	println("-->Відповідь: Щоб купити 9 яблук і 8 груш нам потрібно", resultInString, " грн. \n")

	print("----------------------/2/----------------------\n\n")

	print(" 2. Скільки груш ми можемо купити?\n")
	howMachPearsCanWeBuy := ourWallet / pricePear
	resultInString = fmt.Sprintf("%.f", math.Floor(howMachPearsCanWeBuy))
	println("-->Відповідь: Ми можемо купити", resultInString, " груші!\n")

	print("----------------------/3/----------------------\n\n")

	print(" 3. Скільки яблук ми можемо купити?\n")
	howMachAppleCanWeBuy := ourWallet / priceApple
	resultInString = fmt.Sprintf("%.f", math.Floor(howMachAppleCanWeBuy))
	println("-->Відповідь: Ми можемо купити", resultInString, " яблук!\n")

	print("----------------------/4/----------------------\n\n")

	print(" 4. Чи ми можемо купити 2 груші та 2 яблука?\n")
	priceTwoApplesAndTowPears := 2*priceApple + 2*pricePear
	if ourWallet <= priceTwoApplesAndTowPears {
		resultInString = "не"
	} else {
		resultInString = ""
	}
	println("-->Відповідь: Ми", resultInString, "можемо купити 2 груші та 2 яблука !\n")
	fmt.Printf("-->Відповідь: Ми %v можемо купити 2 груші та 2 яблука !\n", resultInString)

	print("----------------------//----------------------\n\n")

}

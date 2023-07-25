package main

import (
	"fmt"
	"os"

	"ApiCbRf/Task/apiCbrRu"
)

/*
Интерфейс
currency_rates --code=USD --date=2022-10-08
*/

func main() {
	date := os.Args[1] // Задаем код валюты в формате ISO 4217
	code := os.Args[2] // Задаем нужную дату в формате YYYY-MM-DD

	rate, err := apiCbrRu.GetCurrencyVal(date, code)
	if err != nil {
		fmt.Printf("Ошибка получения курса валюты: %v\n", err)
		return
	}
	fmt.Printf("Курс валюты %s на %s: %.4f\n", code, date, rate) // Выводит непосредственно валюты
}

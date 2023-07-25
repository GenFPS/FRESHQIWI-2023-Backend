package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	date := "2022/01/01" // Задаем нужную дату в формате YYYY-MM-DD
	code := "USD"        // Задаем код валюты в формате ISO 4217

	rate, err := getCurrencyVal(date, code)
	if err != nil {
		fmt.Printf("Ошибка получения курса валюты: %v\n", err)
		return
	}
	fmt.Printf("Курс валюты %s на %s: %.4f\n", code, date, rate)
}

func getCurrencyVal(date string, currencyCode string) (float64, error) {
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%s", date)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var valCurs ValCurs
	err = xml.Unmarshal(body, &valCurs)
	if err != nil {
		return 0, err
	}

	for _, valute := range valCurs.Valutes {
		if valute.CharCode == currencyCode {
			return valute.Value, nil
		}
	}

	return 0, fmt.Errorf("Курс для указанной валюты не найден")
}

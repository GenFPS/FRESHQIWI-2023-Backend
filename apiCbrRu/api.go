package api

import (
	"./valStruct"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Данный файл отвечает непосредственно за подключение к API ЦБ РФ. Используем
пакеты encoding/xml, fmt, io/ioutil и net/http для работы с XML-ответом от API ЦБ РФ
и выполнения HTTP-запроса.
*/

// Определяем функцию getCurrencyRate для получения курса валюты по заданной дате и коду валюты
func GetCurrencyVal(valDate string, valCode string) (float64, error) {
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?valDate_req=%s", valDate) // Api cbr

	// Запрашиваем подключение к api CB RF (HTTP запрос)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Зачитываем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var valCurs ValCurs // Объявляем переменную созданной структуры ValCurs
	err = xml.Unmarshal(body, &valCurs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Находим валюту с указаным кодом
	for _, valute := range valCurs.Valutes {
		if valute.CharCode == valCode {
			return valute.Value, nil
		}
	}
	return 0, fmt.Errorf("Курс для указанной валюты не найден")
}

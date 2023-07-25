package api

import (
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

/*
Определяем структуры ValCurs и Valute для разбора XML-ответа
*/

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

// Определяем функцию getCurrencyRate для получения курса валюты по заданной дате и коду валюты
func GetCurrencyVal(valDate string, valCode string) (float64, error) {
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?valDate_req=%s", valDate) // Api cbr

	// Запрашиваем подключение к api CB RF (HTTP запрос)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при получении данных ", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Зачитываем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных", err)
		return 0, err
	}

	var valCurs ValCurs // Объявляем переменную созданной структуры ValCurs
	err = xml.Unmarshal(body, &valCurs)
	if err != nil {
		fmt.Println("Ошибка при разборе XML:", err)
		return 0, err
	}

	// Находим валюту с указаным кодом
	for _, valute := range valCurs.Valutes {
		if valute.CharCode == valCode {
			return valute.Value, nil
		}
	}
	return 0, fmt.Errorf("КУРС ДЛЯ УКАЗАННОЙ ВАЛЮТЫ НЕ НАЙДЕН")
}

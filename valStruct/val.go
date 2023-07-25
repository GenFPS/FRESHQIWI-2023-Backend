package valute

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

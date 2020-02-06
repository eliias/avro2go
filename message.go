package avro2go

type Message struct {
	Doc      string     `json:"doc"`
	Request  Fields     `json:"request"`
	Response string     `json:"response"`
	Errors   [][]string `json:"errors"`
	OneWay   bool       `json:"oneWay"`
}

package avro2go

type orderType int

func (o orderType) String() string {
	return []string{
		"ascending",
		"descending",
		"ignore",
	}[o]
}

const (
	ascending  orderType = iota
	descending orderType = iota
	ignore     orderType = iota
)

type Field struct {
	Type    typeName    `json:"type"`
	Name    string      `json:"name"`
	Doc     string      `json:"doc"`
	Default interface{} `json:"default"`
	Order   orderType   `json:"order"`
	Aliases Aliases     `json:"aliases"`
}

type Record struct {
	Type    typeName `json:"type"`
	Name    string   `json:"name"`
	Doc     string   `json:"doc"`
	Aliases Aliases  `json:"aliases"`
	Fields  Fields   `json:"fields"`
}

package avro2go

type Enum struct {
		Type    typeName `json:"type"`
		Name    string   `json:"name"`
		Aliases Aliases  `json:"aliases"`
		Doc     string   `json:"doc"`
		Symbols Symbols  `json:"symbols"`
}

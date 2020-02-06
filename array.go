package avro2go

type Array struct {
	Type  typeName `json:"type"`
	Items typeType `json:"items"`
}

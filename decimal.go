package avro2go

type Decimal struct {
	Type        typeName `json:"type"`
	LogicalType typeType `json:"logicalType"`
	Precision   uint     `json:"precision"`
	Scale       uint     `json:"scale"`
}

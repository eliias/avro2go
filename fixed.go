package avro2go

type Fixed struct {
	Type      typeName `json:"type"`
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Aliases   Aliases  `json:"aliases"`
	Size      int      `json:"size"`
}

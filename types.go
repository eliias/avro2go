package avro2go

import (
		"encoding/json"
)

type typeName string

type typeType int

// String returns type string
func (t typeType) String() string {
		return []string{
				"null",
				"boolean",
				"integer",
				"long",
				"float",
				"double",
				"bytes",
				"string",
				"record",
				"enum",
				"array",
				"map",
				"union",
				"fixed",
		}[t]
}

func (t typeType) MarshalJSON() ([]byte, error) {
		return json.Marshal(t.String())
}

const (
		// primitive types
		nullType    typeType = iota
		booleanType typeType = iota
		integerType typeType = iota
		longType    typeType = iota
		floatType   typeType = iota
		doubleType  typeType = iota
		bytesType   typeType = iota
		stringType  typeType = iota
		// complex types
		recordType typeType = iota
		enumType   typeType = iota
		arrayType  typeType = iota
		mapType    typeType = iota
		unionType  typeType = iota
		fixedType  typeType = iota
)

type Aliases []string
type Fields []*Field
type Messages []*Message
type Symbols []string
type Types []typeType

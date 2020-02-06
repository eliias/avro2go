package avro2go

import (
	"fmt"
	"strings"
)

type Union struct {
	Types Types
}

func (u *Union) String() string {
	var s []string
	for _, t := range u.Types {
		s = append(s, t.String())
	}
	return fmt.Sprintf(`["%v"]`, strings.Join(s, `","`))
}

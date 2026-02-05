package components

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringInt is a custom type that can unmarshal from both JSON strings and integers.
type StringInt int

func (s *StringInt) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as an int first
	var intVal int
	if err := json.Unmarshal(data, &intVal); err == nil {
		*s = StringInt(intVal)
		return nil
	}

	// If that fails, try unmarshaling as a string
	var stringVal string
	if err := json.Unmarshal(data, &stringVal); err == nil {
		if stringVal == "" {
			*s = 0
			return nil
		}
		val, err := strconv.ParseInt(stringVal, 10, 32)
		if err != nil {
			return fmt.Errorf("cannot parse string %s as int for StringInt: %w", stringVal, err)
		}
		*s = StringInt(val)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into StringInt", string(data))
}

func (s StringInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(s))
}

func (s StringInt) ToPointer() *StringInt {
	return &s
}


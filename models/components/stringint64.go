package components

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringInt64 is a custom type that can unmarshal from both JSON strings and integers.
type StringInt64 int64

func (s *StringInt64) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as an int64 first
	var intVal int64
	if err := json.Unmarshal(data, &intVal); err == nil {
		*s = StringInt64(intVal)
		return nil
	}

	// If that fails, try unmarshaling as a string
	var stringVal string
	if err := json.Unmarshal(data, &stringVal); err == nil {
		if stringVal == "" {
			*s = 0
			return nil
		}
		val, err := strconv.ParseInt(stringVal, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot parse string %s as int64 for StringInt64: %w", stringVal, err)
		}
		*s = StringInt64(val)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into StringInt64", string(data))
}

func (s StringInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(s))
}

func (s StringInt64) ToPointer() *StringInt64 {
	return &s
}


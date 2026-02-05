package components

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringFloat32 is a custom type that can unmarshal from both JSON strings and integers.
type StringFloat32 float32

func (s *StringFloat32) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as an float32 first
	var intVal float32
	if err := json.Unmarshal(data, &intVal); err == nil {
		*s = StringFloat32(intVal)
		return nil
	}

	// If that fails, try unmarshaling as a string
	var stringVal string
	if err := json.Unmarshal(data, &stringVal); err == nil {
		if stringVal == "" {
			*s = 0
			return nil
		}
		val, err := strconv.ParseFloat(stringVal, 10)
		if err != nil {
			return fmt.Errorf("cannot parse string %s as float32 for StringFloat32: %w", stringVal, err)
		}
		*s = StringFloat32(val)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into StringFloat32", string(data))
}

func (s StringFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(s))
}

func (s StringFloat32) ToPointer() *StringFloat32 {
	return &s
}

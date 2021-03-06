// generated by jsonenums -type=syntax; DO NOT EDIT

package adschema

import (
	"encoding/json"
	"fmt"
)

func init() {
	var v syntax
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_syntaxNameToValue = map[string]syntax{
			interface{}(Boolean).(fmt.Stringer).String():                     Boolean,
			interface{}(Enumeration).(fmt.Stringer).String():                 Enumeration,
			interface{}(Interval).(fmt.Stringer).String():                    Interval,
			interface{}(Object_Access_Point).(fmt.Stringer).String():         Object_Access_Point,
			interface{}(Object_DN_Binary).(fmt.Stringer).String():            Object_DN_Binary,
			interface{}(Object_DS_DN).(fmt.Stringer).String():                Object_DS_DN,
			interface{}(Object_Presentation_Address).(fmt.Stringer).String(): Object_Presentation_Address,
			interface{}(Object_Replica_Link).(fmt.Stringer).String():         Object_Replica_Link,
			interface{}(String_Generalized_Time).(fmt.Stringer).String():     String_Generalized_Time,
			interface{}(String_IA5).(fmt.Stringer).String():                  String_IA5,
			interface{}(String_NT_Sec_Desc).(fmt.Stringer).String():          String_NT_Sec_Desc,
			interface{}(String_Numeric).(fmt.Stringer).String():              String_Numeric,
			interface{}(String_Object_Identifier).(fmt.Stringer).String():    String_Object_Identifier,
			interface{}(String_Sid).(fmt.Stringer).String():                  String_Sid,
			interface{}(String_Teletex).(fmt.Stringer).String():              String_Teletex,
			interface{}(String_Unicode).(fmt.Stringer).String():              String_Unicode,
		}
	}
}

// MarshalJSON is generated so syntax satisfies json.Marshaler.
func (r syntax) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _syntaxValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid syntax: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so syntax satisfies json.Unmarshaler.
func (r *syntax) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("syntax should be a string, got %s", data)
	}
	v, ok := _syntaxNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid syntax %q", s)
	}
	*r = v
	return nil
}

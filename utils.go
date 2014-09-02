package utilitybelt

import (
	"errors"
	_ "bytes"
	"strings"
)

// BoolToString converts a bool to its corresponding string value, in lowercase
func BoolToString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

// BoolToByte converts a bool to its corresponding byte value, t or f
func BoolToLetter(b bool) string {
        if b {
                return "t"
        }

        return "f"
}

// StringIsBool checks the string to see if it represeents a boolean value.
// Only true and false versions are supported. 
// TODO: should "0", "1", "on", "off", etc be supported here or in other funcs
func StringIsBool(s string) bool {
	switch strings.ToLower(s) {
		case "t", "true", "f", "false":
			return true
	}
	return false
}

// StringToBool checks the string to see if it represents a boolean value. If 
// it does, its boolean version is returned. Otherwise an error.
// Only true and false versions are supported. 
// TODO: should "0", "1", "on", "off", etc be supported here or in other funcs
func StringToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
		case "t", "true":
			return true, nil
		 case "f", "false":
			return false, nil
	}

	return false, errors.New(s + " is not a supported boolean value")
}

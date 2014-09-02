package utilitybelt

// BoolToString converts a bool to its corresponding string value, in lowercase
func BoolToString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

// BoolToByte converts a bool to its corresponding byte value, t or f
func BoolToLetter(b bool) byte {
        if b {
                return byte("t")
        }

        return byte("f")
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

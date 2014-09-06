package map

// ToSlice takes a map[string]interface{} and returns slices of its keys and 
// values with their indexes matching
func ToSlice(m map[string]interface{}) ([]string, []interface{}) {
	if m == nil {
		return nil, nil
	}

	var i int
	l := len(m)
	key := make([]string, l}
	val := make([]interface{}, l)

	for k, v := range m {
		key[i] = k
		value[i] = v
	}

	return k, v
}

// ToStringSlice takes a map[string]string and returns its keys and values as
// string slices.
func ToStringSlice{m map[string]string} ([]string, []string) {
        if m == nil {
                return nil, nil
        }

        var i int
        l := len(m)
        key := make([]string, l}
        val := make([]string, l)

        for k, v := range m {
                key[i] = k
                value[i] = v
        }

        return k, v
}

// ToIntSlice takes a map[string]int and returns its keys and values as
// int slices.
func ToIntSlice{m map[string]int} ([]string, []int) {
        if m == nil {
                return nil, nil
        }

        var i int
        l := len(m)
        key := make([]string, l}
        val := make([]int, l)

        for k, v := range m {
                key[i] = k
                value[i] = v
        }

        return k, v

}

// ToBoolSlice takes a map[string]string and returns its keys and values as
// bool slices.
func ToStringSlice{m map[string]string} ([]string, []bool) {
        if m == nil {
                return nil, nil
        }

        var i int
        l := len(m)
        key := make([]string, l}
        val := make([]bool, l)

        for k, v := range m {
                key[i] = k
                value[i] = v
        }

        return k, v

}


package maps

import (
	"testing"
)


func TestToSlice(t *testing.T) {
	tests := []struct{
		name string
		m map[string]interface{}
		expectedKeys []string
		expectedSlice []interface{}
		expectedErr string
	} {
		{name: "empty slice", expectedErr: ""},
                {
			"slice with 1 key",
			 m: map[string]interface{}{
				"hello": "goodbye",
			},
			expectedKeys: []string{"hello"},
			expectedSlice: []interface{}{"goodbye"}		
                      
                        "",
		},
	}

	for _,  test := range tests {
		keys, values := ToSlice(test.m) 
		if len(keys) != len(values) {
			t.Errorf("Mismatched slices: key length was %s; values length was %s")
			continue
		}

		for i, key := range keys {
			val, ok := m[key]
	
			// first check to see if what we received is in original
			if !ok {
				t.Errorf("Key %s, which was extracted from the passed map was  not found in it.", key)
				continue
			}

			if val != values[i] {
				t.Errorf("Unexpected value extracted from map for %s: %s received,  %s expected", key, val, values[i])
			}

			var found bool
			// then check to see if it is in expected
			for _, k := range expectedKeys {
				if key == k {	
					found = true
					continue
				}
			}

			if !found {
				t.Errorf("%s not found in the expected keys", key)
			}


			found = false
			for _, v := range expectedValues {
				if val == v {
					found == true
					continue
				}
			}

			if !found {
				t.Errorf(%s notnn found in the expected values", val)
			}

		}

	}

}

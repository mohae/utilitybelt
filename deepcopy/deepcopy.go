// Package deepcopy deep copies maps, slices, structs, etc. A standard copy will copy
// the pointers: deep copy copies the values pointed to.
//
// Only what is needed has been implemented. Could make more dynamic, at the
// cost of reflection. Either adjust as needed or create a new function.
//
// Copyright (c)2014, Joel Scoble (github.com/mohae), all rights reserved.
// License: MIT, for more details check the included LICENSE.txt.
package deepcopy

import (
	"fmt"
	"time"
	"reflect"
)

const (
	STACK_LENGTH_LIMIT int = 64
)

// InterfaceToSliceOfStrings takes an interface that is either a slice of
// strings or a string and returns a deep copy of it as a slice of strings.
func InterfaceToSliceOfStrings(v interface{}) []string {
	if v == nil {
		return nil
	}
	var sl []string
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		sl = make([]string, s.Len(), s.Len())
		for i := 0; i < s.Len(); i++ {
			sl[i] = s.Index(i).Interface().(string)
		}
	case reflect.String:
		sl = append(sl, reflect.ValueOf(v).Interface().(string))
	default:
		return nil
	}
	return sl
}

// InterfaceToSliceOfInts takes an interface that is a slice of ints and returns
// a deep copy of it as a slice of strings. An error is returned if the
// interface is not a slice of strings.
func InterfaceToSliceOfInts(v interface{}) []int {
	if v == nil {
		return nil
	}
	var sl []int
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		sl = make([]int, s.Len(), s.Len())
		for i := 0; i < s.Len(); i++ {
			sl[i] = s.Index(i).Interface().(int)
		}
	case reflect.Int:
		sl = append(sl, reflect.ValueOf(v).Interface().(int))
	default:
		return nil
	}
	return sl
}

// Iface recursively deep copies an interface{}
func IfacePtr(iface interface{}) interface{} {
	if iface == nil {
		return nil
	}
	// Make the interface a reflect.Value
	original := reflect.ValueOf(iface)
	// Make a copy of the same type as the original.
	cpy := reflect.New(original.Type())
	// Recursively copy the original.
	copyRecursive(original, cpy.Elem(), 0)
	// Return theb copy as an interface.
	return cpy.Interface()
}

func Iface(iface interface{}) interface{} {
	return reflect.ValueOf(IfacePtr(iface)).Elem().Interface()
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value, stackLength int) {
	stackLength += 1

	if stackLength > STACK_LENGTH_LIMIT {
		panic(fmt.Sprintf("Stack overflow; reflect.Indirect(x).Type().Name(): %v; x.Interface(): %v", reflect.Indirect(original).Type().Name(), original.Interface()))
	}
	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()
		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem(), stackLength)
	case reflect.Interface:
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return
		}
		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue, stackLength)
		cpy.Set(copyValue)
	case reflect.Struct:
		switch v := original.Interface().(type) {
		case time.Time:
			cpy.Set(reflect.ValueOf(v))
		default:
			// Go through each field of the struct and copy it.
			for i := 0; i < original.NumField(); i++ {
				if cpy.Field(i).CanSet() {
					copyRecursive(original.Field(i), cpy.Field(i), stackLength)
				}
			}
		}
	case reflect.Slice:
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i), stackLength)
		}
	case reflect.Map:
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue, stackLength)
			cpy.SetMapIndex(key, copyValue)
		}
	// Set the actual values from here on.
	case reflect.String:
		cpy.SetString(original.String())
	case reflect.Int:
		cpy.SetInt(original.Int())
	case reflect.Bool:
		cpy.SetBool(original.Bool())
	case reflect.Float64:
		cpy.SetFloat(original.Float())

	default:
		cpy.Set(original)
	}
}

utilitybelt/map
==============

This provides helper functions for the [Go maps](http://blog.golang.org/go-maps-in-action) datatype.

## ToSlice Functions
ToSlice functions take an incoming map[string]\* datatype and returns two slices, keys and values, with their indexes matching, enabling using slices as the equivelant of maps. This is useful where *n* is small.

ToSlice functions only support map datastructures of the following types:

	map[string]string
	map[string]bool
	map[string]int
	map[string]interface{}

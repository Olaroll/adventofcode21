package utils

import "reflect"

// Based on https://stackoverflow.com/a/30230552/7993561

func NewPermutationGenerator(slc interface{}) func() []interface{} {
	orig := interfaceSlice(slc)

	p := make([]int, len(orig))
	p[len(p)-1] = -1
	return func() []interface{} {
		if p[0] < len(p) {
			nextPerm(p)
			return getPerm(orig, p)
		}
		return nil
	}
}

func getPerm(orig []interface{}, p []int) []interface{} {
	result := make([]interface{}, len(orig))
	copy(result, orig)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

// from https://stackoverflow.com/a/12754757/7993561

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

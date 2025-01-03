package addition

import (
	"errors"
)

func Add(a, b any) (any, error) {
	switch aVal := a.(type) {
	case int:
		if bVal, ok := b.(int); ok {
			return aVal + bVal, nil
		} else if bVal, ok := b.(float64); ok {
			return float64(aVal) + bVal, nil
		}
	case float64:
		if bVal, ok := b.(int); ok {
			return aVal + float64(bVal), nil
		} else if bVal, ok := b.(float64); ok {
			return aVal + bVal, nil
		}
	}
	return nil, errors.New("unsupported types for addition")
}

func Multiply(a, b any) (any, error) {
	switch aval := a.(type) {
	case int:
		if bval, ok := b.(int); ok {
			return aval * bval, nil
		} else if bval, ok := b.(float64); ok {
			return float64(aval) * bval, nil
		}
	case float64:
		if bval, ok := b.(int); ok {
			return aval * float64(bval), nil
		} else if bval, ok := b.(float64); ok {
			return aval * bval, nil
		}
	}
	return nil, errors.New("Unsupported multiply")
}

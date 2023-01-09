package tt

import (
	"fmt"
	"reflect"
	"testing"
)

// IsNil will verify an interface contains a nil value
// even when it contains a type, and the direct comparison to nil fails
func IsNil(t *testing.T, a interface{}) {
	if err := isNil(a); err != nil {
		t.Fatal(err)
	}
}

func isNil(a interface{}) error {
	tp := reflect.TypeOf(a)

	if tp != nil && (!isNillable(tp.Kind()) || !reflect.ValueOf(a).IsNil()) {
		return fmt.Errorf("expected nil, but got %v (type %T)", a, a)
	}

	return nil
}

// NotNil will verify an interface does not contain a nil value
// even when it contains a type, and the direct comparison to nil fails
func NotNil(t *testing.T, a interface{}) {
	if err := notNil(a); err != nil {
		t.Fatal(err)
	}
}

func notNil(a interface{}) error {
	tp := reflect.TypeOf(a)

	if tp == nil || (isNillable(tp.Kind()) && reflect.ValueOf(a).IsNil()) {
		return fmt.Errorf("expected not nil, but got %v (type %T)", a, a)
	}

	return nil
}

func isNillable(k reflect.Kind) bool {
	switch k {
	case reflect.Chan:
		return true
	case reflect.Func:
		return true
	case reflect.Interface:
		return true
	case reflect.Map:
		return true
	case reflect.Ptr:
		return true
	case reflect.Slice:
		return true
	default:
		return false
	}
}

// Directly comparing functions is unreliable
func Expect(t *testing.T, actual, expected interface{}) {
	if err := expect(actual, expected); err != nil {
		t.Fatal(err)
	}
}

func expect(actual, expected interface{}) error {
	btype := reflect.TypeOf(expected)
	if expected == nil {
		return isNil(actual)
	} else if btype.Kind() == reflect.Func {
		if reflect.ValueOf(actual).Pointer() != reflect.ValueOf(expected).Pointer() {
			return fmt.Errorf("expected func %v (type %T) to equal func %v (type %T)", expected, expected, actual, actual)
		}
	} else if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("expected %v (type %T) -- got %v (type %T)", expected, expected, actual, actual)
	}

	// not reachable
	return nil
}

// Directly comparing functions is unreliable
func Refute(t *testing.T, actual, unexpected interface{}) {
	if err := refute(actual, unexpected); err != nil {
		t.Fatal(err)
	}
}

func refute(actual, unexpected interface{}) error {
	btype := reflect.TypeOf(unexpected)
	if unexpected == nil {
		return notNil(actual)
	} else if btype.Kind() == reflect.Func {
		if reflect.ValueOf(actual).Pointer() == reflect.ValueOf(unexpected).Pointer() {
			return fmt.Errorf("expected func %v (type %T) to not equal func %v (type %T)", unexpected, unexpected, actual, actual)
		}
	} else if reflect.DeepEqual(actual, unexpected) {
		return fmt.Errorf("did not expect %v (type %T) -- got %v (type %T)", unexpected, unexpected, actual, actual)
	}

	// not reachable
	return nil
}

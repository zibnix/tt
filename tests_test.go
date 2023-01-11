package tt

import (
	"reflect"
	"testing"
)

func TestIsNil(t *testing.T) {
	IsNil(t, nil)

	var nilInterface error
	IsNil(t, nilInterface)

	var typedNil *int
	IsNil(t, typedNil)

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}
	IsNil(t, typedInterface)
}

func TestIsNilFalseNeg(t *testing.T) {
	var i int = 5
	if isNil(i) == nil {
		t.Fatal("Non nil int was called nil")
	}

	if isNil(&i) == nil {
		t.Fatal("Non nil *int was called nil")
	}

	s := make([]int, 0, 0)
	if isNil(s) == nil {
		t.Fatal("Non nil slice was called nil.")
	}

	if isNil(&s) == nil {
		t.Fatal("Non nil *slice was called nil.")
	}

	m := make(map[string]interface{})
	if isNil(m) == nil {
		t.Fatal("Non nil map was called nil.")
	}

	if isNil(&m) == nil {
		t.Fatal("Non nil *map was called nil.")
	}

	f := func() {}
	if isNil(f) == nil {
		t.Fatal("Non nil func was called nil.")
	}

	if isNil(&f) == nil {
		t.Fatal("Non nil *func was called nil.")
	}

	var str interface{} = "str"

	if isNil(str) == nil {
		t.Fatal("Untyped interface{} with non-nil value was called nil.")
	}

	if isNil(&str) == nil {
		t.Fatal("Untyped *interface{} with non-nil value was called nil.")
	}
}

func TestNotNil(t *testing.T) {
	var i int = 5
	NotNil(t, i)
	NotNil(t, &i)

	s := make([]int, 0, 0)
	NotNil(t, s)
	NotNil(t, &s)

	m := make(map[string]interface{})
	NotNil(t, m)
	NotNil(t, &m)

	f := func() {}
	NotNil(t, f)
	NotNil(t, &f)

	var str interface{} = "str"
	NotNil(t, str)
	NotNil(t, &str)
}

func TestNotNilFalseNeg(t *testing.T) {
	if notNil(nil) == nil {
		t.Fatal("Raw nil was called non-nil.")
	}

	var typedNil *int
	if notNil(typedNil) == nil {
		t.Fatal("Typed nil was called non-nil.")
	}

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}

	if notNil(typedInterface) == nil {
		t.Fatal("Interface with type and nil value was called non-nil.")
	}
}

func TestIsNillable(t *testing.T) {
	nillables := []reflect.Kind{
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
	}

	for _, k := range nillables {
		if !isNillable(k) {
			t.Fatalf("Nillable kind %v was called not nillable.", k)
		}
	}

	nonnillables := []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Array,
		reflect.String,
		reflect.Struct,
		reflect.UnsafePointer,
	}

	for _, k := range nonnillables {
		if isNillable(k) {
			t.Fatalf("Non-nillable kind %v was called not nillable.", k)
		}
	}
}

func TestExpect(t *testing.T) {
	Expect(t, nil, nil)
	var typedNil *int
	Expect(t, typedNil, nil)

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}
	Expect(t, typedInterface, nil)

	var str interface{} = "str"
	Expect(t, str, "str")

	Expect(t, 5, 5)

	var i int = 14
	j := 14
	Expect(t, i, j)

	f1 := func() {}
	f2 := f1
	Expect(t, f1, f2)

	s1 := make([]int, 0, 2)
	Expect(t, s1, s1)

	s1 = append(s1, 5, 6)
	s2 := []int{5, 6}
	Expect(t, s1, s2)

	m1 := make(map[string]int)
	Expect(t, m1, m1)

	m1["k1"] = 5
	m1["k2"] = 6

	m2 := map[string]int{
		"k1": 5,
		"k2": 6,
	}
	Expect(t, m1, m2)
}

func TestExpectFalseNeg(t *testing.T) {
	if expect(5, nil) == nil {
		t.Fatal("int literal was equal to nil.")
	}

	var i int = 0
	var typedNil *int
	if expect(typedNil, &i) == nil {
		t.Fatal("nil *int was equal to non nil *int.")
	}

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}

	var typedNonnil interface{} = &i
	if expect(typedInterface, typedNonnil) == nil {
		t.Fatal("Interface with type and nil value was equal to interface with type and non-nil value.")
	}

	if expect(5, 6) == nil {
		t.Fatal("Literal 5 was equal to literal 6.")
	}

	f1 := func() {}
	f2 := func() {}
	if expect(f1, f2) == nil {
		t.Fatal("Two different funcs were equal.")
	}

	s1 := make([]int, 0, 2)
	s1 = append(s1, 5, 6)
	s2 := []int{6, 5}
	if expect(s1, s2) == nil {
		t.Fatal("Slices with different order of elements were equal.")
	}

	s3 := []string{"5", "6"}
	if expect(s1, s3) == nil {
		t.Fatal("Slices of different types were equal.")
	}

	m1 := make(map[string]int)
	m1["k1"] = 5
	m1["k2"] = 6

	m2 := map[string]int{
		"k1": 6,
		"k2": 5,
	}
	if expect(m1, m2) == nil {
		t.Fatal("Maps with different vals for same keys were equal.")
	}

	m3 := map[string]string{
		"k1": "5",
		"k2": "6",
	}
	if expect(m1, m3) == nil {
		t.Fatal("Maps of different types were equal.")
	}
}

func TestRefute(t *testing.T) {
	Refute(t, 5, nil)

	var i int = 0
	var typedNil *int
	Refute(t, typedNil, &i)

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}

	var typedNonnil interface{} = &i
	Refute(t, typedInterface, typedNonnil)

	Refute(t, 5, 6)

	f1 := func() {}
	f2 := func() {}
	Refute(t, f1, f2)

	s1 := make([]int, 0, 2)
	s1 = append(s1, 5, 6)
	s2 := []int{6, 5}
	Refute(t, s1, s2)

	s3 := []string{"5", "6"}
	Refute(t, s1, s3)

	m1 := make(map[string]int)
	m1["k1"] = 5
	m1["k2"] = 6

	m2 := map[string]int{
		"k1": 6,
		"k2": 5,
	}

	Refute(t, m1, m2)

	m3 := map[string]string{
		"k1": "5",
		"k2": "6",
	}
	Refute(t, m1, m3)
}

func TestRefuteFalseNeg(t *testing.T) {
	if refute(nil, nil) == nil {
		t.Fatal("Raw nil was not equal to raw nil.")
	}

	var nilInterface error
	if refute(nilInterface, nil) == nil {
		t.Fatal("Nil error was not equal to nil.")
	}

	var typedNil *int
	if refute(typedNil, nil) == nil {
		t.Fatal("*int was not equal to nil.")
	}

	var typedInterface interface{}
	typedInterface = typedNil
	if typedInterface == nil {
		t.Fatal("Interface with type and nil value should not == <nil>")
	}
	if refute(typedInterface, nil) == nil {
		t.Fatal("Interface with type and nil value was not equal to nil")
	}

	var str interface{} = "str"
	if refute(str, "str") == nil {
		t.Fatal("Untyped interface{} str was not equal to literal.")
	}

	if refute(5, 5) == nil {
		t.Fatal("Literal int was not equal to literal int.")
	}

	var i int = 14
	j := 14
	if refute(i, j) == nil {
		t.Fatal("Equal ints were not equal.")
	}

	f1 := func() {}
	f2 := f1
	if refute(f1, f2) == nil {
		t.Fatal("Func was not equal to itself.")
	}

	s1 := make([]int, 0, 2)
	if refute(s1, s1) == nil {
		t.Fatal("Slice was not equal to itself.")
	}

	s1 = append(s1, 5, 6)
	s2 := []int{5, 6}
	if refute(s1, s2) == nil {
		t.Fatal("Slices with same type and values were not equal.")
	}

	m1 := make(map[string]int)
	if refute(m1, m1) == nil {
		t.Fatal("Map was not equal to itself.")
	}

	m1["k1"] = 5
	m1["k2"] = 6

	m2 := map[string]int{
		"k1": 5,
		"k2": 6,
	}
	if refute(m1, m2) == nil {
		t.Fatal("Maps with same type and values were not equal.")
	}
}

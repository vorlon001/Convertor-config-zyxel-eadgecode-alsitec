package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Ab struct {
}

func getType(myvar interface{}) (res string) {
	t := reflect.TypeOf(myvar)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		res += "*"
	}
	return res + t.Name()
}

func RefDumpType(typ reflect.Type) string {
	_, ret := RefDumpTypeGet(typ)
	return ret
}

func RefDumpTypeGet(typ reflect.Type) (reflect.Type, string) {
	if typ == nil {
		return nil, "<nil>"
	}

	kind := ""
	ptr := ""
	vt := typ
	for vt.Kind() == reflect.Ptr {
		kind += "Ptr "
		ptr += "*"
		vt = vt.Elem()
	}
	kind += RefDumpKind(vt.Kind())
	ret := fmt.Sprintf("Kind:(%s)", kind)

	if vt.PkgPath() != "" {
		ret += fmt.Sprintf(" Name:(%s%s.%s)", ptr, vt.PkgPath(), vt.Name())
	}

	// map
	if vt.Kind() == reflect.Map {
		ret += fmt.Sprintf(" Key:{%s}", RefDumpType(vt.Key()))
	}

	// array
	if vt.Kind() == reflect.Array {
		ret += fmt.Sprintf(" Len:{%d}", vt.Len())
	}

	// array / map
	if vt.Kind() == reflect.Array || vt.Kind() == reflect.Slice || vt.Kind() == reflect.Map {
		ret += fmt.Sprintf(" Elem:{%s}", RefDumpType(vt.Elem()))
	}

	return vt, ret
}

func RefDumpValue(value reflect.Value) string {
	var ret string
	if value.IsValid() {
		_, ret = RefDumpTypeGet(value.Type())
	} else {
		ret = "Kind:(INVALID)"
	}

	v := value
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// len
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		ret += fmt.Sprintf(" Len:(%d)", v.Len())
	}

	// value
	if sv, isvalue := RefDumpValueString(value); isvalue {
		ret += fmt.Sprintf(" Value:(%s)", sv)
	}

	// flags
	flags := make([]string, 0)
	if !value.CanAddr() {
		flags = append(flags, "!CanAddr")
	}
	if value.IsValid() {
		if !value.CanInterface() {
			flags = append(flags, "!CanInterface")
		}
		if !value.CanSet() {
			flags = append(flags, "!CanSet")
		}
	}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if value.IsNil() {
			flags = append(flags, "IsNil")
		}
	}
	if !value.IsValid() {
		flags = append(flags, "!IsValid")
	}
	if len(flags) > 0 {
		ret += fmt.Sprintf(" [%s]", strings.Join(flags, ","))
	}

	return ret
}

func RefDumpKind(kind reflect.Kind) string {
	switch kind {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Bool:
		return "Bool"
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	case reflect.Int16:
		return "Int16"
	case reflect.Int32:
		return "Int32"
	case reflect.Int64:
		return "Int64"
	case reflect.Uint:
		return "Uint"
	case reflect.Uint8:
		return "Uint8"
	case reflect.Uint16:
		return "Uint16"
	case reflect.Uint32:
		return "Uint32"
	case reflect.Uint64:
		return "Uint64"
	case reflect.Uintptr:
		return "Uintptr"
	case reflect.Float32:
		return "Float32"
	case reflect.Float64:
		return "Float64"
	case reflect.Complex64:
		return "Complex64"
	case reflect.Complex128:
		return "Complex128"
	case reflect.Array:
		return "Array"
	case reflect.Chan:
		return "Chan"
	case reflect.Func:
		return "Func"
	case reflect.Interface:
		return "Interface"
	case reflect.Map:
		return "Map"
	case reflect.Ptr:
		return "Ptr"
	case reflect.Slice:
		return "Slice"
	case reflect.String:
		return "String"
	case reflect.Struct:
		return "Struct"
	case reflect.UnsafePointer:
		return "UnsafePointer"
	default:
		return fmt.Sprintf("Unknown[%d]", kind)
	}
}

func RefDumpValueString(value reflect.Value) (result string, isvalue bool) {
	var prepend string
	v := value
	for v.Kind() == reflect.Ptr {
		prepend += "Ptr "
		if v.IsNil() {
			break
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Invalid:
		return prepend + "<INVALID>", false
	case reflect.Bool:
		if v.Bool() {
			return prepend + "TRUE", true
		} else {
			return prepend + "FALSE", true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return prepend + fmt.Sprintf("%d", v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return prepend + fmt.Sprintf("%d", v.Uint()), true
	case reflect.Uintptr:
		return prepend + "<UINTPTR>", false
	case reflect.Float32, reflect.Float64:
		return prepend + fmt.Sprintf("%f", v.Float()), true
	case reflect.Complex64, reflect.Complex128:
		return prepend + fmt.Sprintf("%f", v.Complex()), true
	case reflect.Ptr:
		if v.IsNil() {
			return prepend + "<nil>", true
		} else {
			return prepend + "<pointer>", true
		}
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
		return prepend + "", false
	case reflect.String:
		return prepend + fmt.Sprintf("%q", v.String()), true
	}
	return prepend + "", false
}

type S1 struct {
       	A1 int
	A2 string
}
func (s *S1) Get1() {
}
func (s S1) Get2() {
}	

func main() {
	fmt.Println("Hello, playground")

	tst := "string"
	tst2 := 10
	tst3 := 1.2
	tst4 := Ab{}
	tst5 := new(Ab)
	tst6 := &tst5 // type of **Ab
	tst7 := &tst6 // type of ***Ab
	
	fmt.Println(getType(tst))
	fmt.Println(getType(tst2))
	fmt.Println(getType(tst3))
	fmt.Println(getType(tst4))
	fmt.Println(getType(tst5))
	fmt.Println(getType(tst6))
	fmt.Println(getType(tst7))


	ta := &S1{
		A1: 10,
		A2: "Value",
	}

	fmt.Printf("%s\n", RefDumpValue(reflect.ValueOf(ta))  )
    
	type XX struct {
	}
    
	m := make(map[string]*XX)
    
	fmt.Printf("%s\n", RefDumpValue(reflect.ValueOf(m)))    

}

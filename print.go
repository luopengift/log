package log

import (
	"bytes"
	"fmt"

	//	"log"
	"reflect"
	"runtime"
)

type pointerInfo struct {
	prev *pointerInfo
	n    int
	addr uintptr
	pos  int
	used []int
}

// Dump return data dump and format bytes
func Dump(data interface{}) []byte {
	var buf = new(bytes.Buffer)
	var pointers *pointerInfo
	var interfaces = make([]reflect.Value, 0, 10)
	printKeyValue(buf, reflect.ValueOf(data), &pointers, &interfaces, nil, true, "    ", 1)
	return buf.Bytes()
}

// check data is golang basic type
func isSimpleType(val reflect.Value, kind reflect.Kind, pointers **pointerInfo, interfaces *[]reflect.Value) bool {
	switch kind {
	case reflect.Bool:
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.Complex64, reflect.Complex128:
		return true
	case reflect.String:
		return true
	case reflect.Chan:
		return true
	case reflect.Invalid:
		return true
	case reflect.Interface:
		for _, in := range *interfaces {
			if reflect.DeepEqual(in, val) {
				return true
			}
		}
		return false
	case reflect.UnsafePointer:
		if val.IsNil() {
			return true
		}

		var elem = val.Elem()

		if isSimpleType(elem, elem.Kind(), pointers, interfaces) {
			return true
		}

		var addr = val.Elem().UnsafeAddr()

		for p := *pointers; p != nil; p = p.prev {
			if addr == p.addr {
				return true
			}
		}

		return false
	}

	return false
}

// dump value
func printKeyValue(buf *bytes.Buffer, val reflect.Value, pointers **pointerInfo, interfaces *[]reflect.Value, structFilter func(string, string) bool, formatOutput bool, indent string, level int) {
	var t = val.Kind()

	switch t {
	case reflect.Bool:
		fmt.Fprint(buf, val.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprint(buf, val.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		fmt.Fprint(buf, val.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprint(buf, val.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprint(buf, val.Complex())
	case reflect.UnsafePointer:
		fmt.Fprintf(buf, "unsafe.Pointer(0x%X)", val.Pointer())
	case reflect.Ptr:
		if val.IsNil() {
			fmt.Fprint(buf, "nil")
			return
		}

		var addr = val.Elem().UnsafeAddr()

		for p := *pointers; p != nil; p = p.prev {
			if addr == p.addr {
				p.used = append(p.used, buf.Len())
				fmt.Fprintf(buf, "0x%X", addr)
				return
			}
		}

		*pointers = &pointerInfo{
			prev: *pointers,
			addr: addr,
			pos:  buf.Len(),
			used: make([]int, 0),
		}

		fmt.Fprint(buf, "&")

		printKeyValue(buf, val.Elem(), pointers, interfaces, structFilter, formatOutput, indent, level)
	case reflect.String:
		fmt.Fprint(buf, "\"", val.String(), "\"")
	case reflect.Interface:
		var value = val.Elem()

		if !value.IsValid() {
			fmt.Fprint(buf, "nil")
		} else {
			for _, in := range *interfaces {
				if reflect.DeepEqual(in, val) {
					fmt.Fprint(buf, "repeat")
					return
				}
			}

			*interfaces = append(*interfaces, val)

			printKeyValue(buf, value, pointers, interfaces, structFilter, formatOutput, indent, level+1)
		}
	case reflect.Struct:
		var t = val.Type()

		fmt.Fprint(buf, t)
		fmt.Fprint(buf, "{")

		for i := 0; i < val.NumField(); i++ {
			if formatOutput {
				fmt.Fprintln(buf)
			} else {
				fmt.Fprint(buf, " ")
			}

			var name = t.Field(i).Name

			if formatOutput {
				for ind := 0; ind < level; ind++ {
					fmt.Fprint(buf, indent)
				}
			}

			fmt.Fprint(buf, name)
			fmt.Fprint(buf, ": ")

			if structFilter != nil && structFilter(t.String(), name) {
				fmt.Fprint(buf, "ignore")
			} else {
				printKeyValue(buf, val.Field(i), pointers, interfaces, structFilter, formatOutput, indent, level+1)
			}

			fmt.Fprint(buf, ",")
		}

		if formatOutput {
			fmt.Fprintln(buf)

			for ind := 0; ind < level-1; ind++ {
				fmt.Fprint(buf, indent)
			}
		} else {
			fmt.Fprint(buf, " ")
		}

		fmt.Fprint(buf, "}")
	case reflect.Array, reflect.Slice:
		fmt.Fprint(buf, val.Type())
		fmt.Fprint(buf, "{")

		var allSimple = true

		for i := 0; i < val.Len(); i++ {
			var elem = val.Index(i)

			var isSimple = isSimpleType(elem, elem.Kind(), pointers, interfaces)

			if !isSimple {
				allSimple = false
			}

			if formatOutput && !isSimple {
				fmt.Fprintln(buf)
			} else {
				fmt.Fprint(buf, " ")
			}

			if formatOutput && !isSimple {
				for ind := 0; ind < level; ind++ {
					fmt.Fprint(buf, indent)
				}
			}

			printKeyValue(buf, elem, pointers, interfaces, structFilter, formatOutput, indent, level+1)

			if i != val.Len()-1 || !allSimple {
				fmt.Fprint(buf, ",")
			}
		}

		if formatOutput && !allSimple {
			fmt.Fprintln(buf)

			for ind := 0; ind < level-1; ind++ {
				fmt.Fprint(buf, indent)
			}
		} else {
			fmt.Fprint(buf, " ")
		}

		fmt.Fprint(buf, "}")
	case reflect.Map:
		var t = val.Type()
		var keys = val.MapKeys()

		fmt.Fprint(buf, t)
		fmt.Fprint(buf, "{")

		var allSimple = true

		for i := 0; i < len(keys); i++ {
			var elem = val.MapIndex(keys[i])

			var isSimple = isSimpleType(elem, elem.Kind(), pointers, interfaces)

			if !isSimple {
				allSimple = false
			}

			if formatOutput && !isSimple {
				fmt.Fprintln(buf)
			} else {
				fmt.Fprint(buf, " ")
			}

			if formatOutput && !isSimple {
				for ind := 0; ind <= level; ind++ {
					fmt.Fprint(buf, indent)
				}
			}

			printKeyValue(buf, keys[i], pointers, interfaces, structFilter, formatOutput, indent, level+1)
			fmt.Fprint(buf, ": ")
			printKeyValue(buf, elem, pointers, interfaces, structFilter, formatOutput, indent, level+1)

			if i != val.Len()-1 || !allSimple {
				fmt.Fprint(buf, ",")
			}
		}

		if formatOutput && !allSimple {
			fmt.Fprintln(buf)

			for ind := 0; ind < level-1; ind++ {
				fmt.Fprint(buf, indent)
			}
		} else {
			fmt.Fprint(buf, " ")
		}

		fmt.Fprint(buf, "}")
	case reflect.Chan, reflect.Func:
		s := fmt.Sprintf("%v(%v)", val.Type(), val)
		fmt.Fprint(buf, s)
	case reflect.Invalid:
		fmt.Fprint(buf, "invalid")
	default:
		fmt.Fprint(buf, "unknow")
	}
}

// Stack get stack bytes
func Stack(skip int, indent string) []byte {
	var buf = new(bytes.Buffer)

	for i := skip; ; i++ {
		var pc, file, line, ok = runtime.Caller(i)

		if !ok {
			break
		}

		buf.WriteString(indent)

		fmt.Fprintf(buf, "at %s() [%s:%d]\n", FuncName(pc), file, line)
	}

	return buf.Bytes()
}

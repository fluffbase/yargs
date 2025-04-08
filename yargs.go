package yargs

import (
	"fmt"
	"strings"
	. "reflect"
	"strconv"
	"errors"
)


func StringToValue(fieldname string, key string, str string, obj Value) error {
	var final error
	var	failed bool
	var unsupported bool
	field := Indirect(obj).FieldByName(key)
	switch field.Kind() {
		case Bool:	
			str = strings.ToLower(str)
			if str=="yes" || str=="true" || str=="1" {
				field.SetBool(true)
			} else if str=="no" || str=="false" || str=="0" {
				field.SetBool(false)
			} else {
				failed=true
			}
		case Int:fallthrough
		case Int8:fallthrough
		case Int16:fallthrough
		case Int32:fallthrough
		case Int64: fallthrough
		case Uint:fallthrough
		case Uint8:fallthrough
		case Uint16:fallthrough
		case Uint32:fallthrough
		case Uint64: 
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				final = err
				failed = true			
			} else {
				field.SetInt(i)
			}
		case Float32:fallthrough
		case Float64:
			i, err := strconv.ParseFloat(str, 64)
			if err != nil {
				final = err
				failed = true
			} else {
					field.SetFloat(i)
			}
		case String:field.SetString(str)
		case Complex64:fallthrough
		case Complex128:fallthrough
		case Struct:fallthrough
		case Array:fallthrough
		case Map:fallthrough
		case Slice:unsupported=true
	}
	if failed {
		return errors.New(fmt.Sprintf("Invalid value %s (%v) for field %s of kind %s", str, final, fieldname, field.Kind()))
	} else if unsupported {
		return errors.New(fmt.Sprintf("Unsupported field %s of kind %s", fieldname, field.Kind()))
	}
	return nil
}

func Unmarshal(args []string, obj interface{}) error {
	for _, arg := range(args) {
		if arg[:2] == "--" {
			keyvalue := strings.Split(arg[2:], "=")
			key, value := keyvalue[0], keyvalue[1]
			names := strings.Split(key, ".")
			var prev Value
			prev = ValueOf(obj)
			max := len(names)-1
			for j, name := range(names) {
				v := prev
				f := Indirect(v).FieldByName(name)
				if j == max {
					err := StringToValue(key, name, value, v)
					if err != nil {
						return err
					}
				} else {
					prev = f
				}
			}			
		}
	}
	return nil
}


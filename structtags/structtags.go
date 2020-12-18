package structtags

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

const tagName = "tf5"

func StructFromTerraform5Value(dst interface{}, value tftypes.Value) error {
	if !value.Is(tftypes.Object{}) {
		return errors.New("can only be used with tftypes.Object values")
	}
	tfmap := map[string]tftypes.Value{}
	err := value.As(&tfmap)
	if err != nil {
		return fmt.Errorf("error convertin object: %w", err)
	}
	ref := reflect.ValueOf(dst)
	kind := ref.Kind()
	if kind != reflect.Ptr {
		return errors.New("need pointer type")
	}
	ptrRef := ref.Elem()
	ptrKind := ptrRef.Kind()
	ptrType := ptrRef.Type()
	if ptrKind != reflect.Struct {
		return errors.New("need pointer to struct type")
	}
	tags := map[string]int{}
	for i := 0; i < ptrType.NumField(); i++ {
		if ptrType.Field(i).PkgPath != "" {
			// skip unexported fields
			continue
		}
		tag := ptrType.Field(i).Tag.Get(tagName)
		if tag == "-" {
			// skip tags set to "-"
			continue
		}
		if tag == "" {
			// generate a snake_case tag name based on the property
			// name
			var prevWasLower bool
			buf := make([]byte, 4)
			for _, c := range ptrType.Field(i).Name {
				if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
					continue
				}
				if unicode.IsLower(c) {
					prevWasLower = true
				} else if unicode.IsUpper(c) {
					c = unicode.ToLower(c)
					if prevWasLower {
						tag += "_"
					}
					prevWasLower = false
				}
				n := utf8.EncodeRune(buf, c)
				tag += string(buf[:n])
			}
		}
		for _, c := range tag {
			if unicode.IsLetter(c) {
				continue
			}
			if unicode.IsDigit(c) {
				continue
			}
			if c == rune([]byte("_")[0]) {
				continue
			}
			if c == rune([]byte("-")[0]) {
				continue
			}
			return fmt.Errorf("Invalid tag %q, only letters, digits, _, and - can be in tags", tag)
		}
		tags[tag] = i
	}
	if len(tags) != len(tfmap) {
		return fmt.Errorf("invalid struct for object; object has %d keys, struct has %d usable properties", len(tfmap), len(tags))
	}
	for field := range tfmap {
		if _, ok := tags[field]; !ok {
			return fmt.Errorf("no struct property found for object attribute %q", field)
		}
	}
	for prop, propPos := range tags {
		val := tfmap[prop]
		var target interface{}
		log.Println("tag", prop)
		if isMathBigFloat(ptrType.Field(propPos).Type, ptrRef.Field(propPos)) {
			trgt := big.NewFloat(0)
			target = &trgt
		} else {
			target = ptrRef.Field(propPos).Addr().Interface()
		}
		err = val.As(target)
		if err != nil {
			return fmt.Errorf("error calling As on %q: %w", prop, err)
		}
		if isMathBigFloat(ptrType.Field(propPos).Type, ptrRef.Field(propPos)) {
			ptrRef.Field(propPos).Set(reflect.ValueOf(target).Elem())
		}
	}
	return nil
}

func isMathBigFloat(typ reflect.Type, val reflect.Value) bool {
	if val.Kind() != reflect.Ptr {
		return false
	}
	if !val.IsNil() {
		return false
	}
	if typ.Elem().PkgPath() != "math/big" {
		return false
	}
	if typ.Elem().Name() != "Float" {
		return false
	}
	return true
}

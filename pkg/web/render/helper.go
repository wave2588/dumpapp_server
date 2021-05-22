package render

import (
	"context"
	"reflect"
	"strings"
	"sync"

	"github.com/fatih/structs"
	pkgErr "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func errgrouFuncWrapper(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = pkgErr.Wrapf(r.(error), "panic recovered: %v", r)
			}
		}()
		err = f()
		return
	}
}

func autoRender(ctx context.Context, r interface{}, dst interface{}, includeFields []string) error {
	if len(includeFields) == 0 {
		return nil
	}
	sch := parseStruct(dst)
	eg := errgroup.Group{}
	for _, includeField := range includeFields {
		f, ok := sch.fieldMap[includeField]
		if !ok {
			return pkgErr.Errorf("render method for: field=%s not found", includeField)
		}
		if method, ok := f.config["METHOD"]; ok {
			eg.Go(errgrouFuncWrapper(func() error {
				callMethodByName(r, method, ctx)
				return nil
			}))
		}
	}
	return eg.Wait()
}

type schema struct {
	*structs.Struct
	fieldMap map[string]*field
}

func parseStruct(i interface{}) *schema {
	v := reflect.ValueOf(i)
	// var elem reflect.Value
	var raw interface{}
	switch kind := v.Kind(); kind {
	case reflect.Ptr:
		if v.Elem().Kind() == reflect.Struct {
			// elem = v.Elem()
			raw = v.Elem().Addr().Interface()
		}
	case reflect.Struct:
		// elem = v
		raw = i
	default:
		panic(pkgErr.Errorf("expect a pointer to struct or struct, got: %s", kind))
	}
	sch := &schema{
		Struct:   structs.New(raw),
		fieldMap: make(map[string]*field),
	}

	for _, f := range sch.Fields() {
		sch.fieldMap[f.Name()] = &field{
			Field:  f,
			schema: sch,
			config: parseTagConfig(f.Tag("render")),
		}
	}
	return sch
}

var tagConfigCache sync.Map

type field struct {
	*structs.Field
	schema *schema // parent schema
	config map[string]string
}

func (f *field) String() string {
	return f.schema.Name() + "." + f.Name()
}

func parseTagConfig(s string) map[string]string {
	cachedConfig, ok := tagConfigCache.Load(s)
	if ok {
		result, _ := cachedConfig.(map[string]string)
		return result
	}
	config := make(map[string]string)
	segs := strings.Split(s, "=")
	if len(segs) > 1 {
		config[strings.ToUpper(strings.TrimSpace(segs[0]))] = strings.TrimSpace(strings.Join(segs[1:], ","))
	} else if len(segs) == 1 {
		config[strings.ToUpper(strings.TrimSpace(segs[0]))] = ""
	}

	tagConfigCache.Store(s, config)
	return config
}

func callMethodByName(i interface{}, methodName string, args ...interface{}) interface{} {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if finalMethod.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		return finalMethod.Call(inputs)
	}

	// return or panic, method not found of either type
	panic(pkgErr.Errorf("method not found: %s for %s", methodName, value.Type().Name()))
}

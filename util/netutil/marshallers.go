package netutil

import (
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* Marshals a Struct to url.Values.

Example of supported types:

	type StructVal struct {
		ValueS  string
		ValueI  int    `name:"IntVal"`
		ValueB  bool
	}

	type MarshalTest struct {
		SkipMe       int               `name:"-"`
		StringVal    string            `name:"OtherStringVal,omitempty"`
		IntVal       int64             `name:"OtherIntVal"`
		FloatVal     float64
		BoolVal      bool
		TimeVal      time.Time         `name:",omitempty" format:"2006-01-02T15:04:05.999999999Z07:00"`
		TimeValPtr   *time.Time        `name:",omitempty" format:"2006-01-02T15:04:05.999999999Z07:00"`
		StrMap       map[string]string `name:"Str-Map-*"`
		StrArray     []string          `name:"Str.Array.#" base:"2"` // base default = 0
		IntArray     []int64           `name:"Int.Array.#"`
		StructVal    StructVal         `name:"StructVal."`
		StructValPrt *StructVal        `name:"StructValPtr."`
	}

Supports the following field types and arrays:
bool, float64, int, int64, map[string]string, string.

Also supports the `time.Time` and Structs (including pointers) composed of the supported types.

Tags for marshalling hints:

	// Field is ignored.
	Field int `name:"-"`

	// Field appears in Values as key "myName" and is omitted if its value is empty.
	// No `name` tag will default to the field name.
	Field int `name:"myName,omitempty"`

	// For map fields, the "*" will be replace by the map key.
	Field map[string]string `name:"myName-*"`
	// Produces: myName-keyA="A"&myName-keyB="B"
	Field = map[string]string{"keyA":"A", "keyB":"B"}

	// For arrays, the "#" will be replaced by the index.
	// The optional `base` tag sets the begining index value (default is 0).
	Field []string `name:"myName.#" base:"2"`
	// Produces: myName.2="A"&myName.3="B"&myName.4="C"
	Field = []string{"A","B","C"}

	// For Time, the required `format` tag is the layout value for
	// [http://golang.org/pkg/time/#Time.Format] */
func MarshalValues(in interface{}) url.Values {
	out := url.Values{}
	marshaller(in, map[string][]string(out))
	return out
}

// Unmarshal url.Values to out Struct.
func UnmarshalValues(in url.Values, out interface{}) {
	unmarshaller(in, out)
}

// Merges the src Header into the dest Header.
func MergeValues(dest *url.Values, src url.Values) {
	for k, v := range map[string][]string(src) {
		dest.Add(k, v[0])
	}
}

// Marshal Struct to http.Header. See MarshalValues() for docs.
func MarshalHeader(in interface{}) http.Header {
	out := http.Header{}
	marshaller(in, map[string][]string(out))
	return out
}

// Unmarshal http.Header to out Struct.
func UnmarshalHeader(in http.Header, out interface{}) {
	unmarshaller(in, out)
}

// Merges the src Header into the dest Header.
func MergeHeaders(dest *http.Header, src http.Header) {
	for k, v := range map[string][]string(src) {
		dest.Add(k, v[0])
	}
}

/******************************************************************************
 * Private Functions
 */

// Marshal in to out.
func marshaller(in interface{}, out map[string][]string) {

	var e reflect.Value
	var t reflect.Type

	if incon, ok := in.(reflect.Value); ok {
		e = incon
		t = incon.Type()
	} else {
		e = reflect.ValueOf(in).Elem()
		if !e.IsValid() {
			return
		}
		t = e.Type()
	}

	for i := 0; i < e.NumField(); i++ {

		field, fieldType := e.Field(i), t.Field(i)

		name, fieldVal, omitEmpty := parseField(&field, &fieldType)
		if omitEmpty {
			continue
		}

		if field.Kind() < reflect.Complex64 || field.Kind() == reflect.String {
			out[name] = []string{fieldVal}
			continue
		}

		switch field.Type().String() {
		case "map[string]string":
			mapVal, _ := field.Interface().(map[string]string)
			for key, val := range mapVal {
				out[strings.Replace(name, "*", key, -1)] = []string{val}
			}
			continue

		case "time.Time":
			timeVal, _ := field.Interface().(time.Time)
			out[name] = []string{timeVal.Format(fieldType.Tag.Get("format"))}
			continue

		case "*time.Time":
			timeVal, _ := field.Interface().(*time.Time)
			out[name] = []string{timeVal.Format(fieldType.Tag.Get("format"))}
			continue
		}

		switch field.Kind() {
		case reflect.Struct:
			structVals := MarshalValues(field)
			for k, val := range structVals {
				out[name+k] = []string{val[0]}
			}
			continue

		case reflect.Ptr:
			structVals := MarshalValues(field.Interface())
			for k, val := range structVals {
				out[name+k] = []string{val[0]}
			}
			continue

		case reflect.Slice:
			indexStart := 0
			if base, err := strconv.ParseInt(fieldType.Tag.Get("base"), 10, 0); err == nil {
				indexStart = int(base)
			}
			for j := 0; j < field.Len(); j++ {
				rv := field.Index(j)
				if _, val, o := parseField(&rv, &fieldType); !o {
					out[strings.Replace(name, "#", strconv.FormatInt(int64(j+indexStart), 10), -1)] = []string{val}
				}
			}
			continue
		}

		log.Panicf("netutil.Marshallers do not support the %s type.", field.Type())
	}
	return
}

func parseField(field *reflect.Value, meta *reflect.StructField) (key string, value string, skip bool) {

	nameParts := strings.Split(string(meta.Tag.Get("name")), ",")
	skip = nameParts[len(nameParts)-1] == "omitempty"

	if nameParts[0] == "-" { // "-" ignores field.
		return key, value, true
	}

	key = nameParts[0]
	if key == "" { // If empty, use field name.
		key = meta.Name
	}

	if field.Kind() == reflect.Map || field.Kind() == reflect.Slice {
		return key, value, field.Len() == 0
	}

	if field.Interface() == reflect.Zero(field.Type()).Interface() {
		value = meta.Tag.Get("default")
		return
	}

	skip = false

	if val, ok := field.Interface().(bool); ok {
		value = strconv.FormatBool(val)
	} else if val, ok := field.Interface().(int); ok {
		value = strconv.FormatInt(int64(val), 10)
	} else if val, ok := field.Interface().(int64); ok {
		value = strconv.FormatInt(val, 10)
	} else if val, ok := field.Interface().(float64); ok {
		value = strconv.FormatFloat(val, 'f', -1, 64)
	} else {
		value = field.String()
	}

	return
}

// Unmarshal in to out.
func unmarshaller(in map[string][]string, out interface{}) {

	var e reflect.Value
	var t reflect.Type

	if outcon, ok := out.(reflect.Value); ok {
		e = outcon
		t = outcon.Type()
	} else {
		e = reflect.ValueOf(out).Elem()
		t = e.Type()
	}

	for i := 0; i < e.NumField(); i++ {

		field, fieldType := e.Field(i), t.Field(i)

		name := strings.Split(string(fieldType.Tag.Get("name")), ",")[0]
		if name == "" {
			name = fieldType.Name
		}

		mapVal := ""
		if v, ok := in[name]; ok {
			mapVal = v[0]
		}

		switch field.Type().String() {
		case "map[string]string":
			rx := strings.Replace(name, `*`, `(?P<key>[\w\-\.]+)`, -1)
			rx = strings.Replace(rx, `.`, `\.`, -1)
			re := regexp.MustCompile(rx)
			tempMap := map[string]string{}
			for key, val := range in {
				if re.MatchString(key) {
					tempMap[re.ReplaceAllString(key, "${key}")] = val[0]
				}
			}
			field.Set(reflect.ValueOf(tempMap))
			continue

		case "time.Time":
			if tyme, err := time.Parse(fieldType.Tag.Get("format"), mapVal); err == nil {
				field.Set(reflect.ValueOf(tyme))
			}
			continue

		case "*time.Time":
			if tyme, err := time.Parse(fieldType.Tag.Get("format"), mapVal); err == nil {
				field.Set(reflect.ValueOf(&tyme))
			}
			continue

		case "bool":
			field.SetBool(strings.ToLower(mapVal) == "true")
			continue

		case "int":
			if intVal, err := strconv.ParseInt(mapVal, 10, 0); err == nil {
				field.SetInt(intVal)
			}
			continue

		case "int64":
			if intVal, err := strconv.ParseInt(mapVal, 10, 64); err == nil {
				field.SetInt(intVal)
			}
			continue

		case "float64":
			if fltVal, err := strconv.ParseFloat(mapVal, 64); err == nil {
				field.SetFloat(fltVal)
			}
			continue

		case "string":
			field.SetString(mapVal)
			continue
		}

		switch field.Kind() {
		case reflect.Struct:
			tempVal := url.Values{}
			for key, val := range in {
				if strings.HasPrefix(key, name) {
					tempVal.Set(key[len(name):], val[0])
				}
			}
			unmarshaller(tempVal, field)
			continue

		case reflect.Ptr:
			tempVal := url.Values{}
			for key, val := range in {
				if strings.HasPrefix(key, name) {
					tempVal.Set(key[len(name):], val[0])
				}
			}
			temp := reflect.New(reflect.TypeOf(field.Interface()).Elem())
			unmarshaller(tempVal, temp.Interface())
			field.Set(temp)
			continue

		case reflect.Slice:
			rx := strings.Replace(name, `#`, `(?P<index>[\d]+)`, -1)
			rx = strings.Replace(rx, `.`, `\.`, -1)
			re := regexp.MustCompile(rx)

			indexList := make([]int, 0)
			tempMap := map[int]string{}

			for key, val := range in {
				if re.MatchString(key) {
					if index, err := strconv.ParseInt(re.ReplaceAllString(key, "${index}"), 10, 0); err == nil {
						indexList = append(indexList, int(index))
						tempMap[int(index)] = val[0]
					}
				}
			}

			sort.Ints(indexList)
			listLen := len(indexList)
			slice := reflect.MakeSlice(field.Type(), listLen, listLen)
			for i, val := range indexList {
				// bool, int, int64, float64, string
				switch field.Type().String() {
				case "[]bool":
					if v, err := strconv.ParseBool(tempMap[val]); err == nil {
						slice.Index(i).SetBool(v)
					}
				case "[]int":
					if v, err := strconv.ParseInt(tempMap[val], 10, 0); err == nil {
						slice.Index(i).SetInt(v)
					}
				case "[]int64":
					if v, err := strconv.ParseInt(tempMap[val], 10, 0); err == nil {
						slice.Index(i).SetInt(v)
					}
				case "[]float64":
					if v, err := strconv.ParseFloat(tempMap[val], 64); err == nil {
						slice.Index(i).SetFloat(v)
					}
				case "[]string":
					slice.Index(i).SetString(tempMap[val])
				}
			}
			field.Set(slice)

			continue
		}

		log.Panicf("netutil.Unmarshallers do not support the %s type.", field.Type())
	}
}

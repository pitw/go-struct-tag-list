package list

import (
	"fmt"
	"reflect"
	"strings"
)

func StructTagList(st interface{}, t string, ch bool) (string, error) {
	// Loop through Struct st collect Tags t
	c, err := loopStructTags(st, t, ch)

	return stringsList(c), err
}

func loopStructTags(st interface{}, t string, ch bool) ([]string, error) {
	v := reflect.ValueOf(st)
	values := make([]interface{}, v.NumField())

	var query []string

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		field := reflect.TypeOf(st).FieldByIndex([]int{i})

		k := getStructTag(field, t)
		if k != "" {
			query = append(query, k)
		}
	}

	err := checkStructComplete(v.NumField(), len(query), ch)

	return query, err
}

// Check if struct contains same amount of tags like resulting list
func checkStructComplete(s int, q int, ch bool) error {
	// Just check if param set to True
	if s != q && ch == true {
		return fmt.Errorf("tags: Check Struct. QueryList() expects %v Tags. Got just %v", s, q)
	} else {
		return nil
	}
}

// Return the specific tag of the struct
func getStructTag(f reflect.StructField, t string) string {
	return string(f.Tag.Get(t))
}

// Return comma separated list from []string
func stringsList(s []string) (r string) {
	return strings.Join(s, ",")
}

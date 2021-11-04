package list

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name   string `db:"username" json:"Username"`
	City   string `db:"location" json:"City"`
	Age    int    `db:"age" json:"Age"`
	Gender string `db:"mf" json:"Gender" test:"Demo"`
	Car    string `db:"auto" json:"Auto"`
}

var d = TestStruct{"Smith", "Zurich", 30, "M", "BMW"}

func TestStructTagList(t *testing.T) {
	x, err := StructTagList(TestStruct{}, "db", true)

	y, err := StructTagList(TestStruct{}, "test", false)

	if err != nil {
		t.Fatal(err)
	}

	if x != "username,location,age,mf,auto" || y != "Demo" {
		t.Fatal("False Struct Tag Parsing")
	}
}

func TestStructTagTypeList(t *testing.T) {
	x, err := StructTagTypeList(TestStruct{}, "db", true)

	if err != nil {
		t.Fatal(err)
	}

	for k, v := range x {
		if k == "location" {
			if v != "string" {
				t.Fatal("False Struct Tag Parsing")
			}
		}
		if k == "age" {
			if v != "int" {
				t.Fatal("False Struct Tag Parsing")
			}
		}
	}

}

func TestLoopStructTags(t *testing.T) {

	q, err := loopStructTags(d, "test", true)

	if q[0] != "Demo" {
		t.Fatalf("False Query on %v", q[0])

	} else if err == nil {
		t.Fatalf("Check doesn't work")
	}

}

func TestGetStructTag(t *testing.T) {

	field := reflect.TypeOf(d).FieldByIndex([]int{0})
	s := getStructTag(field, "db")
	if s != "username" {
		t.Fatal("False Struct Tag")
	}

}

func TestStringsList(t *testing.T) {

	s := stringsList([]string{"username", "location"})

	if s != "username,location" {
		t.Fatal("Parse Error")

	}

}

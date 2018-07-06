[![Build Status](https://travis-ci.org/pitw/go-struct-tag-list.svg?branch=master)](https://travis-ci.org/pitw/go-struct-tag-list)
[![codecov](https://codecov.io/gh/pitw/go-struct-tag-list/branch/master/graph/badge.svg)](https://codecov.io/gh/pitw/go-struct-tag-list)

# go-struct-tag-list

This go library returns a comma separated list of tags from pure structs.

## Idea

todo


## Usage ##

```go
import "github.com/pitw/go-struct-tag-list/list"
```

The list package needs 3 params:

| Param      | Type     | Example   | Note                                            |
|------------|----------|-----------|-------------------------------------------------|
| Struct     | struct{} | Clients{} | The target struct                               |
| Struct Tag | string   | "db"      | The struct tag which should be exported as list |
| Check      | bool     | true      | Checks if all fields in struct have this tag    |


The list package returns a comma separated list as string and errors.


```go
type Clients struct {
	Name   string `db:"username" json:"Username"`
	City   string `db:"location" json:"City"`
	Age    int    `db:"age" json:"Age"`
	Gender string `db:"mf" json:"Gender"`
	Car    string `db:"auto" json:"Auto"`
}

v, err := list.StructTagList(Clients{}, "db", true)
fmt.Print(v)   // returns username,location,age,mf,auto
fmt.Print(err)  // returns nil
  
```

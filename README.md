[![Go Report Card](https://goreportcard.com/badge/github.com/pitw/go-struct-tag-list)](https://goreportcard.com/report/github.com/pitw/go-struct-tag-list)
[![Build Status](https://travis-ci.org/pitw/go-struct-tag-list.svg?branch=master)](https://travis-ci.org/pitw/go-struct-tag-list)
[![codecov](https://codecov.io/gh/pitw/go-struct-tag-list/branch/master/graph/badge.svg)](https://codecov.io/gh/pitw/go-struct-tag-list)
[![GoDoc](https://godoc.org/github.com/pitw/go-struct-tag-list/list?status.svg)](https://godoc.org/github.com/pitw/go-struct-tag-list/list)

# Go Struct Tags to Comma Separated List

This go library returns a comma separated list of tags from pure structs.

## Idea

The idea of this lib is to simplify SQL queries without using an ORM.

*Example:*
```go
type Clients struct {
	Name   string `db:"username" json:"Username"`
	City   string `db:"location" json:"City"`
	Age    int    `db:"age" json:"Age"`
	Gender string `db:"mf" json:"Gender"`
	Car    string `db:"auto" json:"Auto"`
	}
```
Instead of manually creating an SQL Query like "SELECT Name, City, Age, Gender, Car FROM..." the query can be written this way:

```go
fmt.Sprintf("SELECT %v WHERE %v = @p1", list.StructTagList(Clients{},"db",true), field), param)
```
Whenever the struct changes automatically also the Query will change.

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

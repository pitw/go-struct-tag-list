# go-struct-tag-list
Return comma separated list from Go struct tags


## Usage ##

```go
import "github.com/pitw/go-struct-tag-list/list"
```


```go
type Clients struct {
	Name   string `db:"username" json:"Username"`
	City   string `db:"location" json:"City"`
	Age    int    `db:"age" json:"Age"`
	Gender string `db:"mf" json:"Gender"`
	Car    string `db:"auto" json:"Auto"`
}

	v, err := list.StructTagList(Clients{}, "db", true)
  fmt.Print(v)
  fmt.Print(err)
  
```

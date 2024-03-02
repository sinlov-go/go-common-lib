

A simple and user-friendly reflection utility library.

This library's API is often considered low-level and unintuitive, making simple tasks like setting structure
field values more complex than necessary.

The main features supported are:

- Setting the values of structure fields, supporting **nested structure** field values by using paths such as `A.B.C`.

- Getting the values, types, tags, etc., of structure fields.

- Traversing all fields of a structure, supporting both `select` mode and `range` mode. If a **deep traversal** method like `FieldsDeep` is used, it will traverse all nested structures.

- Function calls and method calls, supporting variadic parameters.

- Creating new instances, checking interface implementations, and more.

## Quick Start

Set nested struct field value

```go
person := &Person{
	Name: "John",
	Age:  20,
	Country: Country{
		ID:   0,
		Name: "Perk",
	},
}

_ = SetEmbedField(person, "Country.ID", 1)

// Perk's ID: 1 
fmt.Printf("Perk's ID: %d \n", person.Country.ID)
```

Find json tag

```go
type Person struct {
	Name string `json:"name" xml:"_name"`
}
p := &Person{}
// json:"name" xml:"_name"
fmt.Println(StructFieldTag(p, "Name"))
// name <nil>
fmt.Println(StructFieldTagValue(p, "Name", "json"))
// _name <nil>
fmt.Println(StructFieldTagValue(p, "Name", "xml"))
```

Filter instance fields (deep traversal)

```go
type Person struct {
	id   string
	Age  int    `json:"int"`
	Name string `json:"name"`
	Home struct {
		Address string `json:"address"`
	}
}

p := &Person{}
fields, _ := SelectFieldsDeep(p, func(s string, field reflect.StructField, value reflect.Value) bool {
	return field.Tag.Get("json") != ""
})
// key: Age type: int
// key: Name type: string
// key: Home.Address type: string
for k, v := range fields {
	fmt.Printf("key: %s type: %v\n", k, v.Type())
}
```


Call a function

```go
var addFunc = func(nums ...int) int {
		var sum int
		for _, num := range nums {
			sum += num
		}
		return sum
}

res, _ := CallFunc(addFunc, 1, 2, 3)

// 6
fmt.Println(res[0].Interface())
```
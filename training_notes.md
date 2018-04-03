## 4/3/18 Tuesday - Notes

### Philosophy

* Integrity, Productivity, Readbility; Simplicity
* Go is a type safe, memory safe language

### Cross-compilation for different environments

GOOS=linux go build

* recommends gorilla toolkit for web
* buffalo framework (like Ruby on Rails)
* doesn't like beego framework because he feels like it gives you much more than you need

### Types

* Storage
* Representation

### Recommended ways to declare variables

```
var d bool //unlike C, this creates a default zero value
cc := 3.1459 //short variable declaration; it infers the type for you
aa := int32(10) // this is a type conversion (not like a C cast)
```

* [Default zero values](http://yourbasic.org/golang/default-zero-value/)
* Can have int, int8, int16, int32, int64 - by default you should probably just use int
* Go doesn't allow unused variables

### Structs

* zero value of a struct is the zero value of all its composite fields

### Pointers

* everything in Go is pass by value
* for every value that exists of type T, another type exists of type *T
* & = address (where is the box?)
* * = value pointer points to (what is in the box?) - dereferences the pointer

### Constants

* iota - increases the value by 1 in a const block

```
// 0, 1, 2
const (
  unknown = iota
  active = iota
  inactive = iota
)

// 1, 2, 3
const (
  unknown = iota + 1
  active
  inactive
)

// bitshift values - golang infers the expression pattern
const (
  Ldate = 1 << iota
  Rdate
  Tdate
)
```

### Arrays

* Contiguous blocks of memory gives us mechanical sympathy
* golang.org/pkg - the builtin link shows you all the built in functions

### Slices

* Slices can get bigger (unlike arrays)
* Slice is a 3 word structure
* First word is a pointer to a backing array
* Second word is the length
* Third word is the capacity of the whole backing array structure
* Any array or slice can be re-sliced
* append is the only method that can grow a slice

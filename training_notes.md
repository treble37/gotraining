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

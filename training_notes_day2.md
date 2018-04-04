## Day 2 4/4/18 Wednesday - Notes

### Maps

* keys are not sorted (this is by design from the Golang team)

```
u, ok := users["Roy"]
//ok will be true if the value exists there
```

### Methods

```
func (u user) notify() {}
```

* "(u user)" is the method receiver
* You can add methods to any type you create (but not to types you didn't create)
* If you need to mutate bill in topics/go/language/methods/example1/example1.go, you should use pointers
* General advice: if you need to mutate something on a type, provide a pointer/receiver method (this advice isn't always true, but generally is)
* Look at the nature of your type. Is that thing something that should be copied?

### Interfaces

* The compiler checks whether your type implements an interface - sort of like duck typing in Ruby
* see [Russ Cox article](https://research.swtch.com/interfaces)

Method Sets

          T     *T
Val Rcvr |y   |y   |
         |----|----|
Ptr Rcvr |n   |y   |

### Design (Interfaces)

* see topics/go/design/composition/assertions/example1/example1.go for type assertion
* type switches

```
switch v := m.(type) {
  case bike:
    fmt.Println("is a bike", v)
  case skateboard:
    fmt.Println("is a skateboard", v)
  default:
    fmt.Println("don't know")
}
```

* everything satisfies an empty interface (see example3.go in design/composition/assertions)

### Embedding

### Exporting

* dep is the current recommendation for package management
* see counters.go

### Composition

* see topics/go/design/grouping
* try to avoid type pollution
* see topics/go/design/mocking

### Day 2 Exercise

* it is virtually impossible to unit test main

#### Setting up GOPATH

Put this gotraining repo in ~/Documents/github_code/src

Set GOPATH=/Users/hallmarklabs/Documents/github_code

* Notice that your GOPATH needs to point to the parent directory of src
## Resources

[Awesome Go](https://awesome-go.com/)

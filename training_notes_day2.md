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

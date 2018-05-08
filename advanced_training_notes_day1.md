## Day 1 5/8/18 Tuesday - Notes

* mechanical sympathy
* NUMA - non-uniform memory access

### Engineering choices

* 3 things to measure: integer, address, word (generic term for memory allocation)
* string is a 2 word data structure
* var is a readability indictator for a "zero" value state
* use the short variable declaration operator for non-zero values

* The below takes up 8 bytes due to the padding of bytes; since int16 is
  2 bytes there's a "1 byte padding" after the bool so that the int16
  ends on a 2 byte boundary

```
type example struct {
  flag bool // 1 byte
  counter int16 // 2 bytes
  pi float 32 // 4 bytes
}
```

* Don't order your structs from largest to smalles unless you need the
  micro-optimization; structure for readability

* return example{} - can use empty literal construction in a return
* e2 := example{} - this is a code smell

### Pointer mechanics

* Gm = goroutine, P = pointer, m = memory
* Gm and m are responsible for keeping track of what's next to execute,
  Gm at the go level and m at the O/S level
* Gm -> P -> m
* pointers in Go are considered literal types
* "*" says value the variable points to - indirect read or write
* "&" says "address of"
* go wants values on the stack as a first priority (rather than heap)
* allocations in go are what is in the heap (not the stack)
* Do not use pointer semantics during construction if you're going to
  assign it to a variable

```
//don't do this
u := &user{
  name: "Bill",
  email: "bill@ardanlabs.com",
}
```

* go build -gcflags "-m -m" - runs an escape analysis report
* no two goroutines can share stackspace; anything that needs to be shared between goroutines has to be put on the heap

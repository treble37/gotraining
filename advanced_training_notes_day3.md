## Day 3 5/10/18 Thursday - Notes

## Synchronization vs Orchestration

* orchestration - when performance is critical prefer orchestration to synchronization

## Concurrency (repo)

* topics/go/concurrency/goroutines/example1/example1.go
* waitgroup is a counting semaphore
* you should be calling "Add()" just once on a waitgroup
* try not to create go routines that work on timer loops because you won't hit the deadlock checking feature
* topics/go/concurrency/goroutines/example2/example2.go
* if you're writing multi-threaded tests, you'll want to use a "p" switch to up the number of CPUs (esp. on a windows machine due to the windows scheduler)

## Data race (repo)

* if the race detector detects a race in 30-60 sec and you don't know what's going on
* topics/go/concurrency/data_race/example2/example2.go
* Your Lock() and Unlock() must be in the same line of sight (same function) otherwise you'll end up with race conditions (you can do lock/defer/unlock)
* the longer the lock is held, the longer the latencies
* topics/go/concurrency/data_race/example5/example5.go
* topics/go/concurrency/data_race/advanced/example1/example1.go

## Channels

* channels are for signaling events
* topics/go/concurrency/channels/README.md
* your job is to write software that can deal with failure
* Bill notices these things are often missing: software needs to be able to start up and shutting down cleanly; you should test this from day 1 consistently
* you might want to send to a nil channel for rate limiting
* topics/go/concurrency/channels/example1/example1.go - see this for patterns
* when you range over a channel you are performing a receive
* no fan out patterns in a web service
* select statement allows you to perform channel sends/receives - helps you create event loops

## Contexts

* topics/go/packages/context/example1/example1.go
* topics/go/packages/context/example5/example5.go
* comments are documentation
* topics/go/concurrency/patterns/logger/main/main.go

## Benchmarking

* merge sort is the worst sorting algorithm to use in Go (allocation hell because of the splitting)
* when running benchmarks, your machine must be idle

## Testing

* topics/go/testing/benchmarks/basic/basic_test.go
* topics/go/testing/tests/example1/example1_test.go
* you should check error values in your tests
* topics/go/testing/tests/example5/example5_test.go - subtesting
* don't mock database tests
* topics/go/testing/tests/example4/example4.go - example webserver
* doc.go - allows you to put overview comments in your documentation; tabs and spaces determine whether you get code blocks
* go test -cover
* go test -coverprofile p.out ; go tool cover -html p.out
* ngrok - testing tool written in Go

## Profiling

* topics/go/profiling/stack_trace/README.md
* topics/go/profiling/trace/trace.go (try "web list find")
* topics/go/profiling/project/search/search.go
* GODEBUG=gctrace=1 ./project > /dev/null - use this to detect memory leaks by running this tool and paying attention to these numbers: (4->4->0 MB)
* scavenger - return any extra heap memory to the operating system
* see /debug/pprof in your web service
* go tool pprof -alloc_space http://localhost:5000/debug/pprof/heap - get some low hanging fruit
  * top 40 -cum
* topics/go/profiling/memcpu/stream.go

## Resources

* Bill will sit on hangout with you if you need help...Reach out to him in slack
* Safari online has 14 hours of this class online
* see Opening.md - Bill will give you a free PDF of Go In Action

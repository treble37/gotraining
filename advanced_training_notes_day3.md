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

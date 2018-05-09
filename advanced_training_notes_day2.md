## Day 2 5/9/18 Wednesday - Notes

* Think of interface types as being "value-less"
* see topics/go/language/interfaces/README.md

### 10:00 AM composition

* topics/go/design/composition/README.md
* since Go apps are built in small units, it's usually better to copy/paste code (non-DRY)
* interfaces should group by "what you do" (behavior) and not "what you are" - so Animal / Business / etc are bad interfaces; Notify / SendEmail / etc are better interfaces
* see time.go in stdlib for an example of a good type - Duration
* see topics/go/design/composition/grouping/example2/example2.go
* Bill believes in prototype driven development, not test driven development because if you start with tests and interfaces first, you are guessing at what API contract you need; it's better to build a working prototype implementation
* interface pollution guidelines - https://play.golang.org/p/s6HAmeT6oT

### 1:00 PM errors

* topics/go/design/error_handling/README.md
* if you ever have code where you can't use "err" for errors (e.g.,
  err1, err2, etc.), you have bad code
* don't use "else" - have as many returns as you possibly need
* see https://play.golang.org/p/EFPJu4mTJA for good contextual errors
* you should use the default error type until it doesn't give you context (e.g., a network error type)
* see golang.org/src/net/net.go
* see https://play.golang.org/p/0AUU_sJsec
* this function header "func fail() ([]byte, *customError)" should read "func fail() ([]byte, error)"
* most go loggers use the empty interface to write to the log
* see example6.go in error_handling
* errors in pkg repo of github
* https://github.com/ardanlabs/service/blob/master/internal/platform/web/response.go

#### 2:30 PM packages

* see topics/go/design/packaging/README.md
* https://www.goinggo.net/2017/02/design-philosophy-on-packaging.html
* https://www.goinggo.net/2017/02/package-oriented-design.html
* fight the urge to create a common type system here
* importing packages should only go down the tree

#### 3:30 PM threads

* cpu bound workload - thread never goes into a waiting state - it is purely calculating (e.g., calculating pi); having more threads here won't help
* i/o bound workload - threads can help here; database latency can mean that 24 threads suddenly wake up and start running, stressing the O/S
* with go, you just have to be aware of whether your workload is CPU and I/O bound
* O/S system scheduler is pre-emptive (can switch threads at any time it wants, undeterministic, and you must write code with the idea that you don't know when a thread will be running)
* go scheduler is a cooperating scheduler - you need code events in your code to have the go scheduler to make a decision - go will handle the spinning up of threads across cores

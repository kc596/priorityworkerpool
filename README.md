## Priority Worker Pool

A worker pool in GoLang which schedules job according to priority.

### Installation

> go get github.com/kc596/PriorityWorkerPool

### Quickstart

```go

var panicHandler = func(alias string, err interface{}) {
	fmt.Println(alias, err) // or use logger
}

pool := New("testPool", 100, panicHandler)

for _, job := range jobs {  // jobs are slices of func()
	pool.Submit(job, 1+rand.Float64())
}
```

### APIs

Method | Return Type | Description
---|---|---
` New(name string, workers int, panicHandler func(alias string, err interface{})`|`*Pool` | Returns a new Worker pool
`Submit(job func(), priority float64)` | `void` | Submit a new job to worker pool
`WaitGroup()` | `*sync.WaitGroup` | Returns waitgroup to wait for all jobs submitted to finish

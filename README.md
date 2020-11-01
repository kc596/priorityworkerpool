## Priority Worker Pool

[![Build Status](https://travis-ci.org/kc596/priorityworkerpool.svg?branch=master)](https://travis-ci.org/kc596/priorityworkerpool)
[![codecov](https://codecov.io/gh/kc596/priorityworkerpool/branch/master/graph/badge.svg?token=4TOHO1P4XV)](https://codecov.io/gh/kc596/priorityworkerpool)

A worker pool in GoLang which schedules job according to priority.

### Installation

> go get github.com/kc596/priorityworkerpool

### Quickstart

```go

var panicHandler = func(alias string, err interface{}) {
	fmt.Println(alias, err) // or use logger
}


const (
	poolName   = "testPool"
	numWorkers = 1000
)

pool := New(poolName, numWorkers, panicHandler)

for _, job := range jobs {  // jobs are slices of func()
	pool.Submit(job, 1+rand.Float64())
}
```

### APIs

Method | Return Type | Description
---|---|---
` New(name string, workers int, panicHandler func(alias string, err interface{})`|`*Pool` | Returns a new worker pool
`Submit(job func(), priority float64)` | `void` | Submit a new job to worker pool
`WaitGroup()` | `*sync.WaitGroup` | Returns waitgroup to wait for all jobs submitted to finish

## Priority Worker Pool

[![Build Status](https://travis-ci.org/kc596/priorityworkerpool.svg?branch=master)](https://travis-ci.org/kc596/priorityworkerpool)
[![codecov](https://codecov.io/gh/kc596/priorityworkerpool/branch/master/graph/badge.svg?token=4TOHO1P4XV)](https://codecov.io/gh/kc596/priorityworkerpool)
[![Go Report Card](https://goreportcard.com/badge/github.com/kc596/priorityworkerpool?kill_cache=1)](https://goreportcard.com/report/github.com/kc596/priorityworkerpool)
[![Maintainability](https://api.codeclimate.com/v1/badges/a51cd48917a2ffc56aba/maintainability)](https://codeclimate.com/github/kc596/priorityworkerpool/maintainability)

A worker pool in GoLang which schedules job according to priority.

### Installation

> go get github.com/kc596/priorityworkerpool

### Quickstart

```go
import "github.com/kc596/priorityworkerpool"

const (
	poolName   = "testPool"
	numWorkers = 1000
)

var panicHandler = func(alias string, err interface{}) {
	fmt.Println(alias, err) // or use logger
}

pool := priorityworkerpool.New(poolName, numWorkers, panicHandler)

job := func() {
	// code to execute
}

pool.Submit(job, 1+rand.Float64())
```

A complete example : [Here](https://goplay.space/#DIM2U6jBjwY)

### APIs

Method | Return Type | Description
---|---|---
` New(name string, workers int, panicHandler func(alias string, err interface{})`|`*Pool` | Returns a new worker pool
`Submit(job func(), priority float64)` | `void` | Submit a new job to worker pool
`WaitGroup()` | `*sync.WaitGroup` | Returns waitgroup to wait for all jobs submitted to finish
`ShutDown()` | `void` | Delete queue and prevents pickup of next job from the queue

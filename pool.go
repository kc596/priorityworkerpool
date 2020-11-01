package priorityworkerpool

import (
	"fmt"
	"github.com/kc596/UGCPriorityQueue/maxpq"
	"sync"
)

// Pool is type for Worker Pool
type Pool struct {
	name         string
	workers      chan int
	jobQueue     *maxpq.PQ
	panicHandler func(alias string, err interface{})
	wg           *sync.WaitGroup
}

/***************************************************************************
* Worker Pool APIs
***************************************************************************/

// New creates a new worker pool to manage goroutines
func New(name string, workers int, panicHandler func(alias string, err interface{})) *Pool {
	pool := &Pool{
		name:         name,
		workers:      make(chan int, workers),
		jobQueue:     maxpq.New(),
		panicHandler: panicHandler,
		wg:           &sync.WaitGroup{},
	}
	for i := 1; i <= workers; i++ {
		pool.workers <- i
	}
	pool.start()
	return pool
}

// Submit a job to worker pool
func (pool *Pool) Submit(job func(), priority float64) {
	defer func() {
		if err := recover(); err != nil {
			pool.panicHandler("SubmitJob", err)
		}
	}()
	node := maxpq.NewNode(job, priority)
	pool.jobQueue.Insert(node)
	pool.wg.Add(1)
}

// WaitGroup to wait for all jobs submitted to finish
// WARNING: would not wait if there are no jobs at the instant
func (pool *Pool) WaitGroup() *sync.WaitGroup {
	return pool.wg
}

/***************************************************************************
* Helper functions
***************************************************************************/

func (pool *Pool) start() {
	go func() {
		for {
			if pool.jobQueue.Size() > 0 {
				pool.schedule()
			}
		}
	}()
}

func (pool *Pool) schedule() {
	defer func() {
		if err := recover(); err != nil {
			pool.panicHandler("JobQueue", err)
		}
	}()
	node, err := pool.jobQueue.Pop()
	if err != nil {
		panic(err)
	}
	work := node.GetFuncValue()
	worker := <-pool.workers
	go pool.doWork(worker, work)
}

func (pool *Pool) doWork(worker int, work func()) {
	defer func() {
		if err := recover(); err != nil {
			pool.panicHandler(fmt.Sprintf("%s-%d", pool.name, worker), err)
		}
	}()
	defer func() { pool.workers <- worker }()
	defer pool.wg.Done()
	work()
}

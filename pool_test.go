package priorityworkerpool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

const (
	poolName   = "testPool"
	numWorkers = 1000
	jobCount   = uint32(100000)
)

func TestPool(t *testing.T) {
	assert := assert.New(t)
	var (
		panicCount   = uint32(0)
		panicHandler = func(alias string, err interface{}) {
			atomic.AddUint32(&panicCount, 1)
			fmt.Println(alias, err)
		}
		pool     = New(poolName, numWorkers, panicHandler)
		jobs     []func()
		executed uint32
	)

	for i := uint32(0); i < jobCount; i++ {
		jobs = append(jobs, func() {
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			atomic.AddUint32(&executed, 1)
		})
	}
	for _, job := range jobs {
		pool.Submit(job, 1+rand.Float64())
	}
	pool.WaitGroup().Wait()
	assert.Equal(jobCount, atomic.LoadUint32(&executed))
	assert.Zero(atomic.LoadUint32(&panicCount))
}

func TestPoolError(t *testing.T) {
	assert := assert.New(t)
	var (
		panicCount   = uint32(0)
		panicHandler = func(alias string, err interface{}) {
			atomic.AddUint32(&panicCount, 1)
			fmt.Println(alias, err)
		}
		pool     = New(poolName, numWorkers, panicHandler)
		executed uint32
	)
	pool.Submit(func() {
		panic("Some erroneous job")
		atomic.AddUint32(&executed, 1)
	}, 1+rand.Float64())
	pool.WaitGroup().Wait()
	assert.Zero(atomic.LoadUint32(&executed))
	assert.Greater(atomic.LoadUint32(&panicCount), uint32(0))
}

func TestPoolError2(t *testing.T) {
	assert := assert.New(t)
	var (
		panicCount   = uint32(0)
		panicHandler = func(alias string, err interface{}) {
			atomic.AddUint32(&panicCount, 1)
			fmt.Println(alias, err)
		}
		pool     = New(poolName, numWorkers, panicHandler)
		executed uint32
	)
	pool.Submit(func() {
		pool.wg.Done()
		atomic.AddUint32(&executed, 1)
	}, 1+rand.Float64())
	time.Sleep(1 * time.Second)
	assert.Greater(atomic.LoadUint32(&panicCount), uint32(0))
}

func TestScheduleError(t *testing.T) {
	assert := assert.New(t)
	var (
		panicCount   = uint32(0)
		panicHandler = func(alias string, err interface{}) {
			atomic.AddUint32(&panicCount, 1)
			fmt.Println(alias, err)
		}
		pool = New(poolName, numWorkers, panicHandler)
	)
	pool.schedule()
	time.Sleep(1 * time.Second)
	assert.Greater(atomic.LoadUint32(&panicCount), uint32(0))
}

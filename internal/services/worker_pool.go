package services

import (
	"log"
)

var workerPool chan func()

func InitWorkerPool(poolSize int) {
	workerPool = make(chan func(), poolSize)
	for i := 0; i < poolSize; i++ {
		go worker()
	}
}

func worker() {
	for task := range workerPool {
		task()
	}
}

func SubmitTask(task func()) {
	select {
	case workerPool <- task:
	default:
		log.Println("Worker pool is full, task discarded")
	}
}

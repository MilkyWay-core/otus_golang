package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorCounter, goroutineCounter int32 = 0, 0
	i := 0
	var wg sync.WaitGroup
	for {
		if int(atomic.LoadInt32(&goroutineCounter)) > n { //если свободные горутины кончились то ждем освобождения
			time.Sleep(10 * time.Nanosecond)
			continue
		}
		if int(atomic.LoadInt32(&errorCounter)) == m || i == len(tasks) { //если выполнили все таски или дошли до указанного значения ошибок то выходим
			break
		}
		i++
		atomic.AddInt32(&goroutineCounter, 1) //занимаем горутину
		wg.Add(1)
		go func(i int) {
			defer atomic.AddInt32(&goroutineCounter, -1) //освобождаем горутину
			defer wg.Done()
			if int(atomic.LoadInt32(&errorCounter)) < m { //если число ошибок превышено другими горутинами, то работу выполнять не нужно, просто завершаем
				err := tasks[i-1]()
				if err != nil {
					atomic.AddInt32(&errorCounter, 1) //сичитаем ошибки
				}
			}
		}(i)
	}
	wg.Wait()
	if int(errorCounter) >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

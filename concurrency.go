package goo

import (
	"errors"
	"sync"
)

func ConcurrentWithLimit[A, B any](data []A, limit int, processFunc func(A) B) []B {
	var back = make([]B, len(data))
	semaphore := make(chan struct{}, limit)
	var wg sync.WaitGroup

	for k, a := range data {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(ik int, ia A) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			back[ik] = processFunc(ia)
		}(k, a)
	}

	wg.Wait()

	return back
}

func ConcurrentWithLimitRetErrs[A, B any](data []A, limit int, processFunc func(A) (B, error)) ([]B, error) {
	var (
		wg sync.WaitGroup

		errs error

		back     []B
		backLock sync.Mutex

		semaphore = make(chan struct{}, limit)
	)

	for _, a := range data {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(ia A) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			ret, err := processFunc(ia)

			if err != nil {
				errs = errors.Join(errs, err)
				return
			}
			backLock.Lock()
			back = append(back, ret)
			backLock.Unlock()
		}(a)
	}

	wg.Wait()

	return back, errs
}

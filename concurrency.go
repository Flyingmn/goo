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

func ConcurrentWithLimitAndExtraParams[A, B any, EK comparable, EV any](data []A, limit int, extra map[EK]EV, processFunc func(A, map[EK]EV) (B, error)) ([]B, error) {
	var (
		wg sync.WaitGroup

		errs error

		back = make([]B, len(data))

		semaphore = make(chan struct{}, limit)
	)

	for k, a := range data {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(ik int, ia A) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			ret, err := processFunc(ia, extra)

			if err != nil {
				errs = errors.Join(errs, err)
				return
			}

			back[ik] = ret
		}(k, a)
	}

	wg.Wait()

	return back, errs
}

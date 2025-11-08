package goo

import (
	"errors"
	"strings"
	"sync"
)

// 并发执行一个func，支持限制并发数，func遍历执行传入的slices的每个元素
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

// 并发执行一个func,支持控制并发数以及错误返回，func遍历执行传入的slices的每个元素
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
				errs = ErrJoin(errs, err)
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

// 分块执行一个func，支持设置分快大小；func遍历执行每个块
func ChunkExec[V any, R any](values []V, chunkNum int, f func(miniVals []V) ([]R, error)) (res []R, errs error) {
	chunks := ArrayChunk(values, chunkNum)

	for _, chunk := range chunks {
		miniRes, minierr := f(chunk)
		if minierr != nil {
			errs = ErrJoin(errs, minierr)
			continue
		}

		res = append(res, miniRes...)
	}

	return res, errs
}

// 拼接多个error
func ErrJoin(errs ...error) error {

	var errsStrList []string

	for _, err := range errs {
		if err != nil {
			errsStrList = append(errsStrList, err.Error())
		}
	}

	if len(errsStrList) == 0 {

		return nil
	}

	return errors.New(strings.Join(errsStrList, ","))
}

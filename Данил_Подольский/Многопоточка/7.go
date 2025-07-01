package main

import (
	"context"
	"errors"
	"fmt"
)

// Задача 7
// ===========================================================

type Result struct {
	result bool
}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
	// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
	// с ошибкой - отдаем последнюю полученную ошибку
	ch := make(chan Result)
	err := make(chan error)

	var lastError error

	for _, sf := range sfs {
		go func(sf SearchFunc) {

			select {
			case <-ctx.Done():
				err <- ctx.Err()
			default:
				res, ok := sf(ctx, query)

				if ok != nil {
					err <- ok
				} else {
					ch <- res
				}
			}
		}(sf)
	}

	for i := 0; i < len(sfs); i++ {
		select {
		case val := <-ch:
			return val, nil
		case er := <-err:
			lastError = er
		}
	}

	// sfs(ctx, "")

	// var res Result
	return Result{result: false}, lastError
}

func serchFunc(ctx context.Context, query string) (Result, error) {

	query = "no found"

	if query == "found" {
		return Result{result: true}, nil
	} else {
		return Result{result: false}, errors.New("eeee")
	}
}

func serchFunc2(ctx context.Context, query string) (Result, error) {

	query = "no found"

	if query == "found" {
		return Result{result: true}, nil
	} else {
		return Result{result: false}, errors.New("eeee")
	}
}

func serchFunc3(ctx context.Context, query string) (Result, error) {

	query = "found"

	if query == "found" {
		return Result{result: true}, nil
	} else {
		return Result{result: false}, errors.New("eeee")
	}
}

func main() {

	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	var funcArr []SearchFunc = []SearchFunc{serchFunc, serchFunc2, serchFunc3}

	res, er := MultiSearch(ctx, "", funcArr)

	fmt.Println(res, er)

	// query := "123"

	// SearchFunc1(ctx, query)
}

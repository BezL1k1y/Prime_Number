package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var (
		flag    string
		data    []int64
		start   int64
		end     int64
		counter int64
	)

	fmt.Println("Введите стартовое нечетное число: ")
	fmt.Scan(&start)
	fmt.Println("Введите конечное число: ")
	fmt.Scan(&end)

	now := time.Now()
	for ; start <= end; start += 2 {
		if isPrime(start) {
			data = append(data, start)
			counter += 1
		}
	}

	fmt.Println("Простые числа: ", data)
	fmt.Println("Кол-во простых чисел: ", counter)
	fmt.Println("Время выполнения кода:", time.Since(now))
	fmt.Scan(&flag)
}

func isPrime(x int64) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for num := int64(3); num <= int64(math.Sqrt(float64(x)))+1; num += 2 {
		if x%num == 0 {
			return false
		}
	}
	return true
}

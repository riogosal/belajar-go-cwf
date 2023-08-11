package main

import (
	"fmt"
	"testing"
)

const numIterations = 1000000

func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := passByValue(numIterations)
		_ = result
	}
}

func BenchmarkPassByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := passByReference(numIterations)
		_ = result
	}
}

func passByValue(iterations int) int {
	value := 0
	for i := 0; i < iterations; i++ {
		value++
	}
	return value
}

func passByReference(iterations int) int {
	value := 0
	for i := 0; i < iterations; i++ {
		tambah1ByReference(&value)
	}
	return value
}

func tambah1ByReference(nomor *int) {
	*nomor = *nomor + 1
}

func main() {
	fmt.Println("Benchmarking pass by value vs pass by reference:")
	fmt.Println("Pass by value:")
	resultByValue := testing.Benchmark(BenchmarkPassByValue)
	fmt.Println(resultByValue)

	fmt.Println("Pass by reference:")
	resultByReference := testing.Benchmark(BenchmarkPassByReference)
	fmt.Println(resultByReference)
}

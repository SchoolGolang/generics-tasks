package main

import (
	"fmt"
	"strconv"
)

//Various sensors from different manufacturers send us
//various data like pressure level along with their ids.
//Some of them send int64, some - int8, float64, bool or string.

//Task: make a collector method that returns a map where key is ID of sensor
//and value is received data.

type DataStruct[T any] struct {
	ID   int64
	Data T
}

func collectData[T any](arr []DataStruct[T]) map[int64]T {
	res := make(map[int64]T)

	for i := 0; i < len(arr); i++ {
		res[arr[i].ID] = arr[i].Data
	}

	return res
}

func main() {
	data := make([]DataStruct[string], 4)

	for i := 0; i < len(data); i++ {
		data[i] = DataStruct[string]{
			ID:   int64(i),
			Data: strconv.Itoa(i * 10),
		}
	}

	res := collectData[string](data)

	fmt.Println(res)
}

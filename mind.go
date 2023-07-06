package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello world")
	//
	//http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
	//	time.Sleep(time.Second * 10)
	//	writer.Write([]byte(`{"name": "marisa"}`))
	//})
	//err := http.ListenAndServe(":1111", nil)
	//fmt.Println(err.Error())

	//st := time.Now()
	//mapa := make(map[int]string, 100_000_000)
	//for i := 0; i < 100_000_000; i++ {
	//	mapa[i] = fmt.Sprintf("nihao %d", i)
	//}
	//fmt.Println("预分配空间，该函数执行完成耗时：", time.Since(st))
	//
	//st = time.Now()
	//mapa = make(map[int]string)
	//for i := 0; i < 100_000_000; i++ {
	//	mapa[i] = fmt.Sprintf("nihao %d", i)
	//}
	//fmt.Println("未预分配空间，该函数执行完成耗时：", time.Since(st))

	start := time.Now()
	mapa := make(map[int]int, 100_000_000)
	var n int
	n = 100_000_000
	for i := 0; i < n; i++ {
		mapa[i] = i
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

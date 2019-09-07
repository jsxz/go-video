package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	//跳过当前测试
	//t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("wrong result of Print1to20")
	}
}
func testPrint2(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("test2 wrong result of Print1to20")
	}
}
func TestAll(t *testing.T) {
	t.Run("TestPrint", testPrint)
	t.Run("TestPrint2", testPrint2)
}
func TestMain(m *testing.M) {
	fmt.Println("test begin...")
	m.Run()
}

//b.N会变化，如果测试方法 所用时间不稳定，会跑不完
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
		//a(n)
	}
}
func a(n int) int {
	for n > 1 {
		n--
	}
	return n
}

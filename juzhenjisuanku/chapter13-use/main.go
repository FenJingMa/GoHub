package main

import (
	"GoBook/code/chapter13/matrix" //引入matrix包
	"fmt"
)

func main() {
	//初始化两个矩阵
	m1 := matrix.Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := matrix.Matrix([][]float64{[]float64{14, 12, 10}, []float64{16, 17, 19}})
	//计算矩阵相加
	m3, _ := m1.Add(m2)
	//输出矩阵相加的计算结果
	fmt.Println("m3=", m3)
}

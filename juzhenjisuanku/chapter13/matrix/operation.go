package matrix

import (
	"errors"
	"math"
)

//获取行向量的维度
//由于行向量的底层数据类型为[]float64切片，那么仅需要计算行向量切片的长度即可
func (rv1 RowVector) GetShape() (s [2]int) {
	s[0] = 1
	s[1] = len(rv1)
	return
}

//计算行向量的标量乘法
//这里我们使用*RowVector指针类型的变量rv1作为指针接收者
//从而在函数运行完毕后改变rv1自身
func (rv1 *RowVector) Mul(c float64) {
	//获取rv1的维度，也就是[]float64切片的长度
	l := rv1.GetShape()[1]
	//创建一个rvTemp变量，用于复制和存放rv1指针变量所指向的底层[]float64切片
	rvTemp := make([]float64, l, l)
	//复制rv1指针对应的切片，赋值给rvTemp变量
	copy(rvTemp, *rv1)
	//遍历rvTemp切片，将每个元素乘以c
	for i := 0; i < l; i++ {
		rvTemp[i] *= c
	}
	//将rvTemp复制给*rv1所对应的底层切片，从而改变rvTemp
	copy(*rv1, rvTemp)
}

//根据给定向量长度l初始化向量，l需要大于0，否则返回error
func NewRowVector(l int) (rv1 RowVector, err error) {
	if l <= 0 {
		//若l小于等于0，则返回error
		err = errors.New("the dimension of row vectors must > 0")
		return
	}
	rv1 = make([]float64, l, l)
	return
}

//定义行向量的加法运算，判断两个行向量shape是否相同
func (rv1 RowVector) Add(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v + rv2[i]
	}
	return
}

//定义行向量的减法运算，判断两个行向量shape是否相同
func (rv1 RowVector) Minus(rv2 RowVector) (rv3 RowVector, err error) {
	//若行向量维度不同，则返回err
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vector are not equal")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v - rv2[i]
	}
	return
}

//定义行向量的点乘运算
func (rv1 RowVector) Dot(rv2 RowVector) (dot float64, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	for i, v := range rv1 {
		dot += v * rv2[i]
	}
	return
}

//由于大于3维的向量叉乘运算较为复杂
//这里简单起见，我们仅实现2维和3维向量的叉乘
func (rv1 RowVector) Cross(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the shape of the two vectors are not equal")
		return
	}
	dim := rv1.GetShape()[1]
	if dim != 2 && dim != 3 {
		err = errors.New("we can only calc the 2 or 3 dimensions row vector")
		return
	}
	rv3, _ = NewRowVector(3)
	switch dim {
	//两个2维向量做叉乘运算，得到的是一个与两个2维向量相垂直的向量
	//可以想象为两个XY平面上的向量做叉乘,得到的向量沿Z轴方向
	case 2:
		rv3[0] = 0
		rv3[1] = 0
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	//两个3维向量做叉乘,计算公式如下:
	//X×Y=[x_2*y_3-x_3*y_2, x_3*y_1-x_1*y_3, x_1*y_2-x_2*y_1 )]
	case 3:
		rv3[0] = rv1[1]*rv2[2] - rv1[2]*rv2[1]
		rv3[1] = rv1[2]*rv2[0] - rv1[0]*rv2[2]
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	}
	return
}

//定义计算行向量的方法
func (rv1 RowVector) Length() (l float64) {
	//l的初值默认就是0.0
	for _, v := range rv1 {
		l += v * v
	}
	l = math.Sqrt(l)
	return
}

//定义行向量的转置运算，输出结果为列向量
func (rv1 RowVector) Transpose() (cv ColumnVector) {
	//Go语言会对ColumnVector和[][1]float64两个类型之间做隐式转换
	cv = make([][1]float64, rv1.GetShape()[1], rv1.GetShape()[1])
	for i, v := range rv1 {
		cv[i][0] = v
	}
	return
}

//定义获取矩阵形状的方法
func (m1 Matrix) GetShape() (s [2]int) {
	s[0] = len(m1)
	s[1] = len(m1[0])
	return
}

//定义矩阵与标量c相乘的运算
func (m1 *Matrix) Mul(c float64) {
	row, col := m1.GetShape()[0], m1.GetShape()[1]
	mTemp := make([][]float64, row, row)
	//将m1矩阵复制到临时变量复制给临时变量mTemp
	copy(mTemp, *m1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mTemp[i][j] *= c
		}
	}
	//将临时变量mTemp矩阵复制回给m1
	copy(*m1, mTemp)
}

//矩阵的初始化方法，r为矩阵的行数，c为矩阵的列数
func NewMatrix(r, c int) (mat Matrix, err error) {
	//行数或列数小于等于零，返回err
	if r <= 0 || c <= 0 {
		err = errors.New("rows and columns of the matrix must >0")
		return
	}
	//初始化mat的行
	mat = make([][]float64, r, r)
	//此处不使用for range是因为我们要改变遍历元素的值
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c, c)
	}
	return
}

//矩阵的加法
func (m1 Matrix) Add(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	//获取矩阵的形状
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return
}

//矩阵的减法
func (m1 Matrix) Minus(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return
}

//两个矩阵相乘
func (m1 Matrix) MatMul(m2 Matrix) (m3 Matrix, err error) {
	r1, c1 := m1.GetShape()[0], m1.GetShape()[1]
	r2, c2 := m2.GetShape()[0], m2.GetShape()[1]
	if c1 != r2 {
		err = errors.New("the shape of the two matrix are not match")
		return
	}
	//初始化计算结果m3
	m3, _ = NewMatrix(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			//取出m1的第i行
			v1 := RowVector(m1[i])
			v2 := make([]float64, r2, r2)
			for k := 0; k < r2; k++ {
				//将m2的第j列值依次放入v2中
				v2[k] = m2[k][j]
			}
			//利用行向量乘法计算m3的矩阵元
			m3[i][j], _ = v1.Dot(v2)
		}
	}
	return
}

//定义矩阵的转置运算，本质上就是将矩阵元素的索引值i和j进行互换
func (m1 Matrix) Transpose() (m2 Matrix) {
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m2, _ = NewMatrix(c, r)
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			m2[i][j] = m1[j][i]
		}
	}
	return
}

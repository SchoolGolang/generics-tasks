package main

type A[T I] struct {
	val  int64
	val2 int64
}

type B[T I] struct {
	val  float64
	val2 float64
}

type I interface {
	int64 | float64
}

type Q[T I] interface {
	GetVal() T
	GetVal2() T
}

func (a A[I]) GetVal() int64 {
	return a.val
}

func (a A[I]) GetVal2() int64 {
	return a.val2
}

func (b B[I]) GetVal() float64 {
	return b.val
}

func (b B[I]) GetVal2() float64 {
	return b.val2
}

func SumValues[T I](obj Q[T]) T {
	return obj.GetVal() + obj.GetVal2()
}

//Task: implement SumValues for any structure like A and B as single utility method

func main() {
	a := A[]{
		val:  10,
		val2: 20,
	}

	b := B[float64]{
		val:  11,
		val2: 22,
	}
	SumValues[I](a)
}

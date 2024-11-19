package utils

type Progress interface {
	Grow(num int, MyStrVal string)
	Done()
	SetTotal(total int)
}

package utils

import "fmt"

type ProgressNotify struct {
	totalCount   int
	currentCount int
}

func (p *ProgressNotify) Grow(num int, MyStrVal string) {
	p.currentCount += num
	fmt.Printf("已完成 %d 步, 一共 %d 步\n", p.currentCount, p.totalCount)
}

func (p *ProgressNotify) Done() {
	fmt.Println("任务完成！")
}

func (p *ProgressNotify) SetTotal(total int) {
	p.totalCount = total
}

func NewProgress() *ProgressNotify {
	fmt.Println("任务开始")

	return &ProgressNotify{
		totalCount:   0,
		currentCount: 0,
	}
}

package factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type Donglaishun struct {
}

func (d *Donglaishun) GetFood() {
	fmt.Println("东来顺的饭菜已经生产完毕，就绪。。。")
}

type Qingfeng struct{}

func (q *Qingfeng) GetFood() {
	fmt.Println("庆丰的包子已经生产完毕，就绪。。。")
}

func NewRestautant(name string) Restaurant {
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}

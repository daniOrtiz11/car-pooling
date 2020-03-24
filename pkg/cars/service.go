package cars

import (
	"log"
)

type car struct {
	id         int
	capacity   int
	seats      int
	seatsTaken int
	journeying bool
}

var cars []car

/*
Service is a
*/
type Service interface {
	ServiceImpl()
}

/*
ServiceImpl is a
*/
func ServiceImpl() {
	log.Println("in cars")
}

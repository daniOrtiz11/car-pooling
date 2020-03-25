package tables

import (
	"log"
)

type table struct {
	id         int
	capacity   int
	seats      int
	seatsTaken int
	eating     bool
}

var tables []table

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
	log.Println("in tables")
}

package bill

import "log"

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
	log.Println("heee22222")
}

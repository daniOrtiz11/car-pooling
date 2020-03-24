package locate

import "github.com/daniOrtiz11/car-pooling/internal/database"

/*
Service is a
*/
type Service interface {
	ServiceImpl()
}

/*
ServiceImpl is a
*/
func ServiceImpl(idToSearch int) int {
	state := database.FindJourneyByID(idToSearch)
	println(state)
	//WP
	return 0
}

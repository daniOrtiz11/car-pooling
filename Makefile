# .PHONY: database-up
# database-up:
# 		docker-compose -f database/docker-compose.yml  up

# .PHONY: database-down
# database-down:
# 		docker-compose -f database/docker-compose.yml  down

.PHONY: table-booking-up
table-booking-up:
		GOOS=linux GOARCH=amd64 go build cmd/table-booking/main.go 
		mv main table-booking
		docker-compose -f docker-compose.yml  up --build
		#docker build -t table-booking:latest .
		#docker run -p 9091:9091 --net=database_default --link db table-booking:latest 
		#docker run -d -p 9091:9091 table-booking:latest 

.PHONY: table-booking-down
table-booking-down:
		docker-compose -f docker-compose.yml  down
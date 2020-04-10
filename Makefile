.PHONY: database-up
database-up:
		docker-compose -f database/docker-compose.yml  up

.PHONY: database-down
database-down:
		docker-compose -f database/docker-compose.yml  down

.PHONY: table-booking-services-up
table-booking-services-up:
		GOOS=linux GOARCH=amd64 go build cmd/table-booking/main.go 
		mv main table-booking
		docker build -t table-booking:latest .
		docker run -p 9091:9091 table-booking:latest 
		#docker run -d -p 9091:9091 table-booking:latest 

.PHONY: table-booking-services-down
table-booking-services-down:
		docker stop table-booking:latest
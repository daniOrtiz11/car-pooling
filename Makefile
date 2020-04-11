# Copyright 2020 Daniel Ortiz @daniOrtiz11 (https://github.com/daniOrtiz11). All rights reserved.
# Code under the MIT License. See LICENSE in the project root for license information.

.PHONY: help
help: ## Show this help.
	@echo Welcome to Table-booking API!
	@echo Prerequisites:
	@echo 1. Golang
	@echo 2. Docker
	@echo 3. Unix based os
	@echo Options:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: table-booking-up
table-booking-up: ## Command to run the application.
		GOOS=linux GOARCH=amd64 go build cmd/table-booking/main.go 
		mv main table-booking
		docker-compose -f docker-compose.yml  up --build

.PHONY: table-booking-down
table-booking-down: ## Command to shut down the application. 
		docker-compose -f docker-compose.yml  down

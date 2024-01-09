run:
	go run main.go

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrate:
	migrate -path migrations -database postgres://picos:picos@localhost:5432/picos?sslmode=disable up
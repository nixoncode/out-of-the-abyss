run:
	@go run main.go serve

migrate\:create:
	 migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate\:up:
	migrate -database "mysql://root:password@tcp(localhost:3306)/out_of_the_abyss" -path internal/db/migrations up



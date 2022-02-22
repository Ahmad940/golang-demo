migrate:
	migrate -database postgres://postgres:admin@localhost:5432/todo -path postgres/migrations up

migration:
	@read -p "Enter your name: " name; \
	migrate create -ext sql -dir postgres/migrations -seq $$name

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://postgres:Aa@123456@localhost:5432/mydb?sslmode=disable" -verbose up
migratedown:
    migrate -path db/migration -database "postgresql://postgres:Aa@123456@localhost:5432/mydb?sslmode=disable"-verbose down

 
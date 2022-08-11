postgres:
	docker run --name bankManagement -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec  -it bankManagement createdb --username=root --owner=root bankManagement

dropdb:
	docker exec  -it bankManagement dropdb bankManagement
#makemigrations up and down
sqlc:
	sqlc generate

.PHONY:postgres createdb dropdb sqlc
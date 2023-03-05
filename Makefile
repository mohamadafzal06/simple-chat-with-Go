postgresget:
		docker pull postgres:15-alpine

postgresinit:
		docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=$POSTGRES_USER -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -d postgres:15-alpine

postgresstat:
		docker exec -it postgres15 psql 

createdb:
		docker exec -it postgres15 createdb --username=$POSTGRES_USER --owner=$POSTGRES_OWNER go-chat


dropdb:
		docker exec -it postgres15 dropdb go-chat

migratecreate:
		migrate create -ext sql -dir internal/db/migration add_user_table
migrateup:
	migrate -path internal/db/migration -database "postgresql://root:password@localhost:5432/go-chat?sslmode=disable" -v up
migratedown:
		migrate create -ext sql -dir internal/db/migration add_user_table


.PHONY: postgresget postgresinit postgresstat createdb dropdb migratecreate migrateup migratedown

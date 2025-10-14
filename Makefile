GO := go
APP := subsTracker
MIGRATIONSDIR := migrations
DBURL ?= postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

run:
	${GO} run cmd/server/main.go

build: 
	$(GO) build -o $(APP) cmd/server/main.go

create:
	migrate create -ext sql -dir ${MIGRATIONSDIR} ${NAME}

up:
	migrate -path ${MIGRATIONSDIR} -database ${DBURL} up

down:
	migrate -path ${MIGRATIONSDIR} -database ${DBURL} down

migrate-status:
	migrate -path ${MIGRATIONSDIR} -database ${DBURL} version
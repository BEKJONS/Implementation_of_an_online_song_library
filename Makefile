CURRENT_DIR = $(shell pwd)

DB_URL := postgres://postgres:123321@localhost:5432/song_library?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '${DB_URL}' -verbose up

mig-down:
	migrate -path migrations -database '${DB_URL}' -verbose down

mig-force:
	migrate -path migrations -database '${DB_URL}' -verbose force 1

mig-create:
	migrate create -ext sql -dir migrations -seq song_library

run_db:
	docker compose build postgres && docker compose up -d migrate && docker compose up -d redis &&make mig-up

run:
	docker compose build app && docker compose up -d app

swag-gen:
	~/go/bin/swag init -g internal/controller/http/router.go -o docs
#   rm -r db/migrations

tidy:
	go mod tidy; go mod vendor
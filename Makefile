

.PHONY: run dcup itdb

run:
	go run main.go

dcup:
	docker compose up

dcdown:
	docker compose down

itdb:
	docker compose exec -it db psql -U rootuser -d anzeigen

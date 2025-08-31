# lintチェック
lint:
	go tool golangci-lint run

db-up:
	docker compose up -d db

db-down:
	docker compose down db

migration/new:
	go tool sql-migrate new --env="local" ${FILE_NAME}

migrate/up:
	make db-up
	sleep 5
	go tool sql-migrate up --env="local"

migrate/down:
	make db-up
	sleep 5
	go tool sql-migrate down --env="local"

xo:
	rm -f models/*.xo.go
	go tool dbtpl schema --src dbtpl -o models postgres://postgres:password@localhost:5438/daredemo-design_local?sslmode=disable \
	--exclude=gorp_migrations

# ローカルDB接続
psql:
	docker compose exec db psql -U postgres -d daredemo-design_local

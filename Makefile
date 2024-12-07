dev:
	docker compose up --build

prod:
	docker compose -f docker-compose.prod.yml up --build -d

dbInit:
	# To initialize database
	# cat <sql_file> | docker exec -i <container_name> psql -U <username> -d <database_name>
	@echo "Initializing database..."
	cat ./server/v1/sql/db.sql | docker exec -i sits_db psql -U postgres -d sits

dbSeed:
	# To seed database
	# cat <sql_file> | docker exec -i <container_name> psql -U <username> -d <database_name>
	@echo "Seeding database..."
	cat ./server/v1/sql/seed.sql | docker exec -i sits_db psql -U postgres -d sits

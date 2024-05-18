### reset db

docker-compose run --rm migrate -path /migrations -database "cockroach://username:roach@db:26257/speed_cube_time?sslmode=disable" drop -f

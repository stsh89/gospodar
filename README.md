# gospodar

## Start docker environment

        docker-compose up

## Create database

        docker exec -it gospodar_db psql -U postgres -c "CREATE DATABASE gospodar_development"

## Run migrations

        docker exec -it gospodar_web \
          goose -dir=db/migrations postgres \
          "user=postgres dbname=gospodar_development host=gospodar_db sslmode=disable" up

## Check basic functionality

        curl localhost:3000
        curl -X POST localhost:3000/registrations -d "email=ssh@anadeainc.com"

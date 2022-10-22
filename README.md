# README.md

This project aims to create an example go/golang web-server.
Any feedback is welcome since this is a public project.

The project is integrated with docker, then running the following commands after setting up a `.env` file at `go-example-webserver/postgres/.env`, should be enough (there is an example env file in the `go-example-webserver/postgres/` directory):
```bash
docker-compose build
docker-compose --env-file ./postgres/.env up
docker-compose --env-file ./postgres/.env up -d # if you want to run the containers in the background
```

To run the functional tests written in python, we have to set up the DB, in order to do that execute the following goose command:
```bash
docker-compose run webserver sh bin/goose_apply_migrations.sh ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD}
```
And then just run:
```bash
docker-compose run python_tests pytest
```

To run the unit tests:
```bash
docker-compose run webserver sh bin/go_test.sh
```
This will generate a coverage report at `./webserver/src/report/coverage.html`.

To run pylint use the following command:
```bash
docker-compose run python_tests sh bin/pylint.sh
```

Every change has to be made via a Pull Request, and CircleCI checks are needed.
Even for the repo administrators.

To connect to the DataBase, use the following command:
```bash
docker-compose run postgres psql --host=postgres --username=${POSTGRES_USERNAME} --dbname=hello_world
```

To create a new DB migration use the following command:
```bash
docker-compose run webserver sh bin/goose_new_migration.sh ${MIGRATION_NAME}
```

To apply the db_migrations use the following command:
```bash
docker-compose run webserver sh bin/goose_apply_migrations.sh ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD}
```

To unapply the db_migrations use the following command:
```bash
docker-compose run webserver sh bin/goose_downgrade_migration.sh ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD}
```

To add new go dependencies we just have to use the following commands. This can be done through your local environment as long as you have go installed.
```bash
go get ${DEPENDENCY}@${VERSION}
go mod tidy
```

# README.md

## Desription
This project aims to create an example go/golang web-server.
The `go-example-webserver/docs` directory includes the explanation of the application being built (the docs will have a broader scope than the actual code at the point you're reading this probably).
Any feedback is welcome since this is a public project.

## Building and running project
The project is integrated with docker, then running the following commands after setting up a `.env` file at `go-example-webserver/.env`, should be enough (there is an example env file in the `go-example-webserver` directory):
```bash
docker-compose build
docker-compose --env-file ./.env up
docker-compose --env-file ./.env up -d # if you want to run the containers in the background
```

## Functional tests and Manual tests
To run the functional tests written in python, or any manual tests, we have to set up the DB, in order to do that execute the following goose command:
```bash
docker-compose run webserver sh bin/goose_apply_migrations.sh ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD}
```
And then just run:
```bash
docker-compose run python_tests pytest
```

## Go Unit tests
To run the unit tests:
```bash
docker-compose run webserver sh bin/go_test.sh
```
This will generate a coverage report at `./webserver/src/report/coverage.html`.

## Go linting
We have some different commands to run Go linters, run the following commands:
```bash
docker-compose run webserver sh bin/go_fmt.sh
docker-compose run webserver sh bin/go_vet.sh
docker-compose run webserver sh bin/staticcheck.sh
```

## Python linting
To run pylint use the following command:
```bash
docker-compose run python_tests sh bin/pylint.sh
```

## Changes and Pull Requests
Every change has to be made via a Pull Request, and CircleCI checks are needed.
Even for the repo administrators.

## Goose and PostgreSQL DB info
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

## Adding or updating Go dependencies
To add new go dependencies we just have to use the following commands inside the `webserver` directory. This can be done through your local environment as long as you have go installed.
```bash
go get ${DEPENDENCY}@${VERSION}
go mod tidy
```

## Adding or updating Python dependencies
So far I've been adding these manually to the `test/requirements.txt` file. After that, we need to re-build the python container.

## Dependabot
This repo uses Dependabot to keep its Go dependencies up to date.
The config was created following this blog article: https://github.blog/2020-06-01-keep-all-your-packages-up-to-date-with-dependabot/
And the github docs: https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuring-dependabot-version-updates

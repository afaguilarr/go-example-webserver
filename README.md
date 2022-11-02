# README.md

## Description

This project aims to create an example go/golang web-server.
The `go-example-webserver/docs` directory includes the explanation of the application being built (the docs will have a broader scope than the actual code at the point you're reading this probably).
Any feedback is welcome since this is a public project.

## Building and running the project

The project is integrated with docker, then running the following commands after setting up a `.env` file at `go-example-webserver/.env`, should be enough (there is an example env file in the root directory):
```bash
docker-compose build
docker-compose --env-file ./.env up
docker-compose --env-file ./.env up -d # if you want to run the containers in the background
```

## Functional tests and Manual tests

To run the functional tests written in python, or any manual tests, we have to set up the DB, in order to do that, execute the following goose command:
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
docker-compose run go_builder sh bin/go_test.sh
```
This will generate a coverage report at `./app/src/report/coverage.html`.

## Go linting

We have some different commands to run Go linters, run the following commands:
```bash
docker-compose run go_builder sh bin/go_fmt.sh
docker-compose run go_builder sh bin/go_vet.sh
docker-compose run go_builder sh bin/staticcheck.sh
```

## Python linting

To run pylint use the following command:
```bash
docker-compose run python_tests sh bin/pylint.sh
```

## Goose and PostgreSQL DB Info

To connect to the DataBase, use the following command (replace ${MICROSERVICE} by the affected microservice):
```bash
docker-compose run postgres_${MICROSERVICE} psql --host=postgres_${MICROSERVICE} --username=${POSTGRES_USERNAME} --dbname=${POSTGRES_DB_NAME}
```

To create a new DB migration use the following command:
```bash
docker-compose run ${MICROSERVICE} sh bin/goose_new_migration.sh ${MIGRATION_NAME}
```

To apply the db_migrations use the following command:
```bash
docker-compose run ${MICROSERVICE} sh bin/goose_apply_migrations.sh ${POSTGRES_CONTAINER} ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD} ${POSTGRES_DB_NAME}
```

To unapply the db_migrations use the following command:
```bash
docker-compose run ${MICROSERVICE} sh bin/goose_downgrade_migration.sh ${POSTGRES_CONTAINER} ${POSTGRES_USERNAME} ${POSTGRES_PASSWORD} ${POSTGRES_DB_NAME}
```

## Adding or updating Go dependencies

To add new go dependencies we just have to use the following commands inside the `app` directory. This can be done through your local environment as long as you have go installed.
```bash
go get ${DEPENDENCY}@${VERSION}
go mod tidy
```

If you are adding an executable go dependency, add it to the `tools.go` file and follow the instructions there (basically the same instructions above).

## Adding or updating Python dependencies

So far I've been adding these manually to the `test/requirements.txt` file. After that, we need to re-build the python container.

## Updating Protos

Update the `app/bin/generate_protos` and then build the project by executing the `docker-compose build` command or execute the following command:
```bash
docker-compose run go_builder sh bin/generate_protos.sh
```

## Call gRPC endpoints manually

Execute the following command:
```bash
docker-compose run go_builder grpcurl -plaintext -d '${REQUEST_BODY}' ${MICROSERVICE}:8080 go_webserver.${MICROSERVICE}.${SERVICE}/${RPC}
```

## Changes and Pull Requests

Every change has to be made via a Pull Request, and CircleCI checks are needed.
Even for the repo administrators.

## Github Releases

To understand how Github releases work for this repository, this documentation should be useful: https://github.com/go-modules-by-example/index/blob/master/009_submodules/README.md.
This section will probably change once a React.js UI is added.

## Dependabot

This repo uses Dependabot to keep its Go dependencies up to date.
The config was created following this blog article: https://github.blog/2020-06-01-keep-all-your-packages-up-to-date-with-dependabot/
And the github docs: https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuring-dependabot-version-updates

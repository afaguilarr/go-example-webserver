# README.md

This project aims to create an example go/golang web-server.
Any feedback is welcome since this is a public project.

It's integrated with docker so just running the following commands should be enough:
```bash
docker-compose build
docker-compose up
```

To run the functional tests written in python just run:
```bash
docker-compose run python_tests pytest
```

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

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

Every change has to be made via a Pull Request, and CircleCI checks are needed.
Even for the repo administrators.

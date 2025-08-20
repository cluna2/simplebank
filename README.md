<h1 align="center"> SimpleBank Project Design Walkthrough And Learnings   </h1>
 This is a simple Golang-based web service to create and manage money transfers between available accounts. 
 
## Simple bank service
This service provides APIs for the following actions:

1. Create and manage bank accounts. Each account is composed of an owner's name, currency type, 
2. Record all balance changes to each account. Every time money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts, wrapped in a transaction.

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/) (or package manager of choice)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

  ```bash
  brew install golang-migrate
  ```

- [DB Docs](https://dbdocs.io/docs)

  ```bash
  npm install -g dbdocs
  dbdocs login
  ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

  ```bash
  npm install -g @dbml/cli
  dbml2sql --version
  ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

  ```bash
  brew install sqlc
  ```

- [Gomock](https://github.com/golang/mock)

  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup infrastructure

#### docker-compose

Run the following ```make``` commands to build/clean dev environment.

- Generate code (SQLC CRUD, mock interfaces, etc.) and run Docker compose file. Note that this automatically starts the server.
```bash
make compose_up
```

- Stop docker containers and remove images
```bash
make compose_down
```

#### Makefile alternative
Alternatively, you can create a dev environment without using ```docker-compose.yaml``` by running the following series of ```make``` commands. 
- Pulls images, creates a Docker network, creates the database, runs migrations, generates code, and run unit tests
  ``` bash
  make build
  ```
  There are targets for each step of the above command. See the ```Makefile``` for details.

- Deletes the db, stops containers, and removes images.
  ``` bash
  make clean
  ```

### Running locally

- Run the server
  ```bash
  make server
  ```
  Make requests to ```localhost:8080``` for HTTP requests and ```localhost:9090``` for GRPC requests.

- Running unit tests
  ```bash
  make test
  ```

## Usage
### 

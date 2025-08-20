<h1 align="center"> SimpleBank  </h1>
 This is a simple Golang-based web service to create and manage money transfers between available accounts. 
 
## Simple bank service
This service provides APIs for the following actions:

1. Create and manage bank accounts. Each account is composed of an owner's name, currency type, and a balance.
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
### Deployed Service
You can make requests to the deployed service. See [api.cris-simplebank.com/swagger](api.cris-simplebank.com/swagger) for Swagger generated API docs. 

The deployed service accepts HTTP (api.cris-simplebank.com/v1) and GRPC (gapi.cris-simplebank.com) protocols. Use any of the following endpoints to make a request:

### Local Service
You can also make requests to your locally deployed service. See localhost:8080/swagger for docs. 

### Examples: 
- Create a user using ```curl```: api.cris-simplebank.com/v1/create_user
```
curl -X 'POST' \
  'https://api.cris-simplebank.com/v1/create_user' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "username": "cris",
  "fullName": "cris luna",
  "email": "cris@email.com",
  "password": "secret"
}'
```

- Using Evans to make a GRPC call to CreateUser/LoginUser on localhost:

```
make evans

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


localhost:9090> package pb

pb@localhost:9090> service Simplebank

pb.Simplebank@localhost:9090> call LoginUser

pb.Simplebank@localhost:9090> call CreateUser
username (TYPE_STRING) => cris
full_name (TYPE_STRING) => cris luna
email (TYPE_STRING) => cris@email.com
password (TYPE_STRING) => secret
{
  "user": {
    "createdAt": "2025-08-20T20:39:31.430788Z",
    "email": "cris@email.com",
    "fullName": "cris luna",
    "passwordChangedAt": "0001-01-01T00:00:00Z",
    "username": "cris"
  }
}

pb.Simplebank@localhost:9090> call LoginUser
username (TYPE_STRING) => cris
password (TYPE_STRING) => secret
{
  "accessToken": "v2.local.SI-pDf7EcsTBavGPjL2JhkV8VIzot8NHPBqXh1oM7LeNNk-Q1sXpwVJ_pWKtqnIC3jsw28JqrDm_AOBf-jDl5CdLfgs73Djz-E6McDg28a_Co2meagx51GX1e_dfiahXyGqHTlZvNtdJ3rNwQr22N67fblj2hXFstqHzg6D_XTORD3Cqx8KQJ97E1vSF7qfCqmVTcYW0IEg3hgSmKopmFT5Es3Y0pvsDOSCcNZj6TFJNmpOqBZDt6Loc2iJUP7Vv0QgE-GffJotJsREj.bnVsbA",
  "accessTokenExpiresAt": "2025-08-20T20:55:07.955485941Z",
  "refreshToken": "v2.local.QRYtzSOevIHwYLOMYE9vOFZLCfpXVqzoMHY91RiW53IHJTAroUncIyghzOjgvzq1H7WDnYEOFUMYg37AYvVhczyIvBztlzMwwAxRwtBhlc8ocHU8q2HNxoCzHHBVFT-zLeU7JzAoiVVD9P8zMq4hpkdM9rRGf-t5XSheVi94pEH-IEcMQLIMyBwQxSfU5OzFyx760CeOL4ydfBV8ftM8EMWTiJDBOg-axcKbkvAUYqt1cpRbwwnYQwkmJQ8a_EuuTOsOZvsCTP87Hbxl.bnVsbA",
  "refreshTokenExpiresAt": "2025-08-21T20:40:07.955575855Z",
  "sessionId": "9f9b5018-3add-4c13-9196-51507a952935",
  "user": {
    "createdAt": "2025-08-20T20:39:31.430788Z",
    "email": "cris@email.com",
    "fullName": "cris luna",
    "passwordChangedAt": "0001-01-01T00:00:00Z",
    "username": "cris"
  }
}
```

# Golang CMS Backend Application
This application is made as a backend application of Go CMS

## Requirement
to be able to run this application locally, you will need the following:
- go version 1.23
- open api code generator ([Installation Docs](https://github.com/oapi-codegen/oapi-codegen))
- docker with docker compose installed ([Installation Docs](https://docs.docker.com/engine/install/))
- GNU make ([Docs](https://www.gnu.org/software/make/))

## DB Architecture
![DB Architecture](ddl.png)

## Default User
This project is included with default user located at `seed.sql`

```
username: admin
password: somepassword
```
**NOTES: please keep in mind that this is a temporary user meant to be deleted before production deployment. please create another user with proper role and ACL then delete the default user**

## Local Development
```
$ docker compose up --build
```

**note: database container will only run migration once. for any schema modification, please take down composed app and remove volume**
```
$ docker compose down --volumes
```

### Extras
if preferred to run local development natively, just use the postgre container and run the app separately. Copy .env.example to .env and fill the value respectfully

```
$ go mod download
...
$ make run
```

## Available Make Command
- `make generate`: generate openapi code generator for echo server and request/response datatypes
- `make run`: run application for local development
- `make generate_mock`: generate mock with mockery
- `make run_test_unit`: run complete unit test

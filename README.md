# Aroundhome Partner API

This API gives you information about partners that are specified in a service.

As of now the serviceName should be `flooring` but other services can be added later.

The location of all our partners is somewhere within or in the proximity to berlin.
When making requests search for a lat/long within berlin to find partners closeby.

Find detailed information on how to use this api in the [openapi spec](https://github.com/ramonamaltan/aroundhome/blob/main/openapi.yaml).
Make it look pretty pasting it into https://editor.swagger.io/.

### Run
Clone project and run `go run cmd/main.go`.
Go to `localhost:8080` and start making requests.

### Dummy Data
On every start of the project a new set of `100 partners` are being inserted to the db.

### Database Setup
You need to set up a local postgres database for the service to work.
If not already install `postgres` e.g. [via homebrew](https://wiki.postgresql.org/wiki/Homebrew)

#### First create the user `pguser` with password `localtest`

#### Create the database
```
create database aroundhome;
```

#### From Project Folder run migrations
```
migrate -database 'postgres://pguser:localtest@localhost:5432/aroundhome?sslmode=disable' -path internal/db/migrations up
```
I use the following package for migrations: https://github.com/golang-migrate/migrate

#### Go to aroundhome db
```
\q
psql aroundhome;
```

#### Check DB entries
```
SELECT * FROM partners;
```

### Notes and Assumptions
- For creating queries and db models I use [sqlc](https://docs.sqlc.dev/en/stable/)
- I made the Assumption that the Square meters of the floor and Phone number (for the partner to contact the customer)
are not relevant for GET requests but for creating customers data e.g. POST `/customer`
- If I had more time I would try to make the nearby condition part of the SQL Query
as now it's not perfect for looping through all partners to check distance (will be relevant for bigger dataset)

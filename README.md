# Articles Service

## Getting Started

Run `make` to get a list of commands.

- Run `make mysql` to download the Docker Image of MySQL:8.0, and execute create-table scripts on startup.
- Run `go run main.go` to start the server on port 8080.
- Run `make test` to run tests.

## Architecture

### Database

Articles Data are stored in SQL, to have transactional commits instead of querying multiple tables in NoSQL.

SQL data was chosen due to Joins at the Database Level (less application logic), and aggregation functions.

NoSQL Databases should be considered for large-scale, or no-need of consistent article <-> tags.


### Routes

Service serves requests through a REST API. Most of the controllers and routing use [net/http](https://golang.org/pkg/net/http/), for parsing path parameters, [httprouter](https://github.com/julienschmidt/httprouter) is used on top of the `net/http` API.


### JSON Parsing

Article data is served through the JSON format. [encoding/json](https://golang.org/pkg/encoding/json/) is used.

- [gojay](https://github.com/francoispqt/gojay) could've been an alternative to generate performant encode/decode struct. Simplicity was chosen instead.


## Assumptions

### Design

- Possible deleting articles feature. By using cascading deletes, deletes on article automatically delete linked tags.
- Possible updating articles features. By using transactions, queries to get articles+tags are consistent.
- Simple system. For a larger-system, usage of more advanced frameworks or libraries may be more useful, but a simpler codebase was prioritized for now using Go's default libraries.
- Assumed a single database will be forever used in the codebase.

### Features

- Assumed last 10 articles entered for a day is the last 10 created within the system for that date.
- Assumed ArticleIDs are Ints.

### Testing

- Testing is minimal. Effort was focused on small Unit Tests to save time.



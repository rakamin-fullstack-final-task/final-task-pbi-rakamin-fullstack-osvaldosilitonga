# Final Task - PBI Rakamin - Fullstack

## DB Migrations
```bash
$ migrate -database "{DB_STRING}" -path database/migrations up
```
Replace {DB_STRING} with your database connection string. This string typically includes information such as the database type (e.g., MySQL, PostgreSQL, SQLite), hostname, port, username, password, and database name.
<br />
Example :
```bash
$ migrate -database "postgres://username:password@(hostname:port)/database_name?sslmode=disable" -path database/migrations up
```
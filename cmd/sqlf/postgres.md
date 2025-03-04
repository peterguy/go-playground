Launch a PostgreSQL Docker container in one terminal.
Leave the terminal running: in it you'll see log messages showing the queries.
(note that it binds to port 5433, not the default PG port)
```sh
docker run --rm --name go_playground_sqlf_postgres \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5433:5432 \
  postgres
```

In another terminal, connect to the PG instance.
```sh
PGPASSWORD=mypassword psql -h localhost -p 5433 -U myuser -d mydatabase
```

Once connected, run these SQL commands to create the table and fill it with test data.
You can stay connected and run other commands, like `SELECT * FROM users;`, or inserting more records.
```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
INSERT INTO users (name) VALUES
('Alice'),
('Bob'),
('Charlie O''Brien'),
('Diana'),
('Eve Johnson'),
('Franklin'),
('Günther'),
('Hélène'),
('Иван'),
('张伟'),
('100% Real'),
('100 more things Real'),
('test_case'),
('testXcase'),
('C:\Users\John'),
('C:\_sers\%John'),
('C:\Users\LittleJohn'),
('C:\Xsers\John');
```

In a third terminal, run the Go program and observe the output
```
go run ./cmd/sqlf
```
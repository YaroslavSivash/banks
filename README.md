# banks

# 1: Install PostgreSQL:
```
sudo apt-get update
sudo apt-get install postgresql
```
##Check installation:
```
sudo -u postgres psql
```
##Create app database:
```
CREATE DATABASE telegram_bot;
```
Or exit from postgres command line and create db with linux command:
```
\q
sudo -u postgres createdb bank
```
Enter to the database:
```
sudo -u postgres psql -d bank
```
Create user of the db from postgres command line:
```
CREATE USER admin WITH PASSWORD '123';
\q
```
Go to the postgres command line and do some customization:
```
sudo -u postgres psql

ALTER ROLE admin SET client_encoding TO 'utf8';

ALTER ROLE admin SET default_transaction_isolation TO 'read committed';

ALTER ROLE admin SET timezone TO 'UTC';

GRANT ALL PRIVILEGES ON DATABASE bank TO admin;

ALTER USER admin CREATEDB;

\q
```

## Customization config.yml:
```
port: "9001"

port_db: "5432"
host: "127.0.0.1"
db_password: ""
db_name: ""
username: ""
timezone: "Europe/Kiev"
```
## Customization dbconf.yml:
```
development:
    driver: postgres
    open: user= dbname= password= host=127.0.0.1  port=5432 sslmode=disable
```
##Install Goose migration tool:
https://bitbucket.org/liamstask/goose/src/master/
```
go get bitbucket.org/liamstask/goose/cmd/goose
```
Inside the project root directory migrate the database:
```
goose up
```
##To delete tables, repeat below command for each table:
```
goose down
```

And you must install Postman, to use request with json
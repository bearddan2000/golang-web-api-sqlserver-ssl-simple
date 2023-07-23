# golang-web-api-sqlserver-ssl-simple

## Description
This returns a data dump of the table `pop`
as a JSON encoded string. An ORM is used to
create and seed the database; even though scripts
are included, `db/sql`, they are for reference only.

Sql server uses self-signed ssl.

## Tech stack
- bash
- golang 1.13

## Docker stack
- alpine:edge
- ubuntu:latest
- mcr.microsoft.com/mssql/server:2017-latest-ubuntu

## Build notes
The service takes about 1 min before callable.

## To run
`sudo ./install.sh -u`
Available at http://localhost:8080

## To stop
`sudo ./install.sh -d`

## To see help
`sudo ./install.sh -h`

## Credits
https://github.com/microsoft/sql-server-samples/blob/master/samples/tutorials/go/orm.go

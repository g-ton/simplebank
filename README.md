Steps to perform a migration inside the project folder

1. in the root of project's folder type:
make postgres (to create the postgres container)

2. make createdb (to create the DB in postgres)

3. make migrateup (to run the UP migrations inside of db/migration)

4. Run unit test:
go test -race $(go list ./... | grep compare) -v -coverprofile coverage.out
go tool cover -html=./coverage.out // to show through browser
Steps to perform a migration inside the project folder

1. in the root of project's folder type:
make postgres (to create the postgres container)

2. make createdb (to create the DB in postgres)

3. make migrateup (to run the UP migrations inside of db/migration)

4. Run unit test:
go test -race $(go list ./... | grep compare) -v -coverprofile coverage.out
go tool cover -html=./coverage.out // to show through browser

+++++---/// SOME DATA ABOUT TO TEST MANUALLY THE SYSTEM ///---+++++

USER 1:
{
    "username": "Tetita2",
    "password": "secret"
}

USER 2:
{
    "username": "Cachito",
    "password": "secret"
}

{
    "transfer": {
        "id": 59,
        "from_account_id": 151,
        "to_account_id": 150,
        "amount": 60,
        "created_at": "2024-04-26T00:31:18.89298Z"
    },
    "from_account": {
        "id": 151,
        "owner": "Cachito",
        "balance": -60,
        "currency": "EUR",
        "created_at": "2024-04-26T00:30:21.650667Z"
    },
    "to_account": {
        "id": 150,
        "owner": "Tetita2",
        "balance": 60,
        "currency": "EUR",
        "created_at": "2024-04-26T00:27:30.315097Z"
    },
    "from_entry": {
        "id": 89,
        "account_id": 151,
        "amount": -60,
        "created_at": "2024-04-26T00:31:18.89298Z"
    },
    "to_entry": {
        "id": 90,
        "account_id": 150,
        "amount": 60,
        "created_at": "2024-04-26T00:31:18.89298Z"
    }
}
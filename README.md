# go-by-example-learning

## Goal of this project is to create a CRUD app for Animal shop.

TODO:
- [X] Animal structure
- [ ] **IS it possible?** Inheritance structure with Dog and Cat structure
- [X] Http server
- [X] JSON handling
- [X] CRUD
- [X] Simple error handling
- [X] Split into Controller, Service and Repository layer
- [X] MySql handling
- [X] Logging user proper logger
- [X] Move JSON/REST responsibilities to Controller; split main.go and controller
- [X] setup some configuration file
- [ ] **WIP** create CLI tool to run service and tests
- [ ] simple unit testing
- [ ] API versioning (look gin documentation)
- [ ] Authorization

### Local setup:

Run commands: 
`msql` -> `use animals;` -> `source db_scripts/init.sql` to initilize DB

And finially run: `go run .`
Or use CLI tool after setup below.

### Testing with curl:

| Command | Result |
| --- | --- |
| curl -X GET http://localhost:8080/animal | get list of animals |
| curl -X GET http://localhost:8080/animal/<id> | get specific animal |
| curl -X DELETE http://localhost:8080/animal/<id> | delete animal from list |
| curl -X POST -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "00000000-0000-0000-0000-000000000001"}' http://localhost:8080/animal | create new animal |
| curl -X PUT -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "<existing ID>"}' http://localhost:8080/animal | update existing animal |


### Setup CLI tool

`serv-runner` file is already present in main directory ready to use.

However, if you are willing to change something feel free to change files in `cli/` folder.
To build it and use, simply run:

`go build . ; mv ./serv-runner ..`

### Run with CLI (TBD) **WIP**

Setup DB credentials with command:

`./serv-runner db-setup` (**WIP**)

Setup a server locally with:

`./serv-runner local` (**WIP**)

Run curl tests above with:

`./serv-runner test` (**TBD**)

# go-by-example-learning

## Goal of this project is to create a CRUD app for Animal shop.

TODO:
- [X] Animal structure
- [ ] Inheritance structure with Dog and Cat structure
- [X] Http server
- [X] JSON handling
- [X] CRUD
- [X] Simple error handling
- [ ] Split into Controller, Service and Repository layer
- [ ] SQLite handling
- [ ] Authorization

Testing with curl:
| Command | Result |
| curl -X GET http://localhost:8080/animal | get list of animals |
| curl -X GET http://localhost:8080/animal/<id> | get specific animal |
| curl -X DELETE http://localhost:8080/animal/<id> | delete animal from list |
| curl -X POST -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "00000000-0000-0000-0000-000000000001"}' http://localhost:8080/animal | create new animal |
| curl -X PUT -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "<existing ID>"}' http://localhost:8080/animal | update existing animal |

To fix:
- [ ] Fix PUT endpoint

# go-by-example-learning

## Goal of this project is to create a CRUD app for Animal shop.

TODO:
- [X] Animal, Dog, Cat structure
- [X] Http server
- [X] JSON handling
- [X] CRUD
- [ ] Error handling
- [ ] SQLite handling
- [ ] Authorization

Testing with curl:
curl -X GET http://localhost:8080/animal             - get list of animals
curl -X GET http://localhost:8080/animal?id=<id>     - get specific animal
curl -X DELETE http://localhost:8080/animal?id=<id>  - delete animal from list (WIP)
curl -X POST -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "00000000-0000-0000-0000-000000000001"}' http://localhost:8080/animal - create new animal
curl -X PUT -H "Content-Type: application/json" --data '{"name": "Test", "age":0, "id": "<existing ID"}' http://localhost:8080/animal - update existing animal

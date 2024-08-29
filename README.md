# go-by-example-learning

## Goal of this project is to create a CRUD app for Animal shop.

TODO:
- [X] Animal, Dog, Cat structure
- [X] Http server
- [X] JSON handling
- [ ] CRUD
- [ ] Error handling
- [ ] SQLite handling
- [ ] Authorization

Testing with curl:
curl -X GET http://localhost:8080/animal             - get list of animals
curl -X GET http://localhost:8080/animal?id=<id>     - get specific animal
curl -X DELETE http://localhost:8080/animal?id=<id>  - delete animal from list (WIP)

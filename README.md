# mvc - REST API on mvc pattern
1. get users at localhost:8000/users/1
2. add benchmark test example to utils
2. add mocking connecting to db via interface to service
3. play with gin-gonic (add xml/json response)


# src - play with github api

# jwtService
jwt_creator - generate jwt-token
api - client for jwt creator
curl localhost:8000 - get token
curl localhost:9000 --header 'Token: <insert-copyied-token>' - get secret information 
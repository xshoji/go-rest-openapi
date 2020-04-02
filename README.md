# go-rest-openapi

Implementation of rest api by go based on OpenAPI Specification.

# Requirement

```
$ go version
go version go1.13.7 darwin/amd64
$ openapi-generator version
4.3.0
```

# Generate code

```
openapi-generator generate \
  -i spec.yaml \
  --git-user-id="xshoji" \
  --git-repo-id "go-rest-openapi/usersapi" \
  --package-name usersapi \
  -o usersapi \
  -g go-server
```

# Implementation

```
// entrypoint
./usersapi/cmd/main/main.go

// buissiness logic
./usersapi/service/usersapi_service.go
```

# Run

```
$ cd usersapi
$ go run ./cmd/main/main.go
2020/04/02 03:51:51 Server started
```

# Call api

## Create

```
curl \
  -vvv \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "john-'$(openssl rand -hex 15 | fold -w 15 |head -n 1)'",
  "nickname": "nick-'$(openssl rand -hex 15 | fold -w 15 |head -n 1)'",
  "age": '$((RANDOM%100+1))',
  "birth": {
    "datetime": "2020-03-31T17:09:09.916Z",
    "weight": '$((RANDOM%100+1))',
    "hospital": "string"
  },
  "addresss": {
    "country": 1,
    "zipCode": 0,
    "state": "string",
    "city": "string"
  },
  "married": true,
  "email": "user@example.com",
  "website": "string",
  "sports": [
    2
  ]
}' \
'http://localhost:8080/v1/users'
```

# Get

```
# get by id
curl http://localhost:8080/v1/users/763416

# get ll
curl http://localhost:8080/v1/users
```

# Very simple JSON CRUD | go_crud

> Beginners JSON CRUD API to create, read, update and delete users. JSON file used as DB

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./go_crud
```

## Endpoints

### Get All Users
``` bash
GET /v1/users
```
### Get Single User
``` bash
GET v1/users/{id}
```

### Delete User
``` bash
DELETE v1/users/{id}
```

### Create User
``` bash
POST v1/users

# Request sample
#    "username": "amulya1",
#    "email": "amulya@gmail.com",
#    "age": 26,
#    "contact": {
#        "mobile": "9559974779",
#        "address": "Malad West, Mumbai, Maharashtra"
#    }
```

### Update User
``` bash
PUT v1/users/{id}

# Request sample
#    "id": "1",
#    "username": "amulya1",
#    "email": "amulya@gmail.com",
#    "age": 26,
#    "contact": {
#        "mobile": "9559974779",
#        "address": "Malad West, Mumbai, Maharashtra"
#    }

```


```

## Project Details

### Author

Amulya Kashyap
(github.com/amulyakashyap09)

### Version

1.0.0

### License

This project is licensed under the MIT License

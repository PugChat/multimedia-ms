# Golang API Server Sample

## Abstract
- Many to Many Web API Server Sample
- Web Framework : Gin (https://github.com/gin-gonic/gin)
- ORM : GORM (https://github.com/jinzhu/gorm)

## API docments

### Entry

#### Response example
```json
{
  "id": 1,
  "userid": 1,
  "name": "entr23",
  "link": "www.alg232o.com",
  "chatid": "12223"
}
```

#### Get all entries
`GET /entries`
[
    {
        "id": 1,
        "userid": 1,
        "name": "entr23",
        "link": "www.alg232o.com",
        "chatid": "12223",
        "tags": null
    },
    {
        "id": 2,
        "userid": 1,
        "name": "entr23",
        "link": "www.alg232o.com",
        "chatid": "12223",
        "tags": null
    }
]
#### Get an entry by id
`GET /entries/:id`

#### Create an entry
`POST /entries`

If you want to create relations with tags, you don't have to include tag id.

Request example
```json
{
  "id": 1,
  "userid": 1,
  "name": "entr23",
  "link": "www.alg232o.com",
  "chatid": "12223"
}
```

#### Update an entry
`PUT /entries/:id`

#### Delete an entry
`DELETE /entries/:id`

#### Response example
```json
{
  "id": 1,
  "userid": 1,
  "name": "entr23",
  "link": "www.alg232o.com",
  "chatid": "12223"
}



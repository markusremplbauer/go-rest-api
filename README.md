# REST-API

This repository contains a REST-API implementation written in [Go](https://github.com/golang) using the [gorilla/mux](https://github.com/gorilla/mux) package.

## Usage

Run `docker-compose up` to run the REST-API on port `8080` and a PostgreSQL database on port `5432`. Run `docker-compose up --build` to fetch changes.

### GET

`GET` is used to request data from the server. The max products limit is set to `10`.

`http://localhost:8080/products`
`http://localhost:8080/products/1`
`http://localhost:8080/products?offset=5&limit=5`

### POST

`POST` is used to send data to a server to create a product.

`localhost:8080/products`

Request body:

```json
{
    "name":"new product",
    "price": 4.56
}
```

### PUT

`PUT` is used to send data to a server to update a product.

`localhost:8080/products/1`

Request body:

```json
{
    "name":"updated product"
}   
```

### DELETE

`DELETE` is used to delete data form the server and remove a product.

`localhost:8080/products/1`
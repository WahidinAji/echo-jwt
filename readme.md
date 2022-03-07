Run on Linux
```bash
bash sample-config.sh
```
`host` : `http://localhost:9000`
# Products
`notes => this is just data dummies`
## `Header`
```
Accept : application/json;charset=UTF-8;
```
### Get All
* url with `NO` jwt-token
```http request
{host}/api/products
```

* response
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": "73e12106-207e-4693-9c0d-3147d6ab606a",
      "name": "Iphone 12 Pro Max",
      "stock": 51,
      "price": 499.91
    },
    {
      "id": "44c22cb3-ff6c-4043-8c79-8a5506ce11e9",
      "name": "Macbook M1 Pro Max",
      "stock": 52,
      "price": 799.92
    },
    {
      "id": "d90f8110-039a-47f4-a164-37d807f77ab5",
      "name": "Ipad Pro Max",
      "stock": 53,
      "price": 450.93
    }
  ]
}
```
### Get By Id
* url with `NO` jwt-token
```http request
{host}/api/products/{id}
```
* response
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "44c22cb3-ff6c-4043-8c79-8a5506ce11e9",
        "name": "Macbook M1 Pro Max",
        "stock": 52,
        "price": 799.92
    }
}
```


# Orders
`notes => this is just data dummies`
## `Header`
```
Accept : application/json;charset=UTF-8;
```
## `WITH Authorization`
```
Bearer Token : <jwt-token>
```
### Get All
* url with jwt-token
```http request
{host}/api/orders
```
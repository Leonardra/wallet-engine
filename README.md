# Wallet-Engine [![Generic badge](https://img.shields.io/badge/Language-Golang-<COLOR>.svg)](https://shields.io/)
Wallet Engine is a simple REST Application built with the Gin-Gonic framework. It allows users to be able to do the following functionalities:

* Create Wallet
* Debit Wallet
* Credit Wallet
* Activate and Deactivate Wallet

﻿# Prerequisites
***
* MongoDb Compass
* Postman
* Go
***
﻿# Installation
```shell
$ git clone https://github.com/Leonardra/wallet-engine
$ cd wallet-engine/
$ go get -v -d
```

***


﻿# Tools
* Testify: assertion library for unit testing
* Gin-Gonic : Web service framework.
* Mongo Db : Database

﻿# Rest API Usage
***

﻿## Create Wallet
> ﻿#### P0ST http://localhost:8080/api/v1/wallet/

﻿##### Parameter
 ```json
  {
  "firstName":"John",
  "lastName":"Doe"
 }
```

﻿## Response
﻿#### 201 Created on successful request

```json
  {
 "status": 201,
    "message": "Success",
    "timestamp": "2022-04-08T12:15:33.7814862+01:00",
    "data": {
        "wallet": {
            "id": "625019556d85e3d56c61561d",
            "firstName": "John",
            "lastName": "Doe",
            "dateCreated": "2022-04-08T11:15:33.778Z",
            "balance": 0,
            "accountNumber": "6196211200",
            "activationStatus": true
        }
    }
 }

```

﻿## Credit Wallet
> ﻿#### P0ST http://localhost:8080/api/v1/wallet/{walletId}/credit

﻿##### Parameter
 ```json
  {
  "amount":200000.00
 }
```

﻿## Response
﻿#### 200 OK on successful request

```json
  {
      "status": 200,
      "message": "Success",
      "timestamp": "2022-04-08T12:41:38.0540302+01:00",
      "data": {
        "wallet": {
          "id": "625019556d85e3d56c61561d",
          "firstName": "John",
          "lastName": "Doe",
          "dateCreated": "2022-04-08T11:15:33.778Z",
          "balance": 200000,
          "accountNumber": "6196211200",
          "activationStatus": true
        }
      }
 }
```

﻿## Debit Wallet
> ﻿#### P0ST http://localhost:8080/api/v1/wallet/{walletId}/debit

﻿##### Parameter
 ```json
  {
  "amount":100000.00
 }
```

﻿## Response
﻿#### 200 OK on successful request

```json
  {
      "status": 200,
      "message": "Success",
      "timestamp": "2022-04-08T12:41:38.0540302+01:00",
      "data": {
        "wallet": {
          "id": "625019556d85e3d56c61561d",
          "firstName": "John",
          "lastName": "Doe",
          "dateCreated": "2022-04-08T11:15:33.778Z",
          "balance": 100000,
          "accountNumber": "6196211200",
          "activationStatus": true
        }
      }
 }
```

﻿## Activate/Deactivate Wallet
> ﻿#### P0ST http://localhost:8080/api/v1/wallet/{walletId}/active/


﻿## Key
* Activate: true
* Deactivate: false

﻿#### Parameter
 ```json
  {
  "active":false
 }
```

﻿## Response
﻿#### 200 OK on successful request

```json
  {
  "status": 200,
  "message": "Success",
  "timestamp": "2022-04-08T12:41:26.0205348+01:00",
  "data": {
    "wallet": {
      "id": "62529c934305ee8dbb82f32b",
      "firstName": "John",
      "lastName": "Doe",
      "dateCreated": "2022-04-10T09:00:03.898Z",
      "balance": 100000,
      "accountNumber": "6899074097",
      "activationStatus": false
    }
  }
 }
```


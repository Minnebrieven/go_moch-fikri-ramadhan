# Electronic Item Management API Documentation
This API is created for an electric store to manage electronic item data. the item it self consist of id, name, description, stock, price and item category
## Authentication
To able to use this API, JWT Token is required which can be placed at request's header with `Authorization` as a Key and JWT Token as its value `Bearer {{Token}}`.

In order to get JWT Token user must be registered first via endpoint `/register`. After user registering token can be taken via endpoint `/login`.

**Register/Create New User**
----
  Create new user data.

* **URL**

  /register

* **Body**

  * **Content : JSON**
  ```json
  {
    "name": "fikri",
    "email" :"fikri1711@gmail.com",
    "password":"123"
  }
  ```
  
* **Method:**

  `POST`
  
* **Url Params**

  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 201 <br />
    **Content : JSON** 
    ```json
    {
      "metadata": {
          "status_code": 201,
          "message": "success create user"
      },
      "data": {
          "id": 2,
          "name": "fikri",
          "email": "fikri1711@gmail.com",
          "password": "********",
          "metadata": {
              "DeletedAt": null,
              "CreatedAt": "2023-05-24T23:22:19.285+07:00",
              "UpdatedAt": "2023-05-24T23:22:19.285+07:00"
          }
       }
    }
    ```
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 400,
          "message": "name can't be blank"
      },
      "data": null
    }
    ```

  OR

  * **Code:** 409 CONFLICT <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 409,
          "message": "duplicate user found"
      },
      "data": null
    }
    ```

* **Sample Call:**

  ```cURL
   curl --location 'localhost:8000/register' \
   --header 'Content-Type: application/json' \
   --data-raw '{
      "name": "fikri",
      "email" :"fikri1711@gmail.com",
      "password":"123"
    }'
  ```

**Login User**
----
  Login using user data to retrieve jwt token.

* **URL**

  /login

* **Body**

  * **Content : JSON**
  ```json
  {
    "email" :"fikri1711@gmail.com",
    "password":"123"
  }
  ```
  
* **Method:**

  `POST`
  
* **Url Params**

  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content : JSON** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "success login"
      },
      "data": {
          "id": 2,
          "name": "fikri",
          "email": "fikri1711@gmail.com",
          "password": "********",
          "metadata": {
              "DeletedAt": null,
              "CreatedAt": "2023-05-24T23:22:19.285+07:00",
              "UpdatedAt": "2023-05-24T23:22:19.285+07:00"
          }
      },
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZpa3JpMTcxMUBnbWFpbC5jb20iLCJleHAiOjE2ODQ5NDk4NjQsInVzZXJJRCI6Mn0.9c2DMhqTABgPQ28MYRQXhPn9q4KsRwyLS36fPCVsnBU"
    }
    ```
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 401,
          "message": "login failed: wrong email"
      },
      "data": null,
      "token": ""
    }
    ```

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 401,
          "message": "login failed: wrong password"
      },
      "data": null,
      "token": ""
    }
    ```

* **Sample Call:**

  ```cURL
   curl --location 'localhost:8000/login' \
   --header 'Content-Type: application/json' \
   --data-raw '{
      "email":"fikri1711@gmail.com",
      "password":"fikri1711"
   }'
  ```

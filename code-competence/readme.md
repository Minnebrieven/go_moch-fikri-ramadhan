# Electronic Item Management API Documentation
This API is created for an electric store to manage electronic item data. the item it self consist of id, name, description, stock, price and item category
## Authentication
To able to use this API, JWT Token is required which can be placed at request's header with `Authorization` as a Key and JWT Token as its value `Bearer {{Token}}`.

In order to get JWT Token user must be registered first via endpoint `/register`. After user registering token can be taken via endpoint `/login`.

----

**Register/Create New User**
----
  Create new user data.

* **URL**

  /register

* **Header**

  * **Content-Type : application/json**

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
  
* **Screenshot:**
  ![success_register](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/5aaaa6b6-a09a-4d07-96fc-a7e1086e536c)
----

**Login User**
----
  Login using user data to retrieve jwt token.

* **URL**

  /login

* **Header**

  * **Content-Type : application/json**

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
  
* **Screenshot:**
  ![success_login](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/48128bbb-c723-4a2f-ae21-6a1ece3e745b)
----

**Get All Items**
----
  retrieve all items data.

* **URL**

  /items?keyword={item_name}&page={page_number}

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  None
  
* **Method:**

  `GET`
  
* **Url Params**

  None

* **Data Params**

  `keyword=[string]`
  this parameter used for retrieve items data based on item's name

  **Required:**
  
  `page=[integer]`
  this parameter used for change page

* **Success Response:**

  * **Code:** 200 <br />
    **Content : JSON** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "success"
      },
      "pagination": {
          "page": 0,
          "data_shown": 10,
          "total_data": 4
      },
      "data": [
          {
              "id": "246d5d65-7e98-47e6-939f-a09d6d7026ec",
              "name": "Radio",
              "category": {
                  "id": 1,
                  "name": "Analog",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "Radio analog klasik",
              "stock": 10,
              "price": 200000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          },
          {
              "id": "ca80c1de-9fe3-4db2-aeec-fc0df7d6c768",
              "name": "Laptop",
              "category": {
                  "id": 2,
                  "name": "Digital",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "televisi digital",
              "stock": 6,
              "price": 2000000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          },
          {
              "id": "cad29a6a-859e-4bf8-8ba1-634fd14a0a7f",
              "name": "Speaker",
              "category": {
                  "id": 1,
                  "name": "Analog",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "pengeras suara",
              "stock": 5,
              "price": 500000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          },
          {
              "id": "f0024bb2-96d1-46a2-9c21-aee52acb5bca",
              "name": "Television",
              "category": {
                  "id": 2,
                  "name": "Digital",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "televisi digital",
              "stock": 3,
              "price": 1000000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          }
      ]
    }
    ```
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "message": "invalid or expired jwt"
    }
    ```
  OR
  * **Code:** 400 BAD REQUEST <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 400,
          "message": "invalid page parameter: pagestring"
      },
      "pagination": {
          "page": 0,
          "data_shown": 0,
          "total_data": 0
      },
      "data": null
    }
    ```
  

* **Sample Call:**

  ```cURL
   curl --location 'localhost:8000/items?keyword={item_name}&page={page_number}' \
   --header 'Authorization: Bearer {Token}'
  ```

* **Screenshot:**
  ![success_getall](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/f9479e43-0775-47c4-a904-75ec5f60cbb7)
  ![success_getbyname](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/9bc6a043-b17f-4fce-9396-3dc26d2e69da)
----


**Get All Items By Category (Category ID)**
----
  retrieve all items data by category id.

* **URL**

  /items/category/:category_id?page={page_number}

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  None
  
* **Method:**

  `GET`
  
* **Url Params**

  `category_id=[integer]`

* **Data Params**

  **Required:**
  
  `page=[integer]`
  this parameter used for change page

* **Success Response:**

  * **Code:** 200 <br />
    **Content : JSON** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "success"
      },
      "pagination": {
          "page": 0,
          "data_shown": 10,
          "total_data": 2
      },
      "data": [
          {
              "id": "ca80c1de-9fe3-4db2-aeec-fc0df7d6c768",
              "name": "Laptop",
              "category": {
                  "id": 2,
                  "name": "Digital",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "televisi digital",
              "stock": 6,
              "price": 2000000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          },
          {
              "id": "f0024bb2-96d1-46a2-9c21-aee52acb5bca",
              "name": "Television",
              "category": {
                  "id": 2,
                  "name": "Digital",
                  "metadata": {
                      "DeletedAt": null,
                      "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                      "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                  }
              },
              "description": "televisi digital",
              "stock": 3,
              "price": 1000000,
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.883+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
              }
          }
      ]
    }
    ```
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 400,
          "message": "invalid page parameter: pagestring"
      },
      "pagination": {
          "page": 0,
          "data_shown": 0,
          "total_data": 0
      },
      "data": null
    }
  OR
  * **Code:** 400 BAD REQUEST <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 400,
          "message": "invalid id parameter asd"
      },
      "pagination": {
          "page": 0,
          "data_shown": 0,
          "total_data": 0
      },
      "data": null
    }
    ```
  

* **Sample Call:**

  ```cURL
   curl --location 'localhost:8000/items/category/:category_id?page={page_number}' \
   --header 'Authorization: Bearer {Token}'
  ```

* **Screenshot:**
  ![success_getbycategory](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/d6418ae0-73b0-4da3-8d4c-476ac50d5f48)
----

**Get an Item**
----
  Retrieve single item data.

* **URL**

  /items/:id

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  None

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[uuid]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "success"
      },
      "data": {
          "id": "cad29a6a-859e-4bf8-8ba1-634fd14a0a7f",
          "name": "Speaker",
          "category": {
              "id": 1,
              "name": "Analog",
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
              }
          },
          "description": "pengeras suara",
          "stock": 5,
          "price": 500000,
          "metadata": {
              "DeletedAt": null,
              "CreatedAt": "2023-05-24T23:47:59.883+07:00",
              "UpdatedAt": "2023-05-24T23:47:59.883+07:00"
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
          "message": "invalid id parameter cad29a6a-859e-4bf8-8b91-634fd14a0akf"
      },
      "data": null
    }
    ```

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "message": "invalid or expired jwt"
    }
    ```

* **Sample Call:**

  ```cURL
    curl --location 'localhost:8000/items/cad29a6a-859e-4bf8-8ba1-634fd14a0a7f' \
    --header 'Authorization: Bearer {Token}'
  ```


* **Screenshot:**
  ![success_get](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/b3eab8e0-9b04-4457-b185-2ac355f3adf4)
----

**Create New Item**
----
  Create new item data.

* **URL**

  /items

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  * **Content : JSON**
  ```json
  {
      "name": "Kipas",
      "category": {
          "id": 1
      },
      "description": "ademmm dinginnn brrr...",
      "stock": 5,
      "price": 300000
  }
  ```

* **Method:**

  `POST`
  
* **URL Params**

  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 201,
          "message": "success create item"
      },
      "data": {
          "id": "bd5b4dac-de43-4211-9387-2cd08ac9ec76",
          "name": "Kipas",
          "category": {
              "id": 1,
              "name": "Analog",
              "metadata": {
                  "DeletedAt": null,
                  "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                  "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
              }
          },
          "description": "ademmm dinginnn brrr...",
          "stock": 5,
          "price": 300000,
          "metadata": {
              "DeletedAt": null,
              "CreatedAt": "2023-05-25T01:02:18.851+07:00",
              "UpdatedAt": "2023-05-25T01:02:18.851+07:00"
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

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "message": "invalid or expired jwt"
    }
    ```

* **Sample Call:**

  ```cURL
    curl --location 'localhost:8000/items' \
    --header 'Authorization: Bearer {Token}' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Kipas",
        "category": {
            "id": 1
        },
        "description": "ademmm dinginnn brrr...",
        "stock": 5,
        "price": 300000
    }'
  ```
  

* **Screenshot:**
  ![success_create](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/7bf1688a-158a-4742-b3e5-9c6e245f750f)
----

**Edit Item Data**
----
  Edit existing item data.

* **URL**

  /items/:id

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  * **Content : JSON**
  ```json
  {
      "name": "Komputer",
      "category": {
          "id": 2
      },
      "description": "komputer spek dewa"
  }
  ```

* **Method:**

  `PUT`
  
* **URL Params**

   **Required:**

  `id=[uuid]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json
       {
        "metadata": {
            "status_code": 200,
            "message": "success edit item data"
        },
        "data": {
            "id": "874bda65-e445-4efd-af7e-206e5ed5cf65",
            "name": "Komputer",
            "category": {
                "id": 2,
                "name": "Digital",
                "metadata": {
                    "DeletedAt": null,
                    "CreatedAt": "2023-05-24T23:47:59.66+07:00",
                    "UpdatedAt": "2023-05-24T23:47:59.66+07:00"
                }
            },
            "description": "komputer spek dewa",
            "stock": 5,
            "price": 300000,
            "metadata": {
                "DeletedAt": null,
                "CreatedAt": "2023-05-25T01:08:58.148+07:00",
                "UpdatedAt": "2023-05-25T01:09:52.903+07:00"
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
          "message": "nothing to change"
      },
      "data": null
    }
    ```

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "message": "invalid or expired jwt"
    }
    ```

* **Sample Call:**

  ```cURL
    curl --location --request PUT 'localhost:8000/items/874bda65-e445-4efd-af7e-206e5ed5cf65' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZpa3JpMTcxMUBnbWFpbC5jb20iLCJleHAiOjE2ODQ5NTQ5MjEsInVzZXJJRCI6Mn0.ZyD0pwOZC6DP4wANkk5N_Co286NziNB613Bzc9n6EXU' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Komputer",
        "category": {
            "id": 2
        },
        "description": "komputer spek dewa"
    }'
  ```
  

* **Screenshot:**
  ![success_edit](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/26008a72-10d5-4117-8845-6d4ed9507124)
----

**Delete Item Data**
----
  Delete existing item data.

* **URL**

  /items/:id

* **Header**

  * **Content-Type : application/json**
  * **Authorization : Bearer {Token}**

* **Body**

  None

* **Method:**

  `DELETE`
  
* **URL Params**

   **Required:**

  `id=[uuid]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "success delete item"
      },
      "data": {
          "item_id": "bd5b4dac-de43-4211-9387-2cd08ac9ec76"
      }
    }
    ```
 
* **Error Response:**

  * **Code:** 200 OK <br />
    **Content:** 
    ```json
    {
      "metadata": {
          "status_code": 200,
          "message": "item not found"
      },
      "data": null
    }
    ```

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** 
    ```json
    {
      "message": "invalid or expired jwt"
    }
    ```

* **Sample Call:**

  ```cURL
    curl --location --request DELETE 'localhost:8000/items/874bda65-e445-4efd-af7e-206e5ed5cf65' \
    --header 'Authorization: Bearer {Token}'
  ```
  
* **Screenshot:**
  ![success_delete](https://github.com/Minnebrieven/go_moch-fikri-ramadhan/assets/30454868/98f2828c-f900-42e4-8ba3-c4d846491dd9)
----

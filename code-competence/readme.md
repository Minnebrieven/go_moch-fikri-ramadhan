# Electronic Item Management API Documentation
This API is created for an electric store to manage electronic item data. the item it self consist of id, name, description, stock, price and item category
## Authentication
To able to use this API, JWT Token is required which can be placed at request's header with `Authorization` as a Key and JWT Token as its value `Bearer {{Token}}`.

In order to get JWT Token user must be registered first via endpoint `/register`. After user registering token can be taken via endpoint `/login`.

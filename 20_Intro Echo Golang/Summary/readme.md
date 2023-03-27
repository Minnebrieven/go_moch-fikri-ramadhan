# Summary Echo Golang
## Echo
Echo golang adalah sebuah web framework golang yang memiliki performance tinggi, dapat di perluas dan minimalis.
- snytax untuk menginstall web framework echo
```bash
go get github.com/labstack/echo/v4
```

## Basic Routing & Controller
Berikut sintax yang digunakan untuk membuat sebuah routing/handler di web framework echo
```golang
//membuat instance echo baru
e : = echo.New()

//Route atau handler
e.GET("path", Controller)

untuk memulai server
e.Start(":port")
```

Berikut sintax yang digunakan untuk membuat sebuah controller di web framework echo
```golang
func Controller(c echo.Context) error {
  //lakukan prosedur apapun disini
  
  //akan mengembalikan response code 200/OK dan "data"
  return c.JSON(HTTP.StatusOK, "data")
}
```
## Retrive Data
- URL Param
```golang
//pada url localhost:8000/users/:id
c.Param("id")
```
- Query Param

```golang
//pada url localhost:8000/users?country=indoneisa
c.QueryParam("country")
```
- Form Value
```golang
//pada json tag "nama" - `json:"nama"`
c.FormValue("nama")
```

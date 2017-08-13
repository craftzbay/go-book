#  Хялбар файл сервер

Дараах жишээнд `/tmp/files` хавтасын агуулгыг 8080 портоор нээлттэй болгож байна.

```go
package main
import "net/http"

func main() {
        http.ListenAndServe(":8080",
		 http.FileServer(http.Dir("/tmp/files")))
}
```

Файл сервер рүү хандахын тулд браузер дээр http://localhost:8080 хаягаар орох хэрэгтэй.

Статик эсвэл нэг хуудасны (single page javascript програм) вэб програмуудыг мөн ийм аргаар ажиллуулж болно.

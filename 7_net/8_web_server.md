#  Хялбар веб сервер

HTTP нь TCP -ээс дээд түвшний протокол юм. Энэ протокол нь браузер болон веб серверийн хооронд хэрхэн мэдээлэл солилцох дүрмийг тодорхойлсон байдаг.

Go хэлний `net/http` пакет нь веб сервер үүсгэх боломжоор хангадаг. Жишээ болгон хамгийн ялбар "Hello world!" веб сервер үүсгэе.

Үүний тулд эхлээд `net/http` санг импортлоно, хүсэлт боловсруулах функц үүсгэнэ, дараа нь веб серверийг `http.ListenAndServe("localhost:8080", nil)` дуудалтаар асаана.

Хүсэлт боловсруулах `HelloServer()` функ дотор хүсэлтийн URL замыг `req.URL.Path` байдлаар ялгаж авч байна.

```go
package main

import (
  "fmt"
  "log"
  "net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Сайн уу, "+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
```

Вэб серверийг асаагаад браузер дээрээс ажиллаж байгаа эсэхийг шалгаж болно. Үүний тулд браузер дээр http://localhost:8080/world хаягийг бичиж үзэх хэрэгтэй.

Веб серверийн үндсэн порт нь 80 байдаг. Гэхдээ бидний жишээнд 8080 порт дээр ажиллуулж байна. Тийм учраас http://localhost:8080 гэж бичнэ, хэрэв 80 буюу үндсэн портоор ажиллуулсан бол зүгээр http://localhost гэж бичихэд болно.

TODO: https://http2.golang.org/

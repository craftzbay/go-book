# Өгөгдлийг шахаж дамжуулах

Ихэнхи вэб, файл серверүүд өгөгдлийг ямар нэг хэлбэрээр шахаж илгээдэг. HTML, Javascript, CSS зэрэг текстэн өгөгдлийг шахаж илгээх нь оновчтой байдаг, харин png зэрэг зургуудын хувьд анхнаасаа шахалттай байдаг.

Өгөгдлийг шахаж илгээснээр сүлжээгээр дамжих өгөгдлийг их хэмжээгээр багасгана, мөн сүлжээгээр нэвтрэх хурд нэмэгдэнэ. Шахах үйлдэл нь CPU-д бага зэрэг ачаалал өгдөг, гэхдээ gzip, snappy, lz4 зэрэг сайн алгоритмын хувьд энэ нь бараг нөлөөгүй байдаг.

Gzip нь хамгийн түгээмэл шахалтын алгоритм юм. Apache, Nginx зэрэг ихэнхи вэб серверүүд gzip шахалтыг дэмждэг. Мөн сүүлийн үеийн бүх браузер дэмжинэ.

Хэрэв өөрийн вэб сервер үүсгэж байгаа бол gzip шахалтыг нэмж програмчлах шаардлагатай болно.

Жишээ болгон `GzipResponseWriter` нэртэй төрөл үүсгээд үүнийгээ `ServeHTTP()` функц дотор ашиглая.

```go
import "compress/gzip"

type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
```

gzip writer обект үүсгэх:

```go
// gzip бичигч
w.Header().Set("Content-Encoding", "gzip")
gz := gzip.NewWriter(w)
defer gz.Close()

handler.ServeHTTP(GzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
```

lz4, snappy зэрэг бусад шахалтын алгоритм ашиглах бол `go get` командаар татаж суулгаад ашиглаж болно. Эдгээр алгоритмын хувьд одоохондоо ихэнхи браузер дэмжихгүй байгаа.

Snappy татах:

```sh
$ go get code.google.com/p/snappy-go/snappy
```

lz4 татах:

```sh
$ go get github.com/bkaradzic/go-lz4
```

#  Өгөгдлийг шахах

`archive/zip` пакет нь ZIP форматаар өгөгдлийг шахах, задлах боломжийг олгоно.
Жишээлбэл дараах програмд `readme.txt`, `todo.txt` файлуудыг динамикаар үүсгэж `readme.zip` нэртэй архив файл уруу шахаж байна.

```go
package main

import (
  "archive/zip"
  "bytes"
  "io/ioutil"
  "log"
)

func main() {
	buf := new(bytes.Buffer)

	// шинэ zip буфер үүсгэх
	w := zip.NewWriter(buf)

	// архивт нэмэх файлууд
	var files = []struct {
		Name, Body string
	} {
	 {"readme.txt", "Энэ файл нь заавар мэдээлэл агуулна"},
	 {"todo.txt", "Энэ файл нь хийх зүйлсийн жагсаалтыг агуулна"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// буферыг хаах
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// zip өгөгдлийг файлд бичих
	ioutil.WriteFile("readme.zip", buf.Bytes(), 0777)
}
```

ZIP файлыг задалж унших програмыг доор харуулав:

```go
package main 

import (
  "archive/zip"
  "log"
  "fmt"
  "io"
  "os"
)

func main() {
    r, err := zip.OpenReader("readme.zip")
    if err != nil {
            log.Fatal(err)
    }
    defer r.Close()

    // Архив дахь файлуудаар давтаж агуулгыг хэвлэх
    for _, f := range r.File {
            fmt.Printf("'%s' файлын агуулга:\n", f.Name)
            rc, err := f.Open()
            if err != nil {
               log.Fatal(err)
            }
	    _, err = io.Copy(os.Stdout, rc)
            if err != nil {
                log.Fatal(err)
            }
            println()
            rc.Close()
    }
}
```

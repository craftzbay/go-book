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
	}{
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

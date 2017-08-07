package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileExt = ".html"

var (
	taskQueue []string
	done      chan bool
)

// хавтасаар самнах
func crawl(folder string, ch chan string) {
	fmt.Println("scanning ", folder)

	files, _ := ioutil.ReadDir(folder)

	for _, f := range files {

		if f.IsDir() {
			// хавтас бол цааш самнах
			crawl(folder+"/"+f.Name(), ch)
		} else {
			//  *.html файл эсэхийг шалгах
			if strings.HasSuffix(f.Name(), fileExt) {

				// мөн бол ажлын дараалалд оруулах
				ch <- folder + "/" + f.Name()
			}

		}

	}
	done <- true
}

// файл дотроос холбоос хайх функц
func findLink(ch chan string) {
	for {
		select {
		case fileName := <-ch:
			fmt.Println(fileName)
		}
	}
}

func main() {
	// эхлэх хавтас
	rootFolder := "/Users/ub/Dev/src/lerp/web"

	// ажлын даалгавар (файлын нэрс) солилцох суваг үүсгэх
	taskChannel := make(chan string)

	// самнаж эхлэх (go функц үүсгэж байна)
	done = make(chan bool)
	go crawl(rootFolder, taskChannel)

	// 3 зэрэг findLink функц ажиллуулах
	//for i := 0; i < 3; i++ {
	go findLink(taskChannel)
	//}

	<-done

	fmt.Println("Дууслаа")

}

package main

import ( "archive/zip";"log";"fmt";"io";"os" )

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
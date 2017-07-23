# Тэгш хэмтэй нууцлалт

Өгөгдлийг энкриптлэх хоёр үндсэн арга байдаг. Эхнийх нь өгөгдлийг кодлох болон буцаан задлахад нэгэн ижил түлхүүрийг ашиглах арга юм. Энэ түлхүүр нь кодлож байгаа болон задалж байгаа талуудад хоёуланд нь мэдэгдэж байх ёстой. Хоёр тал түлхүүрээ хэрхэн солилцох нь энд яригдахгүй.

Хайшлэхтэй адил өөр олон энкрипт алгоритм байдаг. Эдгээрийн зарим нь сул болсон, ерөнхийдөө компютер илүү хурдтай, хүчтэй болох тусам алгоритмууд сул болж байдаг. Go нь Blowfish, DES гэх зэрэг хэд хэдэн тэгш хэмт алгоритмууд дэмждэг.

The algorithms are block algorithms. That is they work on blocks of data. If you data is not aligned to the block size, then you will have to pad it with extra blanks at the end.

Each algorith is represented by a Cipher object. This is created by NewCipher in the appropriate package, and takes the symmetric key as parameter.  
Once you have a cipher, you can use it to encrypt and decrypt blocks of data. The blocks have to be 8-bit blocks for Blowfish. A program to illustrate this is

```go
// Blowfish.go
package main

import (
        "bytes"
        "code.google.com/p/go.crypto/blowfish"
        "fmt"
)

func main() {
        key := []byte("my key")
        cipher, err := blowfish.NewCipher(key)
        if err != nil {
                fmt.Println(err.Error())
        }
        src := []byte("hello\n\n\n")
        var enc [512]byte

        cipher.Encrypt(enc[0:], src)

        var decrypt [8]byte
        cipher.Decrypt(decrypt[0:], enc[0:])
        result := bytes.NewBuffer(nil)
        result.Write(decrypt[0:8])
        fmt.Println(string(result.Bytes()))
}
```

Blowfish is not in the Go 1 distribution. Instead it is on the [http://code.google.com/p/](http://code.google.com/p/) site. You have to install it by running "go get" in a directory where you have source that needs to use it.


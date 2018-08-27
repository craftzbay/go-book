#  Нийтийн түлхүүртэй нууцлалт

Нийтийн түлхүүртэй нууцлалд хоёр түлхүүр ашигладаг: нэг түлхүүр нь энкрипт хийхэд, нөгөө нь декрипт хийхэд ашиглагдана. Нийтийн түлхүүртэй нууцлал гэж ярьж заншсан ч түүнд нийтийн болон хувийн гэсэн хоёр түлхүүр хэрэглэгддэг гэдгийг санаарай. Энкриптлэх түлхүүр нь нийтэд нээлттэй байдаг, тийм учраас ямар ч хүн уг түлхүүрээр мессежээ энкриптлэж чадна. Декрипт түлхүүр нь харин нууц байх ёстой, үгүй бол бүх хүн өөрт чинь үлдээсэн мессежийг тайлж чаддаг болно! Нийтийн түлхүүртэй нууцлал нь ассеметрик байдаг, ө.х өөр өөр хэрэглээнд зориулж өөр өөр түлхүүр ашиглана.

Нийтийн түлхүүртэй нууцлал нь айлын үүдэнд байдаг шуудангийн хайрцагтай төстэй юм. Шуудан зөөгч, сонин борлуулагч, сурталчилгааны ажилтан, татварын байцаагч, сахилгагүй хүүхэд гээд хэн ч өөрийн хүссэн зүйлээ шуудангийн хайрцагны жижигхэн завсараар шургуулж болох боловч нэгэнт орсон бол буцаж гарч ирэх боломжгүй болно. Зөвхөн хайрцагны эзэн өөрийн түлхүүрээр арын том тагийг онгойлгож авах боломжтой. Нууцлалын алгоритмд урд талын жижиг завсарыг нийтийн түлхүүр, ар талын том тагийг онгойлгох түлхүүрийг хувийн түлхүүр гэж нэрлэж байгаа юм.

Go хэл дэмждэг нийтийн түлхүүртэй нууцлалын алгоритм олон бий. Эдгээрээс хамгийн түгээмэл нь RSA юм.

Дараах програмд нийтийн болон хувийн RSA түлхүүр үүсгэж байна.

```go
// GenRSAKeys.go
package main

import (
        "crypto/rand"
        "crypto/rsa"
        "crypto/x509"
        "encoding/gob"
        "encoding/pem"
        "fmt"
        "os"
)

func main() {
        reader := rand.Reader
        bitSize := 512
        key, err := rsa.GenerateKey(reader, bitSize)
        checkError(err)

        fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
        fmt.Println("Private key exponent", key.D.String())

        publicKey := key.PublicKey
        fmt.Println("Public key modulus", publicKey.N.String())
        fmt.Println("Public key exponent", publicKey.E)

        saveGobKey("private.key", key)
        saveGobKey("public.key", publicKey)

        savePEMKey("private.pem", key)
}

func saveGobKey(fileName string, key interface{}) {
        outFile, err := os.Create(fileName)
        checkError(err)
        encoder := gob.NewEncoder(outFile)
        err = encoder.Encode(key)
        checkError(err)
        outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

        outFile, err := os.Create(fileName)
        checkError(err)

        var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
                Bytes: x509.MarshalPKCS1PrivateKey(key)}

        pem.Encode(outFile, privateKey)

        outFile.Close()
}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
                os.Exit(1)
        }
}
```

Мөн сертификатыг `gob` аргаар хадгалж байна. Үүнийг дараах програмаар буцаан уншиж болно:

```go
// LoadRSAKeys.go
package main

import (
        "crypto/rsa"
        "encoding/gob"
        "fmt"
        "os"
)

func main() {
        var key rsa.PrivateKey
        loadKey("private.key", &key)

        fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
        fmt.Println("Private key exponent", key.D.String())

        var publicKey rsa.PublicKey
        loadKey("public.key", &publicKey)

        fmt.Println("Public key modulus", publicKey.N.String())
        fmt.Println("Public key exponent", publicKey.E)
}

func loadKey(fileName string, key interface{}) {
        inFile, err := os.Open(fileName)
        checkError(err)
        decoder := gob.NewDecoder(inFile)
        err = decoder.Decode(key)
        checkError(err)
        inFile.Close()
}

// func checkError(err error)
```


TODO: интернэт худалдаанд private, public түлхүүр ашиглах тухай оруулах

Мерчант checkout хуудсан дээрээ нэхэмжлэлээ тамгалж байршуулж болно (hash код үүсгэх). 

Үүний тулд банк, мерчант хоорондоо public, private key тохирох ёстой. Мерчант private key-ээ өөртөө үлдээж нууцлана, public key-ээ банкинд дамжуулна.

Мерчант өөрийн private түлхүүрээр нэхэмжлэлээ тамгалахдаа барааны үнэ, код зэрэг нууцлалтай мэдээллийг оролцуулах (checksum үүсгэх) хэрэгтэй. 

Дараах жишээнд `product`, `price`, `priceCurrency` мэдээллүүдийг оролцуулан 'ef23r12vef' код/дардас үүссэн байна.

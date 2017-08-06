# RPC

RPC \(Remote Procedure Call\) нь сервер талд байгаа функцийг сүжээгээр дамжуулан дуудах арга юм. RPC нь ердийн функц дуудаж ажиллуулж байгаа юм шиг харагддаг. Гэхдээ нэг програм \(клиент\), цаашлаад компютерээс хальж сүлжээний өөр нэг компютер дээр ажиллаж байгаа програмын \(сервер\) функцийг ажиллуулж байгаа хэрэг юм.

RPC хэрэгжүүлэх хоёр нийтлэг арга байдаг. Эхнийх нь IDL буюу интерфэйс  тодорхойлох тусгай хэл ашигладаг RPC системүүд бий. Жишээлбэл CORBA, Google RPC, Java RMI, dotNET Remoting зэрэг нь энэ төрөлд хамаарна.

Клиент тал нь мессежийг форматад оруулж "савлаад" сервер рүү илгээнэ. Сервер нь мессежийг задлаад сервер дээр байгаа харгалзах функцийн дуудалт болгон хувиргах буюу функцийг дуудаж ажиллуулна. Дуудалтын үр дүнг дахин форматад оруулж клиент руу буцаахад бэлтгэнэ.![](/7_net/res/rpc.png)

TODO: RPC-н Go, HTTP, TCP, JSON хувилбаруудыг зураг дээр оруулах

Хоёр дахь нь тусгай клиент API гаргаж ашиглах арга юм. Энэ API-г клиент болон сервер талууд мэдэх хэрэгтэй, хоёр тал функц, параметр, мессежийн формат, бүтцээ сайн мэдэж байх хэрэгтэй байдаг. Жишээлбэл Web Service \(SOAP, JSON, REST\), Go RPC  зэрэг нь энэ төрлийн RPC систем юм.

## Go RPC

Go-н RPC нь бусад RPC системээс ялгаатай байдаг. Зөвхөн Go програмууд л хоорондоо харьцахад ашиглагддаг. Мөн өгөгдөл дамжуулахдаа Gob  хувиргалтыг ашигладаг.

RPC системүүд ихэвчлэн дуудагдах функцүүд дээр хязгаарлалт тавьдаг. Үүний ачаар солилцох мессежийн бүтэц нь тодорхой болж зөв ойлголцох боломжтой болдог.

Go RPC хязгаарлалтууд нь:

* функц нь нээлттэй хандалттай байх ёстой \(нэр нь том үсгээр эхлэх\);
* яг хоёр аргументтай байна, тухайлбал дараах хэлбэртэй бичигдэнэ:
  ```go
  F(&T1, &T2) error
  ```
* эхнийх нь клиентээс ирсэн хүсэлт, хоёр дахь нь клиентэд буцаах хариу;
  буцах утга нь `error` төрөлтэй байх

Аргументын тоо хязгаартай учраас клиентээс ирсэн хүсэлтийг нийлмэл төрлөөр тодорхойлох хэрэгтэй болно. Хүсэлт, хариуг Gob ашиглан хувиргаж дамжуулах учраас уг төрөл нь Gob дэмжих хэрэгтэй.

Жишээ болгон арифметик тооцоолол хийх `Arith` нэртэй RPC сервер үүсгэе. Энэ сервер нь хоёр бүхэл тоог үржүүлэх \(`Multiply`\), хуваах \(`Divide`\) хоёр функтэй.

Хүсэлт болох хоёр тоо нь дараах бүтцээр ирнэ гэж тооцоё:

```go
type Args struct {
    A, B int
}
```

Үржүүлэх функцийн хариу нь бүхэл тоо байна, харин хуваах үйлдлийн хариу нь ноогдвор, үлдэгдэл гэсэн хоёр тоог буцаах хэрэгтэй болно. Тэгэхлээр дараах шинэ бүтэц үүсгэж болно.

```go
type Quotient struct {
    Quo, Rem int
}
```

RPC сервер үүсгэхийн тулд `Multiply`, `Divide` функцүүд агуулсан `Arith` нэртэй шинэ төрөл тодорхойлох хэрэгтэй.

```go
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
        *reply = args.A * args.B
        return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
        if args.B == 0 {
                return error.String("divide by zero")
        }
        quo.Quo = args.A / args.B
        quo.Rem = args.A % args.B
        return nil
}
```

Энд `Arith`-н төрөл нь `int` гэж өгөгдсөн байна. Энэ нь дурын төрөл байж болно, тийм ч чухал биш.

Ингээд `Arith` төрлийн обект үүсгэж `rpc.Register()` функцээр бүртгэснээр RPC сервер ажиллагаатай болж клиентээс дуудахад бэлэн болно.

## HTTP RPC Server

Ямар ч RPC системд өгөгдлийг сүлжээгээр тээвэрлэх механизм буюу протокол шаардлагатай болно. Go RPC-н хувьд HTTP, TCP протоколыг ашиглах боломжтой байдаг.  Бэлэн HTTP сангийн боломжийг ашиглан HTTP протоколоор өгөгдөл солилцох RPC серверийг хялбархан хийж болно. Үүний тулд RPC боловсруулагчаа үүсгээд `HandleHTTP()` функцээр HTTP давхарга ашиглана гэдгийг тохируулна, ингээд HTTP серверээ асаана. Бүрэн кодыг доор харуулав.

```go
// ArithServer.go
package main

import ("fmt";"net/rpc";"errors";"net/http")

// Args, Quotient, Arith төрлүүд энд байна
// Multiply, Divide функцүүд энд байна

func main() {
        arith := new(Arith)
        rpc.Register(arith)
        rpc.HandleHTTP()

        err := http.ListenAndServe(":1234", nil)
        if err != nil {
                fmt.Println(err.Error())
        }
}
```

## HTTP RPC клиент

Клиент нь RPC сервер рүү HTTP холболт үүсгэх хэрэгтэй. Мөн серверт илгээх утгуудыг багтааж бүтэц үүсгэх хэрэгтэй, дараа нь ирэх хариуг хадгалахад зориулсан хувьсагч үүсгэх хэрэгтэй, тэгээд аргументуудаа дамжуулаад `Call()` функцийг дуудна:

* Алсад байгаа дуудагдах функцийн нэр
* Илгээх утгууд
* Ирэх хариуг хадгалах хувьсагч

Арифметик серверийн функцүүдийг дуудсан клиент програмыг доор харуулав

```go
// ArithClient.go
package main

import ("net/rpc";"fmt";"log";"os")

// Args, Quotient төрлүүд энд байна

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "server")
                os.Exit(1)
        }
        serverAddress := os.Args[1]

        client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
        if err != nil {
                log.Fatal("dialing:", err)
        }
        // Synchronous call
        args := Args{17, 8}
        var reply int
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

        var quot Quotient
        err = client.Call("Arith.Divide", args, &quot)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d/%d=%d remainder %d\n",
            args.A, args.B, quot.Quo, quot.Rem)

}
```

## TCP RPC сервер

Арифметик серверийн TCP хувилбар нь дараах байдалтай болно.

```go
func main() {
        arith := new(Arith)
        rpc.Register(arith)

        tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)

        for {
                conn, err := listener.Accept()
                if err != nil {
                        continue
                }
                rpc.ServeConn(conn)
        }
}

// checkError функц энд байна
```

Энд `Accept()` дуудалт дээр клиентээс холболт хийхийг хүлээж блок үүсгэж байгааг хараарай. Хэрэв серверийг блоклохгүйгээр өөр ажил зэрэг хийлгэх бол энэ дуудалтыг go функц дотор тавьж болно.

## TCP RPC клиент

Арифметик TCP сервер рүү дуудалт хийх клиентийг доор харуулав

```go
package main

import ("net/rpc"; "fmt"; "log"; "os")

// Args, Quotient төрлүүд

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "server:port")
                os.Exit(1)
        }
        service := os.Args[1]

        client, err := rpc.Dial("tcp", service)
        if err != nil {
                log.Fatal("dialing:", err)
        }
        // Synchronous call
        args := Args{17, 8}
        var reply int
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

        var quot Quotient
        err = client.Call("Arith.Divide", args, &quot)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d/%d=%d remainder %d\n",
            args.A, args.B, quot.Quo, quot.Rem)

}
```

Клиент талд байгаа аргументийн нэрс нь сервер талд байгаа аргументийн нэрсээс өөр байж болно. Жишээлбэл, сервер талд байгаа `Args` нэртэй бүтцэд харгалзах`Values` нэртэй бүтэц үүсгээд клиент талд  ашиглаж болно.

```go
type Values struct {
        C, B int
}
```

Нэр нь чухал биш, харин бүтцийн талбарууд нь таарч байх ёстой. Нэрний зөрүүг RPC дамжуулалт, хувиргалтын явцад тоохгүй өнгөрөөдөг. Гэхдээ будлиан үүсгэхгүйн тулд клиент ба сервер талд байгаа мессеж бүтцүүдийн нэршилийг адилхан байлгавал зүгээр.

## JSON RPC

JSON RPC нь өмнөхөөс ялгаатай зүйл бараг байхгүй, ганц ялгаа нь тээвэрлэх бүтэц нь Gob-н оронд JSON байх юм. Тийм болохоор бусад хэл дээр бичсэн клиент програмаас хандаж болно гэсэн үг.

### JSON RPC client

Арифметик серверийн хоёр функцийг дуудах клиент програм

```go
// JSONArithCLient.go
package main

import ("net/rpc/jsonrpc";"fmt";"log";"os")

// Args, Quotient төрлүүд

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "server:port")
                log.Fatal(1)
        }
        service := os.Args[1]

        client, err := jsonrpc.Dial("tcp", service)
        if err != nil {
                log.Fatal("dialing:", err)
        }
        // Synchronous call
        args := Args{17, 8}
        var reply int
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

        var quot Quotient
        err = client.Call("Arith.Divide", args, &quot)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d/%d=%d remainder %d\n",
            args.A, args.B, quot.Quo, quot.Rem)

}
```

### JSON RPC server

JSON энкодинг ашигласан сервер хувилбар

```go
// JSONArithServer.go
package main

import ("fmt";"net/rpc";"net/rpc/jsonrpc";"os";"net";"errors")

// Args, Quotient, Arith төрлүүд

// Multiply, Divide функцүүд

func main() {
        arith := new(Arith)
        rpc.Register(arith)

        tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)

        /* This works:
        rpc.Accept(listener)
        */
        /* and so does this:
         */
        for {
                conn, err := listener.Accept()
                if err != nil {
                        continue
                }
                jsonrpc.ServeConn(conn)
        }

}

// checkError функц
```




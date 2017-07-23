# RPC

RPC \(Remote Procedure Call\) нь сервер талд байгаа функцийг сүжээгээр дамжуулан дуудах арга юм. RPC нь ердийн функц дуудаж ажиллуулж байгаа юм шиг харагддаг. Гэхдээ нэг програм \(клиент\), цаашлаад компютерээс хальж сүлжээний өөр нэг компютер дээр ажиллаж байгаа програмын \(сервер\) функцийг ажиллуулж байгаа хэрэг юм.

RPC хэрэгжүүлэх хоёр нийтлэг арга байдаг. Эхнийх нь IDL буюу интерфэйс  тодорхойлох тусгай хэл ашигладаг RPC системүүд бий. Жишээлбэл CORBA, Google RPC, Java RMI, NET Remoting зэрэг нь энэ төрөлд хамаарна.

The client-side will package this into a network message and transfer it to the server. The server will unpack this and turn it back into a procedure call on the server side. The results of this call will be packaged up for return to the client.

TODO: RPC зураг оруулах: Stub, Proxy үүсгэх тухай

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

Ямар ч RPC системд өгөгдлийг сүлжээгээр тээвэрлэх механизм шаардлагатай болно. Go RPC-н хувьд HTTP, TCP протокол ашиглах боломжтой байдаг. The advantage of the HTTP mechanism is that it can leverage off the HTTP suport library. You need to add an RPC handler to the HTTP layer which is done using HandleHTTP and then start an HTTP server. The complete code is

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

## HTTP RPC client

The client needs to set up an HTTP connection to the RPC server. It needs to prepare a structure with the values to be sent, and the address of a variable to store the results in. Then it can make a Call with arguments:  
The name of the remote function to execute  
The values to be sent  
The address of a variable to store the result in  
A client that calls both functions of the arithmetic server is

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

## TCP RPC server

Arith серверийн TCP хувилбар нь дараах байдалтай болно.

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

Note that the call to Accept is blocking, and just handles client connections. If the server wishes to do other work as well, it should call this in a goroutine.

## TCP RPC client

A client that uses the TCP server and calls both functions of the arithmetic server is

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

1. Matching values
   We note that the types of the value arguments are not the same on the client and server. In the server, we have used Values while in the client we used Args. That doesn't matter, as we are following the rules of gobserialisation, and the names an types of the two structures' fields match. Better programming practise would say that the names should be the same!
   However, this does point out a possible trap in using Go RPC. If we change the structure in the client to be, say,

```go
type Values struct {
        C, B int
}
```

then gob has no problems: on the server-side the unmarshalling will ignore the value of C given by the client, and use the default zero value for A.  
Using Go RPC will require a rigid enforcement of the stability of field names and types by the programmer. We note that there is no version control mechanism to do this, and no mechanism in gob to signal any possible mismatches.

## JSON RPC

JSON RPC нь өмнөхөөс ялгаатай зүйл бараг байхгүй, ганц ялгаа нь тээвэрлэх бүтэц нь Gob-н оронд JSON байх юм. Тийм болохоор бусад хэл дээр бичсэн клиент програмаас хандаж болно гэсэн үг.

### JSON RPC client

A client that calls both functions of the arithmetic server is

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

A version of the server that uses JSON encoding is

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




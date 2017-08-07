package main 
import ( "fmt" ;"net" ; "os") 

func main() { 
  service := ":8088" 
  listener, err := net.Listen("tcp", service) 
  checkError(err) 
 
  for { 
	conn, err := listener.Accept() 
	if err != nil { 
	 continue 
	} 
	// клиентэд хариу илгээх
	conn.Write([]byte("Амжилттай холбогдлоо. Таны хаяг: " + conn.RemoteAddr().String() + "\n" ))
	conn.Close()
  } 
}

func checkError(err error) { 
  if err != nil { 
	fmt.Fprintf(os.Stderr, "Алдаа: %s", err.Error()) 
	os.Exit(1) 
  } 
} 
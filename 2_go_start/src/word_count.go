package main

func wordCount(s string) int {
    c:=0

    for i:=0; i<len(s); i++ {
	    for i<len(s) && s[i] == ' ' { // сул зайн тэмдэгт
	        i++; 
	    }
	    for i<len(s) && s[i] != ' ' { // бусад тэмдэгт
	        i++; 
	    }
	    c++;
    }

    return c
}

func main() {
    str:="Hello World!  "

    println(str)
    println(wordCount(str), "үг")
}

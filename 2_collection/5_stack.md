# Стек

Стек нь “эхэлж орсон нь эхэлж гарах” зарчим дээр тулгуурласан жагсаалт төрлийн өгөгдлийн бүтэц юм. Стек бүтэц дээр дараах хоёр үйлдэл хийгдэнэ:

* Push – стекийн орой дээр элемент нэмэх
* Pop – стекийн оройноос элемент сугалах

Дараах жишээнд стек хэрхэн үүсгэж, ашиглахыг харуулав.

```go
package main
import "fmt"

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
}

// стекийн урт
func (s *Stack) Len() int {
	return s.size
}

// стекийн оройд элемент нэмэх
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// стекийн оройгоос элемент авах
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func main() {
	stack := new(Stack)

	stack.Push("Things")
	stack.Push("and")
	stack.Push("Stuff")

	for stack.Len() > 0 {
		fmt.Printf("%s ", stack.Pop().(string))
	}
	fmt.Println()
}
```

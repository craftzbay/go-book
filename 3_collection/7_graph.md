# Граф

Граф нь орой болон ирмэгээс тогтох олонлог юм. Ирмэг бүр хоёр оройг холбоно.

Бидний эргэн тойронд маш олон граф байдаг. Граф бүтцийг сүлжээ болон траффик судлах, зам, маршрут тооцоолох, нийгэм буюу хүмүүсийн харилцаа холбоо судлах, мэдлэгийн сан бүрдүүлэх, халдварт өвчний тархалт судлах, удамшил генийн судалгаа хийх зэрэг маш олон төрлийн асуудал шийдвэрлэхэд ашигладаг.

|Граф|Орой|Ирмэг|
|--|--|--|
|Компютерийн сүлжээ|Компютер, төхөөрөмжүүд|Кабел утас|
|Интернэт|Вэб хуудас|Гипер холбоос|
|Санхүү|Данс|Гүйлгээ|
|Ус суваг|Худаг, орон сууц|Дамжуулах хоолой|
|Цахилгаан түгээх|Хэрэглэгч, трансформатор|Цахилгааны шугам|
|Нийгэм|Хүмүүс|Найз, нөхөрлөл, харилцаа|
|Халдварт өвчин|Хүмүүс|Халдварлалт|
|Автобус маршрут|Зогсоол|Тээврийн чиглэл, зам|

Граф бүтцийг майп, жагсаалт, массив гэх мэт бидний өмнө үзсэн өгөгдлийн бүтцүүдийг ашиглан зохион байгуулж болно.

Жишээ болгон давхар майп буюу Key → (Key → Value) бүтэцтэй граф үүсгэе.

![](res/graph_sample.svg)

Давхар майпаар дүрсэлвэл дараах байдалтай болно.

```
A -> ( B -> 2 )
A -> ( C -> 0 )
A -> ( G -> 3 )
A -> ( H -> 5 )
B -> ( A -> 2 )
C -> ( A -> 0 )
G -> ( A -> 3 )
G -> ( C -> 0 )
H -> ( A -> 5 )
```

Давхар майп ашиглан граф үүсгэх програм:

```go
package main

import (
  "bytes"
  "fmt"
)

// граф бүтэц
type Graph struct {
	data map[string]map[string]int
}

func NewGraph() *Graph {
	g := &Graph{}
	g.data = make(map[string]map[string]int)
	return g
}

// оройн жагсаалт
func (g *Graph) V() []string {
	v := make([]string, 0)
	for k, _ := range g.data {
		v = append(v, k)
	}
	return v
}
// ирмэг нэмэх
func (g *Graph) addEdge(source, target string, weight int) {
	map2, ok := g.data[source]
	if !ok {
		map2 = make(map[string]int)
	}
	map2[target] = weight
	g.data[source] = map2
}
// графыг хэвлэх
func (g *Graph) String() string {
	var strBuf bytes.Buffer
	for k, map2 := range g.data {
		strBuf.WriteString(k + "\n")
		for k2, w := range map2 {
		  strBuf.WriteString(fmt.Sprintf("-%s-%d\n", k2, w))
		}
	}
	return strBuf.String()
}

func main() {
	g := NewGraph()
	g.addEdge("A", "B", 2)
	g.addEdge("A", "C", 0)
	g.addEdge("A", "G", 3)
	g.addEdge("A", "H", 5)
	g.addEdge("B", "A", 2)
	g.addEdge("C", "A", 0)
	g.addEdge("G", "A", 3)
	g.addEdge("G", "C", 0)
	g.addEdge("H", "A", 5)

	fmt.Println(g)
}
```

# Эрэмбэлэлт

Go хэлний `sort` пакет нь төрөл бүрийн өгөгдлийн олонлогийг эрэмбэлэхэд зориулагдсан байдаг. Тоон массив зэрэг энгийн өгөгдлийн олонлогийг эрэмбэлэхэд `sort.Sort()` эсвэл `sort.Slice()` функцыг шууд ашиглаж болно.

TODO: sort.Slice() функцийн тухай оруулах

Хэрэв нийлмэл төрлийн өгөгдлийн олонлогийг эрэмбэлэх бол хэрхэн харьцуулах, эрэмбэлэх тухайгаа тохируулж өгөх шаардлагатай.

Тухайлбал `Sort()` функц нь эрэмбэлэгдэх төрөл дээр  `Len`, `Less`, `Swap` гэсэн методууд байхыг шаардана.

Жишээ болгон `Person` төрлийн олонлогыг эрэмбэлэе.

```go
type Person struct {
    Name string
    Age int
}
```

Person төрлийн массивыг нэрээр нь (`Name` талбараар) эрэмбэлэхэд зориулж `ByName` нэртэй шинэ бүтэц үүсгэе.

```go
type ByName []Person

func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}
```

Ингээд `sort.Sort()` функцийг дараах байдлаар ашиглаж болно.

```go
func main() {
    kids := []Person{
        {"Бат",9},
        {"Болд",10},
        {"Амар",9},
    }
    sort.Sort(ByName(kids))
    fmt.Println(kids)
}
```

Дээрхтэй төстэйгээр мөн `Age` талбараар нь хүмүүсийг эрэмбэлэх бол шинэ `ByAge` төрөл үүсгэж болно.

```go
type ByAge []Person
func (this ByAge) Len() int {
    return len(this)
}
func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}
```

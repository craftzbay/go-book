# Эрэмбэлэлт

Go хэлний `sort` пакет нь төрөл бүрийн өгөгдлийн олонлогийг эрэмбэлэхэд зориулагдсан байдаг. Тоон массив зэрэг энгийн өгөгдлийн олонлогийг эрэмбэлэхэд `sort.Sort()` эсвэл `sort.Slice()` функцыг шууд ашиглаж болно.

TODO: sort.Slice() функцийн тухай оруулах

```go
sort.Slice(s, func(i, j int) bool {
    if s[i].Foo != s[j].Foo {
        return s[i].Foo < s[j].Foo
    }
    return s[i].Bar < s[j].Bar
})
```

Шинэ эрэмбэлэгч төрөл үүсгэх замаар мөн өгөгдлийн олонлогыг эрэмбэлж болно. Жишээ болгон `Person` төрлийн олонлогыг эрэмбэлэх `ByName` нэртэй эрэмбэлэгч төрөл үүсгэе.

```go
type Person struct {
    Name string
    Age int
}
```

Эрэмбэлэгч төрөлд хэрхэн харьцуулах, эрэмбэлэх тухай `Len`, `Less`, `Swap` гэсэн методуудад тодорхойлж өгдөг. Дараах `ByName` эрэмбэлэгч нь Person төрлийн нэрээр (`Name` талбараар) харьцуулалт хийж байна.

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

Ингээд `sort.Sort()` функцийг `ByName` эрэмбэлэгчтэй хослуулан дараах байдлаар ашиглаж болно.

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

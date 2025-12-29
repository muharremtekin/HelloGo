# Go Cheatsheet - Ezberle!

## Range Dönen Değerler

```go
for i, v := range slice    // index, value
for k, v := range map      // key, value
for i, r := range "string" // index, rune
for v := range channel     // value (tek)
```

## Comma OK Idiom

```go
val, ok := map[key]       // value, exists?
val, ok := x.(Type)       // value, success? (type assertion)
val, ok := <-channel      // value, open?
```

## Zero Values (Varsayılan Değerler)

| Tip | Zero Value |
|-----|------------|
| `int`, `float64` | `0` |
| `string` | `""` (boş) |
| `bool` | `false` |
| `pointer`, `slice`, `map`, `channel`, `func`, `interface` | `nil` |
| `struct` | tüm field'lar zero value |

## nil vs make

```go
var m map[string]int      // nil - yazma PANIC!
m := make(map[string]int) // boş ama kullanılabilir

var s []int               // nil - ama append çalışır
s := make([]int, 5)       // len=5, cap=5
s := make([]int, 0, 10)   // len=0, cap=10
```

## Kısa Tanımlamalar

```go
x := 5                    // var x int = 5
a, b := 1, 2              // çoklu
a, b = b, a               // swap
```

## Fonksiyon Dönüş Tipleri

```go
func f() int              // tek dönüş
func f() (int, error)     // çoklu dönüş
func f() (result int)     // named return
```

## Pointer

```go
x := 5
p := &x    // p = x'in adresi
*p = 10    // x artık 10
```

## Struct

```go
type T struct { Name string }

t := T{Name: "ali"}       // literal
t := T{"ali"}             // sırayla (önerilmez)
t := new(T)               // pointer döner (*T)
```

## Embedded vs Nested

```go
type A struct {
    B           // embedded - a.Field direkt erişim
    C C         // nested - a.C.Field
}
```

## Method Receiver

```go
func (t T) Method()       // value - kopyayı değiştirir
func (t *T) Method()      // pointer - orijinali değiştirir
```

## Görünürlük (Export)

```go
Name    // Büyük harf = public (dışarı açık)
name    // Küçük harf = private (paket içi)
```

## Format Verbleri

```go
%v      // value (varsayılan)
%T      // type
%d      // int
%f      // float (%.2f = 2 ondalık)
%s      // string
%q      // quoted string
%t      // bool
%p      // pointer
%%      // literal %
```

## Kontrol Yapıları

```go
// if - parantez yok!
if x > 5 { }
if val, ok := m[k]; ok { }  // if with statement

// switch - break otomatik!
switch x {
case 1, 2: // birden fazla
case 3:
    fallthrough  // devam et
default:
}

// for - tek döngü tipi
for i := 0; i < 5; i++ { } // klasik
for x < 10 { }              // while
for { }                     // sonsuz
for i, v := range arr { }   // range
```

## Sık Hatalar

```go
// 1. nil map'e yazma
var m map[string]int
m["x"] = 1  // PANIC!

// 2. Kullanılmayan değişken
x := 5  // HATA: x declared but not used
_ = x   // OK: _ ile yoksay

// 3. Kullanılmayan import
import "fmt"  // HATA: imported but not used

// 4. Döngüde pointer
for _, v := range items {
    results = append(results, &v)  // Hepsi aynı adresi gösterir!
}
// Çözüm:
for _, v := range items {
    v := v  // shadow variable
    results = append(results, &v)
}
```

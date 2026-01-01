# Go Öğrenme İlerlemesi

## Temel Konular

- [x] Değişkenler (variables)
  - [x] `var` ile tanımlama
  - [x] `:=` shorthand
  - [x] Çoklu değişken tanımlama
  - [x] Gruplu değişken tanımlama (`var ( ... )`)
- [x] Sabitler (`const`)
- [x] Temel veri tipleri (`int`, `float64`, `string`, `bool`)
- [x] Formatlama (`fmt.Printf`, format verbleri)
  - [x] Genel: `%v`, `%T`, `%%`
  - [x] Integer: `%d`, `%b`, `%o`, `%x`
  - [x] String: `%s`, `%q`
  - [x] Float: `%f`, `%.2f`

## Veri Yapıları

- [x] Diziler (Arrays)
  - [x] Sabit boyutlu dizi tanımlama
  - [x] `len()` ve `cap()`
- [x] Dilimler (Slices)
  - [x] Slice oluşturma
  - [x] `append()`
  - [x] Slice kesme (`slice[2:5]`)
  - [x] `make()` ve `copy()`
- [x] Struct'lar
  - [x] Struct tanımlama
  - [x] Struct oluşturma yöntemleri
  - [x] Field erişimi
  - [x] Method tanımlama (value receiver)
  - [x] Pointer receiver
  - [x] Embedded struct (composition)
  - [x] Nested struct
  - [x] Struct pointer
  - [x] Anonymous struct
- [x] Maps
  - [x] Map oluşturma (`make`, literal)
  - [x] Okuma/yazma
  - [x] Key var mı kontrolü (comma ok idiom)
  - [x] `delete()` ile silme
  - [x] `range` ile döngü
  - [x] Nested map
  - [x] Map with struct
  - [x] Map as set

## Kontrol Yapıları

- [x] If / Else
  - [x] Basit if/else
  - [x] If with statement
- [x] Switch
  - [x] Basit switch
  - [x] Switch with statement
  - [x] Switch without expression
  - [x] Fallthrough
  - [x] Type switch
- [x] For döngüsü
  - [x] Klasik for
  - [x] While tarzı for
  - [x] Sonsuz döngü
  - [x] Range ile for (slice, map, string)
  - [x] Continue ve Break
  - [x] Break with label

## Fonksiyonlar

- [x] Temel fonksiyon tanımlama
- [x] Multiple return values
- [x] Named return values
- [x] Variadic functions (`...`)
- [x] Anonymous functions
- [x] Closures
- [x] Functions as parameters (higher-order)
- [x] Defer (LIFO sırası)
- [x] Function types

## İleri Seviye

- [x] Pointers
  - [x] & (address of) ve * (dereference)
  - [x] nil pointer
  - [x] Pointer ve fonksiyonlar (pass by reference)
  - [x] Pointer ve struct
  - [x] new() fonksiyonu
  - [x] Pointer to pointer
- [ ] Interfaces
- [ ] Error handling
- [ ] Goroutines (concurrency)
- [ ] Channels
- [ ] Packages ve Modules
- [ ] Testing

## Notlar

- Go'da class yok, struct var
- Inheritance yok, composition (embedding) var
- `=` atama, `:=` tanımlama + atama

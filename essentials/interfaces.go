package essentials

import (
	"fmt"
	"math"
)

// Interfaces in Go
// Interface, bir tipin hangi metodlara sahip olması gerektiğini tanımlar
// C#'taki interface ile aynı mantık ama IMPLICIT (otomatik) implement edilir

func Interfaces() {
	fmt.Println("========== INTERFACES ==========")

	// ==================== TEMEL INTERFACE ====================
	fmt.Println("\n--- Temel Interface ---")

	// Shape interface'ini implemente eden tipler
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	// Her ikisi de Shape - polymorphism!
	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Rectangle area:", rectangle.Area())

	// Interface tipinde değişken
	var shape Shape
	shape = circle
	fmt.Println("Shape (circle) area:", shape.Area())

	shape = rectangle
	fmt.Println("Shape (rectangle) area:", shape.Area())

	// ==================== IMPLICIT IMPLEMENTATION ====================
	fmt.Println("\n--- Implicit Implementation ---")

	// Go'da "implements" keyword'ü YOK!
	// Metod imzaları eşleşiyorsa otomatik implement olur

	// C#'ta:
	// class Circle : IShape { ... }

	// Go'da:
	// Circle struct'ı Area() metoduna sahip -> otomatik olarak Shape

	fmt.Println("Circle, Shape interface'ini implemente ediyor mu?")
	fmt.Println("Evet! Area() metodu var, yeterli.")

	// ==================== INTERFACE İLE FONKSİYON ====================
	fmt.Println("\n--- Interface ile Fonksiyon ---")

	shapes := []Shape{
		Circle{Radius: 2},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 3},
	}

	totalArea := calculateTotalArea(shapes)
	fmt.Printf("Toplam alan: %.2f\n", totalArea)

	// Her şekli yazdır
	for i, s := range shapes {
		printShapeInfo(i, s)
	}

	// ==================== EMPTY INTERFACE ====================
	fmt.Println("\n--- Empty Interface (any) ---")

	// interface{} veya any = her tipi kabul eder
	// C#'taki object gibi

	var anything interface{}

	anything = 42
	fmt.Printf("int: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("string: %v (type: %T)\n", anything, anything)

	anything = Circle{Radius: 5}
	fmt.Printf("Circle: %v (type: %T)\n", anything, anything)

	// any = interface{} için alias (Go 1.18+)
	var anything2 any = 3.14
	fmt.Printf("any: %v (type: %T)\n", anything2, anything2)

	// Slice of any
	mixedSlice := []any{1, "two", 3.0, true, Circle{Radius: 1}}
	fmt.Println("Mixed slice:", mixedSlice)

	// ==================== TYPE ASSERTION ====================
	fmt.Println("\n--- Type Assertion ---")

	var val interface{} = "hello"

	// Type assertion: interface'ten gerçek tipe dönüştür
	str := val.(string) // Eğer string değilse PANIC!
	fmt.Println("String değeri:", str)

	// Güvenli type assertion (comma ok idiom)
	if str, ok := val.(string); ok {
		fmt.Println("String olarak:", str)
	}

	if num, ok := val.(int); ok {
		fmt.Println("Int olarak:", num)
	} else {
		fmt.Println("val bir int değil")
	}

	// ==================== TYPE SWITCH ====================
	fmt.Println("\n--- Type Switch ---")

	testValues := []interface{}{42, "hello", 3.14, true, Circle{Radius: 2}}

	for _, v := range testValues {
		describeType(v)
	}

	// ==================== MULTIPLE INTERFACES ====================
	fmt.Println("\n--- Multiple Interfaces ---")

	// Bir tip birden fazla interface'i implemente edebilir

	dog := Dog{Name: "Karabaş"}

	// Dog hem Speaker hem Mover
	var speaker Speaker = dog
	var mover Mover = dog

	fmt.Println(speaker.Speak())
	fmt.Println(mover.Move())

	// Composed interface
	var animal Animal = dog
	fmt.Println("Animal speak:", animal.Speak())
	fmt.Println("Animal move:", animal.Move())

	// ==================== INTERFACE COMPOSITION ====================
	fmt.Println("\n--- Interface Composition ---")

	// Interface'ler birleştirilebilir (embedding)
	// Animal = Speaker + Mover

	cat := Cat{Name: "Tekir"}
	makeAnimalAct(cat)

	// ==================== STRINGER INTERFACE ====================
	fmt.Println("\n--- Stringer Interface (fmt) ---")

	// fmt.Stringer interface'i: String() string metodu
	// fmt.Println otomatik olarak bunu çağırır

	book := Book{Title: "Go Programming", Author: "John Doe", Pages: 300}
	fmt.Println(book) // String() metodu çağrılır

	// ==================== POINTER VS VALUE RECEIVER ====================
	fmt.Println("\n--- Pointer vs Value Receiver ---")

	// Value receiver: hem value hem pointer çalışır
	// Pointer receiver: sadece pointer çalışır (interface için)

	counter := Counter{count: 0}

	// Value olarak geçemez çünkü Increment() pointer receiver
	// var inc Incrementer = counter // HATA!
	var inc Incrementer = &counter // OK - pointer gerekli
	inc.Increment()
	inc.Increment()
	fmt.Println("Counter:", counter.count)

	// ==================== NIL INTERFACE ====================
	fmt.Println("\n--- Nil Interface ---")

	var s Shape // nil interface
	fmt.Println("Nil interface:", s)
	fmt.Println("s == nil:", s == nil)

	// DİKKAT: nil pointer != nil interface
	var c *Circle = nil
	s = c
	fmt.Println("s (nil *Circle):", s)
	fmt.Println("s == nil:", s == nil) // false! çünkü tip bilgisi var

	// ==================== BEST PRACTICES ====================
	fmt.Println("\n--- Best Practices ---")
	fmt.Println("1. Küçük interface'ler yaz (1-3 metod)")
	fmt.Println("2. Interface'i kullanan tarafta tanımla (consumer)")
	fmt.Println("3. Accept interfaces, return structs")
	fmt.Println("4. io.Reader, io.Writer gibi stdlib interface'leri kullan")
}

// ==================== TYPES & INTERFACES ====================

// Shape interface
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Circle implements Shape (implicit)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Rectangle implements Shape (implicit)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Speaker interface
type Speaker interface {
	Speak() string
}

// Mover interface
type Mover interface {
	Move() string
}

// Animal = Speaker + Mover (composed interface)
type Animal interface {
	Speaker
	Mover
}

// Dog struct
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Hav hav!"
}

func (d Dog) Move() string {
	return d.Name + " is running"
}

// Cat struct
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says: Miyav!"
}

func (c Cat) Move() string {
	return c.Name + " is walking"
}

// Book struct with Stringer interface
type Book struct {
	Title  string
	Author string
	Pages  int
}

// String() metodu = fmt.Stringer interface
func (b Book) String() string {
	return fmt.Sprintf("📚 %s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

// Counter with pointer receiver
type Counter struct {
	count int
}

type Incrementer interface {
	Increment()
}

func (c *Counter) Increment() {
	c.count++
}

// ==================== HELPER FUNCTIONS ====================

// Interface alan fonksiyon
func calculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func printShapeInfo(index int, s Shape) {
	fmt.Printf("Shape %d: area = %.2f\n", index, s.Area())
}

// Type switch ile tip kontrolü
func describeType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("%v is an int (doubled: %d)\n", val, val*2)
	case string:
		fmt.Printf("%v is a string (length: %d)\n", val, len(val))
	case float64:
		fmt.Printf("%v is a float64\n", val)
	case bool:
		fmt.Printf("%v is a bool\n", val)
	case Shape:
		fmt.Printf("%v is a Shape (area: %.2f)\n", val, val.Area())
	default:
		fmt.Printf("%v is unknown type: %T\n", val, val)
	}
}

// Composed interface kullanan fonksiyon
func makeAnimalAct(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.Move())
}

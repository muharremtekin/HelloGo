package essentials

import (
	"errors"
	"fmt"
	"strings"
)

// Functions in Go
// Go'da fonksiyonlar first-class citizen (değişkene atanabilir, parametre olarak geçilebilir)

func Functions() {
	fmt.Println("========== FUNCTIONS ==========")

	// ==================== MULTIPLE RETURN VALUES ====================
	fmt.Println("\n--- Multiple Return Values ---")

	// Go'da fonksiyonlar birden fazla değer dönebilir
	// En yaygın kullanım: (result, error) pattern

	sum, diff := addAndSubtract(10, 3)
	fmt.Printf("10 + 3 = %d, 10 - 3 = %d\n", sum, diff)

	// Error handling pattern
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Hata durumu
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Hata:", err)
	}

	// Sadece birini almak istersen _ kullan
	onlySum, _ := addAndSubtract(5, 3)
	fmt.Println("Sadece toplam:", onlySum)

	// ==================== NAMED RETURN VALUES ====================
	fmt.Println("\n--- Named Return Values ---")

	// Dönüş değerlerine isim verebilirsin
	// "Naked return" yapabilirsin ama önerilmez (okunabilirlik düşer)

	w, h, a := getRectangleInfo(5, 3)
	fmt.Printf("Genişlik: %d, Yükseklik: %d, Alan: %d\n", w, h, a)

	// ==================== VARIADIC FUNCTIONS ====================
	fmt.Println("\n--- Variadic Functions ---")

	// ... ile sınırsız parametre alabilirsin
	// C#'taki params keyword'ü gibi

	fmt.Println("Sum:", sumAll(1, 2, 3, 4, 5))
	fmt.Println("Sum:", sumAll(10, 20))
	fmt.Println("Sum:", sumAll()) // 0 parametre de OK

	// Slice'ı variadic'e geçme
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice sum:", sumAll(numbers...)) // ... ile aç

	// Printf da variadic!
	fmt.Printf("Merhaba %s, yaşın %d\n", "Ali", 25)

	// ==================== ANONYMOUS FUNCTIONS ====================
	fmt.Println("\n--- Anonymous Functions ---")

	// Fonksiyonu değişkene atama
	greet := func(name string) string {
		return "Merhaba " + name
	}
	fmt.Println(greet("Dünya"))

	// Hemen çağırma (IIFE - Immediately Invoked Function Expression)
	result2 := func(a, b int) int {
		return a * b
	}(3, 4)
	fmt.Println("3 * 4 =", result2)

	// ==================== CLOSURES ====================
	fmt.Println("\n--- Closures ---")

	// Closure: Dış scope'taki değişkenlere erişebilen fonksiyon
	// Counter örneği - state tutuyor

	counter := makeCounter()
	fmt.Println("Count:", counter()) // 1
	fmt.Println("Count:", counter()) // 2
	fmt.Println("Count:", counter()) // 3

	// Her makeCounter yeni bir counter oluşturur
	counter2 := makeCounter()
	fmt.Println("Counter2:", counter2()) // 1 (sıfırdan başlar)

	// Multiplier örneği
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("Double 5:", double(5)) // 10
	fmt.Println("Triple 5:", triple(5)) // 15

	// ==================== FUNCTIONS AS PARAMETERS ====================
	fmt.Println("\n--- Functions as Parameters ---")

	// Fonksiyonu parametre olarak geçme
	nums := []int{1, 2, 3, 4, 5}

	// Her elemana fonksiyon uygula
	doubled := mapInts(nums, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled:", doubled)

	squared := mapInts(nums, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)

	// Filter örneği
	evens := filterInts(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	// ==================== DEFER ====================
	fmt.Println("\n--- Defer ---")

	// defer: Fonksiyon bitince çalıştır (finally gibi)
	// LIFO sırası - son defer ilk çalışır

	deferExample()

	// Pratik kullanım: Resource cleanup
	fmt.Println("\nDefer ile kaynak yönetimi simülasyonu:")
	processFile()

	// ==================== FUNCTION TYPES ====================
	fmt.Println("\n--- Function Types ---")

	// Fonksiyon tipini tanımlama
	type MathOp func(int, int) int

	var op MathOp
	op = func(a, b int) int { return a + b }
	fmt.Println("Add:", op(5, 3))

	op = func(a, b int) int { return a * b }
	fmt.Println("Multiply:", op(5, 3))

	// Calculator örneği
	fmt.Println("Calculator add:", calculate(10, 5, "add"))
	fmt.Println("Calculator sub:", calculate(10, 5, "sub"))
	fmt.Println("Calculator mul:", calculate(10, 5, "mul"))
}

// ==================== HELPER FUNCTIONS ====================

// Multiple return values
func addAndSubtract(a, b int) (int, int) {
	return a + b, a - b
}

// Error dönen fonksiyon (Go idiom)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("sıfıra bölme hatası")
	}
	return a / b, nil
}

// Named return values
func getRectangleInfo(width, height int) (w int, h int, area int) {
	w = width
	h = height
	area = width * height
	return // naked return - değişken isimleri zaten tanımlı
}

// Variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Closure - counter factory
func makeCounter() func() int {
	count := 0 // Bu değişken closure'da yaşar
	return func() int {
		count++
		return count
	}
}

// Closure - multiplier factory
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// Higher-order function - map
func mapInts(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}

// Higher-order function - filter
func filterInts(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// Defer örneği
func deferExample() {
	fmt.Println("Başladı")

	defer fmt.Println("Defer 1 - İlk eklenen, son çalışır")
	defer fmt.Println("Defer 2 - İkinci eklenen")
	defer fmt.Println("Defer 3 - Son eklenen, ilk çalışır")

	fmt.Println("Fonksiyon gövdesi")
	// Çıktı sırası: Başladı -> Gövde -> Defer 3 -> Defer 2 -> Defer 1
}

// Defer ile kaynak yönetimi simülasyonu
func processFile() {
	fmt.Println("Dosya açıldı")
	defer fmt.Println("Dosya kapatıldı") // Ne olursa olsun kapanacak

	fmt.Println("Dosya işleniyor...")
	// Hata olsa bile defer çalışır
}

// Calculator with function map
func calculate(a, b int, op string) int {
	operations := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
	}

	if f, ok := operations[op]; ok {
		return f(a, b)
	}
	return 0
}

// Bonus: strings.Builder ile verimli string birleştirme
func joinStrings(strs ...string) string {
	var builder strings.Builder
	for i, s := range strs {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(s)
	}
	return builder.String()
}

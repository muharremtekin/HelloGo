package essentials

import "fmt"

// Pointers in Go
// Pointer, bir değişkenin bellek adresini tutan değişkendir
// C#'taki ref/out parametrelerine benzer ama daha güçlü

func Pointers() {
	fmt.Println("========== POINTERS ==========")

	// ==================== TEMEL POINTER ====================
	fmt.Println("\n--- Temel Pointer ---")

	x := 42

	// & = adres operatörü (address of)
	p := &x // p, x'in adresini tutuyor

	fmt.Println("x'in değeri:", x)
	fmt.Println("x'in adresi:", &x)
	fmt.Println("p'nin değeri (adres):", p)

	// * = dereference operatörü (değere ulaş)
	fmt.Println("p'nin gösterdiği değer:", *p)

	// Pointer ile değer değiştirme
	*p = 100
	fmt.Println("x'in yeni değeri:", x) // 100 oldu!

	// ==================== POINTER SYNTAX ====================
	fmt.Println("\n--- Pointer Syntax ---")

	var a int = 10
	var ptr *int // int pointer (henüz nil)

	fmt.Println("ptr (nil):", ptr)

	ptr = &a
	fmt.Println("ptr (atandı):", ptr)
	fmt.Println("*ptr:", *ptr)

	// Kısa syntax
	b := 20
	bPtr := &b // otomatik tip: *int
	fmt.Printf("bPtr tipi: %T\n", bPtr)

	// ==================== NIL POINTER ====================
	fmt.Println("\n--- Nil Pointer ---")

	var nilPtr *int
	fmt.Println("Nil pointer:", nilPtr)
	fmt.Println("nil mi?:", nilPtr == nil)

	// DİKKAT: nil pointer dereference = PANIC!
	// fmt.Println(*nilPtr) // PANIC: runtime error

	// Güvenli kullanım - fonksiyon ile gösterelim
	fmt.Println("Nil pointer kontrolü:", safeDeref(nilPtr))
	fmt.Println("Normal pointer kontrolü:", safeDeref(&x))

	// ==================== POINTER VE FONKSİYONLAR ====================
	fmt.Println("\n--- Pointer ve Fonksiyonlar ---")

	// Go'da her şey "pass by value" - kopya geçirilir
	// Pointer ile orijinali değiştirebiliriz

	num := 5
	fmt.Println("Önce:", num)

	// Value ile - orijinal değişmez
	incrementValue(num)
	fmt.Println("incrementValue sonrası:", num) // hala 5

	// Pointer ile - orijinal değişir
	incrementPointer(&num)
	fmt.Println("incrementPointer sonrası:", num) // 6

	// ==================== POINTER VE STRUCT ====================
	fmt.Println("\n--- Pointer ve Struct ---")

	type User struct {
		Name string
		Age  int
	}

	user := User{Name: "Ali", Age: 25}
	fmt.Println("User:", user)

	// Struct pointer
	userPtr := &user

	// Go'da -> yok, . ile erişirsin (otomatik dereference)
	fmt.Println("Pointer ile Name:", userPtr.Name)  // (*userPtr).Name yazmaya gerek yok
	fmt.Println("Pointer ile Age:", userPtr.Age)

	// Pointer ile değiştir
	userPtr.Age = 30
	fmt.Println("Değiştirilmiş user:", user)

	// new() ile pointer oluşturma
	userPtr2 := new(User) // zero value ile *User döner
	userPtr2.Name = "Ayşe"
	userPtr2.Age = 28
	fmt.Println("new ile oluşturulan:", *userPtr2)

	// ==================== POINTER VE SLICE/MAP ====================
	fmt.Println("\n--- Pointer ve Slice/Map ---")

	// Slice ve Map zaten referans tipi gibi davranır
	// İçlerinde pointer var, bu yüzden fonksiyona geçince orijinal değişir

	slice := []int{1, 2, 3}
	fmt.Println("Slice önce:", slice)
	modifySlice(slice)
	fmt.Println("Slice sonra:", slice) // değişti!

	m := map[string]int{"a": 1}
	fmt.Println("Map önce:", m)
	modifyMap(m)
	fmt.Println("Map sonra:", m) // değişti!

	// Ama slice'ın kendisini değiştirmek istersen pointer lazım
	fmt.Println("\nSlice append örneği:")
	slice2 := []int{1, 2}
	fmt.Println("Slice2 önce:", slice2)
	appendToSlice(&slice2, 3, 4, 5)
	fmt.Println("Slice2 sonra:", slice2)

	// ==================== POINTER ARİTMETİĞİ ====================
	fmt.Println("\n--- Pointer Aritmetiği ---")

	// Go'da pointer aritmetiği YOK (C'den farklı)
	// p++ veya p+1 yapamazsın
	// Bu güvenlik için kasıtlı kaldırılmış

	arr := [3]int{10, 20, 30}
	arrPtr := &arr[0]
	fmt.Println("İlk eleman:", *arrPtr)
	// arrPtr++ // HATA: invalid operation

	// Eğer gerçekten lazımsa unsafe paketi var (önerilmez)

	// ==================== POINTER TO POINTER ====================
	fmt.Println("\n--- Pointer to Pointer ---")

	val := 100
	p1 := &val   // *int
	p2 := &p1    // **int (pointer to pointer)

	fmt.Println("val:", val)
	fmt.Println("p1 (val'in adresi):", p1)
	fmt.Println("p2 (p1'in adresi):", p2)
	fmt.Println("*p1:", *p1)
	fmt.Println("**p2:", **p2)

	// ==================== PRATİK KULLANIM ====================
	fmt.Println("\n--- Pratik: Optional/Nullable Değer ---")

	// Go'da nullable int yok, pointer kullanılır
	timeout := 30
	config1 := Config{Timeout: &timeout, Name: "Server1"}
	config2 := Config{Timeout: nil, Name: "Server2"} // timeout ayarlanmamış

	printConfig(config1)
	printConfig(config2)

	// ==================== WHEN TO USE POINTERS ====================
	fmt.Println("\n--- Ne Zaman Pointer Kullanmalı? ---")
	fmt.Println("KULLAN:")
	fmt.Println("1. Fonksiyonda orijinal değeri değiştirmek istiyorsan")
	fmt.Println("2. Büyük struct'ları kopyalamamak için (performans)")
	fmt.Println("3. nil/optional değer ifade etmek için")
	fmt.Println("4. Method receiver'da struct'ı değiştirmek için")
	fmt.Println("\nKULLANMA:")
	fmt.Println("1. Küçük struct'lar (kopyalamak daha hızlı olabilir)")
	fmt.Println("2. Sadece okuma yapıyorsan gerek yok")
	fmt.Println("3. Slice, map, channel zaten referans - pointer gereksiz")
}

// ==================== HELPER FUNCTIONS ====================

// Value ile geçme - orijinal değişmez
func incrementValue(n int) {
	n++
	fmt.Println("Fonksiyon içinde (value):", n)
}

// Pointer ile geçme - orijinal değişir
func incrementPointer(n *int) {
	*n++
	fmt.Println("Fonksiyon içinde (pointer):", *n)
}

// Slice içeriğini değiştir (pointer gerekmez)
func modifySlice(s []int) {
	s[0] = 999
}

// Map içeriğini değiştir (pointer gerekmez)
func modifyMap(m map[string]int) {
	m["a"] = 999
}

// Slice'a eleman ekle (pointer gerekir)
func appendToSlice(s *[]int, vals ...int) {
	*s = append(*s, vals...)
}

// Güvenli dereference
func safeDeref(p *int) string {
	if p != nil {
		return fmt.Sprintf("değer: %d", *p)
	}
	return "nil pointer"
}

// Config struct
type Config struct {
	Timeout *int
	Name    string
}

// Config yazdır (nil kontrolü)
func printConfig(c Config) {
	if c.Timeout != nil {
		fmt.Printf("Config %s: timeout=%d\n", c.Name, *c.Timeout)
	} else {
		fmt.Printf("Config %s: timeout=default\n", c.Name)
	}
}

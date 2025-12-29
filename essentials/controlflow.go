package essentials

import "fmt"

// Control Flow - If, Switch, For
// Go'da kontrol yapıları C#'a benzer ama parantez yok

func ControlFlow() {
	fmt.Println("========== CONTROL FLOW ==========")

	// ==================== IF / ELSE ====================
	fmt.Println("\n--- If / Else ---")

	age := 18

	// Basit if (parantez YOK, süslü parantez ZORUNLU)
	if age >= 18 {
		fmt.Println("Reşit")
	}

	// if-else
	if age >= 18 {
		fmt.Println("Oy kullanabilir")
	} else {
		fmt.Println("Oy kullanamaz")
	}

	// if-else if-else
	score := 75
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	// If with statement (Go'ya özgü!)
	// Değişkeni if içinde tanımlayıp hemen kullanabilirsin
	fmt.Println("\n--- If with Statement ---")
	if num := 10; num > 5 {
		fmt.Println(num, "5'ten büyük")
	}
	// num burada erişilemez - sadece if scope'unda geçerli

	// Map ile kullanım (comma ok idiom)
	ages := map[string]int{"Ali": 25, "Veli": 0}
	if val, ok := ages["Ali"]; ok {
		fmt.Println("Ali'nin yaşı:", val)
	}

	// ==================== SWITCH ====================
	fmt.Println("\n--- Switch ---")

	day := 3

	// Basit switch (break yazmaya GEREK YOK - otomatik)
	switch day {
	case 1:
		fmt.Println("Pazartesi")
	case 2:
		fmt.Println("Salı")
	case 3:
		fmt.Println("Çarşamba")
	case 4:
		fmt.Println("Perşembe")
	case 5:
		fmt.Println("Cuma")
	case 6, 7: // Birden fazla değer
		fmt.Println("Hafta sonu")
	default:
		fmt.Println("Geçersiz gün")
	}

	// Switch with statement
	fmt.Println("\n--- Switch with Statement ---")
	switch today := 5; today {
	case 1, 2, 3, 4, 5:
		fmt.Println("Hafta içi")
	case 6, 7:
		fmt.Println("Hafta sonu")
	}

	// Switch without expression (if-else yerine kullanılabilir)
	fmt.Println("\n--- Switch without Expression ---")
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Günaydın")
	case hour < 18:
		fmt.Println("İyi günler")
	default:
		fmt.Println("İyi akşamlar")
	}

	// Fallthrough - bir sonraki case'e düşmek için
	fmt.Println("\n--- Fallthrough ---")
	num := 1
	switch num {
	case 1:
		fmt.Println("Bir")
		fallthrough // Bir sonraki case'e devam et
	case 2:
		fmt.Println("İki")
		// fallthrough olmadığı için burada durur
	case 3:
		fmt.Println("Üç")
	}

	// Type switch (interface ile)
	fmt.Println("\n--- Type Switch ---")
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("%v bir int\n", v)
		case string:
			fmt.Printf("%v bir string\n", v)
		case bool:
			fmt.Printf("%v bir bool\n", v)
		default:
			fmt.Printf("%v bilinmeyen tip\n", v)
		}
	}
	checkType(42)
	checkType("merhaba")
	checkType(true)

	// ==================== FOR ====================
	fmt.Println("\n--- For Döngüsü ---")

	// 1. Klasik for (C tarzı)
	fmt.Println("\nKlasik for:")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. While tarzı for (Go'da while yok!)
	fmt.Println("\nWhile tarzı:")
	count := 0
	for count < 3 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// 3. Sonsuz döngü
	fmt.Println("\nSonsuz döngü (break ile):")
	counter := 0
	for {
		if counter >= 3 {
			break
		}
		fmt.Print(counter, " ")
		counter++
	}
	fmt.Println()

	// 4. Range ile for - Array/Slice
	fmt.Println("\nRange ile slice:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// Sadece value (index'i atla)
	fmt.Println("\nSadece value:")
	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Println()

	// Sadece index
	fmt.Println("\nSadece index:")
	for index := range numbers {
		fmt.Print(index, " ")
	}
	fmt.Println()

	// 5. Range ile for - Map
	fmt.Println("\nRange ile map:")
	colors := map[string]string{"red": "kırmızı", "blue": "mavi"}
	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}

	// 6. Range ile for - String (rune olarak döner)
	fmt.Println("\nRange ile string:")
	for index, char := range "Go!" {
		fmt.Printf("index: %d, char: %c\n", index, char)
	}

	// 7. Continue - döngüde atla
	fmt.Println("\nContinue örneği (çift sayıları atla):")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Çift sayıları atla
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 8. Break with label - iç içe döngüden çıkış
	fmt.Println("\nBreak with label:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // Dış döngüden çık
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()
}

package essentials

import "fmt"

// Maps in Go
// Map, key-value çiftlerini tutan veri yapısı
// C#'taki Dictionary<TKey, TValue> gibi

func Maps() {
	fmt.Println("========== MAPS ==========")

	// 1. Map oluşturma yöntemleri
	fmt.Println("\n--- Map Oluşturma ---")

	// Yöntem 1: make() ile
	ages := make(map[string]int)
	ages["Ali"] = 25
	ages["Ayşe"] = 30
	ages["Mehmet"] = 28
	fmt.Println("Ages:", ages)

	// Yöntem 2: Literal (direkt değerlerle)
	colors := map[string]string{
		"red":   "kırmızı",
		"blue":  "mavi",
		"green": "yeşil",
	}
	fmt.Println("Colors:", colors)

	// Yöntem 3: Boş map (nil map - DİKKAT!)
	var emptyMap map[string]int // nil, yazma yapılamaz!
	fmt.Println("Empty map:", emptyMap, "nil mi?:", emptyMap == nil)
	// emptyMap["test"] = 1 // PANIC! nil map'e yazılamaz

	// 2. Değer okuma ve yazma
	fmt.Println("\n--- Okuma & Yazma ---")
	fmt.Println("Ali'nin yaşı:", ages["Ali"])

	// Olmayan key okuma - zero value döner
	fmt.Println("Veli'nin yaşı:", ages["Veli"]) // 0 döner (int'in zero value'su)

	// Key var mı kontrolü (comma ok idiom)
	age, exists := ages["Veli"]
	if exists {
		fmt.Println("Veli'nin yaşı:", age)
	} else {
		fmt.Println("Veli bulunamadı!")
	}

	// Kısa hali
	if age, ok := ages["Ali"]; ok {
		fmt.Println("Ali bulundu, yaşı:", age)
	}

	// 3. Eleman silme
	fmt.Println("\n--- Eleman Silme ---")
	fmt.Println("Silmeden önce:", ages)
	delete(ages, "Mehmet")
	fmt.Println("Sildikten sonra:", ages)

	// Olmayan key silme - hata vermez
	delete(ages, "Veli") // sorun yok

	// 4. Map üzerinde döngü
	fmt.Println("\n--- Range ile Döngü ---")
	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}

	// Sadece key'ler
	fmt.Println("\nSadece key'ler:")
	for key := range colors {
		fmt.Println(key)
	}

	// DİKKAT: Map'te sıra garantisi YOK! Her çalıştırmada farklı sıra olabilir

	// 5. Map uzunluğu
	fmt.Println("\n--- Map Uzunluğu ---")
	fmt.Println("Colors uzunluğu:", len(colors))

	// 6. Nested map (map içinde map)
	fmt.Println("\n--- Nested Map ---")
	users := map[string]map[string]string{
		"user1": {
			"name":  "Ali",
			"email": "ali@test.com",
		},
		"user2": {
			"name":  "Ayşe",
			"email": "ayse@test.com",
		},
	}
	fmt.Println("Users:", users)
	fmt.Println("User1 name:", users["user1"]["name"])

	// Nested map'e yeni eleman eklerken dikkat!
	// İç map'i önce oluşturmalısın
	users["user3"] = make(map[string]string)
	users["user3"]["name"] = "Zeynep"
	users["user3"]["email"] = "zeynep@test.com"
	fmt.Println("User3:", users["user3"])

	// 7. Map with struct (daha yaygın kullanım)
	fmt.Println("\n--- Map with Struct ---")

	type User struct {
		Name  string
		Email string
		Age   int
	}

	userMap := map[string]User{
		"admin": {Name: "Admin", Email: "admin@test.com", Age: 35},
		"guest": {Name: "Guest", Email: "guest@test.com", Age: 0},
	}

	fmt.Println("Admin:", userMap["admin"])
	fmt.Println("Admin name:", userMap["admin"].Name)

	// Struct field'ını direkt değiştiremezsin!
	// userMap["admin"].Age = 36 // HATA! cannot assign

	// Çözüm: Tüm struct'ı güncelle
	admin := userMap["admin"]
	admin.Age = 36
	userMap["admin"] = admin
	fmt.Println("Updated admin:", userMap["admin"])

	// Veya pointer kullan
	userMapPtr := map[string]*User{
		"admin": {Name: "Admin", Email: "admin@test.com", Age: 35},
	}
	userMapPtr["admin"].Age = 40 // Bu çalışır!
	fmt.Println("Pointer admin age:", userMapPtr["admin"].Age)

	// 8. Map as set (küme olarak kullanım)
	fmt.Println("\n--- Map as Set ---")

	// Go'da Set yok, map[T]bool veya map[T]struct{} kullanılır
	visited := map[string]bool{
		"ankara":   true,
		"istanbul": true,
		"izmir":    true,
	}

	city := "ankara"
	if visited[city] {
		fmt.Println(city, "ziyaret edildi")
	}

	// Bellek açısından daha verimli: struct{} (0 byte)
	visitedEfficient := map[string]struct{}{
		"ankara":   {},
		"istanbul": {},
	}

	if _, ok := visitedEfficient["ankara"]; ok {
		fmt.Println("ankara set'te var")
	}
}

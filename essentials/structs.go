package essentials

import "fmt"

// Structs in Go
// Struct, birden fazla farklı tipteki veriyi bir arada tutan yapılardır
// C#'taki class'lara benzer ama Go'da class yok, struct var

type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

// Struct'lara method ekleyebiliriz
// (p Person) -> receiver, bu method Person struct'ına ait demek
func (p Person) Greet() string {
	return fmt.Sprintf("Merhaba, ben %s ve %d yaşındayım!", p.Name, p.Age)
}

// Pointer receiver - struct'ı değiştirmek istediğimizde kullanılır
func (p *Person) HaveBirthday() {
	p.Age++
}

// Nested struct örneği
type Address struct {
	City    string
	Country string
}

type Employee struct {
	Person   // Embedded struct (inheritance gibi)
	JobTitle string
	Salary   float64
	Address  Address // Nested struct
}

// Employee için method
func (e Employee) GetInfo() string {
	return fmt.Sprintf("%s - %s, Maaş: %.2f TL, Şehir: %s",
		e.Name, e.JobTitle, e.Salary, e.Address.City)
}

func Structs() {
	fmt.Println("========== STRUCTS ==========")

	// 1. Struct oluşturma yöntemleri
	fmt.Println("\n--- Struct Oluşturma ---")

	// Yöntem 1: Field sırasıyla
	person1 := Person{"Ali", 25, "ali@email.com", false}
	fmt.Println("Person 1:", person1)

	// Yöntem 2: Field isimleriyle (önerilen yöntem)
	person2 := Person{
		Name:    "Ayşe",
		Age:     30,
		Email:   "ayse@email.com",
		IsAdmin: true,
	}
	fmt.Println("Person 2:", person2)

	// Yöntem 3: Boş struct, sonra değer atama
	var person3 Person
	person3.Name = "Mehmet"
	person3.Age = 28
	fmt.Println("Person 3:", person3)

	// 2. Struct field'larına erişim
	fmt.Println("\n--- Field Erişimi ---")
	fmt.Println("İsim:", person2.Name)
	fmt.Println("Yaş:", person2.Age)
	fmt.Println("Admin mi?:", person2.IsAdmin)

	// 3. Method kullanımı
	fmt.Println("\n--- Method Kullanımı ---")
	fmt.Println(person1.Greet())
	fmt.Println(person2.Greet())

	// 4. Pointer receiver ile değer değiştirme
	fmt.Println("\n--- Pointer Receiver ---")
	fmt.Println("Yaş (önce):", person1.Age)
	person1.HaveBirthday()
	fmt.Println("Yaş (sonra):", person1.Age)

	// 5. Nested ve Embedded struct
	fmt.Println("\n--- Nested & Embedded Struct ---")
	employee := Employee{
		Person:   Person{Name: "Zeynep", Age: 35, Email: "zeynep@company.com", IsAdmin: true},
		JobTitle: "Software Engineer",
		Salary:   75000,
		Address: Address{
			City:    "İstanbul",
			Country: "Türkiye",
		},
	}

	fmt.Println(employee.GetInfo())
	// Embedded struct sayesinde Person field'larına direkt erişebiliriz
	fmt.Println("Çalışan adı:", employee.Name) // employee.Person.Name yazmaya gerek yok
	fmt.Println(employee.Greet())              // Person'ın methodunu da kullanabiliriz

	// 6. Struct pointer
	fmt.Println("\n--- Struct Pointer ---")
	personPtr := &person1
	fmt.Println("Pointer ile isim:", personPtr.Name)
	personPtr.Age = 50
	fmt.Println("Değiştirilmiş yaş:", person1.Age)

	// 7. Anonymous struct (tek seferlik kullanım için)
	fmt.Println("\n--- Anonymous Struct ---")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)
}

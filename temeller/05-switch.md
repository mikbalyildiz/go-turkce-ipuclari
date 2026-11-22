# 5. switch

Go'da `switch` her case'ten sonra otomatik durur, `break` yazmaya gerek yoktur. İfadesiz `switch`, uzun `if-else` zincirlerinin okunaklı bir alternatifidir.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

func main() { // program girişi
	gun := 6 // haftanın günü (1 = Pazartesi)

	switch gun { // gun değerine göre dallan
	case 1, 2, 3, 4, 5: // bir case'te birden fazla değer olabilir
		fmt.Println("Hafta içi") // eşleşirse çalışır
	case 6, 7: // 6 ya da 7 ise
		fmt.Println("Hafta sonu") // bu çalışır; break yazmaya gerek yok
	default: // hiçbiri eşleşmezse
		fmt.Println("Geçersiz gün") // varsayılan dal
	} // switch sonu

	sicaklik := 28 // örnek sıcaklık
	switch {       // ifadesiz switch, uzun if-else zinciri yerine kullanılır
	case sicaklik < 10: // ilk doğru olan case çalışır
		fmt.Println("Soğuk") // 10'dan küçükse
	case sicaklik < 25: // 10 ile 24 arası
		fmt.Println("Ilık") // ılık hava
	default: // 25 ve üstü
		fmt.Println("Sıcak") // 28 için burası çalışır
	} // switch sonu

	switch x := 1; x { // switch içinde de kısa tanımlama yapılabilir
	case 1: // x 1 ise
		fmt.Println("bir") // yazdır
		fallthrough        // bir sonraki case'i de koşulsuz çalıştır (nadiren kullanılır)
	case 2: // fallthrough ile buraya da girer
		fmt.Println("iki") // yazdır
	} // switch sonu
} // program sonu
```

**Çıktı**

```text
Hafta sonu
Sıcak
bir
iki
```

---

[← for Döngüsü (ve while)](04-for-dongusu.md) · [İçindekiler](../README.md#go-temelleri) · [Fonksiyonlar →](06-fonksiyonlar.md)

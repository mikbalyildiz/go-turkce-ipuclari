# 10. Goroutine ve Kanallar

`go` anahtar kelimesiyle bir fonksiyon eşzamanlı çalıştırılır. Goroutine'ler arasında veri **kanal** (`chan`) ile taşınır. Hepsinin bitmesini beklemek için `sync.WaitGroup` kullanılır.

```go
package main // çalıştırılabilir program

import ( // paketler
	"fmt"  // yazdırma paketi
	"sync" // WaitGroup için
) // import sonu

func kareAl(n int, sonuc chan<- int) { // chan<- yalnızca gönderim yapılabilen kanal
	sonuc <- n * n // hesaplanan değeri kanala gönder
} // fonksiyon sonu

func main() { // program girişi
	// 1) Kanal ile sonuç toplama
	kanal := make(chan int)   // int taşıyan kanal oluştur
	for i := 1; i <= 3; i++ { // 3 iş başlat
		go kareAl(i, kanal) // go: fonksiyonu ayrı bir goroutine'de eşzamanlı çalıştır
	} // döngü sonu
	toplam := 0              // sonuçların toplamı
	for i := 0; i < 3; i++ { // 3 sonuç bekle
		toplam += <-kanal // <-kanal: kanaldan bir değer gelene kadar bekler ve alır
	} // döngü sonu
	fmt.Println("Kareler toplamı:", toplam) // 1+4+9 = 14

	// 2) WaitGroup ile tüm işlerin bitmesini bekleme
	var wg sync.WaitGroup     // bekleme sayacı
	for i := 1; i <= 3; i++ { // 3 goroutine
		wg.Add(1)        // sayacı bir artır
		go func(n int) { // değeri parametre olarak geçir
			defer wg.Done() // iş bitince sayacı azalt
			_ = n * 2       // burada asıl iş yapılır
		}(i) // i'nin değerini gönder
	} // döngü sonu
	wg.Wait()                      // hepsi bitene kadar bekle
	fmt.Println("Tüm işler bitti") // güvenle devam et
} // program sonu
```

**Çıktı**

```text
Kareler toplamı: 14
Tüm işler bitti
```

---

[← Hata Yönetimi](09-hata-yonetimi.md) · [İçindekiler](../README.md#go-temelleri)

# 6. Fonksiyonlar

Fonksiyonlar `func` ile tanımlanır. Go'nun ayırt edici özelliği bir fonksiyonun **birden fazla değer döndürebilmesidir**. `...` ile değişken sayıda argüman alınabilir.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

func topla(a int, b int) int { // iki int alır, bir int döndürür
	return a + b // sonucu geri gönder
} // fonksiyon sonu

func bolVeKalan(a, b int) (int, int) { // aynı tipteki parametreler kısaltılabilir, iki değer döner
	return a / b, a % b // bölüm ve kalanı birlikte döndür
} // fonksiyon sonu

func ortalama(sayilar ...float64) float64 { // ... ile istenen sayıda argüman alınır
	toplam := 0.0               // toplamı tut
	for _, s := range sayilar { // _ ile kullanılmayan indeksi yok say
		toplam += s // her sayıyı ekle
	} // döngü sonu
	return toplam / float64(len(sayilar)) // ortalamayı döndür
} // fonksiyon sonu

func main() { // program girişi
	fmt.Println(topla(3, 4))          // 7
	bolum, kalan := bolVeKalan(17, 5) // iki dönüş değerini iki değişkene al
	fmt.Println(bolum, kalan)         // 3 2
	fmt.Println(ortalama(70, 80, 90)) // 80

	kare := func(x int) int { return x * x } // isimsiz fonksiyonu değişkene ata
	fmt.Println(kare(5))                     // 25
} // program sonu
```

**Çıktı**

```text
7
3 2
80
25
```

---

[← switch](05-switch.md) · [İçindekiler](../README.md#go-temelleri) · [Slice ve Map →](07-slice-ve-map.md)

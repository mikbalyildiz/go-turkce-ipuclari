# 7. Slice ve Map

Dizilerin (array) boyutu sabittir. Pratikte neredeyse her zaman boyutu büyüyebilen **slice** kullanılır. **map** ise anahtar-değer çiftlerini saklayan sözlük yapısıdır.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

func main() { // program girişi
	// Dizi (array): boyutu sabittir, pek sık kullanılmaz
	var gunler [3]string     // 3 elemanlı dizi
	gunler[0] = "Pzt"        // ilk elemana değer ata
	fmt.Println(len(gunler)) // 3

	// Slice: boyutu büyüyebilen liste, Go'da en çok kullanılan yapı
	notlar := []int{70, 85}          // başlangıç değerleriyle slice
	notlar = append(notlar, 90)      // sona eleman ekle, sonucu geri ata
	fmt.Println(notlar, len(notlar)) // [70 85 90] 3
	fmt.Println(notlar[1:3])         // dilimleme: 1. indeksten 3'e kadar (3 hariç)

	// Map: anahtar-değer sözlüğü
	baskent := map[string]string{ // string anahtar, string değer
		"Türkiye": "Ankara", // ilk kayıt
		"Japonya": "Tokyo",  // ikinci kayıt
	} // map sonu
	baskent["İtalya"] = "Roma"      // yeni kayıt ekle
	fmt.Println(baskent["Türkiye"]) // Ankara

	sehir, varMi := baskent["Mısır"] // ikinci değer anahtarın olup olmadığını söyler
	fmt.Println(sehir == "", varMi)  // true false (yoksa sıfır değer döner)

	delete(baskent, "Japonya") // kaydı sil
	fmt.Println(len(baskent))  // 2
} // program sonu
```

**Çıktı**

```text
3
[70 85 90] 3
[85 90]
Ankara
true false
2
```

---

[← Fonksiyonlar](06-fonksiyonlar.md) · [İçindekiler](../README.md#go-temelleri) · [Struct ve Metotlar →](08-struct-ve-metot.md)

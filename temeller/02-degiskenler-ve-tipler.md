# 2. Değişkenler ve Tipler

Go'da değişken `var` ile ya da kısaca `:=` ile tanımlanır. Temel tipler `int`, `float64`, `string` ve `bool`'dur. Değer verilmeyen değişkenler **sıfır değer** alır (0, "", false). Farklı tipler arasında otomatik dönüşüm yoktur.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

const pi = 3.14 // const: değeri asla değişmeyen sabit

func main() { // program girişi
	var ad string = "Ayşe" // var ile tipi açıkça belirterek tanımlama
	var yas int            // değer verilmezse sıfır değer alır: int için 0
	yas = 20               // sonradan değer atama
	boy := 1.68            // := kısa tanımlama, tip otomatik çıkarılır (float64)
	ogrenci := true        // bool: true ya da false

	fmt.Println(ad, yas, boy, ogrenci)     // değerleri yan yana yazdır
	fmt.Printf("%s %d yaşında\n", ad, yas) // %s metin, %d tam sayı yer tutucusu
	fmt.Printf("boy'un tipi: %T\n", boy)   // %T değişkenin tipini yazar
	fmt.Println("Daire alanı:", pi*2*2)    // sabit hesaplamada kullanılabilir

	var x int = 7       // tam sayı
	y := float64(x) / 2 // Go otomatik tip dönüşümü yapmaz, açıkça çevirmek gerekir
	fmt.Println(y)      // 3.5
} // program sonu
```

**Çıktı**

```text
Ayşe 20 1.68 true
Ayşe 20 yaşında
boy'un tipi: float64
Daire alanı: 12.56
3.5
```

---

[← Merhaba Dünya](01-merhaba-dunya.md) · [İçindekiler](../README.md#go-temelleri) · [if / else →](03-if-else.md)

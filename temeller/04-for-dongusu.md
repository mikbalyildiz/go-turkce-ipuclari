# 4. for Döngüsü (ve while)

Go'da tek döngü anahtar kelimesi `for`'dur. **`while` yoktur**, onun yerine sadece koşul yazılan `for` kullanılır. `for` tek başına sonsuz döngü kurar, `range` ise listeler üzerinde gezinir.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

func main() { // program girişi
	// 1) Klasik for: başlangıç; koşul; artış
	for i := 1; i <= 3; i++ { // i 1'den başlar, 3'e kadar birer artar
		fmt.Println("tur", i) // her turda çalışır
	} // döngü sonu

	// 2) while gibi for: Go'da while yoktur, sadece koşul yazılır
	sayac := 3      // geri sayım başlangıcı
	for sayac > 0 { // koşul doğru olduğu sürece döner
		fmt.Println("geri sayım:", sayac) // değeri yazdır
		sayac--                           // koşulu sonunda bozacak adım
	} // while benzeri döngü sonu

	// 3) Sonsuz döngü + break
	n := 0 // sayaç
	for {  // koşulsuz for sonsuza kadar döner
		n++         // her turda bir artır
		if n == 2 { // 2'yi atla
			continue // döngünün kalanını geç, sonraki tura başla
		} // if sonu
		if n > 4 { // çıkış koşulu
			break // döngüden tamamen çık
		} // if sonu
		fmt.Println("n =", n) // 1, 3, 4 yazılır
	} // sonsuz döngü sonu

	// 4) range: dizi, slice, map ve string üzerinde gezinme
	meyveler := []string{"elma", "armut"} // örnek liste
	for indeks, meyve := range meyveler { // indeks ve değer birlikte gelir
		fmt.Println(indeks, meyve) // 0 elma, 1 armut
	} // range sonu
} // program sonu
```

**Çıktı**

```text
tur 1
tur 2
tur 3
geri sayım: 3
geri sayım: 2
geri sayım: 1
n = 1
n = 3
n = 4
0 elma
1 armut
```

▶ Çalıştır: `go run ./ornekler/04-for-dongusu`

---

[← if / else](03-if-else.md) · [İçindekiler](../README.md#go-temelleri) · [switch →](05-switch.md)

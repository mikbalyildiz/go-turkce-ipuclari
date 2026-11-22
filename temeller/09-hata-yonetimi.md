# 9. Hata Yönetimi

Go'da `try/catch` yoktur. Hatalar fonksiyonun son dönüş değeri olarak `error` tipiyle döner. Hatayı her zaman `if err != nil` ile kontrol edin.

```go
package main // çalıştırılabilir program

import ( // birden fazla paket parantezle içe aktarılır
	"errors"  // yeni hata oluşturmak için
	"fmt"     // yazdırma paketi
	"strconv" // metin-sayı dönüşümü
) // import sonu

func bol(a, b float64) (float64, error) { // son dönüş değeri olarak error döndürmek Go geleneğidir
	if b == 0 { // sıfıra bölme kontrolü
		return 0, errors.New("sıfıra bölünemez") // hata durumunda anlamlı mesaj
	} // if sonu
	return a / b, nil // nil: hata yok demektir
} // fonksiyon sonu

func main() { // program girişi
	sonuc, err := bol(10, 0) // sonucu ve hatayı birlikte al
	if err != nil {          // her zaman önce hatayı kontrol et
		fmt.Println("Hata:", err) // Hata: sıfıra bölünemez
	} else { // hata yoksa
		fmt.Println(sonuc) // sonucu kullan
	} // if sonu

	sayi, err := strconv.Atoi("42") // metni int'e çevir, hata dönebilir
	if err != nil {                 // dönüşüm başarısızsa
		fmt.Println("Geçersiz sayı") // kullanıcıya bildir
		return                       // fonksiyondan çık
	} // if sonu
	fmt.Println(sayi + 1) // 43
} // program sonu
```

**Çıktı**

```text
Hata: sıfıra bölünemez
43
```

---

[← Struct ve Metotlar](08-struct-ve-metot.md) · [İçindekiler](../README.md#go-temelleri) · [Goroutine ve Kanallar →](10-goroutine-ve-kanal.md)

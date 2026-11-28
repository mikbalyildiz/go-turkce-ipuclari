# 8. Struct ve Metotlar

Go'da sınıf yoktur. Bunun yerine alanları gruplayan `struct` ve bu tiplere bağlanan **metotlar** vardır. Nesneyi değiştiren metotlar işaretçi alıcı (`*Tip`) kullanır.

```go
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

type Ogrenci struct { // struct: ilgili alanları bir arada tutan özel tip
	Ad  string // öğrencinin adı
	Not int    // öğrencinin notu
} // struct sonu

func (o Ogrenci) Durum() string { // metot: Ogrenci tipine bağlı fonksiyon (değer alıcı)
	if o.Not >= 50 { // alanlara nokta ile erişilir
		return o.Ad + " geçti" // metin birleştirme
	} // if sonu
	return o.Ad + " kaldı" // koşul tutmazsa
} // metot sonu

func (o *Ogrenci) NotEkle(puan int) { // işaretçi alıcı: asıl nesneyi değiştirir
	o.Not += puan // orijinal öğrencinin notu artar
} // metot sonu

func main() { // program girişi
	ali := Ogrenci{Ad: "Ali", Not: 45} // alan adlarıyla nesne oluştur
	fmt.Println(ali.Durum())           // Ali kaldı
	ali.NotEkle(10)                    // Go &ali'yi otomatik alır
	fmt.Println(ali.Not, ali.Durum())  // 55 Ali geçti
	fmt.Printf("%+v\n", ali)           // %+v alan adlarıyla birlikte yazdırır
} // program sonu
```

**Çıktı**

```text
Ali kaldı
55 Ali geçti
{Ad:Ali Not:55}
```

▶ Çalıştır: `go run ./ornekler/08-struct-ve-metot`

---

[← Slice ve Map](07-slice-ve-map.md) · [İçindekiler](../README.md#go-temelleri) · [Hata Yönetimi →](09-hata-yonetimi.md)

# Go Türkçe İpuçları 🐹

Go'da en sık yapılan **10 hata** ve doğru kullanımları. Her madde için kısa bir **yanlış** ve **doğru** örnek var, kodun her satırı Türkçe açıklamalı.

> Öğrenciler, Go'ya yeni başlayanlar ve kod inceleme yapanlar için hazırlanmış kısa bir başvuru kaynağı.

## İçindekiler

1. [nil map'e yazmak](#1-nil-mape-yazmak)
2. [Döngü değişkenini goroutine içinde kullanmak](#2-döngü-değişkenini-goroutine-içinde-kullanmak)
3. [main'in goroutine'leri beklemeyeceğini unutmak](#3-mainin-goroutineleri-beklemeyeceğini-unutmak)
4. [defer sırasını yanlış bilmek](#4-defer-sırasını-yanlış-bilmek)
5. [append sonrası slice'ların aynı diziyi paylaşması](#5-append-sonrası-sliceların-aynı-diziyi-paylaşması)
6. [Hataları yok saymak](#6-hataları-yok-saymak)
7. [interface nil tuzağı](#7-interface-nil-tuzağı)
8. [range ile kopya üzerinde değişiklik yapmak](#8-range-ile-kopya-üzerinde-değişiklik-yapmak)
9. [map'e aynı anda birden fazla goroutine'den yazmak](#9-mape-aynı-anda-birden-fazla-goroutineden-yazmak)
10. [Türkçe karakterli string'i byte ile işlemek](#10-türkçe-karakterli-stringi-byte-ile-işlemek)

---

## 1. nil map'e yazmak

`var m map[string]int` ile tanımlanan map `nil` olur. Okumak sorun çıkarmaz ama yazmak programı çökertir.

❌ Yanlış

```go
var notlar map[string]int // map tanımlandı ama oluşturulmadı, değeri nil
notlar["ali"] = 90        // panic: assignment to entry in nil map
```

✅ Doğru

```go
notlar := make(map[string]int) // make ile map bellekte oluşturulur
notlar["ali"] = 90             // artık güvenle yazılabilir
```

---

## 2. Döngü değişkenini goroutine içinde kullanmak

**Go 1.21 ve öncesinde** döngü değişkeni tüm turlarda aynı değişkendir. Goroutine'ler çalıştığında döngü bitmiş olabilir, bu yüzden hepsi son değeri görebilir. Go 1.22 ile her tur yeni bir değişken oluşturur ve sorun ortadan kalkar. Yine de eski sürümlerle çalışan kodlarda bu hataya sık rastlanır.

❌ Yanlış (Go ≤ 1.21)

```go
for i := 0; i < 3; i++ { // i tüm turlarda aynı değişken
	go func() {           // goroutine i'yi dışarıdan yakalıyor
		fmt.Println(i)    // çoğu zaman 3 3 3 basar
	}()                   // fonksiyon hemen çağrılıyor
}
```

✅ Doğru (her sürümde çalışır)

```go
for i := 0; i < 3; i++ { // döngü her zamanki gibi
	go func(n int) {      // değeri parametre olarak al
		fmt.Println(n)    // her goroutine kendi kopyasını basar
	}(i)                  // i'nin o anki değeri gönderilir
}
```

---

## 3. main'in goroutine'leri beklemeyeceğini unutmak

`main` fonksiyonu bitince program kapanır ve çalışmakta olan goroutine'ler de yarıda kesilir.

❌ Yanlış

```go
func main() {                        // program buradan başlar
	go fmt.Println("merhaba")        // goroutine başlatıldı
}                                    // main hemen bitti, çıktı görünmeyebilir
```

✅ Doğru

```go
func main() {                        // program buradan başlar
	var wg sync.WaitGroup            // bekleme sayacı
	wg.Add(1)                        // bir goroutine bekleneceğini bildir
	go func() {                      // goroutine başlat
		defer wg.Done()              // iş bitince sayacı bir azalt
		fmt.Println("merhaba")       // asıl iş
	}()                              // fonksiyon hemen çağrılıyor
	wg.Wait()                        // sayaç sıfırlanana kadar bekle
}                                    // artık güvenle bitebilir
```

---

## 4. defer sırasını yanlış bilmek

`defer` ile ertelenen çağrılar **son giren ilk çıkar** (LIFO) sırasıyla çalışır. Ayrıca `defer` fonksiyondan çıkarken çalışır, döngünün her turunun sonunda değil.

❌ Yanlış beklenti

```go
for i := 1; i <= 3; i++ { // 1, 2, 3 için döner
	defer fmt.Println(i)   // "1 2 3" beklenir ama "3 2 1" basılır
}
```

❌ Döngüde dosya açıp defer ile kapatmak

```go
for _, ad := range dosyalar { // çok sayıda dosya dolaşılıyor
	f, _ := os.Open(ad)        // dosya açılıyor
	defer f.Close()            // fonksiyon bitene kadar hiçbiri kapanmaz
}
```

✅ Doğru

```go
for _, ad := range dosyalar {       // her dosya için
	func() {                         // her tura ayrı bir fonksiyon kapsamı
		f, err := os.Open(ad)        // dosyayı aç
		if err != nil {              // açılamadıysa
			return                   // bu dosyayı atla
		}
		defer f.Close()              // bu iç fonksiyon bitince kapanır
		// ... dosyayı işle           // asıl iş burada yapılır
	}()                              // iç fonksiyon hemen çağrılıyor
}
```

---

## 5. append sonrası slice'ların aynı diziyi paylaşması

Bir slice'ı başka bir slice'tan kestiğinizde ikisi aynı alttaki diziyi kullanır. Kapasite yeterliyse `append` diğer slice'ın verisinin üzerine yazar.

❌ Yanlış

```go
a := []int{1, 2, 3, 4} // len 4, cap 4
b := a[:2]             // b = [1 2], ama kapasitesi 4 ve a ile aynı diziyi kullanıyor
b = append(b, 99)      // yer olduğu için a[2]'nin üzerine yazar
fmt.Println(a)         // [1 2 99 4]  (a beklenmedik şekilde değişti)
```

✅ Doğru

```go
a := []int{1, 2, 3, 4}  // orijinal veri
b := make([]int, 2)     // b için yeni bir dizi ayır
copy(b, a[:2])          // ilk iki elemanı kopyala
b = append(b, 99)       // artık a etkilenmez
fmt.Println(a)          // [1 2 3 4]
```

> Kısa yol: `b := a[:2:2]` yazarak kapasiteyi 2 ile sınırlayabilirsiniz. Bu durumda `append` yeni bir dizi oluşturur.

---

## 6. Hataları yok saymak

Go'da hatalar dönüş değeri olarak gelir. `_` ile atlamak hatayı ortadan kaldırmaz, sadece gizler.

❌ Yanlış

```go
sayi, _ := strconv.Atoi("12a") // hata yok sayıldı
fmt.Println(sayi * 2)          // 0 basar ve hiçbir uyarı vermez
```

✅ Doğru

```go
sayi, err := strconv.Atoi("12a")                  // hatayı değişkene al
if err != nil {                                   // hata var mı kontrol et
	return fmt.Errorf("sayı okunamadı: %w", err)  // bağlam ekleyip yukarı ilet
}
fmt.Println(sayi * 2)                             // sadece geçerli değerle devam et
```

---

## 7. interface nil tuzağı

Bir interface değeri **tip ve değer** çiftinden oluşur. İçinde `nil` bir işaretçi tutan interface `nil` değildir.

❌ Yanlış

```go
type BenimHatam struct{}                          // özel hata tipi
func (e *BenimHatam) Error() string { return "x" } // error interface'ini karşılar

func islem() error {          // dönüş tipi error interface'i
	var h *BenimHatam = nil   // nil işaretçi
	return h                  // interface'e sarıldı, tipi *BenimHatam
}

// çağıran taraf
if err := islem(); err != nil { // koşul true olur
	fmt.Println("hata var!")    // hata olmadığı hâlde burası çalışır
}
```

✅ Doğru

```go
func islem() error {          // dönüş tipi error interface'i
	var h *BenimHatam         // hata oluşabilecek değişken
	if h == nil {             // gerçekten hata yoksa
		return nil            // açıkça nil interface döndür
	}
	return h                  // hata varsa döndür
}
```

---

## 8. range ile kopya üzerinde değişiklik yapmak

`for _, v := range dizi` içindeki `v` elemanın bir **kopyasıdır**. Kopyayı değiştirmek asıl slice'ı etkilemez.

❌ Yanlış

```go
type Ogrenci struct{ Not int }                 // basit yapı
liste := []Ogrenci{{50}, {60}}                 // iki öğrenci
for _, o := range liste {                      // o her turda bir kopya
	o.Not += 10                                // sadece kopya değişir
}
fmt.Println(liste)                             // [{50} {60}]  (hiçbir şey değişmedi)
```

✅ Doğru

```go
for i := range liste {                         // indeks üzerinden dön
	liste[i].Not += 10                         // asıl eleman değişir
}
fmt.Println(liste)                             // [{60} {70}]
```

---

## 9. map'e aynı anda birden fazla goroutine'den yazmak

Go'daki map yapısı eşzamanlı yazmaya karşı güvenli değildir. Program `fatal error: concurrent map writes` hatasıyla çöker.

❌ Yanlış

```go
sayac := map[string]int{}       // paylaşılan map
for i := 0; i < 100; i++ {      // 100 goroutine
	go func() {                  // hepsi aynı anda çalışıyor
		sayac["ziyaret"]++       // eşzamanlı yazma: çökme riski
	}()                          // fonksiyon hemen çağrılıyor
}
```

✅ Doğru

```go
var (
	mu    sync.Mutex             // map'i korumak için kilit
	sayac = map[string]int{}     // paylaşılan map
	wg    sync.WaitGroup         // goroutine'leri beklemek için
)
for i := 0; i < 100; i++ {       // 100 goroutine
	wg.Add(1)                    // bekleme sayacını artır
	go func() {                  // goroutine başlat
		defer wg.Done()          // bitince sayacı azalt
		mu.Lock()                // kilidi al
		sayac["ziyaret"]++       // güvenli yazma
		mu.Unlock()              // kilidi bırak
	}()                          // fonksiyon hemen çağrılıyor
}
wg.Wait()                        // tüm goroutine'leri bekle
```

> İpucu: Bu tür hataları bulmak için testleri `go test -race` ile çalıştırın.

---

## 10. Türkçe karakterli string'i byte ile işlemek

Go'da string'ler UTF-8 byte dizisidir. `ç, ğ, ı, ö, ş, ü` gibi harfler birden fazla byte kaplar. `len` ve indeksleme karakter değil **byte** sayar.

❌ Yanlış

```go
s := "Işık"           // 4 harf
fmt.Println(len(s))  // 6 basar (byte sayısı)
fmt.Println(s[0:2])  // "I\xc5" gibi bozuk bir parça basar
```

✅ Doğru

```go
s := "Işık"                              // 4 harf
r := []rune(s)                           // karakterlere (rune) çevir
fmt.Println(len(r))                      // 4
fmt.Println(string(r[0:2]))              // "Iş"
fmt.Println(utf8.RuneCountInString(s))   // kopya oluşturmadan sayım: 4
```

> Ek tuzak: `strings.ToUpper("i")` sonucu `"I"` olur, Türkçedeki `"İ"` değil. Türkçe büyük/küçük harf için `strings.ToUpperSpecial(unicode.TurkishCase, s)` kullanın.

---

## Katkı

Eksik gördüğünüz bir hata ya da daha iyi bir örnek varsa issue veya pull request açabilirsiniz.

## Lisans

[MIT](LICENSE) © 2026 Muhammet İkbal Yıldız

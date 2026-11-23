# 1. Merhaba Dünya

Her Go programı bir `package` ile başlar. Çalıştırılabilir programlar `main` paketinde olur ve `main()` fonksiyonundan başlar.

Çalıştırmak için dosyayı `main.go` olarak kaydedip terminalde `go run main.go` yazın.

```go
package main // çalıştırılabilir programlar main paketinde olur

import "fmt" // ekrana yazdırmak için fmt paketini içe aktar

func main() { // program buradan başlar
	fmt.Println("Merhaba, Dünya!") // ekrana yazı yazar ve alt satıra geçer
} // main bitince program sona erer
```

**Çıktı**

```text
Merhaba, Dünya!
```

▶ Çalıştır: `go run ./ornekler/01-merhaba-dunya`

---

[İçindekiler](../README.md#go-temelleri) · [Değişkenler ve Tipler →](02-degiskenler-ve-tipler.md)

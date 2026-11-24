// Konu anlatımı: temeller/03-if-else.md
// Çalıştırmak için: go run ./ornekler/03-if-else
package main // çalıştırılabilir program

import "fmt" // yazdırma paketi

func main() { // program girişi
	not := 72 // örnek sınav notu

	if not >= 85 { // koşul parantezsiz yazılır, süslü parantez zorunludur
		fmt.Println("Pekiyi") // koşul doğruysa çalışır
	} else if not >= 50 { // ilk koşul yanlışsa bu denenir
		fmt.Println("Geçti") // 72 için burası çalışır
	} else { // hiçbir koşul tutmazsa
		fmt.Println("Kaldı") // varsayılan durum
	} // if bloğu sonu

	if kalan := not % 2; kalan == 0 { // if içinde kısa tanımlama yapılabilir
		fmt.Println("Not çift sayı") // kalan sadece bu if bloğunda geçerlidir
	} else { // kalan 1 ise
		fmt.Println("Not tek sayı") // else bloğu da kalan'ı görebilir
	} // kalan burada artık tanımsız

	yas, ehliyet := 19, true  // iki değişken birden tanımla
	if yas >= 18 && ehliyet { // && ve, || veya, ! değil
		fmt.Println("Araç kullanabilir") // iki koşul da doğru
	} // if sonu
} // program sonu

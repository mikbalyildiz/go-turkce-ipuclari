// Konu anlatımı: temeller/01-merhaba-dunya.md
// Çalıştırmak için: go run ./ornekler/01-merhaba-dunya
package main // çalıştırılabilir programlar main paketinde olur

import "fmt" // ekrana yazdırmak için fmt paketini içe aktar

func main() { // program buradan başlar
	fmt.Println("Merhaba, Dünya!") // ekrana yazı yazar ve alt satıra geçer
} // main bitince program sona erer

package main

import (
	"fmt"
	"os"
)

func main() {
	// Komut satırı argümanlarını al
	args := os.Args

	// Eğer komut yoksa hata ver
	if len(args) < 2 {
		fmt.Println("Lütfen bir komut girin: add, list veya delete")
		return
	}

	// Kullanıcının girdiği komutu al
	command := args[1]

	// Komutu kontrol et
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Lütfen eklemek istediğiniz görevi girin!")
			return
		}
		task := args[2]
		fmt.Println("Yeni görev eklendi:", task)
	case "list":
		fmt.Println("Tüm görevleri listele (bu kısmı geliştirebilirsin)")
	case "delete":
		if len(args) < 3 {
			fmt.Println("Silmek istediğiniz görevin numarasını girin!")
			return
		}
		fmt.Println("Görev silindi:", args[2])
	default:
		fmt.Println("Bilinmeyen komut:", command)
	}
}
package main

import (
	"fmt"
)

func main() {
	// tampilan biar kece ahay
	fmt.Println("===== Fizz Buzz =====")
	fmt.Println("[+] Masukkan Angka")
	fmt.Println("[+] Angka 0 Untuk Keluar Dari Program")

	// untuk perulangan agar program terus berjalan, (jika user memasukkan input 0 maka program akan berhenti)
	for {
		// variable input bertipe integer (angka)
		var input int

		// input untuk user bisa memasukkan angka
		fmt.Print("\n[>] ")
		fmt.Scan(&input)

		// kalo bukan 0 bakal terus lanjut cuyh
		if input == 0 {
			fmt.Println("[+] Byee!")
			break
		} else {
			fizzBuzz(input)
		}
	}
}

// fungsi untuk menjalankan program fizz buzz
func fizzBuzz(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("[!] FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("[!] Fizz")
	} else if num%5 == 0 {
		fmt.Println("[!] Buzz")
	} else {
		fmt.Printf("[-] %d\n", num)
	}
}

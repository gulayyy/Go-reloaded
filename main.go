package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	dosya, err := os.OpenFile("sample.txt", os.O_RDONLY, 0o755)
	kontrol(err)
	defer dosya.Close()

	dosya2, err := os.Create("result.txt")
	kontrol(err)
	defer dosya2.Close()

	tarayici := bufio.NewScanner(dosya)

	okuyucu := bufio.NewWriter(dosya2)
	defer okuyucu.Flush()

	for tarayici.Scan() {
		satir := tarayici.Text()
		satir = fonksiyon(satir)
		satir = Bos(satir)
		fmt.Println(satir)
		okuyucu.WriteString(satir + "\n")
	}
}

func kontrol(err error) {
	if err != nil {
		panic(err)
	}
}

func fonksiyon(input string) string {
	kelimeler := strings.Fields(input)

	for i := len(kelimeler) - 1; i >= 0; i-- {
		if kelimeler[i] == "(hex)" {
			sayi, _ := strconv.ParseInt(kelimeler[i-1], 16, 64)
			kelimeler[i-1] = strconv.FormatInt(sayi, 10)
			kelimeler = append(kelimeler[:i], kelimeler[i+1:]...)
		}

		if kelimeler[i] == "(bin)" {
			sayi, _ := strconv.ParseInt(kelimeler[i-1], 2, 64)
			kelimeler[i-1] = strconv.FormatInt(sayi, 10)
			kelimeler = append(kelimeler[:i], kelimeler[i+1:]...)
		}
		// Eğer mevcut kelime "(up)" ise
		if kelimeler[i] == "(up)" {
			// Önceki kelimeyi büyük harfe çevir
			kelimeler[i-1] = strings.ToUpper(kelimeler[i-1])
			// "(up)" kelimesini ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i], kelimeler[i+1:]...)
			i-- // Döngüde bir geri git
		}
		// Eğer mevcut kelime içinde ')' karakteri bulunuyorsa ve bir önceki kelime "(up," ise
		if strings.Contains(kelimeler[i], ")") && kelimeler[i-1] == "(up," {
			// ')' karakterini silerek sayıyı elde et
			nStr := strings.Trim(kelimeler[i], ")")
			n, _ := strconv.Atoi(nStr)

			// Belirtilen sayıdaki önceki kelimeleri büyük harfe çevir
			for a := 0; a < n; a++ {
				kelimeler[i-n+a-1] = strings.ToUpper(kelimeler[i-n+a-1])
			}

			// "(up, n)" kısmını ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i-1], kelimeler[i+1:]...)
			i--      // Döngüde bir geri git
			continue // Bir sonraki iterasyona geç
		}
		// Eğer mevcut kelime "(low)" ise
		if kelimeler[i] == "(low)" {
			// Önceki kelimeyi küçük harfe çevir
			kelimeler[i-1] = strings.ToLower(kelimeler[i-1])
			// "(low)" kelimesini ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i], kelimeler[i+1:]...)
			i-- // Döngüde bir geri git
		}
		// Eğer mevcut kelime içinde ')' karakteri bulunuyorsa ve bir önceki kelime "(low," is),
		if strings.Contains(kelimeler[i], ")") && kelimeler[i-1] == "(low," {
			// ')' karakterini silerek sayıyı elde et
			nStr := strings.Trim(kelimeler[i], ")")
			n, _ := strconv.Atoi(nStr)

			// Belirtilen sayıdaki önceki kelimeleri küçük harfe çevir
			for a := 0; a < n; a++ {
				kelimeler[i-n+a-1] = strings.ToLower(kelimeler[i-n+a-1])
			}

			// "(low, n)" kısmını ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i-1], kelimeler[i+1:]...)
			i--      // Döngüde bir geri git
			continue // Bir sonraki iterasyona geç
		}
		if kelimeler[i] == "(cap)" {
			// Önceki kelimeyi küçük harfe çevir
			kelimeler[i-1] = strings.Title(kelimeler[i-1])
			// "(low)" kelimesini ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i], kelimeler[i+1:]...)
			i-- // Döngüde bir geri git
		}
		// Eğer mevcut kelime içinde ')' karakteri bulunuyorsa ve bir önceki kelime "(low," is),
		if strings.Contains(kelimeler[i], ")") && kelimeler[i-1] == "(cap," {
			// ')' karakterini silerek sayıyı elde et
			nStr := strings.Trim(kelimeler[i], ")")
			n, _ := strconv.Atoi(nStr)

			// Belirtilen sayıdaki önceki kelimeleri küçük harfe çevir
			for a := 0; a < n; a++ {
				kelimeler[i-n+a-1] = strings.Title(kelimeler[i-n+a-1])
			}

			// "(low, n)" kısmını ve bir önceki kelimeyi dilimden çıkar
			kelimeler = append(kelimeler[:i-1], kelimeler[i+1:]...)
			i--      // Döngüde bir geri git
			continue // Bir sonraki iterasyona geç
		} else if kelimeler[i] == "a" {
			// Bir sonraki kelimenin ilk harfi
			firstLetter := kelimeler[i+1][0]
			// Eğer bir sonraki kelimenin ilk harfi ünlü veya 'h' ise 'an' ile değiştir
			if strings.ContainsAny(string(firstLetter), "aeiouhAEIOUH") {
				kelimeler[i] = "an"
			}
		} else if kelimeler[i] == "A" {
			// Bir sonraki kelimenin ilk harfi
			firstLetter := kelimeler[i+1][0]
			// Eğer bir sonraki kelimenin ilk harfi ünlü veya 'h' ise 'An' ile değiştir
			if strings.ContainsAny(string(firstLetter), "aeiouhAEIOUH") {
				kelimeler[i] = "An"
			}
		} else if strings.Contains(kelimeler[i], ")") && kelimeler[i-1] == "(up," {
			nStr := strings.Trim(kelimeler[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				kelimeler[i-n+a-1] = strings.ToUpper(kelimeler[i-n+a-1])
			}
			kelimeler = append(kelimeler[:i-1], kelimeler[i+1:]...)
			i--
			continue
		}
	}
	return strings.Join(kelimeler, " ")
}

func Bos(deger string) string {
	bos := []string{",", ".", "!", "?", ";", ":", "'", "\"", "(", ")", "-"}

	for _, k := range bos {
		deger = strings.Replace(deger, " "+k, k, -1) // tüm dizi için yapmamızı sağlıyor
		deger = strings.Replace(deger, k+" ", k, -1)
	}

	r := []rune(deger)
	dizi := []rune{}
	for i := 0; i < len(r)-1; i++ {
		if unicode.IsPunct(r[i]) && unicode.IsLetter(r[i+1]) {
			if r[i] == '\'' && unicode.IsLetter(r[i+1]) || r[i] == '-' && unicode.IsLetter(r[i+1]) || r[i] == '"' {
				dizi = append(dizi, r[i])
			} else {
				dizi = append(dizi, r[i], ' ')
			}
		} else if (r[i] == ':' || r[i] == ';') && (r[i+1] == '\'' || r[i+1] == '"') {
			dizi = append(dizi, r[i], ' ')
		} else {
			dizi = append(dizi, r[i])
		}
	}
	dizi = append(dizi, r[len(r)-1])

	return string(dizi)
}

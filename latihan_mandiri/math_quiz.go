package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type desc map[string]int
type Levels map[string]desc

func get_random_num(min, max int) int {
	if (min > max) {min-=max;max+=min}
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}

func main(){
	var skor, salah, jawaban int
	var a, b, s int
	var waktu float64
	var level *string = flag.String("level", "2", "enter your level")
	flag.Parse()
	
	levels := Levels{
		"1" : desc{
			"min": 1,
			"max": 10,
		},
		"2" : desc{
			"min": 10,
			"max": 100,
		},
		"3" : desc{
			"min": 100,
			"max": 1000,
		},
	}
	
	fmt.Println("======= Game matematika =======")
	fmt.Println("level", *level)
	fmt.Println("min:", levels[*level]["min"], "\nmax:",levels[*level]["max"])
	
	jalan := true
	nomor := 1
	start := time.Now()
	now   := time.Now()
	for jalan{
		fmt.Println()
		fmt.Printf("nyawa: %d \n", 3 - salah)
		// fmt.Printf("question no.%d \n", nomor)
		
		a = get_random_num(levels[*level]["min"], levels[*level]["max"])
		b = get_random_num(levels[*level]["min"], levels[*level]["max"])

		fmt.Printf("%d.) %d + %d = ", nomor, a, b)
		fmt.Scanln(&jawaban)
		if jawaban == a + b {
			fmt.Println("Benar...")
			s, _ = strconv.Atoi(*level)
			// skor += s * s
			skor += s * s 
			nomor += 1
			} else {
				salah += 1
				fmt.Println("Salah!!!")
			}
			now = time.Now()
			waktu += now.Sub(start).Seconds()
			start = time.Now()
			
		if salah == 3 {
			fmt.Println("\nQuiz Selesai...")
			break
		}
	}
	fmt.Printf("skor : %d \n", skor)
	fmt.Println("time:", waktu, "detik")
}
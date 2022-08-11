package main

import (
	"fmt"
	"time"
	"runtime"
)

func main(){
	sekarang := time.Now()
	fmt.Printf("sekarang: %v\n", sekarang)
	
	//                     tahun, bulan, tanggal,  jam, menit, detik, milisecond, format waktu
	waktu_itu := time.Date(2022 ,   07  ,    01  ,   00 ,   00 , 00 ,     01    ,time.UTC)
	fmt.Println("other time: ",waktu_itu)

	// parsing time dari string
	var date, err = time.Parse(time.RFC822, "01 Jul 06 08:00 WIB")

	if err != nil{
		fmt.Println("Error", err.Error())
	} 
	fmt.Println("other time2:", date)

	fmt.Println("Program di hentikan selama 3 detik")
	time.Sleep(3)
	fmt.Println("Program kembali berjalan")

	
}
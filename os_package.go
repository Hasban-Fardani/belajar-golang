package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func write(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	file.WriteString(text)
	return err
}

func main() {
	var text string

	ls_output, _ := exec.Command("ls").Output()
	fmt.Println("\n>> ls\n", string(ls_output))

	time.Sleep(2 * time.Second)
	fileName := "hello.txt"
	_, err := os.Stat(fileName)
	if err != nil {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("error", err.Error())
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("berhasil membuat file", fileName)
			time.Sleep(2 * time.Second)
		}
		fmt.Print("ketik sesuatu: ")
		fmt.Scanln(&text)
		err = write(fileName, text)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("berhasil menulis:", text, "\ndi file: ", fileName)
		defer file.Close()
	} else {
		fmt.Println("file sudah ada")
		fmt.Println("mencoba untuk menghapus file")
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("berhasil menghapus file", fileName)
			time.Sleep(2 * time.Second)
		}
		defer main()
	}

}

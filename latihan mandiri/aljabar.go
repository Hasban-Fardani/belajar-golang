// script gabut

package main

import "fmt"

func get_superscript_int(number int) string {
	// https://www.compart.com/en/unicode/block/U+2070
	if number == 0 {
		return "\u2070"
	} else if number == 1 {
		return "\u00B9"
	} else if number == 2 {
		return "\u00B2"
	} else if number == 3 {
		return "\u00B3"
	} else if number == 4 {
		return "\u2074"
	} else if number == 5 {
		return "\u2075"
	} else if number == 6 {
		return "\u2076"
	} else if number == 7 {
		return "\u2077"
	} else if number == 8 {
		return "\u2078"
	} else if number == 9 {
		return "\u2079"
	} else {
		return ""
	}
}

type Variabel struct {
	koefisien int
	str       string
	pangkat   int
}

type Aljabar struct {
	variabel []Variabel
}

func (Var Variabel) strf() string {
	if Var.koefisien != 0 {
		return fmt.Sprintf("%d.%s%s", Var.koefisien, Var.str, get_superscript_int((Var.pangkat)))
	} else {
		return ""
	}
}

func (aljabar Aljabar) sum(other Aljabar) Aljabar {
	for i, v := range aljabar.variabel {
		if i < len(other.variabel) {
			if v == other.variabel[i] {
				aljabar.variabel[i].koefisien += other.variabel[i].koefisien
			}
		}
	}
	return aljabar
}

func main() {
	x := Variabel{1, "x", 2}
	y := Variabel{1, "y", 1}
	al := Aljabar{}
	al.variabel = []Variabel{x, y}
	fmt.Println(al)

	x2 := Variabel{1, "x", 2}
	al2 := Aljabar{}
	al2.variabel = []Variabel{x2}
	al.sum(al2)
	fmt.Println(al)
}

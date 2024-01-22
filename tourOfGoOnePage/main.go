package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

// Variables
var c, python, java bool
var var1, var2 int = 1, 2

const Pi = 3.14

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// BUyuk sayilar...
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

type Vertex struct {
	X int
	Y int
}

// structin diğer bir yazim bicimi
// type Vertex struct{
// 	X, Y int
// }

var (
	vervar1 = Vertex{1, 2}  // has type vertex
	vervar2 = Vertex{X: 1}  // y is implict
	vervar3 = Vertex{}      // x 0 y 0
	vervarp = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// page 1 ln yenı satıra yazar
	fmt.Println("Benim fav numaram ", rand.Intn(10))

	// page 2 printf aynı satira yazar
	fmt.Printf("Problem %g \n", math.Sqrt(7))

	// kutuphaneden fonksiyon cagirilmasi eger yok ise undefined veriyor..
	// fmt.Println(math.pi)

	// toplama fonksiyonu
	fmt.Println(add(42, 13))

	// degiskenleri birbirine atamak
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// split
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(var1, var2, c, python, java)

	var var3, var4 int = 1, 2
	k := 3
	c, python, java = true, false, "no!"
	fmt.Println(var3, var4, k, c, python, java)

	// golang komleks sayilar
	fmt.Printf("Type %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v \n", z, z)

	// golang default atanan deger
	var ni int
	var nf float64
	var nb bool
	var ns string
	fmt.Printf("%v %v %v %q\n", ni, nf, nb, ns)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// golang tip type
	var5 := 12.2
	fmt.Printf("v'nin type'i %T\n", var5)
	var6 := 12
	fmt.Printf("v'nin type'i %T\n", var6)

	const World = "aliveli"
	fmt.Println("Hello ", World)
	fmt.Println("Happy ", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// BUyuk sayilari shift kaydirma
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	/*
		ILK BOLUM SONU
		You finished this lesson!

		You can go back to the list of modules to find what to learn next, or continue with the next lesson.
	*/
	// 2. BOLUM
	// FOR LOOP
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// sartsiz for loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// golang for while  döngüsü
	sum = 1
	for sum < 2000 {
		sum += sum
	}
	fmt.Println(sum)

	// sonsuz for loop
	// for {
	// }

	// golang if şarti
	fmt.Println(sqrt(2), sqrt(-4))

	// if kisa şart1
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	// golang switch case
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("OSX\n")
	case "linux":
		fmt.Printf("LİNUX\n")
	default:
		//
		fmt.Printf("%s \n", os)
	}

	// switch degerlendirme
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Bugun")
	case today + 1:
		fmt.Println("Yarin")
	case today + 2:
		fmt.Println("2 gun ıcınde")
	default:
		fmt.Println("Daha cok var...")

	}

	// switch şartsiz olma durumu
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 12:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

	// // defer
	// defer fmt.Println("world")
	// fmt.Println("hello2")

	// Tum  sonraki değerlerden sonra ekrana yazdirilir bu yuzden comment out edildi
	// fmt.Println("COUNTING")
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

	// İşaretçiler --- Pointers https://go.dev/tour/moretypes/1 devam
	p1, p2 := 42, 2701

	p3 := &p1
	fmt.Println(*p3) // 42

	*p3 = 21
	fmt.Println(p1) // 21

	p3 = &p2
	*p3 = *p3 / 37
	fmt.Println(p2) // 73

	// STRUCT

	fmt.Println(Vertex{1, 2})
	ver := Vertex{1, 2}
	ver.X = 4
	fmt.Println(ver.X)

	//  struct ve pointer
	// yapilar ve pointer
	verp := Vertex{1, 2}
	verp1 := &verp
	verp1.X = 1e9
	fmt.Println(verp)

	// struct literals
	fmt.Println(vervar1, vervarp, vervar2, vervar3)

	// arrays diziler
	var arrayone [2]string
	arrayone[0] = "hello"
	arrayone[1] = "world"

	fmt.Println(arrayone[0], arrayone[1])
	fmt.Println(arrayone)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// array parcalama
	primesarray := [6]int{2, 3, 5, 7, 11, 13}
	var primesarrayonefour []int = primesarray[1:4]
	fmt.Println(primesarrayonefour) // 3 5 7

	// string slices string parcalama
	namesarray := [4]string{
		"ali",
		"veli",
		"jale",
		"halil",
	}
	fmt.Println(namesarray) // [ali veli jale halil]
	aa := namesarray[0:2] 
	bb := namesarray[1:3]
	fmt.Println(aa, bb) // [ali veli] [veli jale]

	bb[0] =  "XXX"
	fmt.Println(aa,bb) // [ali XXX] [XXX jale]
	fmt.Println(namesarray) // [ali XXX jale halil]

}

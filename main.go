package main

import (
	"fmt"
	"breadlibrary_go/testMain"
)

func main() {
    testMain.TestMain() // Call the TestMain function from the test package
    fmt.Println() // Print a newline for better output formatting
}




// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var k string
// 	fmt.Scan(&k)
// 	costom_Print(k)
// }


// func costom_Print(numer string) {
//   fmt.Print(numer)
//   k=32
//   fmt.Print(string(k)) 
// }



// package main

// import (
// 	"fmt"
// )

// func main() {
//  var num string
//  fmt.Scan(&num)
//  for i:=len(num) - 1; i >= 0; i-- {
// 	 fmt.Print(string(num[i]))
//  }
// }




// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var num string
// 	fmt.Scan(&num)


// 	sum := 0
// 	for i := 0; i < len(num); i++ {
// 		sum += int(num[i] - '0')
// 	}
// 	fmt.Print(sum)
// }





// package main
// import "fmt"
// func main()  {
//   var len_array int
//   var counter int = 0
//     fmt.Scan(&len_array)
// 	var array []int = make([]int, len_array)
// 	var a int
//     for i:=0; i < len(array); i++{
// 		fmt.Scan(&a)
// 		array[i] = a
// 	}
//     for i:=0; i < len(array);i++{
//         if (array[i] > 0 ){
// 			counter++
//         }
//     }
// 	fmt.Printf("%d", counter)
// }







// package main
// import "fmt"
// func main()  {
//   var len_array int
//     fmt.Scan(&len_array)
// 	var array []int = make([]int, len_array)
// 	var a int
//     for i:=0; i < len(array); i++{
// 		fmt.Scan(&a)
// 		array[i] = a
// 	}
//     for i:=0; i < len(array);i++{
//         if (i % 2 ) == 0{
// 			fmt.Printf("%d ", array[i])
//         }
//     }
// }







// package main

// import "fmt"

// func main() {
// 	var baseSlise []int = []int{1,2,3,4}
// 	fmt.Printf(" массив,: %v\n Длина: %d\n  Ёмкость: %d\n \n", baseSlise, len(baseSlise), cap(baseSlise))
// 	baseSlise = append(baseSlise, baseSlise...)
// 	fmt.Printf(" новый массив,: %v\n Длина: %d\n  Ёмкость: %d\n", baseSlise, len(baseSlise), cap(baseSlise))
// }



// package main

// import (
// 	"fmt"
// )

// func main(){
//     a:=0
// 	fmt.Scan(&a)
// 	for i=0; i<10; i--{
		
// 		fmt.Scan(&a)
// 	}
    

// }



// package main
// import "fmt"
// func main(){
//   var a int
//   fmt.Scan(&a) // считаем переменную 'a' с консоли
//   b	:= string(a)
//   d := len(b)
  
  


//   fmt.Println()
// }







// package main

// import "fmt"

// func main(){
//     var a int
//     fmt.Scan(&a)
//     a*= 2
// 	a+=100
// 	fmt.Print(a)
// }








// func main(){
//   input_continue := "";
// 	fmt.Println(string("Please enter your text the  continue the text"))
// 	fmt.Println(string("programm read 1-st your word"))
// 	fmt.Scan(&input_continue) 
// 	lowerInput := strings.ToLower(input_continue)
// 	git(lowerInput)


// 	}

// func git(text_continue string){
// 		fmt.Print(string(72))
// 		fmt.Print(string(73))
// 		fmt.Print(" ")
// 		fmt.Print(string(71))
// 		fmt.Print(string(73))
// 		fmt.Print(string(84))
// 		if text_continue !="ler" &&  text_continue !="лер" && text_continue !="lеr" && text_continue !="lер" && text_continue !="лer" &&  text_continue !="лeр" {
// 			upper_text_continue := strings.ToUpper(text_continue)
// 			fmt.Print(string(upper_text_continue))
// 		}

// }

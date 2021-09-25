package main

import "fmt"
import "strconv"

func main(){
  var input string
  defer fmt.Println("The last Hero")
  fmt.Print("Masukkan angka :")
  fmt.Scanln(&input)

  var number int
  var err error
  number, err=strconv.Atoi(input)

  if err==nil{
    fmt.Println(number,"Ini adalah angka")
  } else{
    fmt.Println(input,"Ini bukanlah angka")
    panic(err.Error())
  }
}

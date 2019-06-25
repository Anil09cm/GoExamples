package main

import(
        "fmt"
)

type S1 struct{
        name string
        age int
}

var  Mymap = make(map[string]S1)

func main(){

        s1 := S1{"Abc", 10}
        Mymap["Abc"] = s1

        s2 := S1{"Bcd", 20}
        Mymap["Bcd"] = s2

        fmt.Println("Mymap :", Mymap)
        x := Mymap["aba"]
        fmt.Println("X : ", x)
        if x.name == ""{
                fmt.Println("Error in Key!")
        }
}

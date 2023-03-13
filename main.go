package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

 func main(){
	creatFile()
	readFile("name.txt")
 }

 func creatFile(){
	//Create a file in go 
	// Asign an error checker to the function
	f, err := os.Create("name.tx")
	if err != nil{
		fmt.Println("An error has occurred:\t",err)
	}
	defer f.Close()
	r := strings.NewReader("Ive been coming but f me")

	io.Copy(f,r)
 }
 func readFile(nf string){
	//Open the direccion of the file
	//This asign an pointer and error variable
	f, err:=os.Open(nf)
	if err!= nil{
		fmt.Println("An error has occured while opening the file:",err)
		return
	}
	//Close the open file to save Memory
	defer f.Close()
	//This return an []byte and error variable
	bs, err := ioutil.ReadAll(f)
	//Check for the error 
	if err!=nil{
		fmt.Println("There has been an error:",err)
	}
	fmt.Println("Now the content is coming")
	//Turn the []byte to a string
	fmt.Println(string(bs))
 }

 //FIXME: Ways to Print Errors 	
 func printErrors_(){
	_, err:=os.Open("adasd")
	if err!= nil{
		// fmt.Println("An error has occured while opening the file:",err)
		//FIXME:// Now with log.Println("",err)
		log.Println("This is the error",err)
		//FIXME://Now with log.Fatalln(err)
		//This one kills the program
		log.Fatalln(err)

		return
	}
 }
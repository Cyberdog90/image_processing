package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	/*2枚の画像のパスをを引数として取る*/
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("引数の数が足りません。")
		os.Exit(1)
	}

	/*1枚目*/
	file1, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	img1, _, err := image.Decode(file1)
	if err != nil {
		log.Fatal(err)
	}

	/*2枚目*/
	file2, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	img2, _, err := image.Decode(file2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", img1)
	fmt.Printf("%T\n", img2)
}

func gray() {
	/*出力画像作成*/
	clone, err := os.Create("output/グレースケール.png")
	if err != nil {
		log.Fatal(err)
	}
	defer clone.Close()
}

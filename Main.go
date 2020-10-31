package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
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
	file2, err := os.Open(flag.Args()[1])
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

	img1Bounds := img1.Bounds()
	img2_bounds := img2.Bounds()

	/*グレースケール化*/
	clone, err := os.Create("output/グレースケール.png")
	if err != nil {
		log.Fatal(err)
	}
	cloneImage, _, err := image.Decode(clone)
	if err != nil {
		log.Fatal(err)
	}

	defer clone.Close()

	i := image.NewGray16(img1Bounds)
	for x := img1Bounds.Min.X; x < img1Bounds.Max.X; x++ {
		for y := img1Bounds.Min.Y; y < img1Bounds.Max.Y; y++ {
			gray, _ := color.Gray16Model.Convert(img1.At(x, y)).(color.Gray16)
			i.Set(x, y, gray)
		}
	}
	_, err = io.Copy(cloneImage, i)
	if err != nil {
		log.Fatal(err)
	}
}

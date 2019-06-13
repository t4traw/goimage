package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	// ベースのTシャツ画像の読み込み
	baseFile, err := os.Open("base.png")
	defer baseFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// プログラム内で使えるようにデコード
	baseImage, _, err := image.Decode(baseFile)
	if err != nil {
		log.Fatal(err)
	}
	// どこに書き込むか場所指定(x0y0からbaseImageサイズ=500,500)
	startPoint := image.Point{0, 0}
	baseRectangle := image.Rectangle{startPoint, startPoint.Add(baseImage.Bounds().Size())}

	// 確認用の出力
	// fmt.Println(baseRectangle)

	// 上に乗っけるロゴ画像の読み込み
	logoFile, err := os.Open("logo.png")
	defer logoFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// プログラム内で使えるようにデコード
	logoImage, _, err := image.Decode(logoFile)
	if err != nil {
		log.Fatal(err)
	}

	// (ベースの横幅 - ロゴ画像の横幅) / 2 で雑に中央に貼るようのx座標を作成
	logoPositionX := (baseImage.Bounds().Dx() - logoImage.Bounds().Dx()) / 2
	// y座標はもう目視。ここは動的でもいいかもしんない
	logoPositionY := 150
	// ロゴ画像の描画始点を作成
	logoPosition := image.Point{logoPositionX, logoPositionY}
	// んで、書き込む場所を(ry
	logoRectangle := image.Rectangle{logoPosition, logoPosition.Add(logoImage.Bounds().Size())}

	// 確認用の出力
	// fmt.Println(logoRectangle)

	// ベースTシャツと同じサイズで新規画像作成
	newImage := image.NewRGBA(baseRectangle)
	// まずベースのTシャツ画像をnewImageに書き込み
	draw.Draw(newImage, baseRectangle, baseImage, image.Point{0, 0}, draw.Src)
	// 次にロゴ画像を書き込む
	draw.Draw(newImage, logoRectangle, logoImage, image.Point{0, 0}, draw.Over)

	outputFile, _ := os.Create("output.png")
	defer outputFile.Close()
	png.Encode(outputFile, newImage)
}

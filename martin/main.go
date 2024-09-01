package main

import (
	"fmt"
	"os"

	"github.com/ReeseHatfield/core"
	"github.com/ReeseHatfield/ffmpeg"
	"github.com/ReeseHatfield/reader"
	"github.com/ReeseHatfield/web"
)

func main() {

	// turn rel path to abs in bash?
	records, err := reader.GetRecords("../examples/albums.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)

	os.Exit(0)

	img, err := web.GetCover("You and Your Friends", "Peach Pit")

	fmt.Println(img.Path)

	// os.Exit(0)

	martinObj, err := ffmpeg.NewMartin("../examples/MartinListensToRealMusic.png")
	if err != nil {
		fmt.Println(err)
	}

	c, err := core.NewCore(martinObj)

	coverObj := &ffmpeg.Image{
		Path: "../examples/cover3.png",
	}
	c.SetCover(*coverObj)

	err = c.GeneratePfp()
	if err != nil {
		fmt.Println(err)
	}

	// for file in dir:
	// c.setCover(cover)
	// img := c.MakePfp()
	// img.save

}

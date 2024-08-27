package main

import (
	"fmt"

	"github.com/ReeseHatfield/core"
	"github.com/ReeseHatfield/ffmpeg"
)

func main() {

	martinObj, err := ffmpeg.NewMartin("../examples/MartinListensToRealMusic.png")
	if err != nil {
		fmt.Println(err)
	}

	c, err := core.NewCore(martinObj)

	coverObj := &ffmpeg.Image{
		Path: "../examples/cover.png",
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

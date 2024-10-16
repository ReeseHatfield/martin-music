package main

import (
	"fmt"
	"os"

	"github.com/ReeseHatfield/core"
	"github.com/ReeseHatfield/ffmpeg"
	"github.com/ReeseHatfield/query"
	"github.com/ReeseHatfield/web"
)

func main() {

	fmt.Println("got here?")
	albumFilePath := os.Args[1]
	recordQueries, err := query.GetQuerys(albumFilePath)
	if err != nil {
		fmt.Println(err)
	}

	imgs := make([]ffmpeg.Image, 0)

	for _, q := range recordQueries {
		cover, err := web.GetCover(q)
		if err != nil {
			fmt.Println("Could not find album cover for query " + q.String())
			os.Exit(1)
		}

		imgs = append(imgs, *cover)
	}

	martinObj, err := ffmpeg.NewMartin("../MartinListensToRealMusic.png")
	if err != nil {
		fmt.Println(err)
	}

	c, err := core.NewCore(martinObj)

	for i, img := range imgs {

		c.SetCover(img)

		err = c.GeneratePfp(recordQueries[i].String())
		if err != nil {
			fmt.Println(err)
		}
	}
}

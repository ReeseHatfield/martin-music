package main

import (
	"fmt"

	"github.com/ReeseHatfield/core"
	"github.com/ReeseHatfield/ffmpeg"
	"github.com/ReeseHatfield/query"
	"github.com/ReeseHatfield/web"
)

func main() {

	// turn rel path to abs in bash?
	recordQueries, err := query.GetQuerys("../examples/albums.txt")
	if err != nil {
		fmt.Println(err)
	}

	imgs := make([]ffmpeg.Image, 0)

	for _, q := range recordQueries {
		cover, err := web.GetCover(q)
		if err != nil {
			fmt.Println("Could not find album cover for query " + q.String())
		}

		imgs = append(imgs, *cover)
	}

	martinObj, err := ffmpeg.NewMartin("../examples/MartinListensToRealMusic.png")
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

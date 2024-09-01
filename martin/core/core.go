package core

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os/exec"
	"path/filepath"

	"github.com/ReeseHatfield/ffmpeg"
)

const (
	COVER_DIMENSIONS_PX    = 920
	MOVE_X                 = 450
	MOVE_Y                 = 1360
	COVER_ROTATION_DEGREES = 18.0

	TEMP_FILE_NAME = "../temp/transparent_background.png" // this may have broke
	OUT_DIR        = "../out/"
)

type Core struct {
	background ffmpeg.Martin
	cover      *ffmpeg.Image
}

func NewCore(martin *ffmpeg.Martin) (*Core, error) {

	return &Core{
		background: *martin,
		cover:      nil,
	}, nil
}

func (c *Core) SetCover(img ffmpeg.Image) {
	c.cover = &img
}

func (c *Core) GeneratePfp() error {

	if c.cover == nil {
		return errors.New("Err: Cover image not set")
	}

	// build behind image
	coverPath, err := filepath.Abs(c.cover.Path)
	if err != nil {
		return err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	//ffmpeg needs radians for rotation
	rotationRadians := COVER_ROTATION_DEGREES * (math.Pi / 180.0)

	// big ffmpeg command to build the 'background' image
	cmd := exec.Command("ffmpeg", "-y", "-i", coverPath,
		"-vf", fmt.Sprintf("scale=%d:%d,pad=%d:%d:%d:%d:0x00000000,rotate=%f:ow=%d:oh=%d:c=0x00000000",
			COVER_DIMENSIONS_PX, COVER_DIMENSIONS_PX, c.background.Width, c.background.Height, MOVE_X, MOVE_Y, rotationRadians, c.background.Width, c.background.Height),
		TEMP_FILE_NAME)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("sum went wrong")
	}

	fmt.Println(stdout.String())
	fmt.Println(stderr.String())

	// build in front image

	//ffmpeg -i MartinListensToRealMusic.png -i output.png -filter_complex "" -c:v png pfp.png

	// this shoudl be a constant somewhere
	martinPath, err := filepath.Abs(c.background.Image.Path)
	if err != nil {
		return err
	}

	// need to make unique name for out file
	cmd = exec.Command("ffmpeg", "-y", "-i", martinPath,
		"-i", TEMP_FILE_NAME, "-filter_complex", "[1:v][0:v]overlay=0:0", "-c:v", "png", OUT_DIR+"pfp.png")

	err = cmd.Run()
	if err != nil {
		fmt.Println("sum went wrong at the end")
	}

	fmt.Println(stdout.String())
	fmt.Println(stderr.String())

	return nil
}

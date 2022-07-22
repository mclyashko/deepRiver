package main

import (
	"flag"
	"github.com/Knetic/govaluate"
	"image"
	"image/png"
	"math"
	"os"
)

func GetPicGenF(size int, pixelF *govaluate.EvaluableExpression) (picGenF func() [][]uint8) {

	picGenF = func() [][]uint8 {
		result := make([][]uint8, size)

		for i := 0; i < size; i++ {
			result[i] = make([]uint8, size)
			for j := 0; j < size; j++ {
				params := make(map[string]interface{})
				params["x"] = j
				params["y"] = i
				tmp, err := pixelF.Evaluate(params)
				if err != nil {
					panic("can't eval")
				}
				result[i][j] = uint8(tmp.(float64))
			}
		}

		return result
	}

	return
}

func PicMake(size int, pixelF *govaluate.EvaluableExpression) *image.NRGBA {

	picGenF := GetPicGenF(size, pixelF)
	data := picGenF()

	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			val := data[y][x]
			ind := y*img.Stride + x*4

			img.Pix[ind] = val
			img.Pix[ind+1] = val
			img.Pix[ind+2] = 255
			img.Pix[ind+3] = 255
		}
	}
	return img
}

func SavePic(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)

	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {
			panic("can't close: " + err.Error())
		}
	}(imgFile)

	if err != nil {
		panic("can't create: " + err.Error())
	}

	err = png.Encode(imgFile, img.SubImage(img.Rect))

	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

func getStdFuncMapToEvalInterface() (functions map[string]govaluate.ExpressionFunction) {
	functions = map[string]govaluate.ExpressionFunction{
		"pow": func(args ...interface{}) (interface{}, error) {
			firstOp := args[0].(float64)
			secondOp := args[1].(float64)
			return math.Pow(firstOp, secondOp), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			firstOp := args[0].(float64)
			return math.Sin(firstOp), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			firstOp := args[0].(float64)
			return math.Cos(firstOp), nil
		},
	}

	return
}

func getExpFromCLI(stringExp string) *govaluate.EvaluableExpression {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(stringExp, getStdFuncMapToEvalInterface())
	if err != nil {
		panic("wrong expression")
	}
	return expression
}

func MyCLI() (fileName string, size int, exp *govaluate.EvaluableExpression) {
	flag.StringVar(&fileName, "fileName", "out.png", "name of the generated pic")
	flag.IntVar(&size, "size", 256, "size of the generated square pic")
	var stringExp string
	flag.StringVar(&stringExp, "mathExp", "x*y", "the pixel value function")
	flag.Parse()

	exp = getExpFromCLI(stringExp)

	return
}

func main() {
	fileName, size, exp := MyCLI()

	SavePic(fileName, PicMake(size, exp))
}

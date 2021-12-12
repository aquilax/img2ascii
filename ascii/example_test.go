package ascii_test

import (
	"encoding/base64"
	"image"
	"os"
	"strings"

	_ "image/png"

	"github.com/aquilax/img2ascii/ascii"
)

func getExampleImage() image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	i, _, _ := image.Decode(reader)
	return i
}

func ExampleImage_Encode() {
	c := ascii.NewImage()
	i := getExampleImage()
	c.Encode(os.Stdout, i)
	// Output:
	// =++++***###%%%%%%%%%######****************######
	// ===+++***###%%%%%%%%%######****************#####
	// -===+++***###%%%%%%%%%######****************####
	// --===+++***###%%%%%%%%%######****************###
	// ---===+++***###%%%%%%%%%######****************##
	// ----===+++***###%%%%%%%%%######****************#
	// -----===+++***###%%%%%%%%%######****************
	// ------===+++***###%%%%%%%%%######***************
	// -------===+++***###%%%%%%%%%######**************
	// ==------===+++***###%%%%%%%%%######*************
	// ===------===+++***###%%%%%%%%%######************
	// ====------===+++***###%%%%%%%%%######***********
	// =====------===+++***###%%%%%%%%%######**********
	// ======------===+++***###%%%%%%%%%######*********
	// =======------===+++***###%%%%@%%%%######********
	// ========------===+++***###%%%@%%%%%######*******
	// =========------===+++***###%%@%%%%%%######******
	// ===#@@===@@*--@#@*%@#+*@%@%#%@%%@@@%%######*****
	// ===@*=====+@--@+=@*+@++@##@##@%%@@%@%%######****
	// ===*%%==#@@@--@--@==@++@**@##@%%@@@@%%@@#####***
	// =====%*=@++@=-@--@-=@=+@*#@**@##@@%%%%%%%#####**
	// ===@@%==#@#@==@--@--@==@%@#**@##%@@@%%%%%%######
	// ===============------==@+++****###%%%%%%%%%#####
	// -===============------=@=++++***###%%%%%%%%%####
	// --=*@@===@=======------====+++***###%%%%%%%%%###
	// ---@==@==@========-------===+++***###%%%%%%%%%##
	// ---@=====@=========-------===+++***###%%%%%%%%%#
	// --@@@=@==@==+@@*====-=%@@--+@@*++@%@%@@@%%%%%%%%
	// ---@--@==@==@*+@=====%*=---@**@++@##@%%@%%%%%%%%
	// :--@--@==@==@@@@=====@=----@==@=+@+*@*#@#%%%%%%%
	// ::-@--@-=@==@*=======%*=---%++@==@++@**@##%%%%%%
	// :::@--@--@==+@@@==@==+%@@--+@@+==@++@**@###%%%%%
	// ::::------===============-------===+++***###%%%%
	// :::::------================------===+++***###%%%
	// ::::::------================------===+++***###%%
	// :::::::------================------===+++***###%
	// ::::::::------================------===+++***###
	// :::::::::------================------===+++***##
	// -:::::::::-------===============------===+++***#
	// --:::::::::------================------===+++***
	// ---:::::::::-------===============------===+++**
	// =---:::::::::-------===============------===+++*
	// ==---:::::::::-------===============------===+++
	// ===----:::::::::------===============------===++
	// +===----:::::::::------===============------===+
	// ++===----:::::::::------===============------===
	// +++====---:::::::::------===============------==
	// *+++====---:::::::::------===============------=
}

const data = `
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
`

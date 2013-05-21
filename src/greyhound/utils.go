package greyhound

import "io/ioutil"
import "log"
import "net/http"

// Method to read a file at a path, and write the proper response
func HandleFile(w http.ResponseWriter, path string) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print("Could not read file! ", err)
	}
	w.Write(contents)
}

// determines the content type of the string based on file type suffix
/* func getContentType(path string) string {
	if strings.HasSuffix(path, ".css") {
		return "text/css"
	} else if strings.HasSuffix(path, ".js") {
		return "text/js"
	}
	return "text/plain"
} */

// mymodule/utils/utils.go
package utils

import (
	"fmt"
	"net/http"
)

func HelperFunction() {
	fmt.Println("This is a helper function from the utils package.")
}

func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

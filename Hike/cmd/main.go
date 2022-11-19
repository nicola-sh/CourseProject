package hike

import (
	"net/http"
)

func main() {

	http.ListenAndServe(":9090", nil)
}

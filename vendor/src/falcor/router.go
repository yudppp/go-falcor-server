package falcor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type router struct {
}

func NewRouter() router {
	return router{}
}

func (r *router) AddRoute(key string, get interface{}) {

}

func FalcorHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Println("---")
	pathsStr := q.Get("paths")
	paths := make([]string, 0, 10)
	json.Unmarshal([]byte(pathsStr), &paths)

	for _, path := range paths {
		v := strings.Split(path, ".")
		fmt.Println(v)
	}
	fmt.Fprintf(w, "Falcor")
}

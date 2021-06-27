package redirector

import (
	"net/http"
	"strconv"
	"strings"
)

// MapHandler redirects links and presumes that a link has a number as its second part
//examples of inputs and outputs:
// /provinces/ -> /provinces/economic/ (depends on the key-value in pathsToUrls)
// /provinces/1 -> /provinces/economic/1 (depends on the key-value in pathsToUrls)
// /provinces/economic/ -> /provinces/economic/
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		searchPath := ""
		if len(pathParts) >= 1 {
			searchPath = "/" + pathParts[0] + "/"
		}
		if dest, ok := pathsToUrls[searchPath]; ok {
			if len(pathParts) == 1 {
				http.Redirect(w, r, dest, http.StatusTemporaryRedirect)
				return
			} else if len(pathParts) == 2 {
				_, err := strconv.Atoi(pathParts[1])
				if err == nil {
					dest = dest + "/" + pathParts[1]
					http.Redirect(w, r, dest, http.StatusTemporaryRedirect)
					return

				}
			}
		}
		fallback.ServeHTTP(w, r)
	}
}

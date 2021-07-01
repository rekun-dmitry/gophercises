package redirector

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
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

func YAMLHandler(ymlPath string, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	yml, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

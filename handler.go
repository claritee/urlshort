package urlshort

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 //    	w.Write([]byte("hello"))
 //  	})

	// See: https://www.alexedwards.net/blog/making-and-using-middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
    		http.Redirect(w, r, url, http.StatusFound)
			return
		}
	    fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("YAML Handler")
	parsedYaml, err := parseYAML(data)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(data []byte) ([]PathUrl, error) {
	var output []PathUrl

    err := yaml.Unmarshal([]byte(data), &output)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    // fmt.Println(output[0].Path)
    // fmt.Println(output[0].Url)
	return output, nil
}

func buildMap(data []PathUrl) map[string]string {
	return make(map[string]string)
}

// https://godoc.org/gopkg.in/yaml.v2
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
type PathUrl struct {
    Path string `yaml:"path"`
    Url string `yaml:"url"`
}
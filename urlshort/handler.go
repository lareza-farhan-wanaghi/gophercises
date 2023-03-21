package urlshort

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

type MappedPath struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		redirectLink, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, redirectLink, http.StatusSeeOther)
			return
		}
		fallback.ServeHTTP(w, r)
	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathMap, err := parseYAML(yml)
	if err != nil {
		log.Fatal(err)
	}
	return MapHandler(pathMap, fallback), nil
}

// parseYAML creates a map from yaml bytes
func parseYAML(yml []byte) (map[string]string, error) {
	var pathMappers []MappedPath

	err := yaml.Unmarshal(yml, &pathMappers)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pathMapper := range pathMappers {
		result[pathMapper.Path] = pathMapper.Url
	}
	return result, nil
}

// JsonHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathMap, err := parseJSON(jsn)
	if err != nil {
		log.Fatal(err)
	}
	return MapHandler(pathMap, fallback), nil
}

// parseJSON creates a map from json bytes
func parseJSON(jsn []byte) (map[string]string, error) {
	var pathMappers []MappedPath

	err := json.Unmarshal(jsn, &pathMappers)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, pathMapper := range pathMappers {
		result[pathMapper.Path] = pathMapper.Url
	}

	return result, nil
}

// DatabaseHandler creates handler based of data stored in the specified database
func DatabaseHandler(fallback http.Handler, db *sql.DB) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathMap, err := databaseParse(db)
	if err != nil {
		log.Fatal(err)
	}
	return MapHandler(pathMap, fallback), nil
}

// databaseParse creates a map based on data returned from the specified database
func databaseParse(db *sql.DB) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `select path,url from pathmapper`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pathMappers := make(map[string]string)
	for rows.Next() {
		var out MappedPath
		err = rows.Scan(
			&out.Path,
			&out.Url,
		)
		if err != nil {
			return nil, err
		}
		pathMappers[out.Path] = out.Url
	}
	return pathMappers, nil
}

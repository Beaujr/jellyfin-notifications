package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var (
	server  = flag.String("jellyfin", "http://localhost:8096", "Jellyfin Server URL. eg http://192.168.1.1:8096")
	library = flag.String("library", "libraryId", "Jellyfin library eg 767bffe4f11c93ef34b805451a696a4e")
)

func main() {
	flag.Parse()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := updateJellyfin(r.Header.Get("X-Mediabrowser-Token"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write([]byte(err.Error())); err != nil {
				logger.Error(err.Error())
			}
			return
		}
		w.WriteHeader(statusCode)
		return
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err.Error())
	}
}

func updateJellyfin(apikey string) (int, error) {
	u, err := url.Parse(*server)
	if err != nil {
		return 0, err
	}
	path := filepath.Join("/Items", *library, "/Refresh?Recursive=true&ImageRefreshMode=Default&MetadataRefreshMode=Default")
	req, err := http.NewRequest(http.MethodPost, u.String()+path, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Mediabrowser Token=\"%s\"", apikey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

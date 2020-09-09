package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIdFromUrl(r *http.Request, pathIndex int) (uint, error) {
	p := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(p[pathIndex])
	if err != nil {
		return 0, err
	} else {
		return uint(id), nil
	}
}


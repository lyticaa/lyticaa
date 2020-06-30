package helpers

import (
	"net/http"
	"strconv"
)

var (
	defaultLength = int64(10)
	defaultStart  = int64(1)
	defaultDraw   = int64(0)
	defaultSort   = int64(0)
	direction     = map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}
)

func DtDraw(r *http.Request) int64 {
	draw := r.URL.Query()["draw"]
	if len(draw) == 0 {
		return defaultDraw
	}

	v, _ := strconv.ParseInt(draw[0], 10, 64)
	return v
}

func DtStart(r *http.Request) int64 {
	start := r.URL.Query()["start"]
	if len(start) == 0 {
		return defaultStart
	}

	v, _ := strconv.ParseInt(start[0], 10, 64)
	return v
}

func DtLength(r *http.Request) int64 {
	length := r.URL.Query()["length"]
	if len(length) == 0 {
		return defaultLength
	}

	v, _ := strconv.ParseInt(length[0], 10, 64)
	return v
}

func DtSort(r *http.Request) int64 {
	_ = r.ParseForm()
	sort := r.Form["order[0][column]"]
	if len(sort) == 0 {
		return defaultSort
	}

	v, _ := strconv.ParseInt(sort[0], 10, 64)
	return v
}

func DtDir(r *http.Request) string {
	_ = r.ParseForm()
	dir := r.Form["order[0][dir]"]
	if len(dir) == 0 {
		return direction["asc"]
	}

	switch dir[0] {
	case "asc":
		return direction[dir[0]]
	case "desc":
		return direction[dir[0]]
	default:
		return direction["asc"]
	}
}

package handlers

import (
	"net/http"
	"testing"
)

func TestAnimeQuotes(t *testing.T) {
	type args struct {
		rw http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AnimeQuotes(tt.args.rw, tt.args.r)
		})
	}
}

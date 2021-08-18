package repository

import (
	"reflect"
	"testing"
)

func TestNewMemoryCache(t *testing.T) {
	tests := []struct {
		name string
		want MemoryCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

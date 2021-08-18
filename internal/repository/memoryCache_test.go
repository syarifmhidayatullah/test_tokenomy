package repository

import (
	"reflect"
	"testing"

	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func TestMemoryCache_GetDataCache(t *testing.T) {
	type fields struct {
		DataMap map[int]entity.Data
	}
	tests := []struct {
		name        string
		fields      fields
		wantDataMap map[int]entity.Data
	}{
		{
			name: "test case success",
			fields: fields{
				DataMap: map[int]entity.Data{
					1: {Id: 1, Name: "A"},
					2: {Id: 2, Name: "B"},
					3: {Id: 3, Name: "C"},
				},
			},
			wantDataMap: map[int]entity.Data{
				1: {Id: 1, Name: "A"},
				2: {Id: 2, Name: "B"},
				3: {Id: 3, Name: "C"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				DataMap: tt.fields.DataMap,
			}
			if gotDataMap := m.GetDataCache(); !reflect.DeepEqual(gotDataMap, tt.wantDataMap) {
				t.Errorf("MemoryCache.GetDataCache() = %v, want %v", gotDataMap, tt.wantDataMap)
			}
		})
	}
}

func TestMemoryCache_SetDataCache(t *testing.T) {
	type fields struct {
		DataMap map[int]entity.Data
	}
	type args struct {
		dataMap map[int]entity.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test case success",
			args: args{
				dataMap: map[int]entity.Data{
					1: {Id: 1, Name: "A"},
					2: {Id: 2, Name: "B"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				DataMap: tt.fields.DataMap,
			}
			m.SetDataCache(tt.args.dataMap)
		})
	}
}

func TestMemoryCache_ConstructData(t *testing.T) {
	type fields struct {
		DataMap map[int]entity.Data
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test case success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				DataMap: tt.fields.DataMap,
			}
			m.ConstructData()
		})
	}
}

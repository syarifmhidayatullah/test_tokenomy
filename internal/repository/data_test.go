package repository

import (
	"reflect"
	"testing"

	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func TestDataRepo_GetDataList(t *testing.T) {
	type fields struct {
		Cache MemoryCache
	}
	tests := []struct {
		name         string
		fields       fields
		wantDataList []entity.Data
	}{
		{
			name: "test case 1 success",
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 2, Name: "B"}, {Id: 3, Name: "C"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataRepo{
				Cache: tt.fields.Cache,
			}
			if gotDataList := d.GetDataList(); !reflect.DeepEqual(gotDataList, tt.wantDataList) {
				t.Errorf("DataRepo.GetDataList() = %v, want %v", gotDataList, tt.wantDataList)
			}
		})
	}
}

func TestDataRepo_GetDataListByParam(t *testing.T) {
	type fields struct {
		Cache MemoryCache
	}
	type args struct {
		param []int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantDataList []entity.Data
	}{
		{
			name: "test case 1 success",
			args: args{
				param: []int{1, 2},
			},
			fields: fields{
				Cache: MemoryCache{
					DataMap: map[int]entity.Data{
						1: {Id: 1, Name: "A"},
						2: {Id: 2, Name: "B"},
						3: {Id: 3, Name: "C"},
					},
				},
			},
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 2, Name: "B"},
			},
		},
		{
			name: "test case 2 some id not found",
			args: args{
				param: []int{1, 2, 5},
			},
			fields: fields{
				Cache: MemoryCache{
					DataMap: map[int]entity.Data{
						1: {Id: 1, Name: "A"},
						2: {Id: 2, Name: "B"},
						3: {Id: 3, Name: "C"},
					},
				},
			},
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 2, Name: "B"},
			},
		},
		{
			name: "test case 3 all id not found",
			args: args{
				param: []int{5, 6},
			},
			fields: fields{
				Cache: MemoryCache{
					DataMap: map[int]entity.Data{
						1: {Id: 1, Name: "A"},
						2: {Id: 2, Name: "B"},
						3: {Id: 3, Name: "C"},
					},
				},
			},
			wantDataList: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataRepo{
				Cache: tt.fields.Cache,
			}
			if gotDataList := d.GetDataListByParam(tt.args.param); !reflect.DeepEqual(gotDataList, tt.wantDataList) {
				t.Errorf("DataRepo.GetDataListByParam() = %v, want %v", gotDataList, tt.wantDataList)
			}
		})
	}
}

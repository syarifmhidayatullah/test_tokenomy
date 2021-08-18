package usecase

import (
	"reflect"
	"testing"

	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
	repo "github.com/syarifmhidayatullah/test_tokenomy/internal/repository"
)

//better to use github.com/golang/mock/mockgen and create mock
//instead of init all repository in unit test
//but the rule is not allow to use external lib
func InitTest() (dataRepo repo.DataRepo) {
	cache := repo.NewMemoryCache()
	cache.ConstructData()
	dataRepo = repo.NewDataRepo(cache)
	return
}

func TestDataCase_GetDataList(t *testing.T) {

	type fields struct {
		DataRepo repo.DataRepo
	}
	tests := []struct {
		name         string
		fields       fields
		wantDataList []entity.Data
	}{
		{
			name: "test case 1 ",
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 2, Name: "B"}, {Id: 3, Name: "C"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataCase{
				DataRepo: tt.fields.DataRepo,
			}
			if gotDataList := d.GetDataList(); !reflect.DeepEqual(gotDataList, tt.wantDataList) {
				t.Errorf("DataCase.GetDataList() = %v, want %v", gotDataList, tt.wantDataList)
			}
		})
	}
}

func TestDataCase_GetDataListByParam(t *testing.T) {
	repoData := InitTest()
	type fields struct {
		DataRepo repo.DataRepo
	}
	type args struct {
		param string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantDataList []entity.Data
		wantErrType  int
		wantErr      bool
	}{
		{
			name: "test case 1 success singel param",
			args: args{
				param: "1",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"},
			},
			wantErrType: 200,
		},
		{
			name: "test case 2 failed singel param bad request",
			args: args{
				param: "xxx",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: []entity.Data{},
			wantErrType:  400,
			wantErr:      true,
		},
		{
			name: "test case 3 failed singel param not found",
			args: args{
				param: "5",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: nil,
			wantErrType:  404,
			wantErr:      true,
		},
		{
			name: "test case 4 success multi param",
			args: args{
				param: "1,2",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 2, Name: "B"},
			},
			wantErrType: 200,
		},
		{
			name: "test case 5 success multi param with one non int value",
			args: args{
				param: "1,s,3",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: []entity.Data{
				{Id: 1, Name: "A"}, {Id: 3, Name: "C"},
			},
			wantErrType: 200,
		},
		{
			name: "test case 6 multi param with all non int value",
			args: args{
				param: "a,b,c",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: []entity.Data{},
			wantErrType:  400,
			wantErr:      true,
		},
		{
			name: "test case 7 multi param with all not found value",
			args: args{
				param: "5,6,7",
			},
			fields: fields{
				DataRepo: repoData,
			},
			wantDataList: nil,
			wantErrType:  404,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataCase{
				DataRepo: tt.fields.DataRepo,
			}
			gotDataList, gotErrType, err := d.GetDataListByParam(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataCase.GetDataListByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDataList, tt.wantDataList) {
				t.Errorf("DataCase.GetDataListByParam() gotDataList = %v, want %v", gotDataList, tt.wantDataList)
			}
			if gotErrType != tt.wantErrType {
				t.Errorf("DataCase.GetDataListByParam() gotErrType = %v, want %v", gotErrType, tt.wantErrType)
			}
		})
	}
}

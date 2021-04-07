//package inmemory contains an inmemory store for our usecase.
// ONLY TO BE USED IN DEVELOPMENT!!
package inmemory

import (
	"reflect"
	"testing"
)

func TestNewInMemoryClient(t *testing.T) {
	tests := []struct {
		name string
		want *InMemoryClient
	}{
		{"New_Client", &InMemoryClient{
			contents: []inMemoryObject{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryClient() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestInMemoryClient_Read(t *testing.T) {
	type fields struct {
		contents []inMemoryObject
	}
	type args struct {
		s string
	}
	fieldsCases := fields{
		contents: []inMemoryObject{
			{
				key:       "first",
				byteArray: []byte("firstcontent"),
			},
			{
				key:       "second",
				byteArray: []byte("secondcontent"),
			},
			{
				key:       "third",
				byteArray: []byte("thirdcontent"),
			},
			{
				key:       "fourth",
				byteArray: []byte("fourthcontent"),
			},
			{
				key:       "fifth",
				byteArray: []byte("fifthcontent"),
			},
			{
				key:       "sixth",
				byteArray: []byte("sixthcontent"),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "ReadFirst",
			fields: fieldsCases,
			args: args{
				s: "first",
			},
			want:    []byte("firstcontent"),
			wantErr: false,
		},
		{
			name:   "ReadWrong",
			fields: fieldsCases,
			args: args{
				s: "seventh",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "ReadFifth",
			fields: fieldsCases,
			args: args{
				s: "fifth",
			},
			want:    []byte("fifthcontent"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryClient{
				contents: tt.fields.contents,
			}
			got, err := i.Read(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryClient.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryClient.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: This test does not actually test if the deleting works, just if errors are returned. Needs to be reworked in the future.
func TestInMemoryClient_Delete(t *testing.T) {
	type fields struct {
		contents []inMemoryObject
	}
	type args struct {
		s string
	}
	fieldsCases := fields{
		contents: []inMemoryObject{
			{
				key:       "first",
				byteArray: []byte("firstcontent"),
			},
			{
				key:       "second",
				byteArray: []byte("secondcontent"),
			},
			{
				key:       "third",
				byteArray: []byte("thirdcontent"),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "DeleteFirst",
			fields: fieldsCases,
			args: args{
				s: "first",
			},
			wantErr: false,
		},
		{
			name:   "DeleteNonExisting",
			fields: fieldsCases,
			args: args{
				s: "seventh",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryClient{
				contents: tt.fields.contents,
			}
			if err := i.Delete(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryClient.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryClient_Create(t *testing.T) {
	type fields struct {
		contents []inMemoryObject
	}
	type args struct {
		s string
		b []byte
	}
	fieldsCases := fields{
		contents: []inMemoryObject{
			{
				key:       "first",
				byteArray: []byte("firstcontent"),
			},
			{
				key:       "second",
				byteArray: []byte("secondcontent"),
			},
			{
				key:       "third",
				byteArray: []byte("thirdcontent"),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "CreateNew",
			fields: fieldsCases,
			args: args{
				s: "newKey",
				b: []byte("bloop"),
			},
			wantErr: false,
		},
		{
			name:   "Create duplicate",
			fields: fieldsCases,
			args: args{
				s: "first",
				b: []byte("bloop"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryClient{
				contents: tt.fields.contents,
			}
			if err := i.Create(tt.args.s, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryClient.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: Does not actually test the output, just the returned error, like delete. Requires rework!
func TestInMemoryClient_UpdateKey(t *testing.T) {
	type fields struct {
		contents []inMemoryObject
	}
	type args struct {
		s string
		b []byte
	}
	fieldsCases := fields{
		contents: []inMemoryObject{
			{
				key:       "first",
				byteArray: []byte("firstcontent"),
			},
			{
				key:       "second",
				byteArray: []byte("secondcontent"),
			},
			{
				key:       "third",
				byteArray: []byte("thirdcontent"),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "UpdateExisting",
			fields: fieldsCases,
			args: args{
				s: "first",
				b: []byte("bloop"),
			},
			wantErr: false,
		},
		{
			name:   "UpdateNonExisting",
			fields: fieldsCases,
			args: args{
				s: "firstsecondthird",
				b: []byte("bloop"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryClient{
				contents: tt.fields.contents,
			}
			if err := i.UpdateKey(tt.args.s, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryClient.UpdateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryClient_findInSlice(t *testing.T) {
	type fields struct {
		contents []inMemoryObject
	}
	type args struct {
		s string
	}
	fieldsCases := fields{
		contents: []inMemoryObject{
			{
				key:       "first",
				byteArray: []byte("firstcontent"),
			},
			{
				key:       "second",
				byteArray: []byte("secondcontent"),
			},
			{
				key:       "third",
				byteArray: []byte("thirdcontent"),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "FindFirst",
			fields: fieldsCases,
			args: args{
				s: "first",
			},
			want:    0,
			wantErr: false,
		},
		{
			name:   "FindThird",
			fields: fieldsCases,
			args: args{
				s: "third",
			},
			want:    2,
			wantErr: false,
		},
		{
			name:   "FindNonExisting",
			fields: fieldsCases,
			args: args{
				s: "thirdsecondfirst",
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryClient{
				contents: tt.fields.contents,
			}
			got, err := i.findInSlice(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryClient.findInSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InMemoryClient.findInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

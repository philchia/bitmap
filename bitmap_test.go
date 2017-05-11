package bitmap

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name: "new",
			args: args{
				10,
			},
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.length); tt.wantNil != (got == nil) {
				t.Errorf("New() = %v, want nil %v", got, tt.wantNil)
			}
		})
	}
}

func TestBitMap_Set(t *testing.T) {
	type fields struct {
		length int
	}
	type args struct {
		idx uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "new",
			fields: fields{
				10,
			},
			args: args{
				10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(tt.fields.length)
			b.Set(tt.args.idx)
			if !b.Exists(tt.args.idx) {
				t.Error("Set index but not exists")
			}
		})
	}
}

func TestBitMap_Exists(t *testing.T) {
	type fields struct {
		length int
	}
	type args struct {
		idx uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "case1",
			fields: fields{
				10,
			},
			args: args{
				9,
			},
			want: true,
		},
		{
			name: "case2",
			fields: fields{
				10,
			},
			args: args{
				2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(tt.fields.length)
			b.Set(tt.args.idx)
			if got := b.Exists(tt.args.idx); got != tt.want {
				t.Errorf("BitMap.Exists(%d) = %v, want %v", tt.args.idx, got, tt.want)
			}
		})
	}
}

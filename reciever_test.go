package fbmessenger

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	file, err := ioutil.ReadFile("example_messages.json")
	if err != nil {
		panic(err)
	}
	filebytes := bytes.NewReader(file)
	


	type args struct {
		body io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example_message",args{filebytes},"is it 1?"},
	}


		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.body).Entry[0].Messaging[0].Message.Text; !reflect.DeepEqual(got,tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMin(t *testing.T) {
	file, err := ioutil.ReadFile("example_messages.json")
	if err != nil {
		panic(err)
	}
	filebytes := bytes.NewReader(file)
	

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example_message",args{filebytes},"is it 1?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseMin(tt.args.body).Message.Text; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

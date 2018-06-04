package fbmessenger

import "testing"

func TestSendMessage(t *testing.T) {
	type args struct {
		message         string
		recipientid     string
		pageaccesstoken string
	}
	tests := []struct {
		name string
		args args
	}{	
		{"Hello World",{"Hello World","123123","123123"},},
		{"Hello World",{"Hello World","123123","123123"},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMessage(tt.args.message, tt.args.recipientid, tt.args.pageaccesstoken)
		})
	}
}

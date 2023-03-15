package http

import (
	"reflect"
	"testing"
)

func Test_fileListFromHTML(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRet []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok",
			args: args{
				content: []byte(`<html>
<head><title>Index of /xx/xx/xx/</title></head>
<body>
<h1>Index of /xx/xx/xx/<h1><hr><pre><a href="../">../</a>
<a href="xx/">xx/</a>    15-Mar-2032 11:23    -
<a href="abc.txt">abc.txt</a>    15-Mar-2032 11:24    -
<a href="abc.tar.gz">abc.tar.gz</a> 15-Mar-2032 11:24    -
</pre><hr></body>
</html>`),
			},
			wantRet: []string{"abc.txt", "abc.tar.gz"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := fileListFromHTML(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileListFromHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("fileListFromHTML() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

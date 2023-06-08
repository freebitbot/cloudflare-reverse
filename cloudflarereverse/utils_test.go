package cloudflarereverse

import (
	"fmt"
	fpclient "github.com/0xF7A4C6/fingerprint-client/fingerprintclient"
	"testing"
)

func Test_formatFingerprint(t *testing.T) {
	brFp, err := fpclient.LoadFingerprint(&fpclient.LoadingConfig{
		FilePath: "./fp.json",
	})

	if err != nil {
		panic(err)
	}

	type args struct {
		fp *fpclient.Fingerprint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test-format-fingerprint",
			args: args{
				fp: brFp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := formatFingerprint(tt.args.fp)
			fmt.Println(res)
		})
	}
}

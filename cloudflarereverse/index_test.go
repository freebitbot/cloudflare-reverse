package cloudflarereverse

import (
	"fmt"
	fp "github.com/0xF7A4C6/fingerprint-client/fingerprintclient"
	"testing"
)

func TestGetCfbm(t *testing.T) {
	brFp, err := fp.LoadFingerprint(&fp.LoadingConfig{
		FilePath: "./fp.json",
	})

	if err != nil {
		panic(err)
	}

	type args struct {
		brFp  *fp.Fingerprint
		proxy string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				brFp:  brFp,
				proxy: "http://cap1217:gLAty@172.252.8.224:61234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCfbm(tt.args.brFp, tt.args.proxy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCfbm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(got)
		})
	}
}

package main

import (
	"fmt"

	cf "github.com/0xF7A4C6/cloudflare-reverse/cloudflarereverse"
	fp "github.com/0xF7A4C6/cloudflare-reverse/pkg/fpclient"
)

func main() {
	brFp, err := fp.LoadFingerprint(&fp.LoadingConfig{
		FilePath: "../assets/fingerprints/macos.json",
	})

	if err != nil {
		panic(err)
	}

	cf.Init()

	for i := 0; i < 5; i++ {
		cookie, err := cf.GetCfbm(brFp, "")

		if err != nil {
			panic(err)
		}

		fmt.Println(cookie)
	}
}

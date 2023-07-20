package main

import (
	"fmt"

	cf "github.com/Implex-ltd/cloudflare-reverse/cloudflarereverse"
	fp "github.com/Implex-ltd/fingerprint-client/fpclient"
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

package cloudflarereverse

import (
	"encoding/json"
	fpclient "github.com/0xF7A4C6/fingerprint-client/fingerprintclient"
)

type FpStruct map[string][]string

func formatFingerprint(fp *fpclient.Fingerprint) string {
	var fps = FpStruct{}

	fps[fp.Navigator.UserAgent] = []string{"n.userAgent"}

	payload, _ := json.Marshal(fps)
	return string(payload)
}

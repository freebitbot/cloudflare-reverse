package cloudflarereverse

type Fingerprint struct {
	Results []string `json:"results"`
	//Timing  int      `json:"timing"`
}

type FingerprintPayload struct {
	Wp string `json:"wp"`
	S  string `json:"s"`

	/*
		* Params removed on current version.
			Src string  `json:"src"`
			M   string  `json:"m"`
			T   float64 `json:"t"`
			Fp Fingerprint `json:"fp"
	*/
}

type CfParams struct {
	R string `json:"r"`
	M string `json:"m"`
}

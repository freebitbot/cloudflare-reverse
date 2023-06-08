package cloudflarereverse

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/0xF7A4C6/clean-http/cleanhttp"
	fp "github.com/0xF7A4C6/fingerprint-client/fingerprintclient"
	http "github.com/bogdanfinn/fhttp"
	"io/ioutil"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	re = regexp.MustCompile(`[0-9]*\.[0-9]+:[0-9]+:`)
)

// Need improvement
func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func randHexString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func GetCfbm(brFp *fp.Fingerprint, proxy string) (string, error) {
	header := http.Header{
		`authority`:          {`discord.com`},
		`accept`:             {`*/*`},
		`accept-language`:    {`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		`content-type`:       {`application/json`},
		`origin`:             {`https://discord.com`},
		`sec-ch-ua`:          {`"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		`sec-ch-ua-mobile`:   {`?0`},
		`sec-ch-ua-platform`: {`"Windows"`},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-origin`},
		`user-agent`:         {`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`},

		http.HeaderOrderKey: {
			"authority",
			"accept",
			"accept-language",
			"content-type",
			"origin",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"user-agent",
		},
	}

	client, err := cleanhttp.NewCleanHttpClient(&cleanhttp.Config{
		Proxy:     proxy,
		BrowserFp: brFp,
	})
	if err != nil {
		return "", err
	}

	resp, err := client.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com",
		Header: header,
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	Cf := &CfParams{
		R: strings.Split(strings.Split(string(response), "r:'")[1], "',m")[0],
		M: strings.Split(strings.Split(string(response), "m:'")[1], "'};")[0],
	}

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/cdn-cgi/challenge-platform/h/g/scripts/jsd/5da7637f/invisible.js",
		Header: header,
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	JsScript := string(response)

	var Pass string
	for _, char := range strings.Split(strings.Split(JsScript, "'.split(';')")[0], ";") {
		if len(char) == 65 {
			Pass = char
		}
	}

	fmt.Println(Pass)

	if Pass == "" {
		return "", fmt.Errorf("cant find encryption password")
	}

	Base := re.FindString(JsScript)
	S := Base + strings.Split(JsScript, Base)[1][:43]

	timing := float64(randInt(100, 400))

	jsonPayload, _ := json.Marshal(FingerprintPayload{
		Src: "worker",
		T:   float64(timing+float64(randInt(300, 500))) + rand.Float64(),
		S:   S,
		Fp: Fingerprint{
			Results: []string{
				randHexString(16),
			},
			Timing: int(timing),
		},
		M:  Cf.M,
		Wp: "lol",
	})

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://discord.com/cdn-cgi/challenge-platform/h/b/cv/result/%s", Cf.R),
		Header: header,
		Body:   strings.NewReader(string(jsonPayload)),
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("cant submit answer")
	}

	for _, c := range client.CookieJar.Cookies(&url.URL{Host: "discord.com"}) {
		if c.Name == "__cf_bm" {
			return c.Value, nil
		}
	}

	return "", fmt.Errorf("no cookie found")
}

/*
IZFYTPD1_lM2VM5nOh0FbrDgjukwqsz7OIcGAdne1Co-1686198323-0-AQEtAZUBbcbv8LNnAxXhS7adfFX/hZEChdL7+8zU4go3RQHswqMDYLTuQBcMbguF/Q==
U66urRyiVu.D1EuofKk6xs2se6nKyXs2BLkczQMzd.4-1686198435-0-ASfuPmAjrRYSwMtIA6d1+jnQEcNGXqPrOOH+j8kL5Y+z27lPS2D6KyB50jng1c7rhA==
0Sfd7tb0gJ_JrWANsOFi_zZ2VnAuGe5Ps4Jbj850pfs-1686196343-0-AcDqhmOFXD00XBl8vMpV14tpYIqa5YhttXni4dPDa8OuBsm+/hlmBYhs/IMUzfN4/g==
*/

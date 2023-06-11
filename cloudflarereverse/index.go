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
	"regexp"
	"strings"
	"time"
)

var (
	re = regexp.MustCompile(`[0-9]*\.[0-9]+:[0-9]+:`)
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randHexString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func GetCfbm(brFp *fp.Fingerprint, proxy string) (string, error) {
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

	if Pass == "" {
		return "", fmt.Errorf("cant find encryption password")
	}

	Base := re.FindString(JsScript)
	S := Base + strings.Split(JsScript, Base)[1][:43]

	timing := float64(randInt(100, 400))

	jsonPayload, _ := json.Marshal(FingerprintPayload{
		Src: "worker",
		T:   float64(timing+float64(randInt(300, 700))) + rand.Float64(),
		S:   S,
		Fp: Fingerprint{
			Results: []string{
				randHexString(16),
				randHexString(16),
			},
			Timing: int(timing),
		},
		M:  Cf.M,
		Wp: strings.Split(Compress(formatFingerprint(brFp), Pass), "===")[0],
	})

	p := string(jsonPayload)
	head := client.GenerateBaseHeaders()

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://discord.com/cdn-cgi/challenge-platform/h/b/cv/result/%s", Cf.R),
		Header: http.Header{
			`accept`:             {`*/*`},
			`accept-encoding`:    {`gzip, deflate, br`},
			`accept-language`:    {head.AcceptLanguage},
			`content-type`:       {`application/json`},
			`cookie`:             {client.FormatCookies()},
			`origin`:             {"https://discord.com"},
			`referer`:            {`https://discord.com/channels/@me`},
			`sec-ch-ua`:          {head.SecChUa},
			`sec-ch-ua-mobile`:   {`?0`},
			`sec-ch-ua-platform`: {head.SecChUaPlatform},
			`sec-fetch-dest`:     {`empty`},
			`sec-fetch-mode`:     {`cors`},
			`sec-fetch-site`:     {`same-origin`},
			`user-agent`:         {brFp.Navigator.UserAgent},

			http.HeaderOrderKey: {
				`accept`,
				`accept-encoding`,
				`accept-language`,
				`content-type`,
				`cookie`,
				`origin`,
				`referer`,
				`sec-ch-ua`,
				`sec-ch-ua-mobile`,
				`sec-ch-ua-platform`,
				`sec-fetch-dest`,
				`sec-fetch-mode`,
				`sec-fetch-site`,
				`user-agent`,
			},
		},
		Body: strings.NewReader(p),
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("cant submit answer")
	}

	for _, c := range client.Cookies {
		if c.Name == "__cf_bm" {
			return c.Value, nil
		}
	}

	return "", fmt.Errorf("no cookie found")
}

/*
flag:  qWktQ9T0fVhNF6JvbS_ivhH9L5dh0yI4CriPg773xK8-1686372156-0-AUQMpvymJKZcjOeXgqzMHMeXjjFC79uoYEAuUccQ9TMrAfh9OcgoSMHR0YFbtnFbVQ==
clean: mRXc.JAC.UuIBz3P6LOtHuhzWSa4CSnloZsAMb_5NcE-1686372122-0-AQ1lIBvkNn8ns3eTf6+ww4ONfvHLXg2e19bjAq2Zw1EiVeov1gu5tXxYwvMBB0v87w==
*/

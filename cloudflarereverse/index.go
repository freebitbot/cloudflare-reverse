package cloudflarereverse

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	fp "github.com/Implex-ltd/fingerprint-client/fpclient"
	http "github.com/bogdanfinn/fhttp"
)

var (
	re      = regexp.MustCompile(`[0-9]*\.[0-9]+:[0-9]+:`)
	version = "5da7637f"
)

func Init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
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

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	Cf := &CfParams{
		R: strings.Split(strings.Split(string(response), "r:'")[1], "',m")[0],
		//	M: strings.Split(strings.Split(string(response), "m:'")[1], "'};")[0],
	}

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    fmt.Sprintf("https://discord.com/cdn-cgi/challenge-platform/h/g/scripts/jsd/%s/invisible.js", version),
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
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

	//timing := float64(randInt(100, 250))

	jsonPayload, _ := json.Marshal(FingerprintPayload{ //nolint: errcheck
		/*Src: "worker",
		T:   float64(timing+float64(randInt(50, 100))) + rand.Float64(),*/
		S: S,
		/*Fp: Fingerprint{
			Results: []string{
				randHexString(16),
			},
			//Timing: int(timing),
		},
		M:  Cf.M,*/
		Wp: strings.Split(Compress(FormatFingerprint(brFp), Pass), "===")[0],
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
			`cookie`:             {head.Cookies},
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
				`content-length`,
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
		return "", fmt.Errorf("cant submit answer status code %d", resp.StatusCode)
	}

	for _, c := range client.Cookies {
		if c.Name == "__cf_bm" {
			return c.Value, nil
		}
	}

	return "", fmt.Errorf("no cookie found")
}

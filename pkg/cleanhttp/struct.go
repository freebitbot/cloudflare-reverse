package cleanhttp

import (
	"io"

	fp "github.com/IIayk122/cloudflare-reverse/pkg/fpclient"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type CleanHTTP struct {
	Config    *Config
	TLSClient tls_client.HttpClient
	Cookies   []*http.Cookie
}

type Config struct {
	Proxy     string
	Timeout   int
	BrowserFp *fp.Fingerprint
	TLSFp     *fp.TLSFingerprint
	Profile   tls_client.ClientProfile
}

type RequestOption struct {
	Method string
	Body   io.Reader
	URL    string
	Header http.Header
}

type UserAgentInfo struct {
	BrowserName    string
	BrowserVersion string
	OSName         string
	OSVersion      string
	UaVersion      string
}

type HeaderBuilder struct {
	SecChUa         string
	SecChUaPlatform string
	secChUaMobile   string
	AcceptLanguage  string
	Cookies         string
}

package fpclient

import (
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/playwright-community/playwright-go"
)

// Load tls fingerprint and return *TLSFingerprint.
func LoadTLSFingerprint(config *LoadingConfig) (*TLSFingerprint, error) {
	var fpStr string

	if config.FilePath != "" {
		data, err := os.ReadFile(config.FilePath)
		if err != nil {
			return nil, err
		}

		fpStr = string(data)
	}

	if config.String != "" {
		fpStr = config.String

		if config.Isbase64String {
			str, err := base64.StdEncoding.DecodeString(fpStr)
			if err != nil {
				return nil, err
			}

			fpStr = string(str)
		}
	}

	var fp TLSFingerprint
	if err := json.Unmarshal([]byte(fpStr), &fp); err != nil {
		return nil, err
	}

	return &fp, nil
}

// Dump TLS fingerprint from chromedriver
func DumpTLSFingerprint(config *DumpTLSConfig) (*TLSFingerprint, error) {
	err := playwright.Install()
	if err != nil {
		return nil, err
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(config.Headless),
	})
	if err != nil {
		return nil, err
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{
		IgnoreHttpsErrors: playwright.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	resp, err := page.Goto("https://tls.peet.ws/api/all")
	if err != nil {
		return nil, err
	}

	response, err := resp.Text()
	if err != nil {
		return nil, err
	}

	var fp TLSFingerprint
	if errU := json.Unmarshal([]byte(response), &fp); errU != nil {
		return nil, errU
	}

	err = os.WriteFile(config.OutputPath, []byte(response), 0644)
	if err != nil {
		return nil, err
	}

	return &fp, nil
}

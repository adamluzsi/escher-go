package signer_test

import (
	"testing"

	escher "github.com/adamluzsi/escher-go"
	"github.com/adamluzsi/escher-go/signer"
	. "github.com/adamluzsi/escher-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestCanonicalizeRequest(t *testing.T) {
	t.Log("CanonicalizeRequest should return with a proper string")
	EachTestConfigFor(t, "signRequest", func(config escher.Config, testConfig TestConfig) bool {
		if testConfig.Expected.CanonicalizedRequest == "" {
			return false
		}

		canonicalizedRequest := signer.New(config).CanonicalizeRequest(testConfig.Request, testConfig.HeadersToSign)
		return assert.Equal(t, canonicalizedRequest, testConfig.Expected.CanonicalizedRequest, "canonicalizedRequest should be eq")
	})
}

func TestGetStringToSign(t *testing.T) {
	t.Log("GetStringToSign should return with a proper string")
	EachTestConfigFor(t, "signRequest", func(config escher.Config, testConfig TestConfig) bool {
		if testConfig.Expected.StringToSign == "" {
			return false
		}

		stringToSign := signer.New(config).GetStringToSign(testConfig.Request, testConfig.HeadersToSign)
		return assert.Equal(t, stringToSign, testConfig.Expected.StringToSign, "stringToSign expected to eq with the test config expectation")
	})
}

func TestGenerateHeader(t *testing.T) {
	t.Log("GenerateHeader should return with a proper string")
	EachTestConfigFor(t, "signRequest", func(config escher.Config, testConfig TestConfig) bool {
		if testConfig.Expected.AuthHeader == "" {
			return false
		}

		authHeader := signer.New(config).GenerateHeader(testConfig.Request, testConfig.HeadersToSign)
		return assert.Equal(t, testConfig.Expected.AuthHeader, authHeader, "authHeader generation failed")
	})
}

func TestSignRequest(t *testing.T) {
	t.Log("SignRequest should return with a properly signed request")
	EachTestConfigFor(t, "signRequest", func(config escher.Config, testConfig TestConfig) bool {
		if testConfig.Expected.Request.Method == "" {
			return false
		}

		request := signer.New(config).SignRequest(testConfig.Request, testConfig.HeadersToSign)
		return assert.Equal(t, testConfig.Expected.Request, request, "Requests should be eq")
	})
}
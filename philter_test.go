package philter

import (
	"crypto/tls"
	"net/http"
	"testing"
)

func TestFilter(t *testing.T) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	filterResponse := Filter("https://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default")

	if filterResponse.FilteredText != "His SSN was {{{REDACTED-ssn}}}." {
		t.Fail()
	}

}

func TestExplain(t *testing.T) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	explainResponse := Explain("https://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default")

	if explainResponse.FilteredText != "His SSN was {{{REDACTED-ssn}}}." {
		t.Fail()
	}

}

func TestStatus(t *testing.T) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	statusResponse := Status("https://localhost:8080")

	if statusResponse.Status != "Healthy" {
		if statusResponse.Status != "Unhealthy" {
			t.Fail()
		}
	}

	if statusResponse.Version == "" {
		t.Fail()
	}

}
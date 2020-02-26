package philter

import (
	"testing"
)

func TestFilter(t *testing.T) {

	filterResponse := Filter("http://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default")

	if filterResponse.FilteredText != "His SSN was {{{REDACTED-ssn}}}." {
		t.Fail()
	}

}

func TestExplain(t *testing.T) {

	explainResponse := Explain("http://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default")

	if explainResponse.FilteredText != "His SSN was {{{REDACTED-ssn}}}." {
		t.Fail()
	}

}

func TestStatus(t *testing.T) {

	statusResponse := Status("http://localhost:8080")

	if statusResponse.Status != "Healthy" {
		if statusResponse.Status != "Unhealthy" {
			t.Fail()
		}
	}

	if statusResponse.Version == "" {
		t.Fail()
	}

}
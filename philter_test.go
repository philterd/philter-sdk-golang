/*
 * Copyright 2023 Philterd, LLC.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

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

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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type FilterResponse struct {
	FilteredText string `json:"filteredText"`
	Context      string `json:"context"`
	DocumentId   string `json:"documentId"`
}

type ExplainResponse struct {
	FilteredText string      `json:"filteredText"`
	Context      string      `json:"context"`
	DocumentId   string      `json:"documentId"`
	Explanation  Explanation `json:"explanation"`
}

type Explanation struct {
	AppliedSpans []Span `json:"appliedSpans"`
	IgnoredSpans []Span `json:"ignoredSpans"`
}

type Span struct {
	Id             string  `json:"id"`
	CharacterStart int     `json:"characterStart"`
	CharacterEnd   int     `json:"characterEnd"`
	FilterType     string  `json:"filterType"`
	Context        string  `json:"context"`
	DocumentId     string  `json:"documentId"`
	Confidence     float64 `json:"confidence"`
	Text           string  `json:"text"`
	Replacement    string  `json:"replacement"`
	Salt           string  `json:"salt"`
	Ignored        bool    `json:"ignored"`
}

func Status(endpoint string) StatusResponse {

	response, err := http.Get(endpoint + "/api/status")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var responseObject StatusResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}

func Filter(endpoint string, input string, context string, documentId string, policy string) FilterResponse {

	var text = []byte(input)

	base, err := url.Parse(endpoint + "/api/filter")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("c", context)
	params.Add("d", documentId)
	params.Add("p", policy)

	base.RawQuery = params.Encode()

	request, err := http.NewRequest("POST", base.String(), bytes.NewReader(text))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header.Add("Content-Type", "text/plain")

	client := &http.Client{}
	response, err := client.Do(request)

	documentId = "empty" // response.Header.Get("x-document-id")

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	response.Body.Close()

	return FilterResponse{FilteredText: string(responseData), Context: context, DocumentId: documentId}

}

func Explain(endpoint string, input string, context string, documentId string, policy string) ExplainResponse {

	var text = []byte(input)

	base, err := url.Parse(endpoint + "/api/explain")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("c", context)
	params.Add("d", documentId)
	params.Add("p", policy)

	base.RawQuery = params.Encode()

	request, err := http.NewRequest("POST", base.String(), bytes.NewReader(text))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header.Add("Content-Type", "text/plain")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var explainResponse ExplainResponse
	json.Unmarshal(responseData, &explainResponse)

	return explainResponse

}

func GetPolicyNames(endpoint string) []string {

	request, err := http.NewRequest("GET", endpoint+"/api/policies", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var responseObject []string
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}

func GetPolicy(endpoint string, name string) string {

	request, err := http.NewRequest("GET", endpoint+"/api/policies/"+name, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)

}

func UploadPolicy(endpoint string, name string, content string) bool {

	var json = []byte(content)
	request, err := http.NewRequest("POST", endpoint+"/api/policies", bytes.NewBuffer(json))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

func DeletePolicy(endpoint string, name string, content string) bool {

	request, err := http.NewRequest("DELETE", endpoint+"/api/policies", nil)

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

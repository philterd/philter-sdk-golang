/*
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

func Filter(endpoint string, input string, context string, documentId string, filterProfile string) FilterResponse {

	var text = []byte(input)

	base, err := url.Parse(endpoint + "/api/filter")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("c", context)
	params.Add("d", documentId)
	params.Add("p", filterProfile)

	base.RawQuery = params.Encode()

	response, err := http.Post(base.String(), "text/plain", bytes.NewBuffer(text))

	if err != nil {
		log.Fatal(err)
	}

	documentId = response.Header.Get("x-document-id")

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	response.Body.Close()

	return FilterResponse{FilteredText:string(responseData), Context:context, DocumentId:documentId}

}

func Explain(endpoint string, input string, context string, documentId string, filterProfile string) ExplainResponse {

	var text = []byte(input)

	base, err := url.Parse(endpoint + "/api/explain")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("c", context)
	params.Add("d", documentId)
	params.Add("p", filterProfile)

	base.RawQuery = params.Encode()

	response, err := http.Post(base.String(), "text/plain", bytes.NewBuffer(text))

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


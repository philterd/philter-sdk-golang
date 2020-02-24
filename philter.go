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
	"os"
)

type StatusResponse struct {
	status string `json:"status"`
	version string `json:"version"`
}

type FilterResponse struct {
	FilteredText string	`json:"filteredText"`
	Context	string `json:"context"`
	DocumentId string `json:"documentId"`
}

type ExplainResponse struct {
	FilteredText string	`json:"filteredText"`
	Context	string `json:"context"`
	DocumentId string `json:"documentId"`
}

func Status() StatusResponse {

	response, err := http.Get("https://localhost:8080/api/status")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject StatusResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}

func Filter(input string, context string, documentId string, filterProfile string) FilterResponse {

	var text = []byte(input)
	response, err := http.Post("https://localhost:8080/api/filter", "text/plain", bytes.NewBuffer(text))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	documentId = response.Header.Get("x-document-id")
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return FilterResponse{FilteredText:string(responseData), Context:context, DocumentId:documentId}

}

func Explain(input string, context string, documentId string, filterProfile string) ExplainResponse {

	var text = []byte(input)
	response, err := http.Post("https://localhost:8080/api/explain", "text/plain", bytes.NewBuffer(text))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject ExplainResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}


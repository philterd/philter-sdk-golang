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

package fpr

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
	Status  string `json:"status"`
	Version string `json:"version"`
}

func GetFilterProfileNames(endpoint string, token string) []string {

	request, err := http.NewRequest("GET", endpoint + "/api/profiles", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if token != "" {
		request.Header.Add("Authorization", "token:" + token)
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

func GetFilterProfile(endpoint string, name string, token string) string {

	request, err := http.NewRequest("GET", endpoint + "/api/profiles/" + name, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if token != "" {
		request.Header.Add("Authorization", "token:" + token)
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

func UploadFilterProfile(endpoint string, name string, content string, token string) bool {

	var json = []byte(content)
	request, err := http.NewRequest("POST", endpoint + "/api/profiles", bytes.NewBuffer(json))

	request.Header.Set("Content-Type", "application/json")

	if token != "" {
		request.Header.Add("Authorization", "token:" + token)
	}

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

func DeleteFilterProfile(endpoint string, name string, content string, token string) bool {

	request, err := http.NewRequest("DELETE", endpoint + "/api/profiles", nil)

	client := &http.Client{}

	if token != "" {
		request.Header.Add("Authorization", "token:" + token)
	}

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
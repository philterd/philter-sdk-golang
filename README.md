# Philter SDK for Golang

The **Philter SDK for Golang** enables Go developers to easily work with Philter. [Philter](https://www.philterd.ai/philter/) identifies and manipulates sensitive information like Protected Health Information (PHI) and personally identifiable information (PII) from natural language text.

Refer to [Philter API](https://docs.philterd.ai/philter/api/) documentation for details on the methods available.

## Installation

`go get github.com/philterd/philter-sdk-golang`

## Example Usage

With an available running instance of Philter, to filter text:

To filter text:

```
filterResponse := Filter("http://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default", "token")
```

To filter with explanation:

```
explainResponse := Explain("http://localhost:8080", "His SSN was 123-45-6789.", "context", "docid", "default", "token")
```

## Release History

* 1.1.0 - Adds authentication support.
* 1.0.0 - Initial release.

## License

This project is licensed under the Apache License, version 2.0.

Copyright 2020-2023 Mountain Fog, Inc.
Philter is a registered trademark of Mountain Fog, Inc.

package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type GraphQLClient struct {
	Headers map[string]string
	Url     string
}

func NewGraphQLClient(headers map[string]string, url string) *GraphQLClient {
	return &GraphQLClient{
		Headers: headers,
		Url:     url,
	}
}

func (c *GraphQLClient) Execute(args *map[string]any, graphQLDocument string) (resp *any, err error) {
	client := &http.Client{}

	reqBody := GraphQLRequest{
		Query: graphQLDocument,
	}
	if args != nil {
		reqBody.Variables = *args
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(reqBytes))

	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var response any
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	return &response, err
}

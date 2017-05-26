package escher

import "net/url"

type RequestHeaders [][2]string

type Request struct {
	Method  string         `json:"method"`
	Url     string         `json:"url"`
	Headers RequestHeaders `json:"headers"`
	Body    string         `json:"body"`
}

func (r Request) Path() (string, error) {
	url, err := url.Parse(r.Url)

	if err != nil {
		return "", err
	}

	return url.Path, err
}

type QueryParts [][2]string

func (r Request) QueryParts() (QueryParts, error) {
	url, err := url.Parse(r.Url)

	if err != nil {
		return nil, err
	}

	queryParts := make(QueryParts, 0)
	for key, values := range url.Query() {
		for _, value := range values {
			queryParts = append(queryParts, [2]string{key, value})
		}
	}

	return queryParts, nil
}
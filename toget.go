package toget

import (
	"errors"
	"net/http"
	"time"
)

type response struct {
	resp *http.Response
	err  error
}

var TimeoutError = errors.New("http get timeout")

func Get(url string, to time.Duration) (*http.Response, error) {

	var (
		tr     = &http.Transport{}
		client = &http.Client{Transport: tr}
		ret    = make(chan response, 1)
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		resp, err := client.Do(req)
		ret <- response{resp, err}
	}()

	select {
	case r := <-ret:
		return r.resp, r.err
	case <-time.After(to):
		tr.CancelRequest(req)
		return nil, TimeoutError
	}

}

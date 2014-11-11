package toget

import (
	"errors"
	"io"
	"net/http"
	"time"
)

var TimeoutError = errors.New("http get timeout")

func Head(url string, timeout time.Duration) (*http.Response, error) {

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	return Do(req, timeout)

}

func Get(url string, timeout time.Duration) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return Do(req, timeout)

}

func Post(url string, body io.Reader, timeout time.Duration) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	return Do(req, timeout)

}

func Put(url string, body io.Reader, timeout time.Duration) (*http.Response, error) {

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	return Do(req, timeout)

}

func Delete(url string, timeout time.Duration) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return Do(req, timeout)

}

func Do(req *http.Request, timeout time.Duration) (*http.Response, error) {

	type response struct {
		resp *http.Response
		err  error
	}

	var (
		tr     = &http.Transport{}
		client = &http.Client{Transport: tr}
		ret    = make(chan response, 1)
	)

	go func() {
		resp, err := client.Do(req)
		ret <- response{resp, err}
	}()

	select {
	case r := <-ret:
		return r.resp, r.err
	case <-time.After(timeout):
		tr.CancelRequest(req)
		return nil, TimeoutError
	}

}

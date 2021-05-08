package primer

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Dto interface {
	SetRequestId(id string)
}

func StrSafeDeref(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func MarshalRequest(request interface{}) (io.Reader, error) {
	if request == nil {
		return nil, nil
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func DecodeResponse(resp *http.Response, apiResp Dto) error {
	requestId := resp.Header.Get("X-Grabkit-Grab-Requestid")
	apiResp.SetRequestId(requestId)

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(apiResp); err != nil {
			return wrapError(err)
		}
		return nil
	case http.StatusNoContent:
		return nil
	default:
		var msg string
		if resp.ContentLength != 0 {
			bb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return wrapError(err)
			}
			msg = string(bb)
		}
		return &Error{
			Status:    resp.StatusCode,
			Message:   msg,
			RequestId: requestId,
		}
	}
}

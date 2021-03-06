package vars

import (
	"strings"
	"testing"

	"github.com/jmartin82/mmock/definition"
)

func TestGetHeaderParam(t *testing.T) {
	rp := Request{}
	header := make(definition.Values)
	header["Authorization"] = []string{"Bearer abc123"}
	req := definition.Request{}
	req.Headers = header
	rp.Request = &req
	v, f := rp.getHeaderParam("Authorization")
	if !f {
		t.Errorf("Header key not found")
	}

	if !strings.EqualFold(v, "Bearer abc123") {
		t.Errorf("Couldn't get the content. Value: %s", v)
	}
}

func TestGetHeaderParamNotFoundHeaderKey(t *testing.T) {
	rp := Request{}
	header := make(definition.Values)
	header["Authorization"] = []string{"Bearer abc123"}
	req := definition.Request{}
	req.Headers = header
	rp.Request = &req
	_, f := rp.getHeaderParam("Authorization2")
	if f {
		t.Errorf("Header key found")
	}
}

func TestGetHeaderParamWithOutHeaderValue(t *testing.T) {
	rp := Request{}
	header := make(definition.Values)
	header["Authorization"] = []string{}
	req := definition.Request{}
	req.Headers = header
	rp.Request = &req
	v, f := rp.getHeaderParam("Authorization")
	if f {
		t.Errorf("Header key found")
	}

	if strings.EqualFold(v, "Bearer abc1235") {
		t.Errorf("Couldn get the content. Value: %s", v)
	}
}

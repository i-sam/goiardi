package util

import (
	"testing"
	"net/http"
)

type testObj struct {
	name string
	urlType string
}

func (to *testObj) GetName() string {
	return to.name
}

func (to *testObj) URLType() string {
	return to.urlType
}

// The strange URLs are because the config doesn't get parsed here, so it ends
// up using the really-really default settings.

func TestObjURL(t *testing.T){
	obj := &testObj{ name: "foo", urlType: "bar" }
	url := ObjURL(obj)
	expectedUrl := "http://:0/bar/foo"
	if url != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, url)
	}
}

func TestCustomObjUrl(t *testing.T){
	obj := &testObj{ name: "foo", urlType: "bar" }
	url := CustomObjURL(obj, "/baz")
	expectedUrl := "http://:0/bar/foo/baz"
	if url != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, url)
	}
}

func TestCustomURL(t *testing.T){
	initUrl := "/foo/bar"
	url := CustomURL(initUrl)
	expectedUrl := "http://:0/foo/bar"
	if url != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, url)
	}
	initUrl = "foo/bar"
	url = CustomURL(initUrl)
	if url != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, url)
	}
}

func TestGerror(t *testing.T){
	errmsg := "foo bar"
	err := Errorf(errmsg)
	if err.Error() != errmsg {
		t.Errorf("expected %s to match %s", err.Error(), errmsg)
	}
	if err.Status() != http.StatusBadRequest {
		t.Errorf("err.Status() did not return expected default")
	}
	err.SetStatus(http.StatusNotFound)
	if err.Status() != http.StatusNotFound {
		t.Errorf("SetStatus did not set Status correctly")
	}
}

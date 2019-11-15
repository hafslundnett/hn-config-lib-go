package hid

import (
	"hafslundnett/x/hn-config-lib/hnhttp"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

const cacheSize int = 100 // TODO: config?
var intspecCache = []*intspec{}

type intspec struct {
	Active bool `json:"active"`
	Token
}

// Introspection exp
func (hid HIDclient) Introspection(t Token) (bool, error) {
	var i int
	var is *intspec

	for i, is = range intspecCache {
		if is.Access == t.Access {
			return is.introspect(), nil
		}
	}

	is, err := remoteIntspec(hid, t)
	if err != nil {
		return false, err
	}

	if i <= cacheSize {
		intspecCache = append(intspecCache, is)
	} else {
		intspecCache[cacheSize-1] = is
	}

	sort.Slice(intspecCache, func(i, j int) bool {
		return intspecCache[i].Expiration < intspecCache[j].Expiration
	})

	return is.introspect(), nil
}

func remoteIntspec(hid HIDclient, t Token) (*intspec, error) {
	client, err := hnhttp.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "while setting up http client")
	}

	target := hid.Host + "/" + hid.Path

	form := url.Values{}
	form.Add("token", t.Access)
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, target, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+hid.Secret)

	var is intspec
	err = client.Do(req, is)
	if err != nil {
		return nil, errors.Wrapf(err, "while introspecting client token")
	}

	return nil, nil
}

func (is *intspec) introspect() bool {
	//if is.Expiration < time.Now() { //TODO?§
	//	is.Active = false
	//}

	return is.Active
}

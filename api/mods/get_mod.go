package mods

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewGetModAPI(t http.RoundTripper) GetMod {
	return func(modID schema.ModID, o ...func(*GetModRequest)) (*http.Response, error) {
		r := new(GetModRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return r.Do(r.ctx, t)
	}
}

type GetMod func(modID schema.ModID, o ...func(*GetModRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod
type GetModRequest struct {
	ctx context.Context

	ModID schema.ModID
}

func (r *GetModRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/mods/%d", schema.BaseUrl, r.ModID)
	)

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return t.RoundTrip(req)
}

func (GetMod) WithContext(ctx context.Context) func(*GetModRequest) {
	return func(o *GetModRequest) {
		o.ctx = ctx
	}
}
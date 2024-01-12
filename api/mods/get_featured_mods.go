package mods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GetFeaturedModsOption func(*GetFeaturedModsRequest)

func NewGetFeaturedModsAPI(t http.RoundTripper) GetFeaturedMods {
	return func(gameID enum.GameID, o ...GetFeaturedModsOption) (*schema.GetFeaturedModsRequestBody, error) {
		r := new(GetFeaturedModsRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetFeaturedModsRequestBody](r.Do(r.ctx, t))
	}
}

type GetFeaturedMods func(gameID enum.GameID, o ...GetFeaturedModsOption) (*schema.GetFeaturedModsRequestBody, error)

// https://docs.curseforge.com/#get-featured-mods
type GetFeaturedModsRequest struct {
	ctx context.Context

	schema.GetFeaturedModsRequestBody
}

func (r *GetFeaturedModsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/mods/featured", schema.BaseUrl)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetFeaturedModsRequestBody); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", "application/json")

	return t.RoundTrip(req)
}

func (GetFeaturedMods) WithContext(ctx context.Context) GetFeaturedModsOption {
	return func(o *GetFeaturedModsRequest) {
		o.ctx = ctx
	}
}

func (GetFeaturedMods) WithExcludedModIDs(modID ...schema.ModID) GetFeaturedModsOption {
	return func(o *GetFeaturedModsRequest) {
		o.ExcludedModIDs = modID
	}
}

func (GetFeaturedMods) WithGameVersionTypeID(gameVersionTypeID enum.GameVersionType) GetFeaturedModsOption {
	return func(o *GetFeaturedModsRequest) {
		o.GameVersionTypeID = gameVersionTypeID
	}
}

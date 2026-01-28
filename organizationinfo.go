// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// OrganizationInfoService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationInfoService] method instead.
type OrganizationInfoService struct {
	Options []option.RequestOption
}

// NewOrganizationInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInfoService(opts ...option.RequestOption) (r *OrganizationInfoService) {
	r = &OrganizationInfoService{}
	r.Options = opts
	return
}

// Get the org specific information
func (r *OrganizationInfoService) Get(ctx context.Context, query OrganizationInfoGetParams, opts ...option.RequestOption) (res *OrganizationTableWithMembers, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// DEPRECATED: Use GET /organization/info instead
func (r *OrganizationInfoService) Deprecated(ctx context.Context, body OrganizationInfoDeprecatedParams, opts ...option.RequestOption) (res *OrganizationInfoDeprecatedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationInfoDeprecatedResponse = interface{}

type OrganizationInfoGetParams struct {
	OrganizationID param.Field[string] `query:"organization_id,required"`
}

// URLQuery serializes [OrganizationInfoGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationInfoDeprecatedParams struct {
	Organizations param.Field[[]string] `json:"organizations,required"`
}

func (r OrganizationInfoDeprecatedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

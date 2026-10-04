# ProviderGithubUserConnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeUrl** | Pointer to **string** | AuthorizeURL is GitHub&#39;s authorization page for the App. GitHub returns the browser to the callback, which parks the answer and sends it on with complete&#x3D;github&amp;grant&#x3D;&lt;id&gt; — to the return page the connect named, or the console&#39;s /connectors — and that page POSTs the grant to /v1/provider/github/user/complete. | [optional] 

## Methods

### NewProviderGithubUserConnectOut

`func NewProviderGithubUserConnectOut() *ProviderGithubUserConnectOut`

NewProviderGithubUserConnectOut instantiates a new ProviderGithubUserConnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubUserConnectOutWithDefaults

`func NewProviderGithubUserConnectOutWithDefaults() *ProviderGithubUserConnectOut`

NewProviderGithubUserConnectOutWithDefaults instantiates a new ProviderGithubUserConnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeUrl

`func (o *ProviderGithubUserConnectOut) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *ProviderGithubUserConnectOut) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *ProviderGithubUserConnectOut) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *ProviderGithubUserConnectOut) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



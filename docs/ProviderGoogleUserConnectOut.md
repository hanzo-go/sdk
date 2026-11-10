# ProviderGoogleUserConnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeUrl** | Pointer to **string** | AuthorizeURL is Google&#39;s consent page. Google returns the browser to the callback, which parks the answer and sends it on with complete&#x3D;google&amp;grant&#x3D;&lt;id&gt; — to the return page the connect named, or the console&#39;s /connectors — and that page POSTs the grant to /v1/provider/google/user/complete. | [optional] 

## Methods

### NewProviderGoogleUserConnectOut

`func NewProviderGoogleUserConnectOut() *ProviderGoogleUserConnectOut`

NewProviderGoogleUserConnectOut instantiates a new ProviderGoogleUserConnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGoogleUserConnectOutWithDefaults

`func NewProviderGoogleUserConnectOutWithDefaults() *ProviderGoogleUserConnectOut`

NewProviderGoogleUserConnectOutWithDefaults instantiates a new ProviderGoogleUserConnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeUrl

`func (o *ProviderGoogleUserConnectOut) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *ProviderGoogleUserConnectOut) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *ProviderGoogleUserConnectOut) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *ProviderGoogleUserConnectOut) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



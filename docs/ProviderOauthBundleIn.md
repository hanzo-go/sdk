# ProviderOauthBundleIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | Access is the access token. | [optional] 
**Account** | Pointer to **string** | Account is the account label the flow reported; sanitized on ingest. | [optional] 
**Refresh** | Pointer to **string** | Refresh is the refresh token. It is sealed and NEVER handed back out. | [optional] 

## Methods

### NewProviderOauthBundleIn

`func NewProviderOauthBundleIn() *ProviderOauthBundleIn`

NewProviderOauthBundleIn instantiates a new ProviderOauthBundleIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderOauthBundleInWithDefaults

`func NewProviderOauthBundleInWithDefaults() *ProviderOauthBundleIn`

NewProviderOauthBundleInWithDefaults instantiates a new ProviderOauthBundleIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ProviderOauthBundleIn) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ProviderOauthBundleIn) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ProviderOauthBundleIn) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ProviderOauthBundleIn) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccount

`func (o *ProviderOauthBundleIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProviderOauthBundleIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProviderOauthBundleIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProviderOauthBundleIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetRefresh

`func (o *ProviderOauthBundleIn) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *ProviderOauthBundleIn) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *ProviderOauthBundleIn) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *ProviderOauthBundleIn) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



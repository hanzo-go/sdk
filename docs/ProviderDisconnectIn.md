# ProviderDisconnectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account names ONE account to disconnect — the label the connection carries (a GitHub login, for the App) or its platform id (a Page, a channel, a board). Empty disconnects every account the org holds for the provider. An account the org does not hold disconnects nothing. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | [optional] 

## Methods

### NewProviderDisconnectIn

`func NewProviderDisconnectIn() *ProviderDisconnectIn`

NewProviderDisconnectIn instantiates a new ProviderDisconnectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderDisconnectInWithDefaults

`func NewProviderDisconnectInWithDefaults() *ProviderDisconnectIn`

NewProviderDisconnectInWithDefaults instantiates a new ProviderDisconnectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProviderDisconnectIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProviderDisconnectIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProviderDisconnectIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProviderDisconnectIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderDisconnectIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderDisconnectIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderDisconnectIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderDisconnectIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



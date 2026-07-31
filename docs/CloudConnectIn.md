# CloudConnectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID is the provider account the credential should be scoped to, for the providers whose Verify needs one (Cloudflare). Ignored by the OAuth path. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | [optional] 
**Token** | Pointer to **string** | Token is the customer&#39;s provider credential. Its PRESENCE — not its value — is what selects the apikey seal over the OAuth flow for a provider that offers both: {\&quot;token\&quot;:\&quot;…\&quot;}, even empty, is an apikey attempt (→ verify, which answers the \&quot;token required\&quot; 400 on an empty value), while a body with no token key (the console Connect button, &#x60;hanzo connector add&#x60; with no --token) starts OAuth. Read on STDIN by the CLI, never argv; never logged or echoed. | [optional] 

## Methods

### NewCloudConnectIn

`func NewCloudConnectIn() *CloudConnectIn`

NewCloudConnectIn instantiates a new CloudConnectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectInWithDefaults

`func NewCloudConnectInWithDefaults() *CloudConnectIn`

NewCloudConnectInWithDefaults instantiates a new CloudConnectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CloudConnectIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CloudConnectIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CloudConnectIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CloudConnectIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProvider

`func (o *CloudConnectIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudConnectIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudConnectIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudConnectIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetToken

`func (o *CloudConnectIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudConnectIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudConnectIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudConnectIn) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



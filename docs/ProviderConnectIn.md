# ProviderConnectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID is the provider account the credential should be scoped to, for the providers whose Verify needs one (Cloudflare). Ignored by the OAuth path. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields are the named inputs of a credential connector whose form is more than one key (the connector&#39;s published fields). Their presence selects the credential path exactly as Token&#39;s does. Never logged or echoed. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s registry id, from the :provider path segment. | [optional] 
**Return** | Pointer to **string** | Return is the page the OAuth callback sends the person back to, with connected&#x3D;&lt;provider&gt;&amp;account&#x3D;&lt;label&gt; or error&#x3D;&lt;provider&gt;&amp;reason&#x3D;&lt;why&gt; added to its query. It must name one of this deployment&#39;s return pages (on api.hanzo.ai: https://hanzo.ai/?at&#x3D;-/settings/integrations, https://console.hanzo.ai/connectors, https://platform.hanzo.ai/connectors); anything else, or nothing, returns to the console&#39;s /connectors. Ignored by the credential path. | [optional] 
**Token** | Pointer to **string** | Token is the customer&#39;s provider credential. Its PRESENCE — not its value — is what selects the apikey seal over the OAuth flow for a provider that offers both: {\&quot;token\&quot;:\&quot;…\&quot;}, even empty, is an apikey attempt (→ verify, which answers the \&quot;token required\&quot; 400 on an empty value), while a body with no token key (the console Connect button, &#x60;hanzo connector add&#x60; with no --token) starts OAuth. Read on STDIN by the CLI, never argv; never logged or echoed. | [optional] 

## Methods

### NewProviderConnectIn

`func NewProviderConnectIn() *ProviderConnectIn`

NewProviderConnectIn instantiates a new ProviderConnectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderConnectInWithDefaults

`func NewProviderConnectInWithDefaults() *ProviderConnectIn`

NewProviderConnectInWithDefaults instantiates a new ProviderConnectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ProviderConnectIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ProviderConnectIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ProviderConnectIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ProviderConnectIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFields

`func (o *ProviderConnectIn) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProviderConnectIn) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProviderConnectIn) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProviderConnectIn) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderConnectIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderConnectIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderConnectIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderConnectIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReturn

`func (o *ProviderConnectIn) GetReturn() string`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *ProviderConnectIn) GetReturnOk() (*string, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *ProviderConnectIn) SetReturn(v string)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *ProviderConnectIn) HasReturn() bool`

HasReturn returns a boolean if a field has been set.

### GetToken

`func (o *ProviderConnectIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProviderConnectIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProviderConnectIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProviderConnectIn) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



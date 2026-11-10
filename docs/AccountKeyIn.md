# AccountKeyIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**AccountKeyBudget**](AccountKeyBudget.md) | Budget caps what the key may spend on models, in cents, per calendar month and in all; a request past either is refused 402. Omitted means no cap. | [optional] 
**Expires** | Pointer to **string** | Expires is when the key stops working, RFC 3339. Omitted means never. | [optional] 
**Limit** | Pointer to **[]string** | Limit narrows what the key may reach, as &#x60;kind:name&#x60; entries: &#x60;model:zen5&#x60;, &#x60;project:acme&#x60;, &#x60;product:commerce&#x60;, or &#x60;model:*&#x60; for a whole kind. It only ever NARROWS — a key never reaches further than the person who made it. Omitted makes an unrestricted key.  Example: {\&quot;name\&quot;: \&quot;ci\&quot;, \&quot;limit\&quot;: [\&quot;model:zen5\&quot;, \&quot;project:acme\&quot;]} | [optional] 
**Name** | Pointer to **string** | Name labels the key. Omitted names it \&quot;key N\&quot;. | [optional] 
**Rate** | Pointer to **int64** | Rate caps the key&#39;s requests a minute; past it a request is refused 429. Omitted means no limit of its own. | [optional] 
**Type** | Pointer to **string** | Type is the key class: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret. | [optional] 

## Methods

### NewAccountKeyIn

`func NewAccountKeyIn() *AccountKeyIn`

NewAccountKeyIn instantiates a new AccountKeyIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountKeyInWithDefaults

`func NewAccountKeyInWithDefaults() *AccountKeyIn`

NewAccountKeyInWithDefaults instantiates a new AccountKeyIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *AccountKeyIn) GetBudget() AccountKeyBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AccountKeyIn) GetBudgetOk() (*AccountKeyBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AccountKeyIn) SetBudget(v AccountKeyBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AccountKeyIn) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetExpires

`func (o *AccountKeyIn) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AccountKeyIn) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AccountKeyIn) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AccountKeyIn) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetLimit

`func (o *AccountKeyIn) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountKeyIn) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountKeyIn) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccountKeyIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *AccountKeyIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountKeyIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountKeyIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountKeyIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRate

`func (o *AccountKeyIn) GetRate() int64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AccountKeyIn) GetRateOk() (*int64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AccountKeyIn) SetRate(v int64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AccountKeyIn) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetType

`func (o *AccountKeyIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountKeyIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountKeyIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountKeyIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



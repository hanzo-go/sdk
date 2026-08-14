# KeyTypeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **[]string** | Limit narrows what the minted key may reach, as &#x60;kind:name&#x60; entries: &#x60;model:zen5&#x60;, &#x60;project:acme&#x60;, &#x60;product:commerce&#x60;, or &#x60;model:*&#x60; for a whole kind. It only ever NARROWS — a key can never reach further than the person who minted it — so an unrecognised kind costs availability, never privilege.  Omitted mints an unrestricted key, because that is what every key in the estate is today and a default that restricted would revoke all of them.  Example: {\&quot;type\&quot;: \&quot;secret\&quot;, \&quot;limit\&quot;: [\&quot;model:zen5\&quot;, \&quot;project:acme\&quot;]} | [optional] 
**Type** | Pointer to **string** | Type is the key class to act on: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. | [optional] 

## Methods

### NewKeyTypeIn

`func NewKeyTypeIn() *KeyTypeIn`

NewKeyTypeIn instantiates a new KeyTypeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyTypeInWithDefaults

`func NewKeyTypeInWithDefaults() *KeyTypeIn`

NewKeyTypeInWithDefaults instantiates a new KeyTypeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *KeyTypeIn) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KeyTypeIn) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KeyTypeIn) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KeyTypeIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetType

`func (o *KeyTypeIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyTypeIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyTypeIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyTypeIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



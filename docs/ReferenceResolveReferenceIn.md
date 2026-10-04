# ReferenceResolveReferenceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | Keys are the values to look up, at most 100 per call: email addresses or domains, IP addresses, card prefixes, user-agent strings, autonomous system numbers, device digests. | [optional] 
**Sets** | Pointer to **[]string** | Sets narrows which sets to consult. Empty consults every set whose matcher can read the keys given. | [optional] 

## Methods

### NewReferenceResolveReferenceIn

`func NewReferenceResolveReferenceIn() *ReferenceResolveReferenceIn`

NewReferenceResolveReferenceIn instantiates a new ReferenceResolveReferenceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceResolveReferenceInWithDefaults

`func NewReferenceResolveReferenceInWithDefaults() *ReferenceResolveReferenceIn`

NewReferenceResolveReferenceInWithDefaults instantiates a new ReferenceResolveReferenceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ReferenceResolveReferenceIn) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReferenceResolveReferenceIn) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReferenceResolveReferenceIn) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReferenceResolveReferenceIn) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSets

`func (o *ReferenceResolveReferenceIn) GetSets() []string`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *ReferenceResolveReferenceIn) GetSetsOk() (*[]string, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *ReferenceResolveReferenceIn) SetSets(v []string)`

SetSets sets Sets field to given value.

### HasSets

`func (o *ReferenceResolveReferenceIn) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



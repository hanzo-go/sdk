# ReferenceSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is when this publisher was current, RFC 3339. | [optional] 
**Basis** | Pointer to **string** | Basis is the KIND of permission this publisher&#39;s data reaches you under: licence (an explicit grant), registry (the registry of record publishing for anyone to consult), operator (an operator&#39;s own machine-readable statement about its own network, published for third parties to filter by — not a licence, and not claimed as one), own (computed here), or none (nothing reaches you: the membership is held by the component that screens against it). It is on the wire so the licence position is an audit you can run. | [optional] 
**Keys** | Pointer to **int32** | Keys is how many members this publisher contributed. | [optional] 
**Origin** | Pointer to **string** | Origin is exactly where it was taken from, so it can be taken again. | [optional] 
**Refusal** | Pointer to **string** | Refusal is why this publisher&#39;s last take failed, if it did. The set keeps its previous version of this source and ages out visibly rather than silently shrinking. | [optional] 
**Source** | Pointer to **string** | Source is the publisher. | [optional] 
**Terms** | Pointer to **string** | Terms is the CITATION that basis points at — the licence identifier, the registry, or the operator publication. A source with no stated terms is not in the catalog. | [optional] 
**Version** | Pointer to **string** | Version is the content digest of what this publisher last supplied. Two refreshes that agree on it took the same data. | [optional] 

## Methods

### NewReferenceSource

`func NewReferenceSource() *ReferenceSource`

NewReferenceSource instantiates a new ReferenceSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceSourceWithDefaults

`func NewReferenceSourceWithDefaults() *ReferenceSource`

NewReferenceSourceWithDefaults instantiates a new ReferenceSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *ReferenceSource) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ReferenceSource) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ReferenceSource) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ReferenceSource) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetBasis

`func (o *ReferenceSource) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *ReferenceSource) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *ReferenceSource) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *ReferenceSource) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetKeys

`func (o *ReferenceSource) GetKeys() int32`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReferenceSource) GetKeysOk() (*int32, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReferenceSource) SetKeys(v int32)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReferenceSource) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetOrigin

`func (o *ReferenceSource) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ReferenceSource) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ReferenceSource) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ReferenceSource) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceSource) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceSource) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceSource) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceSource) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetSource

`func (o *ReferenceSource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReferenceSource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReferenceSource) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReferenceSource) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTerms

`func (o *ReferenceSource) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *ReferenceSource) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *ReferenceSource) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *ReferenceSource) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceSource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceSource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceSource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceSource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



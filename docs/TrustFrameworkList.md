# TrustFrameworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frameworks** | Pointer to [**[]TrustFrameworkRow**](TrustFrameworkRow.md) | Frameworks is each framework and how many clauses it publishes. | [optional] 

## Methods

### NewTrustFrameworkList

`func NewTrustFrameworkList() *TrustFrameworkList`

NewTrustFrameworkList instantiates a new TrustFrameworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustFrameworkListWithDefaults

`func NewTrustFrameworkListWithDefaults() *TrustFrameworkList`

NewTrustFrameworkListWithDefaults instantiates a new TrustFrameworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrameworks

`func (o *TrustFrameworkList) GetFrameworks() []TrustFrameworkRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *TrustFrameworkList) GetFrameworksOk() (*[]TrustFrameworkRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *TrustFrameworkList) SetFrameworks(v []TrustFrameworkRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *TrustFrameworkList) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



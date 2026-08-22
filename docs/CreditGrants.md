# CreditGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Grants** | Pointer to [**[]CreditGrant**](CreditGrant.md) |  | [optional] 

## Methods

### NewCreditGrants

`func NewCreditGrants() *CreditGrants`

NewCreditGrants instantiates a new CreditGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditGrantsWithDefaults

`func NewCreditGrantsWithDefaults() *CreditGrants`

NewCreditGrantsWithDefaults instantiates a new CreditGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CreditGrants) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreditGrants) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreditGrants) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreditGrants) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGrants

`func (o *CreditGrants) GetGrants() []CreditGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *CreditGrants) GetGrantsOk() (*[]CreditGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *CreditGrants) SetGrants(v []CreditGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *CreditGrants) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



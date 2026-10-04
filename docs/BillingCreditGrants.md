# BillingCreditGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Grants** | Pointer to [**[]BillingCreditGrant**](BillingCreditGrant.md) |  | [optional] 

## Methods

### NewBillingCreditGrants

`func NewBillingCreditGrants() *BillingCreditGrants`

NewBillingCreditGrants instantiates a new BillingCreditGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCreditGrantsWithDefaults

`func NewBillingCreditGrantsWithDefaults() *BillingCreditGrants`

NewBillingCreditGrantsWithDefaults instantiates a new BillingCreditGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BillingCreditGrants) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingCreditGrants) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingCreditGrants) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingCreditGrants) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGrants

`func (o *BillingCreditGrants) GetGrants() []BillingCreditGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *BillingCreditGrants) GetGrantsOk() (*[]BillingCreditGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *BillingCreditGrants) SetGrants(v []BillingCreditGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *BillingCreditGrants) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



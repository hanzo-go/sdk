# PromoIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is the master switch: false parks the offer without deleting it. | [optional] 
**End** | Pointer to **string** | End is when the offer closes (RFC3339). | [optional] 
**PercentOff** | Pointer to **int32** | PercentOff is the discount, 0-100. | [optional] 
**Plans** | Pointer to **[]string** | Plans are the plan ids the offer applies to. | [optional] 
**Start** | Pointer to **string** | Start is when the offer opens (RFC3339). | [optional] 

## Methods

### NewPromoIn

`func NewPromoIn() *PromoIn`

NewPromoIn instantiates a new PromoIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoInWithDefaults

`func NewPromoInWithDefaults() *PromoIn`

NewPromoInWithDefaults instantiates a new PromoIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PromoIn) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PromoIn) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PromoIn) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PromoIn) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnd

`func (o *PromoIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PromoIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PromoIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PromoIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPercentOff

`func (o *PromoIn) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *PromoIn) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *PromoIn) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *PromoIn) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetPlans

`func (o *PromoIn) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PromoIn) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PromoIn) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PromoIn) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetStart

`func (o *PromoIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PromoIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PromoIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *PromoIn) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AdminPromoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentOff** | Pointer to **float64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**Plans** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdminPromoUpdate

`func NewAdminPromoUpdate() *AdminPromoUpdate`

NewAdminPromoUpdate instantiates a new AdminPromoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPromoUpdateWithDefaults

`func NewAdminPromoUpdateWithDefaults() *AdminPromoUpdate`

NewAdminPromoUpdateWithDefaults instantiates a new AdminPromoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentOff

`func (o *AdminPromoUpdate) GetPercentOff() float64`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *AdminPromoUpdate) GetPercentOffOk() (*float64, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *AdminPromoUpdate) SetPercentOff(v float64)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *AdminPromoUpdate) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### GetStart

`func (o *AdminPromoUpdate) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AdminPromoUpdate) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AdminPromoUpdate) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AdminPromoUpdate) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *AdminPromoUpdate) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AdminPromoUpdate) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AdminPromoUpdate) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AdminPromoUpdate) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPlans

`func (o *AdminPromoUpdate) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AdminPromoUpdate) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AdminPromoUpdate) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AdminPromoUpdate) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetActive

`func (o *AdminPromoUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdminPromoUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdminPromoUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AdminPromoUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



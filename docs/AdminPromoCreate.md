# AdminPromoCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentOff** | **float64** |  | 
**Start** | **time.Time** |  | 
**End** | **time.Time** |  | 
**Plans** | Pointer to **[]string** | Plan ids; empty &#x3D; all plans. | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewAdminPromoCreate

`func NewAdminPromoCreate(percentOff float64, start time.Time, end time.Time, ) *AdminPromoCreate`

NewAdminPromoCreate instantiates a new AdminPromoCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPromoCreateWithDefaults

`func NewAdminPromoCreateWithDefaults() *AdminPromoCreate`

NewAdminPromoCreateWithDefaults instantiates a new AdminPromoCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentOff

`func (o *AdminPromoCreate) GetPercentOff() float64`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *AdminPromoCreate) GetPercentOffOk() (*float64, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *AdminPromoCreate) SetPercentOff(v float64)`

SetPercentOff sets PercentOff field to given value.


### GetStart

`func (o *AdminPromoCreate) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AdminPromoCreate) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AdminPromoCreate) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *AdminPromoCreate) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AdminPromoCreate) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AdminPromoCreate) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetPlans

`func (o *AdminPromoCreate) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AdminPromoCreate) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AdminPromoCreate) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AdminPromoCreate) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetActive

`func (o *AdminPromoCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdminPromoCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdminPromoCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AdminPromoCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



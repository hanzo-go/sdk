# AdminPromo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Server-assigned id. | 
**PercentOff** | **float64** | Percent off list price (50 &#x3D; 50% off). | 
**Start** | **time.Time** | Window start, inclusive (UTC). | 
**End** | **time.Time** | Window end, exclusive (UTC). | 
**Plans** | Pointer to **[]string** | Plan ids the promo applies to; empty &#x3D; all plans. | [optional] 
**Active** | **bool** | Admin on/off toggle; the discount applies only when this is true and now is within [start, end). | 

## Methods

### NewAdminPromo

`func NewAdminPromo(id string, percentOff float64, start time.Time, end time.Time, active bool, ) *AdminPromo`

NewAdminPromo instantiates a new AdminPromo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPromoWithDefaults

`func NewAdminPromoWithDefaults() *AdminPromo`

NewAdminPromoWithDefaults instantiates a new AdminPromo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminPromo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminPromo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminPromo) SetId(v string)`

SetId sets Id field to given value.


### GetPercentOff

`func (o *AdminPromo) GetPercentOff() float64`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *AdminPromo) GetPercentOffOk() (*float64, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *AdminPromo) SetPercentOff(v float64)`

SetPercentOff sets PercentOff field to given value.


### GetStart

`func (o *AdminPromo) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AdminPromo) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AdminPromo) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *AdminPromo) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AdminPromo) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AdminPromo) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetPlans

`func (o *AdminPromo) GetPlans() []string`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AdminPromo) GetPlansOk() (*[]string, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AdminPromo) SetPlans(v []string)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AdminPromo) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetActive

`func (o *AdminPromo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdminPromo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdminPromo) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



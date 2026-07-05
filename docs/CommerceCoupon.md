# CommerceCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Used** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommerceCoupon

`func NewCommerceCoupon() *CommerceCoupon`

NewCommerceCoupon instantiates a new CommerceCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceCouponWithDefaults

`func NewCommerceCouponWithDefaults() *CommerceCoupon`

NewCommerceCouponWithDefaults instantiates a new CommerceCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceCoupon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceCoupon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceCoupon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceCoupon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CommerceCoupon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommerceCoupon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommerceCoupon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommerceCoupon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CommerceCoupon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommerceCoupon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommerceCoupon) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommerceCoupon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *CommerceCoupon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CommerceCoupon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CommerceCoupon) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CommerceCoupon) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *CommerceCoupon) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CommerceCoupon) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CommerceCoupon) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CommerceCoupon) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAmount

`func (o *CommerceCoupon) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommerceCoupon) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommerceCoupon) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommerceCoupon) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStartDate

`func (o *CommerceCoupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CommerceCoupon) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CommerceCoupon) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CommerceCoupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CommerceCoupon) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CommerceCoupon) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CommerceCoupon) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CommerceCoupon) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLimit

`func (o *CommerceCoupon) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CommerceCoupon) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CommerceCoupon) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CommerceCoupon) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetUsed

`func (o *CommerceCoupon) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *CommerceCoupon) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *CommerceCoupon) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *CommerceCoupon) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



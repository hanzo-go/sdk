# LevelSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L1Cents** | Pointer to **int32** | L1Cents is lifetime commission accrued to DIRECT referrers, in cents. | [optional] 
**L2Cents** | Pointer to **int32** | L2Cents is lifetime commission accrued one step above the direct referrer, in cents, at the platform-wide level-2 rate. | [optional] 
**L3Cents** | Pointer to **int32** | L3Cents is lifetime commission accrued two steps above, in cents. Nothing accrues past level 3, so l1+l2+l3 is the whole accrual. | [optional] 

## Methods

### NewLevelSplit

`func NewLevelSplit() *LevelSplit`

NewLevelSplit instantiates a new LevelSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLevelSplitWithDefaults

`func NewLevelSplitWithDefaults() *LevelSplit`

NewLevelSplitWithDefaults instantiates a new LevelSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL1Cents

`func (o *LevelSplit) GetL1Cents() int32`

GetL1Cents returns the L1Cents field if non-nil, zero value otherwise.

### GetL1CentsOk

`func (o *LevelSplit) GetL1CentsOk() (*int32, bool)`

GetL1CentsOk returns a tuple with the L1Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL1Cents

`func (o *LevelSplit) SetL1Cents(v int32)`

SetL1Cents sets L1Cents field to given value.

### HasL1Cents

`func (o *LevelSplit) HasL1Cents() bool`

HasL1Cents returns a boolean if a field has been set.

### GetL2Cents

`func (o *LevelSplit) GetL2Cents() int32`

GetL2Cents returns the L2Cents field if non-nil, zero value otherwise.

### GetL2CentsOk

`func (o *LevelSplit) GetL2CentsOk() (*int32, bool)`

GetL2CentsOk returns a tuple with the L2Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Cents

`func (o *LevelSplit) SetL2Cents(v int32)`

SetL2Cents sets L2Cents field to given value.

### HasL2Cents

`func (o *LevelSplit) HasL2Cents() bool`

HasL2Cents returns a boolean if a field has been set.

### GetL3Cents

`func (o *LevelSplit) GetL3Cents() int32`

GetL3Cents returns the L3Cents field if non-nil, zero value otherwise.

### GetL3CentsOk

`func (o *LevelSplit) GetL3CentsOk() (*int32, bool)`

GetL3CentsOk returns a tuple with the L3Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3Cents

`func (o *LevelSplit) SetL3Cents(v int32)`

SetL3Cents sets L3Cents field to given value.

### HasL3Cents

`func (o *LevelSplit) HasL3Cents() bool`

HasL3Cents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Decision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Defaults** | Pointer to [**Grants**](Grants.md) |  | [optional] 
**Gpc** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Notice** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewDecision

`func NewDecision() *Decision`

NewDecision instantiates a new Decision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionWithDefaults

`func NewDecisionWithDefaults() *Decision`

NewDecisionWithDefaults instantiates a new Decision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaults

`func (o *Decision) GetDefaults() Grants`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *Decision) GetDefaultsOk() (*Grants, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *Decision) SetDefaults(v Grants)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *Decision) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGpc

`func (o *Decision) GetGpc() bool`

GetGpc returns the Gpc field if non-nil, zero value otherwise.

### GetGpcOk

`func (o *Decision) GetGpcOk() (*bool, bool)`

GetGpcOk returns a tuple with the Gpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpc

`func (o *Decision) SetGpc(v bool)`

SetGpc sets Gpc field to given value.

### HasGpc

`func (o *Decision) HasGpc() bool`

HasGpc returns a boolean if a field has been set.

### GetMode

`func (o *Decision) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Decision) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Decision) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Decision) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNotice

`func (o *Decision) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *Decision) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *Decision) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *Decision) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### GetRegion

`func (o *Decision) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Decision) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Decision) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Decision) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVersion

`func (o *Decision) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Decision) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Decision) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Decision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# O11yO11yFunnelStepsUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description replaces the funnel&#39;s description. Empty leaves it as it was. | [optional] 
**FunnelId** | Pointer to **string** | FunnelID is the funnel to update. | [optional] 
**FunnelName** | Pointer to **string** | Name replaces the funnel&#39;s name. Empty leaves it as it was. | [optional] 
**Steps** | Pointer to [**[]O11yFunnelStep**](O11yFunnelStep.md) | Steps are the funnel&#39;s steps, in order. At least two are needed before any analytics read will answer. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp is when the change was made, as a millisecond epoch. | [optional] 

## Methods

### NewO11yO11yFunnelStepsUpdateIn

`func NewO11yO11yFunnelStepsUpdateIn() *O11yO11yFunnelStepsUpdateIn`

NewO11yO11yFunnelStepsUpdateIn instantiates a new O11yO11yFunnelStepsUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFunnelStepsUpdateInWithDefaults

`func NewO11yO11yFunnelStepsUpdateInWithDefaults() *O11yO11yFunnelStepsUpdateIn`

NewO11yO11yFunnelStepsUpdateInWithDefaults instantiates a new O11yO11yFunnelStepsUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yFunnelStepsUpdateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yFunnelStepsUpdateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yFunnelStepsUpdateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yFunnelStepsUpdateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunnelId

`func (o *O11yO11yFunnelStepsUpdateIn) GetFunnelId() string`

GetFunnelId returns the FunnelId field if non-nil, zero value otherwise.

### GetFunnelIdOk

`func (o *O11yO11yFunnelStepsUpdateIn) GetFunnelIdOk() (*string, bool)`

GetFunnelIdOk returns a tuple with the FunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelId

`func (o *O11yO11yFunnelStepsUpdateIn) SetFunnelId(v string)`

SetFunnelId sets FunnelId field to given value.

### HasFunnelId

`func (o *O11yO11yFunnelStepsUpdateIn) HasFunnelId() bool`

HasFunnelId returns a boolean if a field has been set.

### GetFunnelName

`func (o *O11yO11yFunnelStepsUpdateIn) GetFunnelName() string`

GetFunnelName returns the FunnelName field if non-nil, zero value otherwise.

### GetFunnelNameOk

`func (o *O11yO11yFunnelStepsUpdateIn) GetFunnelNameOk() (*string, bool)`

GetFunnelNameOk returns a tuple with the FunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelName

`func (o *O11yO11yFunnelStepsUpdateIn) SetFunnelName(v string)`

SetFunnelName sets FunnelName field to given value.

### HasFunnelName

`func (o *O11yO11yFunnelStepsUpdateIn) HasFunnelName() bool`

HasFunnelName returns a boolean if a field has been set.

### GetSteps

`func (o *O11yO11yFunnelStepsUpdateIn) GetSteps() []O11yFunnelStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *O11yO11yFunnelStepsUpdateIn) GetStepsOk() (*[]O11yFunnelStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *O11yO11yFunnelStepsUpdateIn) SetSteps(v []O11yFunnelStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *O11yO11yFunnelStepsUpdateIn) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yFunnelStepsUpdateIn) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yFunnelStepsUpdateIn) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yFunnelStepsUpdateIn) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yFunnelStepsUpdateIn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



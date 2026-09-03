# O11yO11yFunnelUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description replaces the funnel&#39;s description. Empty leaves it as it was. | [optional] 
**FunnelName** | Pointer to **string** | Name replaces the funnel&#39;s name. Empty leaves it as it was. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp is when the change was made, as a millisecond epoch. | [optional] 

## Methods

### NewO11yO11yFunnelUpdateIn

`func NewO11yO11yFunnelUpdateIn() *O11yO11yFunnelUpdateIn`

NewO11yO11yFunnelUpdateIn instantiates a new O11yO11yFunnelUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFunnelUpdateInWithDefaults

`func NewO11yO11yFunnelUpdateInWithDefaults() *O11yO11yFunnelUpdateIn`

NewO11yO11yFunnelUpdateInWithDefaults instantiates a new O11yO11yFunnelUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yFunnelUpdateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yFunnelUpdateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yFunnelUpdateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yFunnelUpdateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunnelName

`func (o *O11yO11yFunnelUpdateIn) GetFunnelName() string`

GetFunnelName returns the FunnelName field if non-nil, zero value otherwise.

### GetFunnelNameOk

`func (o *O11yO11yFunnelUpdateIn) GetFunnelNameOk() (*string, bool)`

GetFunnelNameOk returns a tuple with the FunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelName

`func (o *O11yO11yFunnelUpdateIn) SetFunnelName(v string)`

SetFunnelName sets FunnelName field to given value.

### HasFunnelName

`func (o *O11yO11yFunnelUpdateIn) HasFunnelName() bool`

HasFunnelName returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yFunnelUpdateIn) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yFunnelUpdateIn) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yFunnelUpdateIn) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yFunnelUpdateIn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



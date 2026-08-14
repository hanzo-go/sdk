# O11yLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**O11yLimitConfig**](O11yLimitConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**O11yLimitMetric**](O11yLimitMetric.md) |  | [optional] 
**Signal** | Pointer to **string** | \&quot;logs\&quot;, \&quot;traces\&quot;, \&quot;metrics\&quot; | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yLimit

`func NewO11yLimit() *O11yLimit`

NewO11yLimit instantiates a new O11yLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLimitWithDefaults

`func NewO11yLimitWithDefaults() *O11yLimit`

NewO11yLimitWithDefaults instantiates a new O11yLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *O11yLimit) GetConfig() O11yLimitConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yLimit) GetConfigOk() (*O11yLimitConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yLimit) SetConfig(v O11yLimitConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yLimit) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yLimit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yLimit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyId

`func (o *O11yLimit) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *O11yLimit) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *O11yLimit) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *O11yLimit) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetMetric

`func (o *O11yLimit) GetMetric() O11yLimitMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *O11yLimit) GetMetricOk() (*O11yLimitMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *O11yLimit) SetMetric(v O11yLimitMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *O11yLimit) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetSignal

`func (o *O11yLimit) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *O11yLimit) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *O11yLimit) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *O11yLimit) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTags

`func (o *O11yLimit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yLimit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yLimit) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yLimit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yLimit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yLimit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yLimit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yLimit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



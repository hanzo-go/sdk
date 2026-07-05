# NexusUsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**TokenCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewNexusUsageInfo

`func NewNexusUsageInfo() *NexusUsageInfo`

NewNexusUsageInfo instantiates a new NexusUsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusUsageInfoWithDefaults

`func NewNexusUsageInfoWithDefaults() *NexusUsageInfo`

NewNexusUsageInfoWithDefaults instantiates a new NexusUsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *NexusUsageInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NexusUsageInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NexusUsageInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NexusUsageInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStartTime

`func (o *NexusUsageInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NexusUsageInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NexusUsageInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NexusUsageInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTokenCount

`func (o *NexusUsageInfo) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *NexusUsageInfo) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *NexusUsageInfo) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *NexusUsageInfo) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



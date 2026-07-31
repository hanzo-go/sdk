# CloudPublishResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** |  | [optional] 
**ExternalIds** | Pointer to **map[string]string** |  | [optional] 
**Results** | Pointer to [**[]CloudChannelResult**](CloudChannelResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPublishResult

`func NewCloudPublishResult() *CloudPublishResult`

NewCloudPublishResult instantiates a new CloudPublishResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPublishResultWithDefaults

`func NewCloudPublishResultWithDefaults() *CloudPublishResult`

NewCloudPublishResultWithDefaults instantiates a new CloudPublishResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *CloudPublishResult) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CloudPublishResult) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CloudPublishResult) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CloudPublishResult) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetExternalIds

`func (o *CloudPublishResult) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *CloudPublishResult) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *CloudPublishResult) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *CloudPublishResult) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetResults

`func (o *CloudPublishResult) GetResults() []CloudChannelResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CloudPublishResult) GetResultsOk() (*[]CloudChannelResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CloudPublishResult) SetResults(v []CloudChannelResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *CloudPublishResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPublishResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPublishResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPublishResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPublishResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



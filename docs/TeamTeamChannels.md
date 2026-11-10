# TeamTeamChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]TeamTeamChannel**](TeamTeamChannel.md) | Channels is every channel of every space the caller&#39;s org owns, each with the work facet it carries. | [optional] 

## Methods

### NewTeamTeamChannels

`func NewTeamTeamChannels() *TeamTeamChannels`

NewTeamTeamChannels instantiates a new TeamTeamChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelsWithDefaults

`func NewTeamTeamChannelsWithDefaults() *TeamTeamChannels`

NewTeamTeamChannelsWithDefaults instantiates a new TeamTeamChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *TeamTeamChannels) GetChannels() []TeamTeamChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TeamTeamChannels) GetChannelsOk() (*[]TeamTeamChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TeamTeamChannels) SetChannels(v []TeamTeamChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TeamTeamChannels) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



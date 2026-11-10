# TeamPublicChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]TeamListed**](TeamListed.md) | Channels is every published channel the query matched, newest-written first. | [optional] 

## Methods

### NewTeamPublicChannels

`func NewTeamPublicChannels() *TeamPublicChannels`

NewTeamPublicChannels instantiates a new TeamPublicChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPublicChannelsWithDefaults

`func NewTeamPublicChannelsWithDefaults() *TeamPublicChannels`

NewTeamPublicChannelsWithDefaults instantiates a new TeamPublicChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *TeamPublicChannels) GetChannels() []TeamListed`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TeamPublicChannels) GetChannelsOk() (*[]TeamListed, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TeamPublicChannels) SetChannels(v []TeamListed)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TeamPublicChannels) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



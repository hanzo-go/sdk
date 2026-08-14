# ChannelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Channel**](Channel.md) | Data is every social channel the caller&#39;s org has connected, disabled ones included (Disabled says which). | [optional] 

## Methods

### NewChannelList

`func NewChannelList() *ChannelList`

NewChannelList instantiates a new ChannelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelListWithDefaults

`func NewChannelListWithDefaults() *ChannelList`

NewChannelListWithDefaults instantiates a new ChannelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChannelList) GetData() []Channel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChannelList) GetDataOk() (*[]Channel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChannelList) SetData(v []Channel)`

SetData sets Data field to given value.

### HasData

`func (o *ChannelList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



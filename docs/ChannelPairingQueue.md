# ChannelPairingQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]ChannelPairingView**](ChannelPairingView.md) | Pending is every unexpired pairing request waiting on an org admin, each carrying the channel, the requesting sender and the code to approve it with. | [optional] 

## Methods

### NewChannelPairingQueue

`func NewChannelPairingQueue() *ChannelPairingQueue`

NewChannelPairingQueue instantiates a new ChannelPairingQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPairingQueueWithDefaults

`func NewChannelPairingQueueWithDefaults() *ChannelPairingQueue`

NewChannelPairingQueueWithDefaults instantiates a new ChannelPairingQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *ChannelPairingQueue) GetPending() []ChannelPairingView`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ChannelPairingQueue) GetPendingOk() (*[]ChannelPairingView, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ChannelPairingQueue) SetPending(v []ChannelPairingView)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ChannelPairingQueue) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



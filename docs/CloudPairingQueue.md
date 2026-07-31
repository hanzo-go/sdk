# CloudPairingQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]CloudPairingView**](CloudPairingView.md) | Pending is every unexpired pairing request waiting on an org admin, each carrying the channel, the requesting sender and the code to approve it with. | [optional] 

## Methods

### NewCloudPairingQueue

`func NewCloudPairingQueue() *CloudPairingQueue`

NewCloudPairingQueue instantiates a new CloudPairingQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPairingQueueWithDefaults

`func NewCloudPairingQueueWithDefaults() *CloudPairingQueue`

NewCloudPairingQueueWithDefaults instantiates a new CloudPairingQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *CloudPairingQueue) GetPending() []CloudPairingView`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CloudPairingQueue) GetPendingOk() (*[]CloudPairingView, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CloudPairingQueue) SetPending(v []CloudPairingView)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CloudPairingQueue) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NotifyNotifySend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Generated 16-byte hex message id (opaque handle). | 
**TaskId** | Pointer to **string** | hanzoai/tasks workflow id in async mode; always empty in the sync fold.  | [optional] 
**Status** | **string** | Terminal delivery status. In the sync fold this is &#x60;sent&#x60; on success or &#x60;failed&#x60; on terminal failure. (The type also defines &#x60;queued&#x60;, &#x60;sending&#x60;, &#x60;delivered&#x60; for the async plane, which is not folded.)  | 
**Error** | Pointer to **string** | Set on terminal failure (sync mode only). | [optional] 
**Items** | [**[]NotifySendResponse**](NotifySendResponse.md) |  | 

## Methods

### NewNotifyNotifySend200Response

`func NewNotifyNotifySend200Response(messageId string, status string, items []NotifySendResponse, ) *NotifyNotifySend200Response`

NewNotifyNotifySend200Response instantiates a new NotifyNotifySend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNotifySend200ResponseWithDefaults

`func NewNotifyNotifySend200ResponseWithDefaults() *NotifyNotifySend200Response`

NewNotifyNotifySend200ResponseWithDefaults instantiates a new NotifyNotifySend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *NotifyNotifySend200Response) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *NotifyNotifySend200Response) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *NotifyNotifySend200Response) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetTaskId

`func (o *NotifyNotifySend200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *NotifyNotifySend200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *NotifyNotifySend200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *NotifyNotifySend200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetStatus

`func (o *NotifyNotifySend200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotifyNotifySend200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotifyNotifySend200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *NotifyNotifySend200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NotifyNotifySend200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NotifyNotifySend200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NotifyNotifySend200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetItems

`func (o *NotifyNotifySend200Response) GetItems() []NotifySendResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NotifyNotifySend200Response) GetItemsOk() (*[]NotifySendResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NotifyNotifySend200Response) SetItems(v []NotifySendResponse)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



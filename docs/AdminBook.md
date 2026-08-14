# AdminBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AdminBookData**](AdminBookData.md) | Data is the book. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; — the operator console&#39;s envelope discriminator. | [optional] 

## Methods

### NewAdminBook

`func NewAdminBook() *AdminBook`

NewAdminBook instantiates a new AdminBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminBookWithDefaults

`func NewAdminBookWithDefaults() *AdminBook`

NewAdminBookWithDefaults instantiates a new AdminBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AdminBook) GetData() AdminBookData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminBook) GetDataOk() (*AdminBookData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminBook) SetData(v AdminBookData)`

SetData sets Data field to given value.

### HasData

`func (o *AdminBook) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *AdminBook) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminBook) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminBook) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminBook) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *AdminBook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminBook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminBook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminBook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



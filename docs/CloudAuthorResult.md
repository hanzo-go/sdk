# CloudAuthorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudAuthorData**](CloudAuthorData.md) | Data carries the author. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;. | [optional] 

## Methods

### NewCloudAuthorResult

`func NewCloudAuthorResult() *CloudAuthorResult`

NewCloudAuthorResult instantiates a new CloudAuthorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthorResultWithDefaults

`func NewCloudAuthorResultWithDefaults() *CloudAuthorResult`

NewCloudAuthorResultWithDefaults instantiates a new CloudAuthorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAuthorResult) GetData() CloudAuthorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAuthorResult) GetDataOk() (*CloudAuthorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAuthorResult) SetData(v CloudAuthorData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAuthorResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudAuthorResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudAuthorResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudAuthorResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudAuthorResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAuthorResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAuthorResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAuthorResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAuthorResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



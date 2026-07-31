# CloudBindOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudBindData**](CloudBindData.md) | Data is the bound signer. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewCloudBindOut

`func NewCloudBindOut() *CloudBindOut`

NewCloudBindOut instantiates a new CloudBindOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBindOutWithDefaults

`func NewCloudBindOutWithDefaults() *CloudBindOut`

NewCloudBindOutWithDefaults instantiates a new CloudBindOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudBindOut) GetData() CloudBindData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudBindOut) GetDataOk() (*CloudBindData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudBindOut) SetData(v CloudBindData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudBindOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudBindOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudBindOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudBindOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudBindOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBindOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBindOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBindOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBindOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



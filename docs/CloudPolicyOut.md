# CloudPolicyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudPolicyData**](CloudPolicyData.md) | Data is the stored policy. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewCloudPolicyOut

`func NewCloudPolicyOut() *CloudPolicyOut`

NewCloudPolicyOut instantiates a new CloudPolicyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPolicyOutWithDefaults

`func NewCloudPolicyOutWithDefaults() *CloudPolicyOut`

NewCloudPolicyOutWithDefaults instantiates a new CloudPolicyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudPolicyOut) GetData() CloudPolicyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudPolicyOut) GetDataOk() (*CloudPolicyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudPolicyOut) SetData(v CloudPolicyData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudPolicyOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudPolicyOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudPolicyOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudPolicyOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudPolicyOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPolicyOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPolicyOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPolicyOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPolicyOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



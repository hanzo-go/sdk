# PolicyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PolicyData**](PolicyData.md) | Data is the stored policy. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewPolicyOut

`func NewPolicyOut() *PolicyOut`

NewPolicyOut instantiates a new PolicyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyOutWithDefaults

`func NewPolicyOutWithDefaults() *PolicyOut`

NewPolicyOutWithDefaults instantiates a new PolicyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PolicyOut) GetData() PolicyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PolicyOut) GetDataOk() (*PolicyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PolicyOut) SetData(v PolicyData)`

SetData sets Data field to given value.

### HasData

`func (o *PolicyOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *PolicyOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PolicyOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PolicyOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *PolicyOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



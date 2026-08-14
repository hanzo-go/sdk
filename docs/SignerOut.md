# SignerOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**SignerData**](SignerData.md) | Data is the bound signer. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewSignerOut

`func NewSignerOut() *SignerOut`

NewSignerOut instantiates a new SignerOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerOutWithDefaults

`func NewSignerOutWithDefaults() *SignerOut`

NewSignerOutWithDefaults instantiates a new SignerOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SignerOut) GetData() SignerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignerOut) GetDataOk() (*SignerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignerOut) SetData(v SignerData)`

SetData sets Data field to given value.

### HasData

`func (o *SignerOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *SignerOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SignerOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SignerOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SignerOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *SignerOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignerOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignerOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SignerOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



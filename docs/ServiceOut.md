# ServiceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceOne**](ServiceOne.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceOut

`func NewServiceOut() *ServiceOut`

NewServiceOut instantiates a new ServiceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOutWithDefaults

`func NewServiceOutWithDefaults() *ServiceOut`

NewServiceOutWithDefaults instantiates a new ServiceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceOut) GetData() ServiceOne`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceOut) GetDataOk() (*ServiceOne, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceOut) SetData(v ServiceOne)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *ServiceOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ServiceOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ServiceOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ServiceOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



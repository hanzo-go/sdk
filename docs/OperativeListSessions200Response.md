# OperativeListSessions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]OperativeSession**](OperativeSession.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewOperativeListSessions200Response

`func NewOperativeListSessions200Response() *OperativeListSessions200Response`

NewOperativeListSessions200Response instantiates a new OperativeListSessions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeListSessions200ResponseWithDefaults

`func NewOperativeListSessions200ResponseWithDefaults() *OperativeListSessions200Response`

NewOperativeListSessions200ResponseWithDefaults instantiates a new OperativeListSessions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *OperativeListSessions200Response) GetSessions() []OperativeSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *OperativeListSessions200Response) GetSessionsOk() (*[]OperativeSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *OperativeListSessions200Response) SetSessions(v []OperativeSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *OperativeListSessions200Response) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetTotal

`func (o *OperativeListSessions200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OperativeListSessions200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OperativeListSessions200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OperativeListSessions200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



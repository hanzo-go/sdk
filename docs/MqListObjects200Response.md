# MqListObjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]MqObjectInfo**](MqObjectInfo.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListObjects200Response

`func NewMqListObjects200Response() *MqListObjects200Response`

NewMqListObjects200Response instantiates a new MqListObjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListObjects200ResponseWithDefaults

`func NewMqListObjects200ResponseWithDefaults() *MqListObjects200Response`

NewMqListObjects200ResponseWithDefaults instantiates a new MqListObjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MqListObjects200Response) GetObjects() []MqObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MqListObjects200Response) GetObjectsOk() (*[]MqObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MqListObjects200Response) SetObjects(v []MqObjectInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *MqListObjects200Response) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetTotal

`func (o *MqListObjects200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListObjects200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListObjects200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListObjects200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



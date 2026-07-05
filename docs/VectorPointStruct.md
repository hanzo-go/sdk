# VectorPointStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**VectorPointStructId**](VectorPointStructId.md) |  | 
**Vector** | [**VectorPointStructVector**](VectorPointStructVector.md) |  | 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVectorPointStruct

`func NewVectorPointStruct(id VectorPointStructId, vector VectorPointStructVector, ) *VectorPointStruct`

NewVectorPointStruct instantiates a new VectorPointStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorPointStructWithDefaults

`func NewVectorPointStructWithDefaults() *VectorPointStruct`

NewVectorPointStructWithDefaults instantiates a new VectorPointStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VectorPointStruct) GetId() VectorPointStructId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VectorPointStruct) GetIdOk() (*VectorPointStructId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VectorPointStruct) SetId(v VectorPointStructId)`

SetId sets Id field to given value.


### GetVector

`func (o *VectorPointStruct) GetVector() VectorPointStructVector`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *VectorPointStruct) GetVectorOk() (*VectorPointStructVector, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *VectorPointStruct) SetVector(v VectorPointStructVector)`

SetVector sets Vector field to given value.


### GetPayload

`func (o *VectorPointStruct) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VectorPointStruct) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VectorPointStruct) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *VectorPointStruct) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



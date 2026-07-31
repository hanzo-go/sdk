# GatewayEmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GatewayEmbeddingResponseDataInner**](GatewayEmbeddingResponseDataInner.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**GatewayEmbeddingResponseUsage**](GatewayEmbeddingResponseUsage.md) |  | [optional] 

## Methods

### NewGatewayEmbeddingResponse

`func NewGatewayEmbeddingResponse() *GatewayEmbeddingResponse`

NewGatewayEmbeddingResponse instantiates a new GatewayEmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayEmbeddingResponseWithDefaults

`func NewGatewayEmbeddingResponseWithDefaults() *GatewayEmbeddingResponse`

NewGatewayEmbeddingResponseWithDefaults instantiates a new GatewayEmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GatewayEmbeddingResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GatewayEmbeddingResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GatewayEmbeddingResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GatewayEmbeddingResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *GatewayEmbeddingResponse) GetData() []GatewayEmbeddingResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GatewayEmbeddingResponse) GetDataOk() (*[]GatewayEmbeddingResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GatewayEmbeddingResponse) SetData(v []GatewayEmbeddingResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GatewayEmbeddingResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModel

`func (o *GatewayEmbeddingResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GatewayEmbeddingResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GatewayEmbeddingResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GatewayEmbeddingResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *GatewayEmbeddingResponse) GetUsage() GatewayEmbeddingResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GatewayEmbeddingResponse) GetUsageOk() (*GatewayEmbeddingResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GatewayEmbeddingResponse) SetUsage(v GatewayEmbeddingResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GatewayEmbeddingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



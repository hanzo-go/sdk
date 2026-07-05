# AiImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]AiImageResponseDataInner**](AiImageResponseDataInner.md) |  | [optional] 

## Methods

### NewAiImageResponse

`func NewAiImageResponse() *AiImageResponse`

NewAiImageResponse instantiates a new AiImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImageResponseWithDefaults

`func NewAiImageResponseWithDefaults() *AiImageResponse`

NewAiImageResponseWithDefaults instantiates a new AiImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AiImageResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiImageResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiImageResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiImageResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *AiImageResponse) GetData() []AiImageResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiImageResponse) GetDataOk() (*[]AiImageResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiImageResponse) SetData(v []AiImageResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AiImageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



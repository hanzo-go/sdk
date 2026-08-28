# OpenaiImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]OpenaiImageResponseDataInner**](OpenaiImageResponseDataInner.md) |  | [optional] 

## Methods

### NewOpenaiImageResponse

`func NewOpenaiImageResponse() *OpenaiImageResponse`

NewOpenaiImageResponse instantiates a new OpenaiImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiImageResponseWithDefaults

`func NewOpenaiImageResponseWithDefaults() *OpenaiImageResponse`

NewOpenaiImageResponseWithDefaults instantiates a new OpenaiImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OpenaiImageResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenaiImageResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenaiImageResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OpenaiImageResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *OpenaiImageResponse) GetData() []OpenaiImageResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenaiImageResponse) GetDataOk() (*[]OpenaiImageResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenaiImageResponse) SetData(v []OpenaiImageResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *OpenaiImageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



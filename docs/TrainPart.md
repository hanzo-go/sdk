# TrainPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]string** | Headers must be sent verbatim; the signature covers them. | [optional] 
**Number** | Pointer to **int64** | Number orders the part from 1. | [optional] 
**Offset** | Pointer to **int64** | Offset is the part&#39;s first byte in the object. | [optional] 
**Size** | Pointer to **int64** | Size is the part&#39;s byte count. | [optional] 
**Url** | Pointer to **string** | URL is the presigned address to PUT the range to. | [optional] 

## Methods

### NewTrainPart

`func NewTrainPart() *TrainPart`

NewTrainPart instantiates a new TrainPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainPartWithDefaults

`func NewTrainPartWithDefaults() *TrainPart`

NewTrainPartWithDefaults instantiates a new TrainPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *TrainPart) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TrainPart) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TrainPart) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TrainPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetNumber

`func (o *TrainPart) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TrainPart) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TrainPart) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TrainPart) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOffset

`func (o *TrainPart) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TrainPart) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TrainPart) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TrainPart) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSize

`func (o *TrainPart) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrainPart) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrainPart) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrainPart) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *TrainPart) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrainPart) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrainPart) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrainPart) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



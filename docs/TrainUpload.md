# TrainUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int64** | Expires is when the grant stops being accepted, unix seconds. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers must be sent verbatim with the single PUT; the signature covers them. | [optional] 
**Method** | Pointer to **string** | Method is PUT. | [optional] 
**Parts** | Pointer to [**[]TrainPart**](TrainPart.md) | Parts are the object&#39;s consecutive byte ranges, each PUT to its own address. | [optional] 
**Url** | Pointer to **string** | URL is the presigned address of a single PUT; empty when Parts is given. | [optional] 

## Methods

### NewTrainUpload

`func NewTrainUpload() *TrainUpload`

NewTrainUpload instantiates a new TrainUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainUploadWithDefaults

`func NewTrainUploadWithDefaults() *TrainUpload`

NewTrainUploadWithDefaults instantiates a new TrainUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *TrainUpload) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TrainUpload) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TrainUpload) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TrainUpload) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHeaders

`func (o *TrainUpload) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TrainUpload) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TrainUpload) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TrainUpload) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *TrainUpload) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TrainUpload) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TrainUpload) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TrainUpload) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParts

`func (o *TrainUpload) GetParts() []TrainPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *TrainUpload) GetPartsOk() (*[]TrainPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *TrainUpload) SetParts(v []TrainPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *TrainUpload) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetUrl

`func (o *TrainUpload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrainUpload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrainUpload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrainUpload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



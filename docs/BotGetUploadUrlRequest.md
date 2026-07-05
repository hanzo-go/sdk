# BotGetUploadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 

## Methods

### NewBotGetUploadUrlRequest

`func NewBotGetUploadUrlRequest(filename string, ) *BotGetUploadUrlRequest`

NewBotGetUploadUrlRequest instantiates a new BotGetUploadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotGetUploadUrlRequestWithDefaults

`func NewBotGetUploadUrlRequestWithDefaults() *BotGetUploadUrlRequest`

NewBotGetUploadUrlRequestWithDefaults instantiates a new BotGetUploadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *BotGetUploadUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BotGetUploadUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BotGetUploadUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *BotGetUploadUrlRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BotGetUploadUrlRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BotGetUploadUrlRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BotGetUploadUrlRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



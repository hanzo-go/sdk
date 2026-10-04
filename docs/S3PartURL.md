# S3PartURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Part** | Pointer to **int64** | Part is the part number. | [optional] 
**Url** | Pointer to **string** | URL is the presigned address to PUT the part&#39;s bytes to, directly. | [optional] 

## Methods

### NewS3PartURL

`func NewS3PartURL() *S3PartURL`

NewS3PartURL instantiates a new S3PartURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3PartURLWithDefaults

`func NewS3PartURLWithDefaults() *S3PartURL`

NewS3PartURLWithDefaults instantiates a new S3PartURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPart

`func (o *S3PartURL) GetPart() int64`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *S3PartURL) GetPartOk() (*int64, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *S3PartURL) SetPart(v int64)`

SetPart sets Part field to given value.

### HasPart

`func (o *S3PartURL) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetUrl

`func (o *S3PartURL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *S3PartURL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *S3PartURL) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *S3PartURL) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



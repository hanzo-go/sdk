# S3PartURLs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int64** | Expiry is how many seconds the URLs stay valid. | [optional] 
**Urls** | Pointer to [**[]S3PartURL**](S3PartURL.md) | URLs are one presigned PUT per part asked for. | [optional] 

## Methods

### NewS3PartURLs

`func NewS3PartURLs() *S3PartURLs`

NewS3PartURLs instantiates a new S3PartURLs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3PartURLsWithDefaults

`func NewS3PartURLsWithDefaults() *S3PartURLs`

NewS3PartURLsWithDefaults instantiates a new S3PartURLs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *S3PartURLs) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *S3PartURLs) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *S3PartURLs) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *S3PartURLs) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetUrls

`func (o *S3PartURLs) GetUrls() []S3PartURL`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *S3PartURLs) GetUrlsOk() (*[]S3PartURL, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *S3PartURLs) SetUrls(v []S3PartURL)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *S3PartURLs) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



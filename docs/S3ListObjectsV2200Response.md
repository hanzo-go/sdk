# S3ListObjectsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to [**[]S3ObjectInfo**](S3ObjectInfo.md) |  | [optional] 
**CommonPrefixes** | Pointer to [**[]S3ListObjectsV2200ResponseCommonPrefixesInner**](S3ListObjectsV2200ResponseCommonPrefixesInner.md) |  | [optional] 
**IsTruncated** | Pointer to **bool** |  | [optional] 
**NextContinuationToken** | Pointer to **string** |  | [optional] 

## Methods

### NewS3ListObjectsV2200Response

`func NewS3ListObjectsV2200Response() *S3ListObjectsV2200Response`

NewS3ListObjectsV2200Response instantiates a new S3ListObjectsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ListObjectsV2200ResponseWithDefaults

`func NewS3ListObjectsV2200ResponseWithDefaults() *S3ListObjectsV2200Response`

NewS3ListObjectsV2200ResponseWithDefaults instantiates a new S3ListObjectsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *S3ListObjectsV2200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3ListObjectsV2200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3ListObjectsV2200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3ListObjectsV2200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *S3ListObjectsV2200Response) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3ListObjectsV2200Response) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3ListObjectsV2200Response) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3ListObjectsV2200Response) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetContents

`func (o *S3ListObjectsV2200Response) GetContents() []S3ObjectInfo`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *S3ListObjectsV2200Response) GetContentsOk() (*[]S3ObjectInfo, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *S3ListObjectsV2200Response) SetContents(v []S3ObjectInfo)`

SetContents sets Contents field to given value.

### HasContents

`func (o *S3ListObjectsV2200Response) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *S3ListObjectsV2200Response) GetCommonPrefixes() []S3ListObjectsV2200ResponseCommonPrefixesInner`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *S3ListObjectsV2200Response) GetCommonPrefixesOk() (*[]S3ListObjectsV2200ResponseCommonPrefixesInner, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *S3ListObjectsV2200Response) SetCommonPrefixes(v []S3ListObjectsV2200ResponseCommonPrefixesInner)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *S3ListObjectsV2200Response) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetIsTruncated

`func (o *S3ListObjectsV2200Response) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *S3ListObjectsV2200Response) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *S3ListObjectsV2200Response) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *S3ListObjectsV2200Response) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetNextContinuationToken

`func (o *S3ListObjectsV2200Response) GetNextContinuationToken() string`

GetNextContinuationToken returns the NextContinuationToken field if non-nil, zero value otherwise.

### GetNextContinuationTokenOk

`func (o *S3ListObjectsV2200Response) GetNextContinuationTokenOk() (*string, bool)`

GetNextContinuationTokenOk returns a tuple with the NextContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextContinuationToken

`func (o *S3ListObjectsV2200Response) SetNextContinuationToken(v string)`

SetNextContinuationToken sets NextContinuationToken field to given value.

### HasNextContinuationToken

`func (o *S3ListObjectsV2200Response) HasNextContinuationToken() bool`

HasNextContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



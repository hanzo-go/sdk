# S3AdminInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SqsARN** | Pointer to **[]string** |  | [optional] 

## Methods

### NewS3AdminInfo200Response

`func NewS3AdminInfo200Response() *S3AdminInfo200Response`

NewS3AdminInfo200Response instantiates a new S3AdminInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AdminInfo200ResponseWithDefaults

`func NewS3AdminInfo200ResponseWithDefaults() *S3AdminInfo200Response`

NewS3AdminInfo200ResponseWithDefaults instantiates a new S3AdminInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *S3AdminInfo200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *S3AdminInfo200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *S3AdminInfo200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *S3AdminInfo200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegion

`func (o *S3AdminInfo200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3AdminInfo200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3AdminInfo200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *S3AdminInfo200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSqsARN

`func (o *S3AdminInfo200Response) GetSqsARN() []string`

GetSqsARN returns the SqsARN field if non-nil, zero value otherwise.

### GetSqsARNOk

`func (o *S3AdminInfo200Response) GetSqsARNOk() (*[]string, bool)`

GetSqsARNOk returns a tuple with the SqsARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqsARN

`func (o *S3AdminInfo200Response) SetSqsARN(v []string)`

SetSqsARN sets SqsARN field to given value.

### HasSqsARN

`func (o *S3AdminInfo200Response) HasSqsARN() bool`

HasSqsARN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



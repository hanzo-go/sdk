# S3BucketPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Statement** | Pointer to [**[]S3BucketPolicyStatementInner**](S3BucketPolicyStatementInner.md) |  | [optional] 

## Methods

### NewS3BucketPolicy

`func NewS3BucketPolicy() *S3BucketPolicy`

NewS3BucketPolicy instantiates a new S3BucketPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketPolicyWithDefaults

`func NewS3BucketPolicyWithDefaults() *S3BucketPolicy`

NewS3BucketPolicyWithDefaults instantiates a new S3BucketPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *S3BucketPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *S3BucketPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *S3BucketPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *S3BucketPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatement

`func (o *S3BucketPolicy) GetStatement() []S3BucketPolicyStatementInner`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *S3BucketPolicy) GetStatementOk() (*[]S3BucketPolicyStatementInner, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *S3BucketPolicy) SetStatement(v []S3BucketPolicyStatementInner)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *S3BucketPolicy) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



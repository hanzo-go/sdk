# KnowledgeFileIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is one of the org&#39;s buckets, by the friendly name /v1/s3/buckets lists. Required. | [optional] 
**Key** | Pointer to **string** | Key is the object&#39;s key in that bucket, as the upload wrote it. Required. | [optional] 
**Project** | Pointer to **string** | Project indexes the file under one project scope. Empty is the org&#39;s own. | [optional] 

## Methods

### NewKnowledgeFileIn

`func NewKnowledgeFileIn() *KnowledgeFileIn`

NewKnowledgeFileIn instantiates a new KnowledgeFileIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileInWithDefaults

`func NewKnowledgeFileInWithDefaults() *KnowledgeFileIn`

NewKnowledgeFileInWithDefaults instantiates a new KnowledgeFileIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KnowledgeFileIn) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KnowledgeFileIn) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KnowledgeFileIn) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KnowledgeFileIn) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *KnowledgeFileIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KnowledgeFileIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KnowledgeFileIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KnowledgeFileIn) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProject

`func (o *KnowledgeFileIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KnowledgeFileIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KnowledgeFileIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KnowledgeFileIn) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



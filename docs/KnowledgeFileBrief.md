# KnowledgeFileBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket and Key are where its bytes are: GET /v1/s3/buckets/{bucket}/objects/{key} answers a signed download URL. | [optional] 
**Id** | Pointer to **string** | ID is the file&#39;s id. | [optional] 
**Key** | Pointer to **string** | Key is the object&#39;s key in Bucket. | [optional] 
**Name** | Pointer to **string** | Name is the file&#39;s name. | [optional] 
**Type** | Pointer to **string** | Type is the file&#39;s media type. | [optional] 

## Methods

### NewKnowledgeFileBrief

`func NewKnowledgeFileBrief() *KnowledgeFileBrief`

NewKnowledgeFileBrief instantiates a new KnowledgeFileBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileBriefWithDefaults

`func NewKnowledgeFileBriefWithDefaults() *KnowledgeFileBrief`

NewKnowledgeFileBriefWithDefaults instantiates a new KnowledgeFileBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KnowledgeFileBrief) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KnowledgeFileBrief) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KnowledgeFileBrief) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KnowledgeFileBrief) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetId

`func (o *KnowledgeFileBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeFileBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeFileBrief) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeFileBrief) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *KnowledgeFileBrief) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KnowledgeFileBrief) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KnowledgeFileBrief) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KnowledgeFileBrief) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeFileBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeFileBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeFileBrief) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnowledgeFileBrief) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *KnowledgeFileBrief) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KnowledgeFileBrief) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KnowledgeFileBrief) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KnowledgeFileBrief) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudPushEventRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**CloudPushEventPusher**](CloudPushEventPusher.md) |  | [optional] 

## Methods

### NewCloudPushEventRepository

`func NewCloudPushEventRepository() *CloudPushEventRepository`

NewCloudPushEventRepository instantiates a new CloudPushEventRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPushEventRepositoryWithDefaults

`func NewCloudPushEventRepositoryWithDefaults() *CloudPushEventRepository`

NewCloudPushEventRepositoryWithDefaults instantiates a new CloudPushEventRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneUrl

`func (o *CloudPushEventRepository) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *CloudPushEventRepository) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *CloudPushEventRepository) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *CloudPushEventRepository) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetName

`func (o *CloudPushEventRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPushEventRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPushEventRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPushEventRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudPushEventRepository) GetOwner() CloudPushEventPusher`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudPushEventRepository) GetOwnerOk() (*CloudPushEventPusher, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudPushEventRepository) SetOwner(v CloudPushEventPusher)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudPushEventRepository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



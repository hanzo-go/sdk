# KmsCreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** |  | 
**OrganizationId** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsCreateProjectRequest

`func NewKmsCreateProjectRequest(projectName string, organizationId string, ) *KmsCreateProjectRequest`

NewKmsCreateProjectRequest instantiates a new KmsCreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateProjectRequestWithDefaults

`func NewKmsCreateProjectRequestWithDefaults() *KmsCreateProjectRequest`

NewKmsCreateProjectRequestWithDefaults instantiates a new KmsCreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *KmsCreateProjectRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KmsCreateProjectRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KmsCreateProjectRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetOrganizationId

`func (o *KmsCreateProjectRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KmsCreateProjectRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KmsCreateProjectRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetSlug

`func (o *KmsCreateProjectRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *KmsCreateProjectRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *KmsCreateProjectRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *KmsCreateProjectRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetType

`func (o *KmsCreateProjectRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsCreateProjectRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsCreateProjectRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsCreateProjectRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplateId

`func (o *KmsCreateProjectRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *KmsCreateProjectRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *KmsCreateProjectRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *KmsCreateProjectRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



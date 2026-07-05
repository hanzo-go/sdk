# KmsCreateIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewKmsCreateIdentityRequest

`func NewKmsCreateIdentityRequest(name string, organizationId string, role string, ) *KmsCreateIdentityRequest`

NewKmsCreateIdentityRequest instantiates a new KmsCreateIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateIdentityRequestWithDefaults

`func NewKmsCreateIdentityRequestWithDefaults() *KmsCreateIdentityRequest`

NewKmsCreateIdentityRequestWithDefaults instantiates a new KmsCreateIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsCreateIdentityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateIdentityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateIdentityRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *KmsCreateIdentityRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KmsCreateIdentityRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KmsCreateIdentityRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetRole

`func (o *KmsCreateIdentityRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KmsCreateIdentityRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KmsCreateIdentityRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



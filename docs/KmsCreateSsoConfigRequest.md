# KmsCreateSsoConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Type** | **string** |  | 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewKmsCreateSsoConfigRequest

`func NewKmsCreateSsoConfigRequest(orgId string, type_ string, ) *KmsCreateSsoConfigRequest`

NewKmsCreateSsoConfigRequest instantiates a new KmsCreateSsoConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateSsoConfigRequestWithDefaults

`func NewKmsCreateSsoConfigRequestWithDefaults() *KmsCreateSsoConfigRequest`

NewKmsCreateSsoConfigRequestWithDefaults instantiates a new KmsCreateSsoConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *KmsCreateSsoConfigRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *KmsCreateSsoConfigRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *KmsCreateSsoConfigRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetType

`func (o *KmsCreateSsoConfigRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsCreateSsoConfigRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsCreateSsoConfigRequest) SetType(v string)`

SetType sets Type field to given value.


### GetIsActive

`func (o *KmsCreateSsoConfigRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *KmsCreateSsoConfigRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *KmsCreateSsoConfigRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *KmsCreateSsoConfigRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



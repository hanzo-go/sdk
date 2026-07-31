# CloudRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the formation was opened. | [optional] 
**Name** | Pointer to **string** | Name is the company name the entity is being formed under. | [optional] 
**Org** | Pointer to **string** | Org is the org whose formation this row projects. | [optional] 
**Stage** | Pointer to **string** | Stage is the formation&#39;s current state — what the platform reads to see which formations are stalled and where. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity being formed: c-corp, llc or dao-llc. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the most recent write to the formation, and the key the register sorts on (newest activity first). | [optional] 

## Methods

### NewCloudRegistration

`func NewCloudRegistration() *CloudRegistration`

NewCloudRegistration instantiates a new CloudRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegistrationWithDefaults

`func NewCloudRegistrationWithDefaults() *CloudRegistration`

NewCloudRegistrationWithDefaults instantiates a new CloudRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudRegistration) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudRegistration) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudRegistration) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudRegistration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *CloudRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRegistration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRegistration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudRegistration) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudRegistration) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudRegistration) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudRegistration) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStage

`func (o *CloudRegistration) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CloudRegistration) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CloudRegistration) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CloudRegistration) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStructure

`func (o *CloudRegistration) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *CloudRegistration) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *CloudRegistration) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *CloudRegistration) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudRegistration) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudRegistration) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudRegistration) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudRegistration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



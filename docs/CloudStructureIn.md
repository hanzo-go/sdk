# CloudStructureIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation: DE or WY. | [optional] 
**Name** | Pointer to **string** | Name is the proposed company name. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewCloudStructureIn

`func NewCloudStructureIn() *CloudStructureIn`

NewCloudStructureIn instantiates a new CloudStructureIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStructureInWithDefaults

`func NewCloudStructureInWithDefaults() *CloudStructureIn`

NewCloudStructureInWithDefaults instantiates a new CloudStructureIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJurisdiction

`func (o *CloudStructureIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *CloudStructureIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *CloudStructureIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *CloudStructureIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *CloudStructureIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudStructureIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudStructureIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudStructureIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStructure

`func (o *CloudStructureIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *CloudStructureIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *CloudStructureIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *CloudStructureIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StructureIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation: DE or WY. | [optional] 
**Name** | Pointer to **string** | Name is the proposed company name. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewStructureIn

`func NewStructureIn() *StructureIn`

NewStructureIn instantiates a new StructureIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructureInWithDefaults

`func NewStructureInWithDefaults() *StructureIn`

NewStructureInWithDefaults instantiates a new StructureIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJurisdiction

`func (o *StructureIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *StructureIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *StructureIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *StructureIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *StructureIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructureIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructureIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StructureIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStructure

`func (o *StructureIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *StructureIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *StructureIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *StructureIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



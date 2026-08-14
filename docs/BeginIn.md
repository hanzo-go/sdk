# BeginIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyIncorporated** | Pointer to **bool** | AlreadyIncorporated declares an org that already has an entity, which takes the import path (POST /v1/company/skip) instead of the formation path. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation: DE or WY. | [optional] 
**Name** | Pointer to **string** | Name is the proposed company name. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity to form: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewBeginIn

`func NewBeginIn() *BeginIn`

NewBeginIn instantiates a new BeginIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeginInWithDefaults

`func NewBeginInWithDefaults() *BeginIn`

NewBeginInWithDefaults instantiates a new BeginIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyIncorporated

`func (o *BeginIn) GetAlreadyIncorporated() bool`

GetAlreadyIncorporated returns the AlreadyIncorporated field if non-nil, zero value otherwise.

### GetAlreadyIncorporatedOk

`func (o *BeginIn) GetAlreadyIncorporatedOk() (*bool, bool)`

GetAlreadyIncorporatedOk returns a tuple with the AlreadyIncorporated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyIncorporated

`func (o *BeginIn) SetAlreadyIncorporated(v bool)`

SetAlreadyIncorporated sets AlreadyIncorporated field to given value.

### HasAlreadyIncorporated

`func (o *BeginIn) HasAlreadyIncorporated() bool`

HasAlreadyIncorporated returns a boolean if a field has been set.

### GetJurisdiction

`func (o *BeginIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *BeginIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *BeginIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *BeginIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *BeginIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BeginIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BeginIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BeginIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStructure

`func (o *BeginIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *BeginIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *BeginIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *BeginIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



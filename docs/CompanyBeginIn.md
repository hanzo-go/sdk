# CompanyBeginIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyIncorporated** | Pointer to **bool** | AlreadyIncorporated declares an org that already has an entity, which takes the import path (POST /v1/company/skip) instead of the formation path. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation: DE or WY. | [optional] 
**Name** | Pointer to **string** | Name is the proposed company name. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity to form: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewCompanyBeginIn

`func NewCompanyBeginIn() *CompanyBeginIn`

NewCompanyBeginIn instantiates a new CompanyBeginIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyBeginInWithDefaults

`func NewCompanyBeginInWithDefaults() *CompanyBeginIn`

NewCompanyBeginInWithDefaults instantiates a new CompanyBeginIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyIncorporated

`func (o *CompanyBeginIn) GetAlreadyIncorporated() bool`

GetAlreadyIncorporated returns the AlreadyIncorporated field if non-nil, zero value otherwise.

### GetAlreadyIncorporatedOk

`func (o *CompanyBeginIn) GetAlreadyIncorporatedOk() (*bool, bool)`

GetAlreadyIncorporatedOk returns a tuple with the AlreadyIncorporated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyIncorporated

`func (o *CompanyBeginIn) SetAlreadyIncorporated(v bool)`

SetAlreadyIncorporated sets AlreadyIncorporated field to given value.

### HasAlreadyIncorporated

`func (o *CompanyBeginIn) HasAlreadyIncorporated() bool`

HasAlreadyIncorporated returns a boolean if a field has been set.

### GetJurisdiction

`func (o *CompanyBeginIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *CompanyBeginIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *CompanyBeginIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *CompanyBeginIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *CompanyBeginIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyBeginIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyBeginIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyBeginIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStructure

`func (o *CompanyBeginIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *CompanyBeginIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *CompanyBeginIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *CompanyBeginIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



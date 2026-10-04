# PrincipalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ein** | Pointer to **string** | EIN is on_file when the org&#39;s W-9 carries an EIN, and absent otherwise. | [optional] 
**Founders** | Pointer to [**[]PrincipalFounder**](PrincipalFounder.md) | Founders are its beneficial owners with their verification status — answered to the org&#39;s admins. | [optional] 
**Imported** | Pointer to **bool** | Imported is an entity that existed before and was imported. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the U.S. state of formation. | [optional] 
**Name** | Pointer to **string** | Name is the company name. | [optional] 
**Stage** | Pointer to **string** | Stage is the formation&#39;s stage; \&quot;company\&quot; is formed. | [optional] 
**Structure** | Pointer to **string** | Structure is c-corp, llc or dao-llc. | [optional] 

## Methods

### NewPrincipalEntity

`func NewPrincipalEntity() *PrincipalEntity`

NewPrincipalEntity instantiates a new PrincipalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntityWithDefaults

`func NewPrincipalEntityWithDefaults() *PrincipalEntity`

NewPrincipalEntityWithDefaults instantiates a new PrincipalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEin

`func (o *PrincipalEntity) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *PrincipalEntity) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *PrincipalEntity) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *PrincipalEntity) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetFounders

`func (o *PrincipalEntity) GetFounders() []PrincipalFounder`

GetFounders returns the Founders field if non-nil, zero value otherwise.

### GetFoundersOk

`func (o *PrincipalEntity) GetFoundersOk() (*[]PrincipalFounder, bool)`

GetFoundersOk returns a tuple with the Founders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounders

`func (o *PrincipalEntity) SetFounders(v []PrincipalFounder)`

SetFounders sets Founders field to given value.

### HasFounders

`func (o *PrincipalEntity) HasFounders() bool`

HasFounders returns a boolean if a field has been set.

### GetImported

`func (o *PrincipalEntity) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *PrincipalEntity) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *PrincipalEntity) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *PrincipalEntity) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetJurisdiction

`func (o *PrincipalEntity) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *PrincipalEntity) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *PrincipalEntity) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *PrincipalEntity) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *PrincipalEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStage

`func (o *PrincipalEntity) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PrincipalEntity) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PrincipalEntity) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PrincipalEntity) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStructure

`func (o *PrincipalEntity) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *PrincipalEntity) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *PrincipalEntity) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *PrincipalEntity) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



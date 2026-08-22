# TariffIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentOfRecord** | Pointer to **bool** | AgentOfRecord puts us on file as the entity&#39;s agent, yearly. | [optional] 
**ExpeditedEin** | Pointer to **bool** | ExpeditedEIN prioritises the EIN. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity being formed: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewTariffIn

`func NewTariffIn() *TariffIn`

NewTariffIn instantiates a new TariffIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffInWithDefaults

`func NewTariffInWithDefaults() *TariffIn`

NewTariffInWithDefaults instantiates a new TariffIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentOfRecord

`func (o *TariffIn) GetAgentOfRecord() bool`

GetAgentOfRecord returns the AgentOfRecord field if non-nil, zero value otherwise.

### GetAgentOfRecordOk

`func (o *TariffIn) GetAgentOfRecordOk() (*bool, bool)`

GetAgentOfRecordOk returns a tuple with the AgentOfRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOfRecord

`func (o *TariffIn) SetAgentOfRecord(v bool)`

SetAgentOfRecord sets AgentOfRecord field to given value.

### HasAgentOfRecord

`func (o *TariffIn) HasAgentOfRecord() bool`

HasAgentOfRecord returns a boolean if a field has been set.

### GetExpeditedEin

`func (o *TariffIn) GetExpeditedEin() bool`

GetExpeditedEin returns the ExpeditedEin field if non-nil, zero value otherwise.

### GetExpeditedEinOk

`func (o *TariffIn) GetExpeditedEinOk() (*bool, bool)`

GetExpeditedEinOk returns a tuple with the ExpeditedEin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpeditedEin

`func (o *TariffIn) SetExpeditedEin(v bool)`

SetExpeditedEin sets ExpeditedEin field to given value.

### HasExpeditedEin

`func (o *TariffIn) HasExpeditedEin() bool`

HasExpeditedEin returns a boolean if a field has been set.

### GetJurisdiction

`func (o *TariffIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *TariffIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *TariffIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *TariffIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetStructure

`func (o *TariffIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *TariffIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *TariffIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *TariffIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpkeepIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentOfRecord** | Pointer to **bool** | AgentOfRecord includes our agent fee, which most entities owe someone. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state the entity stands in. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity kind: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewUpkeepIn

`func NewUpkeepIn() *UpkeepIn`

NewUpkeepIn instantiates a new UpkeepIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpkeepInWithDefaults

`func NewUpkeepInWithDefaults() *UpkeepIn`

NewUpkeepInWithDefaults instantiates a new UpkeepIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentOfRecord

`func (o *UpkeepIn) GetAgentOfRecord() bool`

GetAgentOfRecord returns the AgentOfRecord field if non-nil, zero value otherwise.

### GetAgentOfRecordOk

`func (o *UpkeepIn) GetAgentOfRecordOk() (*bool, bool)`

GetAgentOfRecordOk returns a tuple with the AgentOfRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOfRecord

`func (o *UpkeepIn) SetAgentOfRecord(v bool)`

SetAgentOfRecord sets AgentOfRecord field to given value.

### HasAgentOfRecord

`func (o *UpkeepIn) HasAgentOfRecord() bool`

HasAgentOfRecord returns a boolean if a field has been set.

### GetJurisdiction

`func (o *UpkeepIn) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *UpkeepIn) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *UpkeepIn) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *UpkeepIn) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetStructure

`func (o *UpkeepIn) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *UpkeepIn) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *UpkeepIn) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *UpkeepIn) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



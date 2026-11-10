# GraphGraphOntologyClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edge** | Pointer to **bool** | Edge is true for the relationship-type classes (CS* and nm_*RelationshipType) that reify a two-row edge sharing a relationshipId in stable_graph. | [optional] 
**Group** | Pointer to **string** | Group is the class family: core, person, attribute, place, event, activity, activity type, transport, media, animal, commerce or edge. | [optional] 
**Id** | Pointer to **string** | ID is the raw code stored in graph.db, such as SB104, CS165 or nm_walletOrder. | [optional] 
**Label** | Pointer to **string** | Label is the human-readable class name, such as person or employment relationship type. | [optional] 
**Mechanism** | Pointer to **string** | Mechanism records whether the daemon mechanism populating the class is implemented, unverified or unlocated. | [optional] 

## Methods

### NewGraphGraphOntologyClass

`func NewGraphGraphOntologyClass() *GraphGraphOntologyClass`

NewGraphGraphOntologyClass instantiates a new GraphGraphOntologyClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphOntologyClassWithDefaults

`func NewGraphGraphOntologyClassWithDefaults() *GraphGraphOntologyClass`

NewGraphGraphOntologyClassWithDefaults instantiates a new GraphGraphOntologyClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdge

`func (o *GraphGraphOntologyClass) GetEdge() bool`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *GraphGraphOntologyClass) GetEdgeOk() (*bool, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *GraphGraphOntologyClass) SetEdge(v bool)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *GraphGraphOntologyClass) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### GetGroup

`func (o *GraphGraphOntologyClass) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GraphGraphOntologyClass) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GraphGraphOntologyClass) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GraphGraphOntologyClass) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *GraphGraphOntologyClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphGraphOntologyClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphGraphOntologyClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphGraphOntologyClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GraphGraphOntologyClass) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GraphGraphOntologyClass) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GraphGraphOntologyClass) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GraphGraphOntologyClass) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMechanism

`func (o *GraphGraphOntologyClass) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *GraphGraphOntologyClass) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *GraphGraphOntologyClass) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *GraphGraphOntologyClass) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



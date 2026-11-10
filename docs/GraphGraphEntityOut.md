# GraphGraphEntityOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant this state was resolved at, RFC 3339. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the valid instant this state was resolved at, RFC 3339. | [optional] 
**Entity** | Pointer to **string** | Entity is the folded entity key the question named. | [optional] 
**Expired** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Expired is every statement about the entity that was known by AsKnown and begun by AsOf but no longer holds at AsOf because its until passed, a retraction ended it, or a later statement superseded it. | [optional] 
**Incoming** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Incoming is every in-force edge pointing at the entity at the point, ordered by source entity and relation key. | [optional] 
**Known** | Pointer to **bool** | Known is true when at least one property, outgoing edge, incoming edge or expired statement was found for the entity. | [optional] 
**Outgoing** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Outgoing is every in-force edge asserted about the entity at the point, ordered by relation key and then winner order. | [optional] 
**Properties** | Pointer to [**[]GraphGraphEntityProperty**](GraphGraphEntityProperty.md) | Properties is every in-force scalar property of the entity at the point, ordered by relation key. | [optional] 
**Scorecard** | Pointer to [**[]GraphGraphKinshipScore**](GraphGraphKinshipScore.md) | Scorecard is the kinship and closeness scores holding on the entity, highest score first. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says the read of outgoing or incoming assertions hit the walk ceiling. | [optional] 
**Type** | Pointer to **string** | Type is the winning &#x60;type&#x60; of the entity at the point. Absent when no type holds. | [optional] 

## Methods

### NewGraphGraphEntityOut

`func NewGraphGraphEntityOut() *GraphGraphEntityOut`

NewGraphGraphEntityOut instantiates a new GraphGraphEntityOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphEntityOutWithDefaults

`func NewGraphGraphEntityOutWithDefaults() *GraphGraphEntityOut`

NewGraphGraphEntityOutWithDefaults instantiates a new GraphGraphEntityOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphEntityOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphEntityOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphEntityOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphEntityOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphEntityOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphEntityOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphEntityOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphEntityOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetEntity

`func (o *GraphGraphEntityOut) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphEntityOut) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphEntityOut) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphGraphEntityOut) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetExpired

`func (o *GraphGraphEntityOut) GetExpired() []GraphWireFact`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *GraphGraphEntityOut) GetExpiredOk() (*[]GraphWireFact, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *GraphGraphEntityOut) SetExpired(v []GraphWireFact)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *GraphGraphEntityOut) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetIncoming

`func (o *GraphGraphEntityOut) GetIncoming() []GraphWireFact`

GetIncoming returns the Incoming field if non-nil, zero value otherwise.

### GetIncomingOk

`func (o *GraphGraphEntityOut) GetIncomingOk() (*[]GraphWireFact, bool)`

GetIncomingOk returns a tuple with the Incoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoming

`func (o *GraphGraphEntityOut) SetIncoming(v []GraphWireFact)`

SetIncoming sets Incoming field to given value.

### HasIncoming

`func (o *GraphGraphEntityOut) HasIncoming() bool`

HasIncoming returns a boolean if a field has been set.

### GetKnown

`func (o *GraphGraphEntityOut) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *GraphGraphEntityOut) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *GraphGraphEntityOut) SetKnown(v bool)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *GraphGraphEntityOut) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetOutgoing

`func (o *GraphGraphEntityOut) GetOutgoing() []GraphWireFact`

GetOutgoing returns the Outgoing field if non-nil, zero value otherwise.

### GetOutgoingOk

`func (o *GraphGraphEntityOut) GetOutgoingOk() (*[]GraphWireFact, bool)`

GetOutgoingOk returns a tuple with the Outgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoing

`func (o *GraphGraphEntityOut) SetOutgoing(v []GraphWireFact)`

SetOutgoing sets Outgoing field to given value.

### HasOutgoing

`func (o *GraphGraphEntityOut) HasOutgoing() bool`

HasOutgoing returns a boolean if a field has been set.

### GetProperties

`func (o *GraphGraphEntityOut) GetProperties() []GraphGraphEntityProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GraphGraphEntityOut) GetPropertiesOk() (*[]GraphGraphEntityProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GraphGraphEntityOut) SetProperties(v []GraphGraphEntityProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GraphGraphEntityOut) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScorecard

`func (o *GraphGraphEntityOut) GetScorecard() []GraphGraphKinshipScore`

GetScorecard returns the Scorecard field if non-nil, zero value otherwise.

### GetScorecardOk

`func (o *GraphGraphEntityOut) GetScorecardOk() (*[]GraphGraphKinshipScore, bool)`

GetScorecardOk returns a tuple with the Scorecard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorecard

`func (o *GraphGraphEntityOut) SetScorecard(v []GraphGraphKinshipScore)`

SetScorecard sets Scorecard field to given value.

### HasScorecard

`func (o *GraphGraphEntityOut) HasScorecard() bool`

HasScorecard returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphEntityOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphEntityOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphEntityOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphEntityOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetType

`func (o *GraphGraphEntityOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphGraphEntityOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphGraphEntityOut) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GraphGraphEntityOut) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



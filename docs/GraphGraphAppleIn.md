# GraphGraphAppleIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **string** | At is the valid instant of the snapshot, RFC 3339. Required. | 
**Expired** | Pointer to [**[]GraphGraphAppleTriple**](GraphGraphAppleTriple.md) | Expired is the withdrawn rows from expired_stable_graph. Each is filed with an until instant so it stays in the bitemporal record and surfaces in graphEntity&#39;s expired list without holding in live resolution. | [optional] 
**Owner** | Pointer to **string** | Owner is the entity key for the device owner (\&quot;(Me)\&quot;). When set, entities flagged with nm_isCurrentUser are linked to Owner via &#x60;same&#x60;, and high-confidence kinship scorecard columns emit edges from Owner. | [optional] 
**Scorecards** | Pointer to [**[]GraphGraphAppleScorecard**](GraphGraphAppleScorecard.md) | Scorecards is the per-person kinship probability rows from entityTagging.db. | [optional] 
**Source** | Pointer to **string** | Source names the device or store this snapshot came from. Absent defaults to apple/intelligence-platform. | [optional] 
**Triples** | Pointer to [**[]GraphGraphAppleTriple**](GraphGraphAppleTriple.md) | Triples is the live rows from stable_graph. Reified two-row edges sharing a relationship_id are paired into edges automatically. | [optional] 

## Methods

### NewGraphGraphAppleIn

`func NewGraphGraphAppleIn(at string, ) *GraphGraphAppleIn`

NewGraphGraphAppleIn instantiates a new GraphGraphAppleIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAppleInWithDefaults

`func NewGraphGraphAppleInWithDefaults() *GraphGraphAppleIn`

NewGraphGraphAppleInWithDefaults instantiates a new GraphGraphAppleIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphAppleIn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphAppleIn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphAppleIn) SetAt(v string)`

SetAt sets At field to given value.


### GetExpired

`func (o *GraphGraphAppleIn) GetExpired() []GraphGraphAppleTriple`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *GraphGraphAppleIn) GetExpiredOk() (*[]GraphGraphAppleTriple, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *GraphGraphAppleIn) SetExpired(v []GraphGraphAppleTriple)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *GraphGraphAppleIn) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetOwner

`func (o *GraphGraphAppleIn) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GraphGraphAppleIn) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GraphGraphAppleIn) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GraphGraphAppleIn) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScorecards

`func (o *GraphGraphAppleIn) GetScorecards() []GraphGraphAppleScorecard`

GetScorecards returns the Scorecards field if non-nil, zero value otherwise.

### GetScorecardsOk

`func (o *GraphGraphAppleIn) GetScorecardsOk() (*[]GraphGraphAppleScorecard, bool)`

GetScorecardsOk returns a tuple with the Scorecards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorecards

`func (o *GraphGraphAppleIn) SetScorecards(v []GraphGraphAppleScorecard)`

SetScorecards sets Scorecards field to given value.

### HasScorecards

`func (o *GraphGraphAppleIn) HasScorecards() bool`

HasScorecards returns a boolean if a field has been set.

### GetSource

`func (o *GraphGraphAppleIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphGraphAppleIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphGraphAppleIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GraphGraphAppleIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTriples

`func (o *GraphGraphAppleIn) GetTriples() []GraphGraphAppleTriple`

GetTriples returns the Triples field if non-nil, zero value otherwise.

### GetTriplesOk

`func (o *GraphGraphAppleIn) GetTriplesOk() (*[]GraphGraphAppleTriple, bool)`

GetTriplesOk returns a tuple with the Triples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriples

`func (o *GraphGraphAppleIn) SetTriples(v []GraphGraphAppleTriple)`

SetTriples sets Triples field to given value.

### HasTriples

`func (o *GraphGraphAppleIn) HasTriples() bool`

HasTriples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GraphGraphPathOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant it was taken at, RFC 3339, the same way. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the instant the path was taken at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Bound** | Pointer to **int64** | Bound is the ceiling the search was held to, the same one a walk is. | [optional] 
**Decay** | Pointer to **float64** | Decay is the product of the edges&#39; confidences: how much certainty survives the whole chain. Absent unless every edge states a confidence above zero — an unstated confidence reads as zero, and a product of numbers nobody gave measures nothing. | [optional] 
**Edges** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Edges is the path from From to To, one assertion per hop, each with its source, evidence and filer. An edge is listed as it was asserted, so under in or both it may point against the direction of travel. | [optional] 
**Found** | Pointer to **bool** | Found is false when no path of in-force edges joins the two within the bound. That is an answer, not an error. | [optional] 
**Hops** | Pointer to **int64** | Hops is how many edges the path crosses: zero when From is To, and zero when nothing was found. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says the bound stopped the search before a path was found, so Found false means none within the bound rather than none at all. | [optional] 
**Weakest** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Weakest is the edge with the lowest confidence, the first of them on a tie: the link the chain is only as strong as. Absent exactly when Decay is. | [optional] 

## Methods

### NewGraphGraphPathOut

`func NewGraphGraphPathOut() *GraphGraphPathOut`

NewGraphGraphPathOut instantiates a new GraphGraphPathOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphPathOutWithDefaults

`func NewGraphGraphPathOutWithDefaults() *GraphGraphPathOut`

NewGraphGraphPathOutWithDefaults instantiates a new GraphGraphPathOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphPathOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphPathOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphPathOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphPathOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphPathOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphPathOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphPathOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphPathOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetBound

`func (o *GraphGraphPathOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphGraphPathOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphGraphPathOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphGraphPathOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetDecay

`func (o *GraphGraphPathOut) GetDecay() float64`

GetDecay returns the Decay field if non-nil, zero value otherwise.

### GetDecayOk

`func (o *GraphGraphPathOut) GetDecayOk() (*float64, bool)`

GetDecayOk returns a tuple with the Decay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecay

`func (o *GraphGraphPathOut) SetDecay(v float64)`

SetDecay sets Decay field to given value.

### HasDecay

`func (o *GraphGraphPathOut) HasDecay() bool`

HasDecay returns a boolean if a field has been set.

### GetEdges

`func (o *GraphGraphPathOut) GetEdges() []GraphWireFact`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GraphGraphPathOut) GetEdgesOk() (*[]GraphWireFact, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GraphGraphPathOut) SetEdges(v []GraphWireFact)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GraphGraphPathOut) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFound

`func (o *GraphGraphPathOut) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *GraphGraphPathOut) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *GraphGraphPathOut) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *GraphGraphPathOut) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetHops

`func (o *GraphGraphPathOut) GetHops() int64`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *GraphGraphPathOut) GetHopsOk() (*int64, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *GraphGraphPathOut) SetHops(v int64)`

SetHops sets Hops field to given value.

### HasHops

`func (o *GraphGraphPathOut) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphPathOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphPathOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphPathOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphPathOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetWeakest

`func (o *GraphGraphPathOut) GetWeakest() GraphWireFact`

GetWeakest returns the Weakest field if non-nil, zero value otherwise.

### GetWeakestOk

`func (o *GraphGraphPathOut) GetWeakestOk() (*GraphWireFact, bool)`

GetWeakestOk returns a tuple with the Weakest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakest

`func (o *GraphGraphPathOut) SetWeakest(v GraphWireFact)`

SetWeakest sets Weakest field to given value.

### HasWeakest

`func (o *GraphGraphPathOut) HasWeakest() bool`

HasWeakest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



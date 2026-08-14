# IngestResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **int32** | Accepted is how many samples this report landed. | [optional] 
**Links** | Pointer to [**[]LinkView**](LinkView.md) | Links is the link row each distinct (machine, provider, account) in the batch refreshed. | [optional] 
**Stored** | Pointer to **bool** | Stored reports whether history was durably written; false means the warehouse was unavailable and only the link rows were refreshed. | [optional] 

## Methods

### NewIngestResp

`func NewIngestResp() *IngestResp`

NewIngestResp instantiates a new IngestResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestRespWithDefaults

`func NewIngestRespWithDefaults() *IngestResp`

NewIngestRespWithDefaults instantiates a new IngestResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *IngestResp) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *IngestResp) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *IngestResp) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *IngestResp) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetLinks

`func (o *IngestResp) GetLinks() []LinkView`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IngestResp) GetLinksOk() (*[]LinkView, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IngestResp) SetLinks(v []LinkView)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IngestResp) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStored

`func (o *IngestResp) GetStored() bool`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *IngestResp) GetStoredOk() (*bool, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *IngestResp) SetStored(v bool)`

SetStored sets Stored field to given value.

### HasStored

`func (o *IngestResp) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



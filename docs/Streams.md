# Streams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | Pointer to [**[]Stream**](Stream.md) | Streams is the page, ordered by name. | [optional] 
**Total** | Pointer to **int32** | Total is the org&#39;s stream count before paging. | [optional] 

## Methods

### NewStreams

`func NewStreams() *Streams`

NewStreams instantiates a new Streams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsWithDefaults

`func NewStreamsWithDefaults() *Streams`

NewStreamsWithDefaults instantiates a new Streams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *Streams) GetStreams() []Stream`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *Streams) GetStreamsOk() (*[]Stream, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *Streams) SetStreams(v []Stream)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *Streams) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTotal

`func (o *Streams) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Streams) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Streams) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Streams) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



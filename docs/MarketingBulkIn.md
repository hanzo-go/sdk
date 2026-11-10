# MarketingBulkIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is tag, untag, move, unsubscribe, clean or delete. | [optional] 
**Filter** | Pointer to [**MarketingFilter**](MarketingFilter.md) | Filter selects by filter, narrowing the segment when both are set. With no ids, segment or filter the action covers the whole list. | [optional] 
**Id** | Pointer to **string** | ID is the list, by id or slug, from the path. | [optional] 
**Ids** | Pointer to **[]int64** | IDs selects subscribers by id. When set, the segment and filter are not read. | [optional] 
**Segment** | Pointer to **string** | Segment selects a saved segment of the list. | [optional] 
**Tag** | Pointer to **string** | Tag is the tag tag and untag add or remove. | [optional] 
**To** | Pointer to **string** | To is the list move moves the selection to, by id or slug. An address already on it keeps its row there and leaves this list. | [optional] 

## Methods

### NewMarketingBulkIn

`func NewMarketingBulkIn() *MarketingBulkIn`

NewMarketingBulkIn instantiates a new MarketingBulkIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBulkInWithDefaults

`func NewMarketingBulkInWithDefaults() *MarketingBulkIn`

NewMarketingBulkInWithDefaults instantiates a new MarketingBulkIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MarketingBulkIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MarketingBulkIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MarketingBulkIn) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MarketingBulkIn) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFilter

`func (o *MarketingBulkIn) GetFilter() MarketingFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MarketingBulkIn) GetFilterOk() (*MarketingFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MarketingBulkIn) SetFilter(v MarketingFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MarketingBulkIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *MarketingBulkIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingBulkIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingBulkIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingBulkIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *MarketingBulkIn) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MarketingBulkIn) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MarketingBulkIn) SetIds(v []int64)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MarketingBulkIn) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSegment

`func (o *MarketingBulkIn) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *MarketingBulkIn) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *MarketingBulkIn) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *MarketingBulkIn) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetTag

`func (o *MarketingBulkIn) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MarketingBulkIn) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MarketingBulkIn) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *MarketingBulkIn) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTo

`func (o *MarketingBulkIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MarketingBulkIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MarketingBulkIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MarketingBulkIn) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



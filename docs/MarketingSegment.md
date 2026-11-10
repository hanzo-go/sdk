# MarketingSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is how many subscribers it selects now; filled on read of one. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when it was saved, unix seconds. | [optional] 
**Filter** | Pointer to [**MarketingFilter**](MarketingFilter.md) | Filter is the selection. | [optional] 
**Id** | Pointer to **string** | ID is the segment&#39;s id (\&quot;seg_\&quot; and 96 random bits in hex). | [optional] 
**List** | Pointer to **string** | List is the list it selects from. | [optional] 
**Name** | Pointer to **string** | Name is what the segment is called; at most 200 characters. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last change, unix seconds. | [optional] 

## Methods

### NewMarketingSegment

`func NewMarketingSegment() *MarketingSegment`

NewMarketingSegment instantiates a new MarketingSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSegmentWithDefaults

`func NewMarketingSegmentWithDefaults() *MarketingSegment`

NewMarketingSegmentWithDefaults instantiates a new MarketingSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MarketingSegment) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MarketingSegment) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MarketingSegment) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MarketingSegment) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingSegment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingSegment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingSegment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFilter

`func (o *MarketingSegment) GetFilter() MarketingFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MarketingSegment) GetFilterOk() (*MarketingFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MarketingSegment) SetFilter(v MarketingFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MarketingSegment) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *MarketingSegment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingSegment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingSegment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *MarketingSegment) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingSegment) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingSegment) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingSegment) HasList() bool`

HasList returns a boolean if a field has been set.

### GetName

`func (o *MarketingSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingSegment) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingSegment) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingSegment) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingSegment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



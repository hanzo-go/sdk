# MarketingSegmentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**MarketingFilter**](MarketingFilter.md) | Filter is the selection. | [optional] 
**Id** | Pointer to **string** | ID is the list, by id or slug, from the path. | [optional] 
**Name** | Pointer to **string** | Name is what the segment is called. Required; at most 200 characters. | [optional] 
**Sid** | Pointer to **string** | SID is the segment to change, from the path; ignored on create. | [optional] 

## Methods

### NewMarketingSegmentIn

`func NewMarketingSegmentIn() *MarketingSegmentIn`

NewMarketingSegmentIn instantiates a new MarketingSegmentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSegmentInWithDefaults

`func NewMarketingSegmentInWithDefaults() *MarketingSegmentIn`

NewMarketingSegmentInWithDefaults instantiates a new MarketingSegmentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MarketingSegmentIn) GetFilter() MarketingFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MarketingSegmentIn) GetFilterOk() (*MarketingFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MarketingSegmentIn) SetFilter(v MarketingFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MarketingSegmentIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *MarketingSegmentIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingSegmentIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingSegmentIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingSegmentIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MarketingSegmentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingSegmentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingSegmentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingSegmentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSid

`func (o *MarketingSegmentIn) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MarketingSegmentIn) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MarketingSegmentIn) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MarketingSegmentIn) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



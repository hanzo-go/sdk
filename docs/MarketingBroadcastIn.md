# MarketingBroadcastIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**MarketingFilter**](MarketingFilter.md) | Filter narrows it further. | [optional] 
**FromAddress** | Pointer to **string** | FromAddress overrides the list&#39;s sender address. | [optional] 
**FromName** | Pointer to **string** | FromName overrides the list&#39;s sender name. | [optional] 
**Html** | Pointer to **string** | HTML is the message as HTML; at most 1 MiB. | [optional] 
**Id** | Pointer to **string** | ID is the broadcast to change, from the path; ignored on create. | [optional] 
**List** | Pointer to **string** | List is the list to send to, by id or slug; on create only. | [optional] 
**Provider** | Pointer to **string** | Provider is ses, sendgrid, or empty for the list&#39;s. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo overrides where replies go. | [optional] 
**ScheduledAt** | Pointer to **int64** | ScheduledAt schedules it: unix seconds. Zero leaves it a draft. | [optional] 
**Segment** | Pointer to **string** | Segment narrows it to a saved segment of the list. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject line. Required; at most 300 characters. | [optional] 
**Text** | Pointer to **string** | Text is the message as plain text; at most 256 KiB. One of html and text is required. | [optional] 

## Methods

### NewMarketingBroadcastIn

`func NewMarketingBroadcastIn() *MarketingBroadcastIn`

NewMarketingBroadcastIn instantiates a new MarketingBroadcastIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBroadcastInWithDefaults

`func NewMarketingBroadcastInWithDefaults() *MarketingBroadcastIn`

NewMarketingBroadcastInWithDefaults instantiates a new MarketingBroadcastIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MarketingBroadcastIn) GetFilter() MarketingFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MarketingBroadcastIn) GetFilterOk() (*MarketingFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MarketingBroadcastIn) SetFilter(v MarketingFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MarketingBroadcastIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFromAddress

`func (o *MarketingBroadcastIn) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MarketingBroadcastIn) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MarketingBroadcastIn) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MarketingBroadcastIn) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromName

`func (o *MarketingBroadcastIn) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MarketingBroadcastIn) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MarketingBroadcastIn) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *MarketingBroadcastIn) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetHtml

`func (o *MarketingBroadcastIn) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MarketingBroadcastIn) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MarketingBroadcastIn) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MarketingBroadcastIn) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *MarketingBroadcastIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingBroadcastIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingBroadcastIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingBroadcastIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *MarketingBroadcastIn) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingBroadcastIn) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingBroadcastIn) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingBroadcastIn) HasList() bool`

HasList returns a boolean if a field has been set.

### GetProvider

`func (o *MarketingBroadcastIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketingBroadcastIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketingBroadcastIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MarketingBroadcastIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *MarketingBroadcastIn) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MarketingBroadcastIn) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MarketingBroadcastIn) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MarketingBroadcastIn) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetScheduledAt

`func (o *MarketingBroadcastIn) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *MarketingBroadcastIn) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *MarketingBroadcastIn) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *MarketingBroadcastIn) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetSegment

`func (o *MarketingBroadcastIn) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *MarketingBroadcastIn) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *MarketingBroadcastIn) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *MarketingBroadcastIn) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetSubject

`func (o *MarketingBroadcastIn) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MarketingBroadcastIn) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MarketingBroadcastIn) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MarketingBroadcastIn) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MarketingBroadcastIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MarketingBroadcastIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MarketingBroadcastIn) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MarketingBroadcastIn) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MarketingBroadcast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **int64** | Audience is how many subscribers it selected when it started. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt and UpdatedAt are unix seconds. | [optional] 
**Error** | Pointer to **string** | Error is why a failed broadcast stopped. | [optional] 
**Filter** | Pointer to [**MarketingFilter**](MarketingFilter.md) | Filter narrows it further, as a subscriber filter does. Whatever it says, only subscribers whose status is subscribed are sent to. | [optional] 
**FinishedAt** | Pointer to **int64** | FinishedAt is when it ended. | [optional] 
**FromAddress** | Pointer to **string** | FromAddress is the sender&#39;s address. | [optional] 
**FromName** | Pointer to **string** | FromName, FromAddress and ReplyTo override the list&#39;s sender. | [optional] 
**Html** | Pointer to **string** | HTML is the message as HTML. | [optional] 
**Id** | Pointer to **string** | ID is the broadcast&#39;s id (\&quot;brd_\&quot; and 96 random bits in hex). | [optional] 
**List** | Pointer to **string** | List is the list it goes to. | [optional] 
**Provider** | Pointer to **string** | Provider is ses or sendgrid; empty takes the list&#39;s. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is where replies go. | [optional] 
**ScheduledAt** | Pointer to **int64** | ScheduledAt is when it is to start, unix seconds. | [optional] 
**Segment** | Pointer to **string** | Segment narrows it to a saved segment of the list; empty is the list. | [optional] 
**StartedAt** | Pointer to **int64** | StartedAt and FinishedAt are when sending began and ended. | [optional] 
**Status** | Pointer to **string** | Status is draft, scheduled, sending, sent, cancelled or failed. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject line. | [optional] 
**Text** | Pointer to **string** | Text is the message as plain text. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last change. | [optional] 

## Methods

### NewMarketingBroadcast

`func NewMarketingBroadcast() *MarketingBroadcast`

NewMarketingBroadcast instantiates a new MarketingBroadcast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBroadcastWithDefaults

`func NewMarketingBroadcastWithDefaults() *MarketingBroadcast`

NewMarketingBroadcastWithDefaults instantiates a new MarketingBroadcast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *MarketingBroadcast) GetAudience() int64`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *MarketingBroadcast) GetAudienceOk() (*int64, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *MarketingBroadcast) SetAudience(v int64)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *MarketingBroadcast) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingBroadcast) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingBroadcast) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingBroadcast) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingBroadcast) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *MarketingBroadcast) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketingBroadcast) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketingBroadcast) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketingBroadcast) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFilter

`func (o *MarketingBroadcast) GetFilter() MarketingFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MarketingBroadcast) GetFilterOk() (*MarketingFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MarketingBroadcast) SetFilter(v MarketingFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MarketingBroadcast) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFinishedAt

`func (o *MarketingBroadcast) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *MarketingBroadcast) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *MarketingBroadcast) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *MarketingBroadcast) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetFromAddress

`func (o *MarketingBroadcast) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MarketingBroadcast) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MarketingBroadcast) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MarketingBroadcast) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromName

`func (o *MarketingBroadcast) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MarketingBroadcast) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MarketingBroadcast) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *MarketingBroadcast) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetHtml

`func (o *MarketingBroadcast) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MarketingBroadcast) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MarketingBroadcast) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MarketingBroadcast) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *MarketingBroadcast) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingBroadcast) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingBroadcast) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingBroadcast) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *MarketingBroadcast) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingBroadcast) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingBroadcast) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingBroadcast) HasList() bool`

HasList returns a boolean if a field has been set.

### GetProvider

`func (o *MarketingBroadcast) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketingBroadcast) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketingBroadcast) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MarketingBroadcast) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *MarketingBroadcast) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MarketingBroadcast) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MarketingBroadcast) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MarketingBroadcast) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetScheduledAt

`func (o *MarketingBroadcast) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *MarketingBroadcast) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *MarketingBroadcast) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *MarketingBroadcast) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetSegment

`func (o *MarketingBroadcast) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *MarketingBroadcast) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *MarketingBroadcast) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *MarketingBroadcast) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetStartedAt

`func (o *MarketingBroadcast) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MarketingBroadcast) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MarketingBroadcast) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MarketingBroadcast) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingBroadcast) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingBroadcast) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingBroadcast) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingBroadcast) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *MarketingBroadcast) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MarketingBroadcast) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MarketingBroadcast) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MarketingBroadcast) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MarketingBroadcast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MarketingBroadcast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MarketingBroadcast) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MarketingBroadcast) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingBroadcast) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingBroadcast) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingBroadcast) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingBroadcast) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



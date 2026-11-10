# EmailMessageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounce** | Pointer to **string** | Bounce is hard or soft, for a message that bounced. | [optional] 
**BouncedAt** | Pointer to **time.Time** | BouncedAt is when it bounced, and Bounce whether hard or soft. | [optional] 
**Campaign** | Pointer to **string** | Campaign is the campaign the message was tagged with. | [optional] 
**ClickedAt** | Pointer to **time.Time** | ClickedAt is when a link in it was first clicked. | [optional] 
**ComplainedAt** | Pointer to **time.Time** | ComplainedAt is when its recipient reported it as spam. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the message was sent, or accepted to be. | [optional] 
**DeliveredAt** | Pointer to **time.Time** | DeliveredAt is when the receiving server first took it. | [optional] 
**Error** | Pointer to **string** | Error says why it failed, bounced, was deferred or was rejected. | [optional] 
**From** | Pointer to **string** | From is the sender address, when known. | [optional] 
**Id** | Pointer to **string** | ID is our id for the message (msg_…): every link in it carries it as from&#x3D;, and the event stream names it by it. | [optional] 
**IdempotencyKey** | Pointer to **string** | IdempotencyKey is the key the message was sent under. | [optional] 
**OpenedAt** | Pointer to **time.Time** | OpenedAt is when it was first opened. | [optional] 
**Org** | Pointer to **string** | Org is the organization owning this message. | [optional] 
**Origin** | Pointer to **string** | Origin is api for a message POST /v1/email sent, provider for one the provider&#39;s events told of, sent from the org&#39;s own tools. | [optional] 
**Provider** | Pointer to **string** | Provider is the rail it went by: ses or sendgrid. | [optional] 
**ProviderId** | Pointer to **string** | ProviderID is the provider&#39;s own id for the message, which its events name it by; empty until the provider answers one. | [optional] 
**Status** | Pointer to **string** | Status is queued, sent, failed, delivered, opened, clicked, bounced or complained. | [optional] 
**Subject** | Pointer to **string** | Subject is the message subject line, when known. | [optional] 
**To** | Pointer to **[]string** | To lists the recipients. | [optional] 
**UnsubscribedAt** | Pointer to **time.Time** | UnsubscribedAt is when its recipient unsubscribed from it. | [optional] 

## Methods

### NewEmailMessageRecord

`func NewEmailMessageRecord() *EmailMessageRecord`

NewEmailMessageRecord instantiates a new EmailMessageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageRecordWithDefaults

`func NewEmailMessageRecordWithDefaults() *EmailMessageRecord`

NewEmailMessageRecordWithDefaults instantiates a new EmailMessageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounce

`func (o *EmailMessageRecord) GetBounce() string`

GetBounce returns the Bounce field if non-nil, zero value otherwise.

### GetBounceOk

`func (o *EmailMessageRecord) GetBounceOk() (*string, bool)`

GetBounceOk returns a tuple with the Bounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounce

`func (o *EmailMessageRecord) SetBounce(v string)`

SetBounce sets Bounce field to given value.

### HasBounce

`func (o *EmailMessageRecord) HasBounce() bool`

HasBounce returns a boolean if a field has been set.

### GetBouncedAt

`func (o *EmailMessageRecord) GetBouncedAt() time.Time`

GetBouncedAt returns the BouncedAt field if non-nil, zero value otherwise.

### GetBouncedAtOk

`func (o *EmailMessageRecord) GetBouncedAtOk() (*time.Time, bool)`

GetBouncedAtOk returns a tuple with the BouncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouncedAt

`func (o *EmailMessageRecord) SetBouncedAt(v time.Time)`

SetBouncedAt sets BouncedAt field to given value.

### HasBouncedAt

`func (o *EmailMessageRecord) HasBouncedAt() bool`

HasBouncedAt returns a boolean if a field has been set.

### GetCampaign

`func (o *EmailMessageRecord) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *EmailMessageRecord) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *EmailMessageRecord) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *EmailMessageRecord) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetClickedAt

`func (o *EmailMessageRecord) GetClickedAt() time.Time`

GetClickedAt returns the ClickedAt field if non-nil, zero value otherwise.

### GetClickedAtOk

`func (o *EmailMessageRecord) GetClickedAtOk() (*time.Time, bool)`

GetClickedAtOk returns a tuple with the ClickedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedAt

`func (o *EmailMessageRecord) SetClickedAt(v time.Time)`

SetClickedAt sets ClickedAt field to given value.

### HasClickedAt

`func (o *EmailMessageRecord) HasClickedAt() bool`

HasClickedAt returns a boolean if a field has been set.

### GetComplainedAt

`func (o *EmailMessageRecord) GetComplainedAt() time.Time`

GetComplainedAt returns the ComplainedAt field if non-nil, zero value otherwise.

### GetComplainedAtOk

`func (o *EmailMessageRecord) GetComplainedAtOk() (*time.Time, bool)`

GetComplainedAtOk returns a tuple with the ComplainedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplainedAt

`func (o *EmailMessageRecord) SetComplainedAt(v time.Time)`

SetComplainedAt sets ComplainedAt field to given value.

### HasComplainedAt

`func (o *EmailMessageRecord) HasComplainedAt() bool`

HasComplainedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailMessageRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailMessageRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailMessageRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailMessageRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *EmailMessageRecord) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *EmailMessageRecord) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *EmailMessageRecord) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *EmailMessageRecord) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetError

`func (o *EmailMessageRecord) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EmailMessageRecord) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EmailMessageRecord) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EmailMessageRecord) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFrom

`func (o *EmailMessageRecord) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessageRecord) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessageRecord) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailMessageRecord) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *EmailMessageRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailMessageRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailMessageRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailMessageRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *EmailMessageRecord) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *EmailMessageRecord) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *EmailMessageRecord) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *EmailMessageRecord) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetOpenedAt

`func (o *EmailMessageRecord) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *EmailMessageRecord) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *EmailMessageRecord) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.

### HasOpenedAt

`func (o *EmailMessageRecord) HasOpenedAt() bool`

HasOpenedAt returns a boolean if a field has been set.

### GetOrg

`func (o *EmailMessageRecord) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *EmailMessageRecord) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *EmailMessageRecord) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *EmailMessageRecord) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOrigin

`func (o *EmailMessageRecord) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EmailMessageRecord) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EmailMessageRecord) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *EmailMessageRecord) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetProvider

`func (o *EmailMessageRecord) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmailMessageRecord) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmailMessageRecord) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmailMessageRecord) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderId

`func (o *EmailMessageRecord) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *EmailMessageRecord) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *EmailMessageRecord) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *EmailMessageRecord) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetStatus

`func (o *EmailMessageRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailMessageRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailMessageRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailMessageRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *EmailMessageRecord) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailMessageRecord) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailMessageRecord) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailMessageRecord) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessageRecord) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessageRecord) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessageRecord) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailMessageRecord) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUnsubscribedAt

`func (o *EmailMessageRecord) GetUnsubscribedAt() time.Time`

GetUnsubscribedAt returns the UnsubscribedAt field if non-nil, zero value otherwise.

### GetUnsubscribedAtOk

`func (o *EmailMessageRecord) GetUnsubscribedAtOk() (*time.Time, bool)`

GetUnsubscribedAtOk returns a tuple with the UnsubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribedAt

`func (o *EmailMessageRecord) SetUnsubscribedAt(v time.Time)`

SetUnsubscribedAt sets UnsubscribedAt field to given value.

### HasUnsubscribedAt

`func (o *EmailMessageRecord) HasUnsubscribedAt() bool`

HasUnsubscribedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



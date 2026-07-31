# NotifySendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** | Destination address(es). Phone number for SMS, email for email. Multiple recipients fan out into one provider call each and produce one response entry per recipient.  | 
**Channel** | Pointer to [**NotifyChannel**](NotifyChannel.md) |  | [optional] 
**Provider** | Pointer to **string** | Pins a specific provider service name (e.g. &#x60;twilio&#x60;, &#x60;plivo&#x60;, &#x60;twilio_email&#x60;, &#x60;mail&#x60;). When omitted the fold picks the org&#39;s default for the channel from configured KMS credentials.  | [optional] 
**Subject** | Pointer to **string** | Raw subject; used verbatim when &#x60;body&#x60; is set (no-template path). | [optional] 
**Body** | Pointer to **string** | Raw message body. When non-empty it wins verbatim and no template is rendered.  | [optional] 
**TemplateId** | Pointer to **string** | Built-in template identifier. When &#x60;body&#x60; is empty, the template is resolved from &#x60;template_id&#x60; or, failing that, &#x60;event&#x60;.  | [optional] 
**TemplateVars** | Pointer to **map[string]interface{}** | Free-form variables the template is rendered against (e.g. &#x60;otp&#x60;, &#x60;app&#x60;, &#x60;recipient&#x60;). A missing &#x60;app&#x60; defaults to &#x60;Hanzo&#x60;.  | [optional] 
**Event** | Pointer to **string** | Event-catalog identifier. Also used as a fallback template id when &#x60;template_id&#x60; is unset (the IAM OTP path sends &#x60;event: iam.otp_sent&#x60; with no &#x60;template_id&#x60;).  | [optional] 
**IdempotencyKey** | Pointer to **string** | Per-tenant deduplication key. Accepted on the wire; the sync fold does not persist a Message row.  | [optional] 
**SendAt** | Pointer to **string** | RFC3339 schedule time for async dispatch. Async is not available in the fold, so a set value has no effect on the sync path.  | [optional] 
**Options** | Pointer to **map[string]interface{}** | Free-form per-provider knob bag. | [optional] 

## Methods

### NewNotifySendRequest

`func NewNotifySendRequest(to []string, ) *NotifySendRequest`

NewNotifySendRequest instantiates a new NotifySendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifySendRequestWithDefaults

`func NewNotifySendRequestWithDefaults() *NotifySendRequest`

NewNotifySendRequestWithDefaults instantiates a new NotifySendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *NotifySendRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NotifySendRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NotifySendRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetChannel

`func (o *NotifySendRequest) GetChannel() NotifyChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *NotifySendRequest) GetChannelOk() (*NotifyChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *NotifySendRequest) SetChannel(v NotifyChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *NotifySendRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetProvider

`func (o *NotifySendRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotifySendRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotifySendRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotifySendRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSubject

`func (o *NotifySendRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotifySendRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotifySendRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotifySendRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *NotifySendRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotifySendRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotifySendRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotifySendRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTemplateId

`func (o *NotifySendRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *NotifySendRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *NotifySendRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *NotifySendRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVars

`func (o *NotifySendRequest) GetTemplateVars() map[string]interface{}`

GetTemplateVars returns the TemplateVars field if non-nil, zero value otherwise.

### GetTemplateVarsOk

`func (o *NotifySendRequest) GetTemplateVarsOk() (*map[string]interface{}, bool)`

GetTemplateVarsOk returns a tuple with the TemplateVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVars

`func (o *NotifySendRequest) SetTemplateVars(v map[string]interface{})`

SetTemplateVars sets TemplateVars field to given value.

### HasTemplateVars

`func (o *NotifySendRequest) HasTemplateVars() bool`

HasTemplateVars returns a boolean if a field has been set.

### GetEvent

`func (o *NotifySendRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotifySendRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotifySendRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *NotifySendRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *NotifySendRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *NotifySendRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *NotifySendRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *NotifySendRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetSendAt

`func (o *NotifySendRequest) GetSendAt() string`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *NotifySendRequest) GetSendAtOk() (*string, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *NotifySendRequest) SetSendAt(v string)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *NotifySendRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetOptions

`func (o *NotifySendRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NotifySendRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NotifySendRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *NotifySendRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



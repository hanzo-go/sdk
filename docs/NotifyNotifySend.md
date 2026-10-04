# NotifyNotifySend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the message text, sent verbatim when present — the no-template path. | [optional] 
**Channel** | Pointer to **string** | Channel selects the delivery channel, sms or email. The per-channel routes (/send/sms, /send/email) pin it, overriding whatever the body names; on the generic route it is required. | [optional] 
**Event** | Pointer to **string** | Event is the event name, which doubles as the template id when TemplateID is empty — the IAM OTP path sends event&#x3D;iam.otp_sent and nothing else. | [optional] 
**Provider** | Pointer to **string** | Provider pins a provider service name (twilio, plivo, twilio_email, mail). Empty picks the one whose org credentials are actually configured in KMS. | [optional] 
**Subject** | Pointer to **string** | Subject is the message subject, carried on the email channel only. | [optional] 
**Sync** | Pointer to **string** | Sync must be exactly \&quot;true\&quot;: delivery here is synchronous, and anything else answers 503 because the queue plane that would run an async dispatch is owned elsewhere. Over REST it rides as ?sync&#x3D;true (the URL binds over the body); a by-name call states it in its arguments. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID selects a built-in template when Body is empty. | [optional] 
**TemplateVars** | Pointer to **interface{}** |  | [optional] 
**To** | Pointer to **[]string** | To is the destination address per recipient — a phone number for sms, an email address for email. Several recipients fan out into one provider call each, and the response shape follows the count (see the items field). | [optional] 

## Methods

### NewNotifyNotifySend

`func NewNotifyNotifySend() *NotifyNotifySend`

NewNotifyNotifySend instantiates a new NotifyNotifySend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNotifySendWithDefaults

`func NewNotifyNotifySendWithDefaults() *NotifyNotifySend`

NewNotifyNotifySendWithDefaults instantiates a new NotifyNotifySend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *NotifyNotifySend) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotifyNotifySend) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotifyNotifySend) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotifyNotifySend) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetChannel

`func (o *NotifyNotifySend) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *NotifyNotifySend) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *NotifyNotifySend) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *NotifyNotifySend) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEvent

`func (o *NotifyNotifySend) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotifyNotifySend) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotifyNotifySend) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *NotifyNotifySend) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetProvider

`func (o *NotifyNotifySend) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotifyNotifySend) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotifyNotifySend) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotifyNotifySend) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSubject

`func (o *NotifyNotifySend) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotifyNotifySend) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotifyNotifySend) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotifyNotifySend) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSync

`func (o *NotifyNotifySend) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *NotifyNotifySend) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *NotifyNotifySend) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *NotifyNotifySend) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetTemplateId

`func (o *NotifyNotifySend) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *NotifyNotifySend) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *NotifyNotifySend) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *NotifyNotifySend) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVars

`func (o *NotifyNotifySend) GetTemplateVars() interface{}`

GetTemplateVars returns the TemplateVars field if non-nil, zero value otherwise.

### GetTemplateVarsOk

`func (o *NotifyNotifySend) GetTemplateVarsOk() (*interface{}, bool)`

GetTemplateVarsOk returns a tuple with the TemplateVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVars

`func (o *NotifyNotifySend) SetTemplateVars(v interface{})`

SetTemplateVars sets TemplateVars field to given value.

### HasTemplateVars

`func (o *NotifyNotifySend) HasTemplateVars() bool`

HasTemplateVars returns a boolean if a field has been set.

### SetTemplateVarsNil

`func (o *NotifyNotifySend) SetTemplateVarsNil(b bool)`

 SetTemplateVarsNil sets the value for TemplateVars to be an explicit nil

### UnsetTemplateVars
`func (o *NotifyNotifySend) UnsetTemplateVars()`

UnsetTemplateVars ensures that no value is present for TemplateVars, not even an explicit nil
### GetTo

`func (o *NotifyNotifySend) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NotifyNotifySend) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NotifyNotifySend) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *NotifyNotifySend) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



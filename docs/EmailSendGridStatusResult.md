# EmailSendGridStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProvider** | Pointer to **string** | ActiveProvider is the provider the org&#39;s mail leaves by when a request pins none. | [optional] 
**AsmGroupId** | Pointer to **string** | ASMGroupID is the optional SendGrid suppression group ID. | [optional] 
**Configured** | Pointer to **bool** | Configured reports whether the org holds the SendGrid credentials a send needs (api-key, from-email). | [optional] 
**CreditsRemaining** | Pointer to **int64** | CreditsRemaining is the credits left in the period. | [optional] 
**CreditsTotal** | Pointer to **int64** | CreditsTotal, CreditsRemaining and CreditsUsed are the account&#39;s sending credits for the period, as SendGrid answers them. | [optional] 
**CreditsUsed** | Pointer to **int64** | CreditsUsed is the credits used in the period. | [optional] 
**DnsRecords** | Pointer to [**[]EmailSESDNSRecord**](EmailSESDNSRecord.md) | DNSRecords lists the required SendGrid Domain Authentication CNAME, Link Branding, SPF, and DMARC DNS records. | [optional] 
**Domain** | Pointer to **string** | Domain is the sending domain used for SendGrid Domain Authentication and DMARC. | [optional] 
**FromEmail** | Pointer to **string** | FromEmail is the configured SendGrid outreach sender email address. | [optional] 
**FromName** | Pointer to **string** | FromName is the configured SendGrid outreach sender display name. | [optional] 
**IpPool** | Pointer to **string** | IPPool is the SendGrid IP pool the credential names, if any. | [optional] 
**Org** | Pointer to **string** | Org is the organization identifier. | [optional] 
**QuotaError** | Pointer to **string** | QuotaError says why the account&#39;s credits could not be read; they are then 0. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is the configured Reply-To address for outreach replies. | [optional] 
**SentLast24Hours** | Pointer to **int64** | SentLast24Hours is what the org&#39;s ledger counts sent by SendGrid in the last 24 hours: through POST /v1/email, and every message SendGrid&#39;s events told of. | [optional] 
**SeparateFromTwilioNotify** | Pointer to **bool** | SeparateFromTwilioNotify confirms that SendGrid outreach uses its own dedicated API key, sender identity, and IP pool distinct from standard Twilio transactional notifications. | [optional] 

## Methods

### NewEmailSendGridStatusResult

`func NewEmailSendGridStatusResult() *EmailSendGridStatusResult`

NewEmailSendGridStatusResult instantiates a new EmailSendGridStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendGridStatusResultWithDefaults

`func NewEmailSendGridStatusResultWithDefaults() *EmailSendGridStatusResult`

NewEmailSendGridStatusResultWithDefaults instantiates a new EmailSendGridStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveProvider

`func (o *EmailSendGridStatusResult) GetActiveProvider() string`

GetActiveProvider returns the ActiveProvider field if non-nil, zero value otherwise.

### GetActiveProviderOk

`func (o *EmailSendGridStatusResult) GetActiveProviderOk() (*string, bool)`

GetActiveProviderOk returns a tuple with the ActiveProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProvider

`func (o *EmailSendGridStatusResult) SetActiveProvider(v string)`

SetActiveProvider sets ActiveProvider field to given value.

### HasActiveProvider

`func (o *EmailSendGridStatusResult) HasActiveProvider() bool`

HasActiveProvider returns a boolean if a field has been set.

### GetAsmGroupId

`func (o *EmailSendGridStatusResult) GetAsmGroupId() string`

GetAsmGroupId returns the AsmGroupId field if non-nil, zero value otherwise.

### GetAsmGroupIdOk

`func (o *EmailSendGridStatusResult) GetAsmGroupIdOk() (*string, bool)`

GetAsmGroupIdOk returns a tuple with the AsmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmGroupId

`func (o *EmailSendGridStatusResult) SetAsmGroupId(v string)`

SetAsmGroupId sets AsmGroupId field to given value.

### HasAsmGroupId

`func (o *EmailSendGridStatusResult) HasAsmGroupId() bool`

HasAsmGroupId returns a boolean if a field has been set.

### GetConfigured

`func (o *EmailSendGridStatusResult) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *EmailSendGridStatusResult) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *EmailSendGridStatusResult) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *EmailSendGridStatusResult) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetCreditsRemaining

`func (o *EmailSendGridStatusResult) GetCreditsRemaining() int64`

GetCreditsRemaining returns the CreditsRemaining field if non-nil, zero value otherwise.

### GetCreditsRemainingOk

`func (o *EmailSendGridStatusResult) GetCreditsRemainingOk() (*int64, bool)`

GetCreditsRemainingOk returns a tuple with the CreditsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsRemaining

`func (o *EmailSendGridStatusResult) SetCreditsRemaining(v int64)`

SetCreditsRemaining sets CreditsRemaining field to given value.

### HasCreditsRemaining

`func (o *EmailSendGridStatusResult) HasCreditsRemaining() bool`

HasCreditsRemaining returns a boolean if a field has been set.

### GetCreditsTotal

`func (o *EmailSendGridStatusResult) GetCreditsTotal() int64`

GetCreditsTotal returns the CreditsTotal field if non-nil, zero value otherwise.

### GetCreditsTotalOk

`func (o *EmailSendGridStatusResult) GetCreditsTotalOk() (*int64, bool)`

GetCreditsTotalOk returns a tuple with the CreditsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsTotal

`func (o *EmailSendGridStatusResult) SetCreditsTotal(v int64)`

SetCreditsTotal sets CreditsTotal field to given value.

### HasCreditsTotal

`func (o *EmailSendGridStatusResult) HasCreditsTotal() bool`

HasCreditsTotal returns a boolean if a field has been set.

### GetCreditsUsed

`func (o *EmailSendGridStatusResult) GetCreditsUsed() int64`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *EmailSendGridStatusResult) GetCreditsUsedOk() (*int64, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *EmailSendGridStatusResult) SetCreditsUsed(v int64)`

SetCreditsUsed sets CreditsUsed field to given value.

### HasCreditsUsed

`func (o *EmailSendGridStatusResult) HasCreditsUsed() bool`

HasCreditsUsed returns a boolean if a field has been set.

### GetDnsRecords

`func (o *EmailSendGridStatusResult) GetDnsRecords() []EmailSESDNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *EmailSendGridStatusResult) GetDnsRecordsOk() (*[]EmailSESDNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *EmailSendGridStatusResult) SetDnsRecords(v []EmailSESDNSRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *EmailSendGridStatusResult) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetDomain

`func (o *EmailSendGridStatusResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailSendGridStatusResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailSendGridStatusResult) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailSendGridStatusResult) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFromEmail

`func (o *EmailSendGridStatusResult) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *EmailSendGridStatusResult) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *EmailSendGridStatusResult) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *EmailSendGridStatusResult) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetFromName

`func (o *EmailSendGridStatusResult) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailSendGridStatusResult) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailSendGridStatusResult) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailSendGridStatusResult) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetIpPool

`func (o *EmailSendGridStatusResult) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *EmailSendGridStatusResult) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *EmailSendGridStatusResult) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *EmailSendGridStatusResult) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetOrg

`func (o *EmailSendGridStatusResult) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *EmailSendGridStatusResult) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *EmailSendGridStatusResult) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *EmailSendGridStatusResult) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetQuotaError

`func (o *EmailSendGridStatusResult) GetQuotaError() string`

GetQuotaError returns the QuotaError field if non-nil, zero value otherwise.

### GetQuotaErrorOk

`func (o *EmailSendGridStatusResult) GetQuotaErrorOk() (*string, bool)`

GetQuotaErrorOk returns a tuple with the QuotaError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaError

`func (o *EmailSendGridStatusResult) SetQuotaError(v string)`

SetQuotaError sets QuotaError field to given value.

### HasQuotaError

`func (o *EmailSendGridStatusResult) HasQuotaError() bool`

HasQuotaError returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailSendGridStatusResult) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailSendGridStatusResult) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailSendGridStatusResult) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailSendGridStatusResult) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSentLast24Hours

`func (o *EmailSendGridStatusResult) GetSentLast24Hours() int64`

GetSentLast24Hours returns the SentLast24Hours field if non-nil, zero value otherwise.

### GetSentLast24HoursOk

`func (o *EmailSendGridStatusResult) GetSentLast24HoursOk() (*int64, bool)`

GetSentLast24HoursOk returns a tuple with the SentLast24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentLast24Hours

`func (o *EmailSendGridStatusResult) SetSentLast24Hours(v int64)`

SetSentLast24Hours sets SentLast24Hours field to given value.

### HasSentLast24Hours

`func (o *EmailSendGridStatusResult) HasSentLast24Hours() bool`

HasSentLast24Hours returns a boolean if a field has been set.

### GetSeparateFromTwilioNotify

`func (o *EmailSendGridStatusResult) GetSeparateFromTwilioNotify() bool`

GetSeparateFromTwilioNotify returns the SeparateFromTwilioNotify field if non-nil, zero value otherwise.

### GetSeparateFromTwilioNotifyOk

`func (o *EmailSendGridStatusResult) GetSeparateFromTwilioNotifyOk() (*bool, bool)`

GetSeparateFromTwilioNotifyOk returns a tuple with the SeparateFromTwilioNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateFromTwilioNotify

`func (o *EmailSendGridStatusResult) SetSeparateFromTwilioNotify(v bool)`

SetSeparateFromTwilioNotify sets SeparateFromTwilioNotify field to given value.

### HasSeparateFromTwilioNotify

`func (o *EmailSendGridStatusResult) HasSeparateFromTwilioNotify() bool`

HasSeparateFromTwilioNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



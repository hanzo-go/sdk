# EmailSESStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProvider** | Pointer to **string** | ActiveProvider is the provider the org&#39;s mail leaves by when a request pins none. | [optional] 
**ConfigurationSet** | Pointer to **string** | ConfigurationSet is the SES configuration set messages go out under, whose event destination publishes their events (hanzo-crm-events unless the credential names another). | [optional] 
**Configured** | Pointer to **bool** | Configured reports whether the org holds the SES credentials a send needs (access-key-id, secret-access-key, from-email). | [optional] 
**DnsRecords** | Pointer to [**[]EmailSESDNSRecord**](EmailSESDNSRecord.md) | DNSRecords lists the Easy DKIM, SPF, MX, and DMARC records for the sending domain. | [optional] 
**Domain** | Pointer to **string** | Domain is the sending domain inspected. | [optional] 
**FromEmail** | Pointer to **string** | FromEmail is the default verified sender address. | [optional] 
**FromName** | Pointer to **string** | FromName is the default sender display name. | [optional] 
**MailFromDomain** | Pointer to **string** | MailFromDomain is the custom MAIL FROM subdomain (bounce.&lt;domain&gt;). | [optional] 
**Max24HourSend** | Pointer to **float64** | Max24HourSend is the account&#39;s 24-hour sending quota, as SES answers it. | [optional] 
**MaxSendRate** | Pointer to **float64** | MaxSendRate is the account&#39;s most messages a second, as SES answers it. | [optional] 
**Org** | Pointer to **string** | Org is the organization identifier. | [optional] 
**ProviderSentLast24Hours** | Pointer to **float64** | ProviderSentLast24Hours is what SES counts sent by the account in the last 24 hours, from anywhere. | [optional] 
**QuotaError** | Pointer to **string** | QuotaError says why the account&#39;s limits could not be read; the limits are then 0. | [optional] 
**Region** | Pointer to **string** | Region is the configured Amazon SES AWS region (defaults to us-east-1). | [optional] 
**Sandbox** | Pointer to **bool** | Sandbox is whether SES says the account is still in its sandbox. | [optional] 
**SendingEnabled** | Pointer to **bool** | SendingEnabled is whether SES says the account may send. | [optional] 
**SentLast24Hours** | Pointer to **int64** | SentLast24Hours is what the org&#39;s ledger counts sent by SES in the last 24 hours: through POST /v1/email, and every message SES&#39;s events told of. | [optional] 

## Methods

### NewEmailSESStatusResult

`func NewEmailSESStatusResult() *EmailSESStatusResult`

NewEmailSESStatusResult instantiates a new EmailSESStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSESStatusResultWithDefaults

`func NewEmailSESStatusResultWithDefaults() *EmailSESStatusResult`

NewEmailSESStatusResultWithDefaults instantiates a new EmailSESStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveProvider

`func (o *EmailSESStatusResult) GetActiveProvider() string`

GetActiveProvider returns the ActiveProvider field if non-nil, zero value otherwise.

### GetActiveProviderOk

`func (o *EmailSESStatusResult) GetActiveProviderOk() (*string, bool)`

GetActiveProviderOk returns a tuple with the ActiveProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProvider

`func (o *EmailSESStatusResult) SetActiveProvider(v string)`

SetActiveProvider sets ActiveProvider field to given value.

### HasActiveProvider

`func (o *EmailSESStatusResult) HasActiveProvider() bool`

HasActiveProvider returns a boolean if a field has been set.

### GetConfigurationSet

`func (o *EmailSESStatusResult) GetConfigurationSet() string`

GetConfigurationSet returns the ConfigurationSet field if non-nil, zero value otherwise.

### GetConfigurationSetOk

`func (o *EmailSESStatusResult) GetConfigurationSetOk() (*string, bool)`

GetConfigurationSetOk returns a tuple with the ConfigurationSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSet

`func (o *EmailSESStatusResult) SetConfigurationSet(v string)`

SetConfigurationSet sets ConfigurationSet field to given value.

### HasConfigurationSet

`func (o *EmailSESStatusResult) HasConfigurationSet() bool`

HasConfigurationSet returns a boolean if a field has been set.

### GetConfigured

`func (o *EmailSESStatusResult) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *EmailSESStatusResult) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *EmailSESStatusResult) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *EmailSESStatusResult) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDnsRecords

`func (o *EmailSESStatusResult) GetDnsRecords() []EmailSESDNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *EmailSESStatusResult) GetDnsRecordsOk() (*[]EmailSESDNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *EmailSESStatusResult) SetDnsRecords(v []EmailSESDNSRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *EmailSESStatusResult) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetDomain

`func (o *EmailSESStatusResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailSESStatusResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailSESStatusResult) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailSESStatusResult) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFromEmail

`func (o *EmailSESStatusResult) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *EmailSESStatusResult) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *EmailSESStatusResult) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *EmailSESStatusResult) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetFromName

`func (o *EmailSESStatusResult) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailSESStatusResult) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailSESStatusResult) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailSESStatusResult) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetMailFromDomain

`func (o *EmailSESStatusResult) GetMailFromDomain() string`

GetMailFromDomain returns the MailFromDomain field if non-nil, zero value otherwise.

### GetMailFromDomainOk

`func (o *EmailSESStatusResult) GetMailFromDomainOk() (*string, bool)`

GetMailFromDomainOk returns a tuple with the MailFromDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFromDomain

`func (o *EmailSESStatusResult) SetMailFromDomain(v string)`

SetMailFromDomain sets MailFromDomain field to given value.

### HasMailFromDomain

`func (o *EmailSESStatusResult) HasMailFromDomain() bool`

HasMailFromDomain returns a boolean if a field has been set.

### GetMax24HourSend

`func (o *EmailSESStatusResult) GetMax24HourSend() float64`

GetMax24HourSend returns the Max24HourSend field if non-nil, zero value otherwise.

### GetMax24HourSendOk

`func (o *EmailSESStatusResult) GetMax24HourSendOk() (*float64, bool)`

GetMax24HourSendOk returns a tuple with the Max24HourSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax24HourSend

`func (o *EmailSESStatusResult) SetMax24HourSend(v float64)`

SetMax24HourSend sets Max24HourSend field to given value.

### HasMax24HourSend

`func (o *EmailSESStatusResult) HasMax24HourSend() bool`

HasMax24HourSend returns a boolean if a field has been set.

### GetMaxSendRate

`func (o *EmailSESStatusResult) GetMaxSendRate() float64`

GetMaxSendRate returns the MaxSendRate field if non-nil, zero value otherwise.

### GetMaxSendRateOk

`func (o *EmailSESStatusResult) GetMaxSendRateOk() (*float64, bool)`

GetMaxSendRateOk returns a tuple with the MaxSendRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSendRate

`func (o *EmailSESStatusResult) SetMaxSendRate(v float64)`

SetMaxSendRate sets MaxSendRate field to given value.

### HasMaxSendRate

`func (o *EmailSESStatusResult) HasMaxSendRate() bool`

HasMaxSendRate returns a boolean if a field has been set.

### GetOrg

`func (o *EmailSESStatusResult) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *EmailSESStatusResult) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *EmailSESStatusResult) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *EmailSESStatusResult) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProviderSentLast24Hours

`func (o *EmailSESStatusResult) GetProviderSentLast24Hours() float64`

GetProviderSentLast24Hours returns the ProviderSentLast24Hours field if non-nil, zero value otherwise.

### GetProviderSentLast24HoursOk

`func (o *EmailSESStatusResult) GetProviderSentLast24HoursOk() (*float64, bool)`

GetProviderSentLast24HoursOk returns a tuple with the ProviderSentLast24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSentLast24Hours

`func (o *EmailSESStatusResult) SetProviderSentLast24Hours(v float64)`

SetProviderSentLast24Hours sets ProviderSentLast24Hours field to given value.

### HasProviderSentLast24Hours

`func (o *EmailSESStatusResult) HasProviderSentLast24Hours() bool`

HasProviderSentLast24Hours returns a boolean if a field has been set.

### GetQuotaError

`func (o *EmailSESStatusResult) GetQuotaError() string`

GetQuotaError returns the QuotaError field if non-nil, zero value otherwise.

### GetQuotaErrorOk

`func (o *EmailSESStatusResult) GetQuotaErrorOk() (*string, bool)`

GetQuotaErrorOk returns a tuple with the QuotaError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaError

`func (o *EmailSESStatusResult) SetQuotaError(v string)`

SetQuotaError sets QuotaError field to given value.

### HasQuotaError

`func (o *EmailSESStatusResult) HasQuotaError() bool`

HasQuotaError returns a boolean if a field has been set.

### GetRegion

`func (o *EmailSESStatusResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EmailSESStatusResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EmailSESStatusResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EmailSESStatusResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSandbox

`func (o *EmailSESStatusResult) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *EmailSESStatusResult) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *EmailSESStatusResult) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *EmailSESStatusResult) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetSendingEnabled

`func (o *EmailSESStatusResult) GetSendingEnabled() bool`

GetSendingEnabled returns the SendingEnabled field if non-nil, zero value otherwise.

### GetSendingEnabledOk

`func (o *EmailSESStatusResult) GetSendingEnabledOk() (*bool, bool)`

GetSendingEnabledOk returns a tuple with the SendingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingEnabled

`func (o *EmailSESStatusResult) SetSendingEnabled(v bool)`

SetSendingEnabled sets SendingEnabled field to given value.

### HasSendingEnabled

`func (o *EmailSESStatusResult) HasSendingEnabled() bool`

HasSendingEnabled returns a boolean if a field has been set.

### GetSentLast24Hours

`func (o *EmailSESStatusResult) GetSentLast24Hours() int64`

GetSentLast24Hours returns the SentLast24Hours field if non-nil, zero value otherwise.

### GetSentLast24HoursOk

`func (o *EmailSESStatusResult) GetSentLast24HoursOk() (*int64, bool)`

GetSentLast24HoursOk returns a tuple with the SentLast24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentLast24Hours

`func (o *EmailSESStatusResult) SetSentLast24Hours(v int64)`

SetSentLast24Hours sets SentLast24Hours field to given value.

### HasSentLast24Hours

`func (o *EmailSESStatusResult) HasSentLast24Hours() bool`

HasSentLast24Hours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



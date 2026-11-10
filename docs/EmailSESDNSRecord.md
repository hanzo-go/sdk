# EmailSESDNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the DNS record host or subdomain name. | [optional] 
**Priority** | Pointer to **int64** | Priority is the MX priority when Type is MX. | [optional] 
**Purpose** | Pointer to **string** | Purpose describes the role of the DNS record (e.g. DKIM, MAIL FROM MX, MAIL FROM SPF, DMARC). | [optional] 
**Type** | Pointer to **string** | Type is the DNS record type (CNAME, MX, TXT). | [optional] 
**Value** | Pointer to **string** | Value is the DNS record target or text value. | [optional] 
**Verified** | Pointer to **bool** | Verified reports whether the DNS record is verified for the sending identity. | [optional] 

## Methods

### NewEmailSESDNSRecord

`func NewEmailSESDNSRecord() *EmailSESDNSRecord`

NewEmailSESDNSRecord instantiates a new EmailSESDNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSESDNSRecordWithDefaults

`func NewEmailSESDNSRecordWithDefaults() *EmailSESDNSRecord`

NewEmailSESDNSRecordWithDefaults instantiates a new EmailSESDNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *EmailSESDNSRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailSESDNSRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailSESDNSRecord) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailSESDNSRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPriority

`func (o *EmailSESDNSRecord) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EmailSESDNSRecord) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EmailSESDNSRecord) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EmailSESDNSRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPurpose

`func (o *EmailSESDNSRecord) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *EmailSESDNSRecord) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *EmailSESDNSRecord) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *EmailSESDNSRecord) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetType

`func (o *EmailSESDNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailSESDNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailSESDNSRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailSESDNSRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *EmailSESDNSRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EmailSESDNSRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EmailSESDNSRecord) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EmailSESDNSRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVerified

`func (o *EmailSESDNSRecord) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EmailSESDNSRecord) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EmailSESDNSRecord) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EmailSESDNSRecord) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MarketingEsignDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | ExternalID is the reference the document was made under: the lead&#39;s. | [optional] 
**Id** | Pointer to **int64** | ID is the document&#39;s id in eSign. | [optional] 
**Recipients** | Pointer to [**[]MarketingEsignRecipient**](MarketingEsignRecipient.md) | Recipients are the document&#39;s signers. | [optional] 

## Methods

### NewMarketingEsignDocument

`func NewMarketingEsignDocument() *MarketingEsignDocument`

NewMarketingEsignDocument instantiates a new MarketingEsignDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEsignDocumentWithDefaults

`func NewMarketingEsignDocumentWithDefaults() *MarketingEsignDocument`

NewMarketingEsignDocumentWithDefaults instantiates a new MarketingEsignDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *MarketingEsignDocument) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MarketingEsignDocument) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MarketingEsignDocument) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MarketingEsignDocument) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *MarketingEsignDocument) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingEsignDocument) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingEsignDocument) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingEsignDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipients

`func (o *MarketingEsignDocument) GetRecipients() []MarketingEsignRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *MarketingEsignDocument) GetRecipientsOk() (*[]MarketingEsignRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *MarketingEsignDocument) SetRecipients(v []MarketingEsignRecipient)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *MarketingEsignDocument) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



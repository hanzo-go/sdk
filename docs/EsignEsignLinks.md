# EsignEsignLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document that went out. | [optional] 
**Recipients** | Pointer to [**[]EsignEsignLink**](EsignEsignLink.md) | Recipients is one link per signing recipient. Nothing is emailed by this call; delivering the links is the caller&#39;s. | [optional] 
**Status** | Pointer to **string** | Status is PENDING — the state a sent document is in until every signer has finished. | [optional] 

## Methods

### NewEsignEsignLinks

`func NewEsignEsignLinks() *EsignEsignLinks`

NewEsignEsignLinks instantiates a new EsignEsignLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignLinksWithDefaults

`func NewEsignEsignLinksWithDefaults() *EsignEsignLinks`

NewEsignEsignLinksWithDefaults instantiates a new EsignEsignLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsignEsignLinks) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignEsignLinks) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignEsignLinks) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignEsignLinks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipients

`func (o *EsignEsignLinks) GetRecipients() []EsignEsignLink`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EsignEsignLinks) GetRecipientsOk() (*[]EsignEsignLink, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EsignEsignLinks) SetRecipients(v []EsignEsignLink)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EsignEsignLinks) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetStatus

`func (o *EsignEsignLinks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignEsignLinks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignEsignLinks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignEsignLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



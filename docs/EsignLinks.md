# EsignLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document that went out. | [optional] 
**Recipients** | Pointer to [**[]EsignLink**](EsignLink.md) | Recipients is one link per signing recipient. Nothing is emailed by this call; delivering the links is the caller&#39;s. | [optional] 
**Status** | Pointer to **string** | Status is PENDING — the state a sent document is in until every signer has finished. | [optional] 

## Methods

### NewEsignLinks

`func NewEsignLinks() *EsignLinks`

NewEsignLinks instantiates a new EsignLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignLinksWithDefaults

`func NewEsignLinksWithDefaults() *EsignLinks`

NewEsignLinksWithDefaults instantiates a new EsignLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsignLinks) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignLinks) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignLinks) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignLinks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipients

`func (o *EsignLinks) GetRecipients() []EsignLink`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EsignLinks) GetRecipientsOk() (*[]EsignLink, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EsignLinks) SetRecipients(v []EsignLink)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EsignLinks) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetStatus

`func (o *EsignLinks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignLinks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignLinks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



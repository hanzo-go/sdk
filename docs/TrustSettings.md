# TrustSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the centre&#39;s display name. Required to publish. | [optional] 
**Nda** | Pointer to **string** | Nda is the text a party must accept before asking for a document. Optional; empty asks for no acceptance. The text in force is copied onto each request as it is accepted, so editing it never changes what anyone already agreed to. | [optional] 
**Publish** | Pointer to **bool** | Publish makes the centre answer at its public address. False withdraws it: the address stops answering while every item, grant and record stays exactly as it was, so withdrawing is reversible and loses nothing. | [optional] 
**Slug** | Pointer to **string** | Slug is the public address to answer at — a lowercase label of letters, digits and hyphens. Required to publish, unique across the deployment, and one org holds one: publishing under a new address MOVES the centre rather than leaving the old one answering. | [optional] 

## Methods

### NewTrustSettings

`func NewTrustSettings() *TrustSettings`

NewTrustSettings instantiates a new TrustSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustSettingsWithDefaults

`func NewTrustSettingsWithDefaults() *TrustSettings`

NewTrustSettingsWithDefaults instantiates a new TrustSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrustSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNda

`func (o *TrustSettings) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *TrustSettings) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *TrustSettings) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *TrustSettings) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetPublish

`func (o *TrustSettings) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *TrustSettings) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *TrustSettings) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *TrustSettings) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetSlug

`func (o *TrustSettings) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TrustSettings) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TrustSettings) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TrustSettings) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TrustEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body replaces the item&#39;s content. | [optional] 
**Document** | Pointer to **string** | Document replaces the file the item points at — this is how a report is superseded by its next edition. The new document must already exist in the caller org&#39;s own store. | [optional] 
**Framework** | Pointer to **string** | Framework replaces the standard it speaks to. | [optional] 
**Id** | Pointer to **string** | ID is the item to change, taken from the path. | [optional] 
**Name** | Pointer to **string** | Name replaces its title. | [optional] 
**Retired** | Pointer to **bool** | Retired withdraws the item, or true→false restores it. A retired item leaves the public centre at once and can no longer be granted; grants already made over it stand, because they are part of the record. | [optional] 
**Summary** | Pointer to **string** | Summary replaces the line about it. | [optional] 
**Tier** | Pointer to **string** | Tier moves it between public and gated. Moving an auditor-signed item to public is refused. | [optional] 

## Methods

### NewTrustEdit

`func NewTrustEdit() *TrustEdit`

NewTrustEdit instantiates a new TrustEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustEditWithDefaults

`func NewTrustEditWithDefaults() *TrustEdit`

NewTrustEditWithDefaults instantiates a new TrustEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TrustEdit) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TrustEdit) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TrustEdit) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TrustEdit) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDocument

`func (o *TrustEdit) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TrustEdit) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TrustEdit) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *TrustEdit) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetFramework

`func (o *TrustEdit) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TrustEdit) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TrustEdit) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TrustEdit) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *TrustEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrustEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetired

`func (o *TrustEdit) GetRetired() bool`

GetRetired returns the Retired field if non-nil, zero value otherwise.

### GetRetiredOk

`func (o *TrustEdit) GetRetiredOk() (*bool, bool)`

GetRetiredOk returns a tuple with the Retired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetired

`func (o *TrustEdit) SetRetired(v bool)`

SetRetired sets Retired field to given value.

### HasRetired

`func (o *TrustEdit) HasRetired() bool`

HasRetired returns a boolean if a field has been set.

### GetSummary

`func (o *TrustEdit) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TrustEdit) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TrustEdit) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TrustEdit) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier

`func (o *TrustEdit) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TrustEdit) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TrustEdit) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TrustEdit) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



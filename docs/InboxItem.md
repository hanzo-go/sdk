# InboxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Extracted** | Pointer to [**Extracted**](Extracted.md) |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | the file hash (&#x3D;&#x3D; a scan&#39;s ScanID) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewInboxItem

`func NewInboxItem() *InboxItem`

NewInboxItem instantiates a new InboxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxItemWithDefaults

`func NewInboxItemWithDefaults() *InboxItem`

NewInboxItemWithDefaults instantiates a new InboxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *InboxItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InboxItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InboxItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InboxItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfidence

`func (o *InboxItem) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *InboxItem) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *InboxItem) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *InboxItem) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InboxItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InboxItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InboxItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InboxItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExtracted

`func (o *InboxItem) GetExtracted() Extracted`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *InboxItem) GetExtractedOk() (*Extracted, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *InboxItem) SetExtracted(v Extracted)`

SetExtracted sets Extracted field to given value.

### HasExtracted

`func (o *InboxItem) HasExtracted() bool`

HasExtracted returns a boolean if a field has been set.

### GetFilename

`func (o *InboxItem) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *InboxItem) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *InboxItem) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *InboxItem) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *InboxItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InboxItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *InboxItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboxItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboxItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InboxItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendor

`func (o *InboxItem) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InboxItem) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InboxItem) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InboxItem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



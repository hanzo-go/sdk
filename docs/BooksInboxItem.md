# BooksInboxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the expense account the scanner proposed, as a chart number — a PROPOSAL, not a posting: nothing is booked until it is accepted. | [optional] 
**Confidence** | Pointer to **string** | Confidence is how sure the scanner is of that reading, and is the signal for whether a person needs to check it before it is booked. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the document was uploaded. | [optional] 
**Extracted** | Pointer to [**BooksExtracted**](BooksExtracted.md) | Extracted is what the scanner read off the document. Absent until it has been scanned, so its absence is \&quot;not read yet\&quot;, never \&quot;nothing on it\&quot;. | [optional] 
**Filename** | Pointer to **string** | Filename is the name the document was uploaded under, for a person to recognise it by. It is not part of the item&#39;s identity. | [optional] 
**Id** | Pointer to **string** | ID is the CONTENT HASH of the uploaded bytes, which is what makes the queue idempotent: re-uploading the same document returns this item rather than adding a second one. It is also the id the scan of this document carries. | [optional] 
**Status** | Pointer to **string** | Status is where the document is in the queue — unsorted until the scanner has read it, and thereafter whether it is waiting on a person or has been booked. | [optional] 
**Vendor** | Pointer to **string** | Vendor is the supplier the scanner identified, surfaced beside the item so a queue renders without opening each document. | [optional] 

## Methods

### NewBooksInboxItem

`func NewBooksInboxItem() *BooksInboxItem`

NewBooksInboxItem instantiates a new BooksInboxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksInboxItemWithDefaults

`func NewBooksInboxItemWithDefaults() *BooksInboxItem`

NewBooksInboxItemWithDefaults instantiates a new BooksInboxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BooksInboxItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BooksInboxItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BooksInboxItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BooksInboxItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfidence

`func (o *BooksInboxItem) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BooksInboxItem) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BooksInboxItem) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *BooksInboxItem) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BooksInboxItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BooksInboxItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BooksInboxItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BooksInboxItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExtracted

`func (o *BooksInboxItem) GetExtracted() BooksExtracted`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *BooksInboxItem) GetExtractedOk() (*BooksExtracted, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *BooksInboxItem) SetExtracted(v BooksExtracted)`

SetExtracted sets Extracted field to given value.

### HasExtracted

`func (o *BooksInboxItem) HasExtracted() bool`

HasExtracted returns a boolean if a field has been set.

### GetFilename

`func (o *BooksInboxItem) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BooksInboxItem) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BooksInboxItem) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *BooksInboxItem) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *BooksInboxItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BooksInboxItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BooksInboxItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BooksInboxItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BooksInboxItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BooksInboxItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BooksInboxItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BooksInboxItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendor

`func (o *BooksInboxItem) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BooksInboxItem) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BooksInboxItem) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *BooksInboxItem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



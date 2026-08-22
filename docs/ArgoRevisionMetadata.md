# ArgoRevisionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the commit author. Always absent: an App CR pins an IMAGE, so this process has no commit to read one from and will not invent one. | [optional] 
**Date** | Pointer to **string** | Date is when the App CR was created, RFC 3339 UTC — the only real timestamp there is here. It is NOT the date of the revision asked for. | [optional] 
**Message** | Pointer to **string** | Message is the revision asked for, echoed back — not a commit message. The empty revision and \&quot;HEAD\&quot; resolve to the image tag the CR declares (spec.image.tag), and anything longer than 256 characters is truncated to it. | [optional] 
**SignatureInfo** | Pointer to **string** | SignatureInfo is the GPG verification result for the revision. Always absent: nothing here verifies a signature, and an empty field says so rather than implying an unsigned commit. | [optional] 
**Tags** | Pointer to **[]string** | Tags are the git tags pointing at the revision. Always absent, for the same reason as Author. | [optional] 

## Methods

### NewArgoRevisionMetadata

`func NewArgoRevisionMetadata() *ArgoRevisionMetadata`

NewArgoRevisionMetadata instantiates a new ArgoRevisionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoRevisionMetadataWithDefaults

`func NewArgoRevisionMetadataWithDefaults() *ArgoRevisionMetadata`

NewArgoRevisionMetadataWithDefaults instantiates a new ArgoRevisionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ArgoRevisionMetadata) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ArgoRevisionMetadata) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ArgoRevisionMetadata) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ArgoRevisionMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDate

`func (o *ArgoRevisionMetadata) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ArgoRevisionMetadata) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ArgoRevisionMetadata) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ArgoRevisionMetadata) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMessage

`func (o *ArgoRevisionMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArgoRevisionMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArgoRevisionMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArgoRevisionMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSignatureInfo

`func (o *ArgoRevisionMetadata) GetSignatureInfo() string`

GetSignatureInfo returns the SignatureInfo field if non-nil, zero value otherwise.

### GetSignatureInfoOk

`func (o *ArgoRevisionMetadata) GetSignatureInfoOk() (*string, bool)`

GetSignatureInfoOk returns a tuple with the SignatureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureInfo

`func (o *ArgoRevisionMetadata) SetSignatureInfo(v string)`

SetSignatureInfo sets SignatureInfo field to given value.

### HasSignatureInfo

`func (o *ArgoRevisionMetadata) HasSignatureInfo() bool`

HasSignatureInfo returns a boolean if a field has been set.

### GetTags

`func (o *ArgoRevisionMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArgoRevisionMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArgoRevisionMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArgoRevisionMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



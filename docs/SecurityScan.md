# SecurityScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the source to scan. It is NEVER stored: what persists is the finding, with a masked preview and a fingerprint. | [optional] 
**Path** | Pointer to **string** | Path is where the file lives, recorded on any finding so a result can be located in the tree it came from. | [optional] 

## Methods

### NewSecurityScan

`func NewSecurityScan() *SecurityScan`

NewSecurityScan instantiates a new SecurityScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanWithDefaults

`func NewSecurityScanWithDefaults() *SecurityScan`

NewSecurityScanWithDefaults instantiates a new SecurityScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SecurityScan) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SecurityScan) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SecurityScan) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SecurityScan) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPath

`func (o *SecurityScan) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecurityScan) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecurityScan) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecurityScan) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



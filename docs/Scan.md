# Scan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the source to scan. It is NEVER stored: what persists is the finding, with a masked preview and a fingerprint. | [optional] 
**Path** | Pointer to **string** | Path is where the file lives, recorded on any finding so a result can be located in the tree it came from. | [optional] 

## Methods

### NewScan

`func NewScan() *Scan`

NewScan instantiates a new Scan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanWithDefaults

`func NewScanWithDefaults() *Scan`

NewScanWithDefaults instantiates a new Scan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Scan) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Scan) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Scan) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Scan) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPath

`func (o *Scan) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Scan) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Scan) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Scan) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



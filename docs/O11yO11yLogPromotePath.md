# O11yO11yLogPromotePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | Pointer to [**[]O11yO11yLogPromoteIndex**](O11yO11yLogPromoteIndex.md) | Indexes are the indexes to put on the path. | [optional] 
**Path** | Pointer to **string** | Path is the body path, e.g. body.user.id on the way in; listed without the body. prefix. | [optional] 
**Promote** | Pointer to **bool** | Promote lifts the path into its own column when true. | [optional] 

## Methods

### NewO11yO11yLogPromotePath

`func NewO11yO11yLogPromotePath() *O11yO11yLogPromotePath`

NewO11yO11yLogPromotePath instantiates a new O11yO11yLogPromotePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPromotePathWithDefaults

`func NewO11yO11yLogPromotePathWithDefaults() *O11yO11yLogPromotePath`

NewO11yO11yLogPromotePathWithDefaults instantiates a new O11yO11yLogPromotePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *O11yO11yLogPromotePath) GetIndexes() []O11yO11yLogPromoteIndex`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *O11yO11yLogPromotePath) GetIndexesOk() (*[]O11yO11yLogPromoteIndex, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *O11yO11yLogPromotePath) SetIndexes(v []O11yO11yLogPromoteIndex)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *O11yO11yLogPromotePath) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetPath

`func (o *O11yO11yLogPromotePath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *O11yO11yLogPromotePath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *O11yO11yLogPromotePath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *O11yO11yLogPromotePath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPromote

`func (o *O11yO11yLogPromotePath) GetPromote() bool`

GetPromote returns the Promote field if non-nil, zero value otherwise.

### GetPromoteOk

`func (o *O11yO11yLogPromotePath) GetPromoteOk() (*bool, bool)`

GetPromoteOk returns a tuple with the Promote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromote

`func (o *O11yO11yLogPromotePath) SetPromote(v bool)`

SetPromote sets Promote field to given value.

### HasPromote

`func (o *O11yO11yLogPromotePath) HasPromote() bool`

HasPromote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



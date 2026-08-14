# O11yO11ySavedViewUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CompositeQuery** | Pointer to [**O11yCompositeQuery**](O11yCompositeQuery.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ExtraData** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourcePage** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**ViewId** | Pointer to **string** | ViewID is the id of the view to replace, taken from the URL. | [optional] 

## Methods

### NewO11yO11ySavedViewUpdateIn

`func NewO11yO11ySavedViewUpdateIn() *O11yO11ySavedViewUpdateIn`

NewO11yO11ySavedViewUpdateIn instantiates a new O11yO11ySavedViewUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySavedViewUpdateInWithDefaults

`func NewO11yO11ySavedViewUpdateInWithDefaults() *O11yO11ySavedViewUpdateIn`

NewO11yO11ySavedViewUpdateInWithDefaults instantiates a new O11yO11ySavedViewUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *O11yO11ySavedViewUpdateIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *O11yO11ySavedViewUpdateIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *O11yO11ySavedViewUpdateIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *O11yO11ySavedViewUpdateIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompositeQuery

`func (o *O11yO11ySavedViewUpdateIn) GetCompositeQuery() O11yCompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *O11yO11ySavedViewUpdateIn) GetCompositeQueryOk() (*O11yCompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *O11yO11ySavedViewUpdateIn) SetCompositeQuery(v O11yCompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *O11yO11ySavedViewUpdateIn) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11ySavedViewUpdateIn) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11ySavedViewUpdateIn) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11ySavedViewUpdateIn) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11ySavedViewUpdateIn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11ySavedViewUpdateIn) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11ySavedViewUpdateIn) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11ySavedViewUpdateIn) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11ySavedViewUpdateIn) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExtraData

`func (o *O11yO11ySavedViewUpdateIn) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *O11yO11ySavedViewUpdateIn) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *O11yO11ySavedViewUpdateIn) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *O11yO11ySavedViewUpdateIn) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetId

`func (o *O11yO11ySavedViewUpdateIn) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11ySavedViewUpdateIn) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11ySavedViewUpdateIn) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11ySavedViewUpdateIn) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11yO11ySavedViewUpdateIn) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11yO11ySavedViewUpdateIn) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *O11yO11ySavedViewUpdateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11ySavedViewUpdateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11ySavedViewUpdateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11ySavedViewUpdateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourcePage

`func (o *O11yO11ySavedViewUpdateIn) GetSourcePage() string`

GetSourcePage returns the SourcePage field if non-nil, zero value otherwise.

### GetSourcePageOk

`func (o *O11yO11ySavedViewUpdateIn) GetSourcePageOk() (*string, bool)`

GetSourcePageOk returns a tuple with the SourcePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePage

`func (o *O11yO11ySavedViewUpdateIn) SetSourcePage(v string)`

SetSourcePage sets SourcePage field to given value.

### HasSourcePage

`func (o *O11yO11ySavedViewUpdateIn) HasSourcePage() bool`

HasSourcePage returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11ySavedViewUpdateIn) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11ySavedViewUpdateIn) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11ySavedViewUpdateIn) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11ySavedViewUpdateIn) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11ySavedViewUpdateIn) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11ySavedViewUpdateIn) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11ySavedViewUpdateIn) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11ySavedViewUpdateIn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11ySavedViewUpdateIn) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11ySavedViewUpdateIn) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11ySavedViewUpdateIn) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11ySavedViewUpdateIn) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetViewId

`func (o *O11yO11ySavedViewUpdateIn) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *O11yO11ySavedViewUpdateIn) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *O11yO11ySavedViewUpdateIn) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *O11yO11ySavedViewUpdateIn) HasViewId() bool`

HasViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



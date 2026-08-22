# Index

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]Op**](Op.md) |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 

## Methods

### NewIndex

`func NewIndex() *Index`

NewIndex instantiates a new Index object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexWithDefaults

`func NewIndexWithDefaults() *Index`

NewIndexWithDefaults instantiates a new Index object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Index) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Index) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Index) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Index) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *Index) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Index) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Index) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Index) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Index) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Index) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Index) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Index) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperations

`func (o *Index) GetOperations() []Op`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Index) GetOperationsOk() (*[]Op, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Index) SetOperations(v []Op)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Index) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetStage

`func (o *Index) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Index) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Index) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Index) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



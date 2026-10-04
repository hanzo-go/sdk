# KnowledgeFileLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** | Direction is out when this file links to the other, in when the other links to this one, and shared when the two name the same entities. | [optional] 
**File** | Pointer to **string** | File is the other file, by id. | [optional] 
**Kind** | Pointer to **string** | Kind is href (a hyperlink), cites (the text names the file), xref (a cross-reference inside one document) or entity (shared mentions). | [optional] 
**Label** | Pointer to **string** | Label is the link&#39;s own text, or the entities the two files share. | [optional] 
**Name** | Pointer to **string** | Name is the other file&#39;s name. | [optional] 
**Weight** | Pointer to **int64** | Weight is how many edges of this kind join the two; for shared entities, how many entities they share. | [optional] 

## Methods

### NewKnowledgeFileLink

`func NewKnowledgeFileLink() *KnowledgeFileLink`

NewKnowledgeFileLink instantiates a new KnowledgeFileLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileLinkWithDefaults

`func NewKnowledgeFileLinkWithDefaults() *KnowledgeFileLink`

NewKnowledgeFileLinkWithDefaults instantiates a new KnowledgeFileLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *KnowledgeFileLink) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KnowledgeFileLink) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KnowledgeFileLink) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *KnowledgeFileLink) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFile

`func (o *KnowledgeFileLink) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeFileLink) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeFileLink) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeFileLink) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetKind

`func (o *KnowledgeFileLink) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KnowledgeFileLink) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KnowledgeFileLink) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KnowledgeFileLink) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *KnowledgeFileLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *KnowledgeFileLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *KnowledgeFileLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *KnowledgeFileLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeFileLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeFileLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeFileLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnowledgeFileLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *KnowledgeFileLink) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *KnowledgeFileLink) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *KnowledgeFileLink) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *KnowledgeFileLink) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



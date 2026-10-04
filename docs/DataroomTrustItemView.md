# DataroomTrustItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attester** | Pointer to **string** | Attester is who vouched for it: self or auditor. | [optional] 
**Body** | Pointer to **string** | Body is the item&#39;s content for the kinds that are text rather than a file. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when it was published, in unix milliseconds. | [optional] 
**Document** | Pointer to **string** | Document is the data-room document holding its bytes, empty when it has none. | [optional] 
**Framework** | Pointer to **string** | Framework is the standard it speaks to, when it speaks to one. | [optional] 
**Id** | Pointer to **string** | ID addresses the item. | [optional] 
**Kind** | Pointer to **string** | Kind is one of report, letter, policy, questionnaire, subprocessor, article or update — the closed set the public centre knows how to draw. | [optional] 
**Name** | Pointer to **string** | Name is the label the centre lists it under. | [optional] 
**Retired** | Pointer to **bool** | Retired is whether it has been withdrawn. A retired item is absent from the public centre and cannot be granted; it is kept because a grant already made over it is part of the record. | [optional] 
**Summary** | Pointer to **string** | Summary is a line about it. | [optional] 
**Tier** | Pointer to **string** | Tier is public or gated. Gated is the default and an auditor-signed item can only ever be gated. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomTrustItemView

`func NewDataroomTrustItemView() *DataroomTrustItemView`

NewDataroomTrustItemView instantiates a new DataroomTrustItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustItemViewWithDefaults

`func NewDataroomTrustItemViewWithDefaults() *DataroomTrustItemView`

NewDataroomTrustItemViewWithDefaults instantiates a new DataroomTrustItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttester

`func (o *DataroomTrustItemView) GetAttester() string`

GetAttester returns the Attester field if non-nil, zero value otherwise.

### GetAttesterOk

`func (o *DataroomTrustItemView) GetAttesterOk() (*string, bool)`

GetAttesterOk returns a tuple with the Attester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttester

`func (o *DataroomTrustItemView) SetAttester(v string)`

SetAttester sets Attester field to given value.

### HasAttester

`func (o *DataroomTrustItemView) HasAttester() bool`

HasAttester returns a boolean if a field has been set.

### GetBody

`func (o *DataroomTrustItemView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DataroomTrustItemView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DataroomTrustItemView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DataroomTrustItemView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataroomTrustItemView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataroomTrustItemView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataroomTrustItemView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataroomTrustItemView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocument

`func (o *DataroomTrustItemView) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DataroomTrustItemView) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DataroomTrustItemView) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *DataroomTrustItemView) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetFramework

`func (o *DataroomTrustItemView) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *DataroomTrustItemView) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *DataroomTrustItemView) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *DataroomTrustItemView) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *DataroomTrustItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomTrustItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomTrustItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomTrustItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DataroomTrustItemView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataroomTrustItemView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataroomTrustItemView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DataroomTrustItemView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DataroomTrustItemView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomTrustItemView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomTrustItemView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomTrustItemView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetired

`func (o *DataroomTrustItemView) GetRetired() bool`

GetRetired returns the Retired field if non-nil, zero value otherwise.

### GetRetiredOk

`func (o *DataroomTrustItemView) GetRetiredOk() (*bool, bool)`

GetRetiredOk returns a tuple with the Retired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetired

`func (o *DataroomTrustItemView) SetRetired(v bool)`

SetRetired sets Retired field to given value.

### HasRetired

`func (o *DataroomTrustItemView) HasRetired() bool`

HasRetired returns a boolean if a field has been set.

### GetSummary

`func (o *DataroomTrustItemView) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DataroomTrustItemView) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DataroomTrustItemView) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DataroomTrustItemView) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier

`func (o *DataroomTrustItemView) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DataroomTrustItemView) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DataroomTrustItemView) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DataroomTrustItemView) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomTrustItemView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomTrustItemView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomTrustItemView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomTrustItemView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



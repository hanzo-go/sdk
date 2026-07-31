# CloudSubjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the subject was recorded. | [optional] 
**HasEmail** | Pointer to **bool** | HasEmail reports whether a contact email is on file, without exposing it. | [optional] 
**Id** | Pointer to **string** | ID is the subject&#39;s opaque id. | [optional] 
**Kind** | Pointer to **string** | Kind is the party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Ref** | Pointer to **string** | Ref is the org&#39;s own opaque external id for this subject. | [optional] 

## Methods

### NewCloudSubjectSummary

`func NewCloudSubjectSummary() *CloudSubjectSummary`

NewCloudSubjectSummary instantiates a new CloudSubjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubjectSummaryWithDefaults

`func NewCloudSubjectSummaryWithDefaults() *CloudSubjectSummary`

NewCloudSubjectSummaryWithDefaults instantiates a new CloudSubjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudSubjectSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSubjectSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSubjectSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSubjectSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHasEmail

`func (o *CloudSubjectSummary) GetHasEmail() bool`

GetHasEmail returns the HasEmail field if non-nil, zero value otherwise.

### GetHasEmailOk

`func (o *CloudSubjectSummary) GetHasEmailOk() (*bool, bool)`

GetHasEmailOk returns a tuple with the HasEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmail

`func (o *CloudSubjectSummary) SetHasEmail(v bool)`

SetHasEmail sets HasEmail field to given value.

### HasHasEmail

`func (o *CloudSubjectSummary) HasHasEmail() bool`

HasHasEmail returns a boolean if a field has been set.

### GetId

`func (o *CloudSubjectSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSubjectSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSubjectSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSubjectSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudSubjectSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudSubjectSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudSubjectSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudSubjectSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRef

`func (o *CloudSubjectSummary) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudSubjectSummary) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudSubjectSummary) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudSubjectSummary) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



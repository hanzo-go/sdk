# VerificationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is an inline subject&#39;s contact email, sealed at rest. | [optional] 
**Kind** | Pointer to **string** | Kind is an inline subject&#39;s party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Name** | Pointer to **string** | Name is an inline subject&#39;s name, sealed at rest. | [optional] 
**Ref** | Pointer to **string** | Ref is the org&#39;s own opaque external id for an inline subject. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID names an existing subject to verify; empty creates one inline. | [optional] 

## Methods

### NewVerificationReq

`func NewVerificationReq() *VerificationReq`

NewVerificationReq instantiates a new VerificationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationReqWithDefaults

`func NewVerificationReqWithDefaults() *VerificationReq`

NewVerificationReqWithDefaults instantiates a new VerificationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VerificationReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VerificationReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VerificationReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VerificationReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *VerificationReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VerificationReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VerificationReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VerificationReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *VerificationReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerificationReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerificationReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VerificationReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRef

`func (o *VerificationReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *VerificationReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *VerificationReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *VerificationReq) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSubjectId

`func (o *VerificationReq) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *VerificationReq) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *VerificationReq) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *VerificationReq) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



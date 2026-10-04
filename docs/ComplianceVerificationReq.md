# ComplianceVerificationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is an inline subject&#39;s contact email, sealed at rest. | [optional] 
**Kind** | Pointer to **string** | Kind is an inline subject&#39;s party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Name** | Pointer to **string** | Name is an inline subject&#39;s name, sealed at rest. | [optional] 
**Ref** | Pointer to **string** | Ref is the org&#39;s own opaque external id for an inline subject. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID names an existing subject to verify; empty creates one inline. | [optional] 

## Methods

### NewComplianceVerificationReq

`func NewComplianceVerificationReq() *ComplianceVerificationReq`

NewComplianceVerificationReq instantiates a new ComplianceVerificationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceVerificationReqWithDefaults

`func NewComplianceVerificationReqWithDefaults() *ComplianceVerificationReq`

NewComplianceVerificationReqWithDefaults instantiates a new ComplianceVerificationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ComplianceVerificationReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ComplianceVerificationReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ComplianceVerificationReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ComplianceVerificationReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *ComplianceVerificationReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComplianceVerificationReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComplianceVerificationReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComplianceVerificationReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ComplianceVerificationReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComplianceVerificationReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComplianceVerificationReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComplianceVerificationReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRef

`func (o *ComplianceVerificationReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ComplianceVerificationReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ComplianceVerificationReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ComplianceVerificationReq) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSubjectId

`func (o *ComplianceVerificationReq) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ComplianceVerificationReq) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ComplianceVerificationReq) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *ComplianceVerificationReq) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



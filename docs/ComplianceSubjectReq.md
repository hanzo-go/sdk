# ComplianceSubjectReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the subject&#39;s contact email, sealed at rest. | [optional] 
**Kind** | Pointer to **string** | Kind is the party type: \&quot;individual\&quot; (KYC) or \&quot;business\&quot; (KYB). | [optional] 
**Name** | Pointer to **string** | Name is the subject&#39;s name, sealed at rest. | [optional] 
**Ref** | Pointer to **string** | Ref is the org&#39;s own opaque external id for this subject. | [optional] 

## Methods

### NewComplianceSubjectReq

`func NewComplianceSubjectReq() *ComplianceSubjectReq`

NewComplianceSubjectReq instantiates a new ComplianceSubjectReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceSubjectReqWithDefaults

`func NewComplianceSubjectReqWithDefaults() *ComplianceSubjectReq`

NewComplianceSubjectReqWithDefaults instantiates a new ComplianceSubjectReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ComplianceSubjectReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ComplianceSubjectReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ComplianceSubjectReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ComplianceSubjectReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *ComplianceSubjectReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComplianceSubjectReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComplianceSubjectReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComplianceSubjectReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ComplianceSubjectReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComplianceSubjectReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComplianceSubjectReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComplianceSubjectReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRef

`func (o *ComplianceSubjectReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ComplianceSubjectReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ComplianceSubjectReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ComplianceSubjectReq) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



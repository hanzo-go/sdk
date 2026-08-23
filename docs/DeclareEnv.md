# DeclareEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** | Public marks a value that may be WRITTEN INTO GIT. Absent, it is false, and the value is sealed into KMS and referenced.  ★ THE DEFAULT IS SECRET, AND THE POLARITY IS THE WHOLE DESIGN. This lane&#39;s output is a commit in a repository replicated to every clone, so a misclassification is not a bug to fix later — it is a credential published forever. A heuristic classifier fails in both directions; what decides is which direction it fails IN. Seal-by-default makes the failure mode \&quot;an operator cannot read back a config value\&quot;, which is a support ticket. Classify-by-shape made it \&quot;a password is in git history\&quot;, which is an incident with no rollback.  It is also the only rule that needs no list. PGPASSWORD, *_PW, a symbol-rich password, a KUBECONFIG, a base32 MFA seed — every one of them slipped a shape classifier, and each miss was a different reason. There is no reason left when the default is to seal. | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewDeclareEnv

`func NewDeclareEnv() *DeclareEnv`

NewDeclareEnv instantiates a new DeclareEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareEnvWithDefaults

`func NewDeclareEnvWithDefaults() *DeclareEnv`

NewDeclareEnvWithDefaults instantiates a new DeclareEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeclareEnv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeclareEnv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeclareEnv) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeclareEnv) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *DeclareEnv) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *DeclareEnv) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *DeclareEnv) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *DeclareEnv) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetValue

`func (o *DeclareEnv) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeclareEnv) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeclareEnv) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeclareEnv) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



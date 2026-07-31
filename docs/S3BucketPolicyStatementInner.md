# S3BucketPolicyStatementInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** |  | [optional] 
**Effect** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **[]string** |  | [optional] 
**Resource** | Pointer to **[]string** |  | [optional] 
**Condition** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewS3BucketPolicyStatementInner

`func NewS3BucketPolicyStatementInner() *S3BucketPolicyStatementInner`

NewS3BucketPolicyStatementInner instantiates a new S3BucketPolicyStatementInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketPolicyStatementInnerWithDefaults

`func NewS3BucketPolicyStatementInnerWithDefaults() *S3BucketPolicyStatementInner`

NewS3BucketPolicyStatementInnerWithDefaults instantiates a new S3BucketPolicyStatementInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSid

`func (o *S3BucketPolicyStatementInner) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *S3BucketPolicyStatementInner) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *S3BucketPolicyStatementInner) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *S3BucketPolicyStatementInner) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetEffect

`func (o *S3BucketPolicyStatementInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *S3BucketPolicyStatementInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *S3BucketPolicyStatementInner) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *S3BucketPolicyStatementInner) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetPrincipal

`func (o *S3BucketPolicyStatementInner) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *S3BucketPolicyStatementInner) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *S3BucketPolicyStatementInner) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *S3BucketPolicyStatementInner) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetAction

`func (o *S3BucketPolicyStatementInner) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *S3BucketPolicyStatementInner) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *S3BucketPolicyStatementInner) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *S3BucketPolicyStatementInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *S3BucketPolicyStatementInner) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *S3BucketPolicyStatementInner) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *S3BucketPolicyStatementInner) SetResource(v []string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *S3BucketPolicyStatementInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCondition

`func (o *S3BucketPolicyStatementInner) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *S3BucketPolicyStatementInner) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *S3BucketPolicyStatementInner) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *S3BucketPolicyStatementInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



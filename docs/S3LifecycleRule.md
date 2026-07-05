# S3LifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to [**S3LifecycleRuleExpiration**](S3LifecycleRuleExpiration.md) |  | [optional] 
**Transition** | Pointer to [**S3LifecycleRuleTransition**](S3LifecycleRuleTransition.md) |  | [optional] 

## Methods

### NewS3LifecycleRule

`func NewS3LifecycleRule() *S3LifecycleRule`

NewS3LifecycleRule instantiates a new S3LifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LifecycleRuleWithDefaults

`func NewS3LifecycleRuleWithDefaults() *S3LifecycleRule`

NewS3LifecycleRuleWithDefaults instantiates a new S3LifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3LifecycleRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3LifecycleRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3LifecycleRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3LifecycleRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *S3LifecycleRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3LifecycleRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3LifecycleRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *S3LifecycleRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrefix

`func (o *S3LifecycleRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3LifecycleRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3LifecycleRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3LifecycleRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetExpiration

`func (o *S3LifecycleRule) GetExpiration() S3LifecycleRuleExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *S3LifecycleRule) GetExpirationOk() (*S3LifecycleRuleExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *S3LifecycleRule) SetExpiration(v S3LifecycleRuleExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *S3LifecycleRule) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetTransition

`func (o *S3LifecycleRule) GetTransition() S3LifecycleRuleTransition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *S3LifecycleRule) GetTransitionOk() (*S3LifecycleRuleTransition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *S3LifecycleRule) SetTransition(v S3LifecycleRuleTransition)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *S3LifecycleRule) HasTransition() bool`

HasTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



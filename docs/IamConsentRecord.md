# IamConsentRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**GrantedScopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamConsentRecord

`func NewIamConsentRecord() *IamConsentRecord`

NewIamConsentRecord instantiates a new IamConsentRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamConsentRecordWithDefaults

`func NewIamConsentRecordWithDefaults() *IamConsentRecord`

NewIamConsentRecordWithDefaults instantiates a new IamConsentRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamConsentRecord) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamConsentRecord) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamConsentRecord) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamConsentRecord) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetGrantedScopes

`func (o *IamConsentRecord) GetGrantedScopes() []string`

GetGrantedScopes returns the GrantedScopes field if non-nil, zero value otherwise.

### GetGrantedScopesOk

`func (o *IamConsentRecord) GetGrantedScopesOk() (*[]string, bool)`

GetGrantedScopesOk returns a tuple with the GrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedScopes

`func (o *IamConsentRecord) SetGrantedScopes(v []string)`

SetGrantedScopes sets GrantedScopes field to given value.

### HasGrantedScopes

`func (o *IamConsentRecord) HasGrantedScopes() bool`

HasGrantedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



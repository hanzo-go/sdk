# IamObjectWebFinger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to [**[]IamObjectWebFingerLink**](IamObjectWebFingerLink.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectWebFinger

`func NewIamObjectWebFinger() *IamObjectWebFinger`

NewIamObjectWebFinger instantiates a new IamObjectWebFinger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectWebFingerWithDefaults

`func NewIamObjectWebFingerWithDefaults() *IamObjectWebFinger`

NewIamObjectWebFingerWithDefaults instantiates a new IamObjectWebFinger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *IamObjectWebFinger) GetAliases() map[string]interface{}`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *IamObjectWebFinger) GetAliasesOk() (*map[string]interface{}, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *IamObjectWebFinger) SetAliases(v map[string]interface{})`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *IamObjectWebFinger) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetLinks

`func (o *IamObjectWebFinger) GetLinks() []IamObjectWebFingerLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IamObjectWebFinger) GetLinksOk() (*[]IamObjectWebFingerLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IamObjectWebFinger) SetLinks(v []IamObjectWebFingerLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IamObjectWebFinger) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProperties

`func (o *IamObjectWebFinger) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IamObjectWebFinger) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IamObjectWebFinger) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IamObjectWebFinger) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSubject

`func (o *IamObjectWebFinger) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IamObjectWebFinger) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IamObjectWebFinger) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IamObjectWebFinger) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



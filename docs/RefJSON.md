# RefJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the short ref name (\&quot;main\&quot;, \&quot;v1.2.0\&quot;), not the full refs/… path. | [optional] 
**Sha** | Pointer to **string** | SHA is the full commit hash the ref resolves to. | [optional] 

## Methods

### NewRefJSON

`func NewRefJSON() *RefJSON`

NewRefJSON instantiates a new RefJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefJSONWithDefaults

`func NewRefJSONWithDefaults() *RefJSON`

NewRefJSONWithDefaults instantiates a new RefJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RefJSON) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RefJSON) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RefJSON) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RefJSON) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha

`func (o *RefJSON) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *RefJSON) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *RefJSON) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *RefJSON) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



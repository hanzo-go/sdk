# IamIdentifierBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | ClientId is the application&#39;s OAuth client id. | [optional] 
**Identifier** | Pointer to **string** | Identifier is the email address the person typed. | [optional] 

## Methods

### NewIamIdentifierBody

`func NewIamIdentifierBody() *IamIdentifierBody`

NewIamIdentifierBody instantiates a new IamIdentifierBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdentifierBodyWithDefaults

`func NewIamIdentifierBodyWithDefaults() *IamIdentifierBody`

NewIamIdentifierBodyWithDefaults instantiates a new IamIdentifierBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IamIdentifierBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamIdentifierBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamIdentifierBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamIdentifierBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIdentifier

`func (o *IamIdentifierBody) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IamIdentifierBody) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IamIdentifierBody) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IamIdentifierBody) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



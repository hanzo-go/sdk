# IamAddressBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the email address or phone number being added. | [optional] 
**ClientId** | Pointer to **string** | ClientId is the application the address will sign in at — the one whose applicationId the code was sent under. | [optional] 
**Code** | Pointer to **string** | Code is the one-time code sent to Address. | [optional] 

## Methods

### NewIamAddressBody

`func NewIamAddressBody() *IamAddressBody`

NewIamAddressBody instantiates a new IamAddressBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAddressBodyWithDefaults

`func NewIamAddressBodyWithDefaults() *IamAddressBody`

NewIamAddressBodyWithDefaults instantiates a new IamAddressBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IamAddressBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IamAddressBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IamAddressBody) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IamAddressBody) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetClientId

`func (o *IamAddressBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamAddressBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamAddressBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamAddressBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCode

`func (o *IamAddressBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamAddressBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamAddressBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamAddressBody) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



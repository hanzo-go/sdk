# CaptableClassHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **int32** | Authorized is how many shares of the class are authorized. | [optional] 
**ClassType** | Pointer to **string** | ClassType is COMMON or PREFERRED. | [optional] 
**Issued** | Pointer to **int32** | Issued is how many shares of the class have been issued. | [optional] 
**Name** | Pointer to **string** | Name is the class name. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the share class. | [optional] 

## Methods

### NewCaptableClassHolding

`func NewCaptableClassHolding() *CaptableClassHolding`

NewCaptableClassHolding instantiates a new CaptableClassHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableClassHoldingWithDefaults

`func NewCaptableClassHoldingWithDefaults() *CaptableClassHolding`

NewCaptableClassHoldingWithDefaults instantiates a new CaptableClassHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CaptableClassHolding) GetAuthorized() int32`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CaptableClassHolding) GetAuthorizedOk() (*int32, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CaptableClassHolding) SetAuthorized(v int32)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CaptableClassHolding) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetClassType

`func (o *CaptableClassHolding) GetClassType() string`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *CaptableClassHolding) GetClassTypeOk() (*string, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *CaptableClassHolding) SetClassType(v string)`

SetClassType sets ClassType field to given value.

### HasClassType

`func (o *CaptableClassHolding) HasClassType() bool`

HasClassType returns a boolean if a field has been set.

### GetIssued

`func (o *CaptableClassHolding) GetIssued() int32`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CaptableClassHolding) GetIssuedOk() (*int32, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CaptableClassHolding) SetIssued(v int32)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *CaptableClassHolding) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetName

`func (o *CaptableClassHolding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableClassHolding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableClassHolding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableClassHolding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableClassHolding) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableClassHolding) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableClassHolding) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableClassHolding) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



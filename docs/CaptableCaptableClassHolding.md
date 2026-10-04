# CaptableCaptableClassHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **int64** | Authorized is how many shares of the class are authorized. | [optional] 
**ClassType** | Pointer to **string** | ClassType is COMMON or PREFERRED. | [optional] 
**Issued** | Pointer to **int64** | Issued is how many shares of the class have been issued. | [optional] 
**Name** | Pointer to **string** | Name is the class name. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID addresses the class this position is for. | [optional] 

## Methods

### NewCaptableCaptableClassHolding

`func NewCaptableCaptableClassHolding() *CaptableCaptableClassHolding`

NewCaptableCaptableClassHolding instantiates a new CaptableCaptableClassHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableClassHoldingWithDefaults

`func NewCaptableCaptableClassHoldingWithDefaults() *CaptableCaptableClassHolding`

NewCaptableCaptableClassHoldingWithDefaults instantiates a new CaptableCaptableClassHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CaptableCaptableClassHolding) GetAuthorized() int64`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CaptableCaptableClassHolding) GetAuthorizedOk() (*int64, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CaptableCaptableClassHolding) SetAuthorized(v int64)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CaptableCaptableClassHolding) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetClassType

`func (o *CaptableCaptableClassHolding) GetClassType() string`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *CaptableCaptableClassHolding) GetClassTypeOk() (*string, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *CaptableCaptableClassHolding) SetClassType(v string)`

SetClassType sets ClassType field to given value.

### HasClassType

`func (o *CaptableCaptableClassHolding) HasClassType() bool`

HasClassType returns a boolean if a field has been set.

### GetIssued

`func (o *CaptableCaptableClassHolding) GetIssued() int64`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CaptableCaptableClassHolding) GetIssuedOk() (*int64, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CaptableCaptableClassHolding) SetIssued(v int64)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *CaptableCaptableClassHolding) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetName

`func (o *CaptableCaptableClassHolding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableCaptableClassHolding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableCaptableClassHolding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableCaptableClassHolding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableCaptableClassHolding) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableCaptableClassHolding) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableCaptableClassHolding) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableCaptableClassHolding) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



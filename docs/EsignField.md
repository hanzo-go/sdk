# EsignField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomText** | Pointer to **string** | CustomText is the value a non-signature field was filled with, empty until it is. A signature&#39;s value is not here: it is stored separately and rendered onto the page at sealing. | [optional] 
**FieldMeta** | Pointer to **interface{}** |  | [optional] 
**Height** | Pointer to **float64** | Height is the field&#39;s height, -1 when the renderer is to choose one. | [optional] 
**Id** | Pointer to **string** | ID is the field id. | [optional] 
**Inserted** | Pointer to **bool** | Inserted is whether this field has been filled in. | [optional] 
**Page** | Pointer to **float64** | Page is the 1-based page the field sits on. | [optional] 
**PositionX** | Pointer to **float64** | PositionX is the field&#39;s horizontal position on that page. | [optional] 
**PositionY** | Pointer to **float64** | PositionY is the field&#39;s vertical position on that page. | [optional] 
**RecipientId** | Pointer to **string** | RecipientID is who must fill this field. It is absent on a signer&#39;s own view of a document, where every field returned is already theirs. | [optional] 
**Type** | Pointer to **string** | Type is what the field collects — SIGNATURE, DATE, NAME, EMAIL, TEXT and the rest. | [optional] 
**Width** | Pointer to **float64** | Width is the field&#39;s width, -1 when the renderer is to choose one. | [optional] 

## Methods

### NewEsignField

`func NewEsignField() *EsignField`

NewEsignField instantiates a new EsignField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignFieldWithDefaults

`func NewEsignFieldWithDefaults() *EsignField`

NewEsignFieldWithDefaults instantiates a new EsignField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomText

`func (o *EsignField) GetCustomText() string`

GetCustomText returns the CustomText field if non-nil, zero value otherwise.

### GetCustomTextOk

`func (o *EsignField) GetCustomTextOk() (*string, bool)`

GetCustomTextOk returns a tuple with the CustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText

`func (o *EsignField) SetCustomText(v string)`

SetCustomText sets CustomText field to given value.

### HasCustomText

`func (o *EsignField) HasCustomText() bool`

HasCustomText returns a boolean if a field has been set.

### GetFieldMeta

`func (o *EsignField) GetFieldMeta() interface{}`

GetFieldMeta returns the FieldMeta field if non-nil, zero value otherwise.

### GetFieldMetaOk

`func (o *EsignField) GetFieldMetaOk() (*interface{}, bool)`

GetFieldMetaOk returns a tuple with the FieldMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMeta

`func (o *EsignField) SetFieldMeta(v interface{})`

SetFieldMeta sets FieldMeta field to given value.

### HasFieldMeta

`func (o *EsignField) HasFieldMeta() bool`

HasFieldMeta returns a boolean if a field has been set.

### SetFieldMetaNil

`func (o *EsignField) SetFieldMetaNil(b bool)`

 SetFieldMetaNil sets the value for FieldMeta to be an explicit nil

### UnsetFieldMeta
`func (o *EsignField) UnsetFieldMeta()`

UnsetFieldMeta ensures that no value is present for FieldMeta, not even an explicit nil
### GetHeight

`func (o *EsignField) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *EsignField) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *EsignField) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *EsignField) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *EsignField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInserted

`func (o *EsignField) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *EsignField) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *EsignField) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *EsignField) HasInserted() bool`

HasInserted returns a boolean if a field has been set.

### GetPage

`func (o *EsignField) GetPage() float64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EsignField) GetPageOk() (*float64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EsignField) SetPage(v float64)`

SetPage sets Page field to given value.

### HasPage

`func (o *EsignField) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPositionX

`func (o *EsignField) GetPositionX() float64`

GetPositionX returns the PositionX field if non-nil, zero value otherwise.

### GetPositionXOk

`func (o *EsignField) GetPositionXOk() (*float64, bool)`

GetPositionXOk returns a tuple with the PositionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionX

`func (o *EsignField) SetPositionX(v float64)`

SetPositionX sets PositionX field to given value.

### HasPositionX

`func (o *EsignField) HasPositionX() bool`

HasPositionX returns a boolean if a field has been set.

### GetPositionY

`func (o *EsignField) GetPositionY() float64`

GetPositionY returns the PositionY field if non-nil, zero value otherwise.

### GetPositionYOk

`func (o *EsignField) GetPositionYOk() (*float64, bool)`

GetPositionYOk returns a tuple with the PositionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionY

`func (o *EsignField) SetPositionY(v float64)`

SetPositionY sets PositionY field to given value.

### HasPositionY

`func (o *EsignField) HasPositionY() bool`

HasPositionY returns a boolean if a field has been set.

### GetRecipientId

`func (o *EsignField) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignField) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignField) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignField) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetType

`func (o *EsignField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsignField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsignField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsignField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *EsignField) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *EsignField) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *EsignField) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *EsignField) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



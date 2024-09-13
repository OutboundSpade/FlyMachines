# FlyHTTPResponseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Pristine** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlyHTTPResponseOptions

`func NewFlyHTTPResponseOptions() *FlyHTTPResponseOptions`

NewFlyHTTPResponseOptions instantiates a new FlyHTTPResponseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyHTTPResponseOptionsWithDefaults

`func NewFlyHTTPResponseOptionsWithDefaults() *FlyHTTPResponseOptions`

NewFlyHTTPResponseOptionsWithDefaults instantiates a new FlyHTTPResponseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *FlyHTTPResponseOptions) GetHeaders() map[string]map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FlyHTTPResponseOptions) GetHeadersOk() (*map[string]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FlyHTTPResponseOptions) SetHeaders(v map[string]map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FlyHTTPResponseOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPristine

`func (o *FlyHTTPResponseOptions) GetPristine() bool`

GetPristine returns the Pristine field if non-nil, zero value otherwise.

### GetPristineOk

`func (o *FlyHTTPResponseOptions) GetPristineOk() (*bool, bool)`

GetPristineOk returns a tuple with the Pristine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPristine

`func (o *FlyHTTPResponseOptions) SetPristine(v bool)`

SetPristine sets Pristine field to given value.

### HasPristine

`func (o *FlyHTTPResponseOptions) HasPristine() bool`

HasPristine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the SetEffectsEntryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetEffectsEntryType{}

// SetEffectsEntryType struct for SetEffectsEntryType
type SetEffectsEntryType struct {
	Name *string `json:"name,omitempty"`
	Id *int32 `json:"id,omitempty"`
	// true if a type is generated from the Api instead of Ankama. In that case, always prefer showing the templated string and omit everything else. The \"name\" field will have an english description of the meta type. An example for such effects are class sets effects.
	IsMeta *bool `json:"is_meta,omitempty"`
	// Affects target or source actively.
	IsActive *bool `json:"is_active,omitempty"`
}

// NewSetEffectsEntryType instantiates a new SetEffectsEntryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetEffectsEntryType() *SetEffectsEntryType {
	this := SetEffectsEntryType{}
	return &this
}

// NewSetEffectsEntryTypeWithDefaults instantiates a new SetEffectsEntryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetEffectsEntryTypeWithDefaults() *SetEffectsEntryType {
	this := SetEffectsEntryType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetEffectsEntryType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntryType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetEffectsEntryType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetEffectsEntryType) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SetEffectsEntryType) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntryType) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SetEffectsEntryType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SetEffectsEntryType) SetId(v int32) {
	o.Id = &v
}

// GetIsMeta returns the IsMeta field value if set, zero value otherwise.
func (o *SetEffectsEntryType) GetIsMeta() bool {
	if o == nil || IsNil(o.IsMeta) {
		var ret bool
		return ret
	}
	return *o.IsMeta
}

// GetIsMetaOk returns a tuple with the IsMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntryType) GetIsMetaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMeta) {
		return nil, false
	}
	return o.IsMeta, true
}

// HasIsMeta returns a boolean if a field has been set.
func (o *SetEffectsEntryType) HasIsMeta() bool {
	if o != nil && !IsNil(o.IsMeta) {
		return true
	}

	return false
}

// SetIsMeta gets a reference to the given bool and assigns it to the IsMeta field.
func (o *SetEffectsEntryType) SetIsMeta(v bool) {
	o.IsMeta = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *SetEffectsEntryType) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntryType) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *SetEffectsEntryType) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *SetEffectsEntryType) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o SetEffectsEntryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetEffectsEntryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsMeta) {
		toSerialize["is_meta"] = o.IsMeta
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	return toSerialize, nil
}

type NullableSetEffectsEntryType struct {
	value *SetEffectsEntryType
	isSet bool
}

func (v NullableSetEffectsEntryType) Get() *SetEffectsEntryType {
	return v.value
}

func (v *NullableSetEffectsEntryType) Set(val *SetEffectsEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableSetEffectsEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableSetEffectsEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetEffectsEntryType(val *SetEffectsEntryType) *NullableSetEffectsEntryType {
	return &NullableSetEffectsEntryType{value: val, isSet: true}
}

func (v NullableSetEffectsEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetEffectsEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



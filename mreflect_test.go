package mschema

import (
	"testing"
	"github.com/pquerna/ffjson/ffjson"
)

type User struct {
	//Id       int       `json:"id" bson:"_id,omitempty" jsonschema:"required,maximum=20,minimum=1"`
	//Age      int       `json:"age" jsonschema:"minimum=18,maximum=120,exclusiveMaximum=true,exclusiveMinimum=true"`
	//Uid      string    `json:"uid,omitempty" bson:"uid,omitempty" jsonschema:"required,minLength=1,maxLength=20"`
	//Pwd      string    `json:"pwd,omitempty" bson:"pwd,omitempty" jsonschema:"required,minLength=1,maxLength=20"`
	//Name     string    `json:"name,omitempty" bson:"name,omitempty" jsonschema:"required"`
	//Phone    float32   `json:"phone,omitempty" bson:"phone,omitempty" jsonschema:"required"`
	Remark   string    `json:"remark,omitempty" bson:"remark,omitempty" jsonschema:"description=test"`
	//RoleId   []int32   `json:"role,omitempty" bson:"role,omitempty" jsonschema:"-"`
	//Account  float64   `json:"account,omitempty" bson:"account,omitempty" jsonschema:"required,maximum=20,minimum=1"`
	//CreateAt time.Time `json:"createat" bson:"createat,omitempty"`
	//Method   string    `json:"method" bson:"method,omitempty" jsonschema:"enum=POST|GET|PUT|DELETE"`
}

func TestReflectFromType(t *testing.T) {
	u := &User{}
	flect := &Reflector{ExpandedStruct: true, RequiredFromJSONSchemaTags: true, AllowAdditionalProperties: true}
	sc := flect.Reflect(u)
	bts, _ := ffjson.Marshal(&sc)
	t.Log(string(bts))
}

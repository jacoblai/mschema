package mschema

import (
	"github.com/bytedance/sonic"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

type User struct {
	Id       int                  `json:"id" bson:"_id,omitempty" jsonschema:"required,maximum=20,minimum=1,description=test1"`
	Age      int                  `json:"age" jsonschema:"minimum=18,maximum=120,exclusiveMaximum=true,exclusiveMinimum=true"`
	Uid      string               `json:"uid,omitempty" bson:"uid,omitempty" jsonschema:"required,minLength=1,maxLength=20"`
	Pwd      string               `json:"pwd,omitempty" bson:"pwd,omitempty" jsonschema:"required,minLength=1,maxLength=20"`
	Name     string               `json:"name,omitempty" bson:"name,omitempty" jsonschema:"required"`
	Phone    float32              `json:"phone,omitempty" bson:"phone,omitempty" jsonschema:"required,description=test1"`
	Remark   string               `json:"remark,omitempty" bson:"remark,omitempty" jsonschema:"description=test2"`
	RoleId   []int32              `json:"role,omitempty" bson:"role,omitempty" jsonschema:"-"`
	Account  float64              `json:"account,omitempty" bson:"account,omitempty" jsonschema:"required,maximum=20,minimum=1,description=test"`
	CreateAt time.Time            `json:"createat" bson:"createat,omitempty"`
	Method   string               `json:"method" bson:"method,omitempty" jsonschema:"enum=POST|GET|PUT|DELETE"`
	Oid      primitive.ObjectID   `json:"oid,omitempty" bson:"oid,omitempty" jsonschema:"required,oid"`
	Oids     []primitive.ObjectID `json:"oids,omitempty" bson:"oids,omitempty" jsonschema:"required,oids"`
}

func TestReflectFromType(t *testing.T) {
	u := &User{}
	flect := &Reflector{ExpandedStruct: true, RequiredFromJSONSchemaTags: true, AllowAdditionalProperties: true}
	sc := flect.Reflect(u)
	bts, _ := sonic.Marshal(&sc)
	t.Log(string(bts))
}

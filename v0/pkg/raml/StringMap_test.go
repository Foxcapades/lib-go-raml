package raml_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
)

func TestStringMap_Put(t *testing.T) {
	Convey("TestStringMap.Put", t, func() {
		var k string = "4ff99860-4e01-4cb9-ba88-6e3174d33ce5"
		var v string = "2222eb87-ad87-468c-a3e0-bbff2ac9e217"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "82425411-1f19-4984-9b37-754f09d50a37"
		var v string = "bcade4b6-b419-4eea-9c1f-7663c532b9e7"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "c4cd55f4-daa1-4336-9657-d68914a24c33"
		var v string = "f1efd8d1-cdba-4bad-8e4f-d4cb9b3c406d"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3c251778-74ee-4318-86c4-e1fcb1116f43"+"e667c8aa-6f2f-4ebf-bac4-04e8705332ff"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "36095a79-2a72-4679-a649-b82c940f2d67"
		var v string = "88d83c5b-bc91-492e-b0bc-85337d35f1ce"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("dc2a5f26-76e2-416a-86ad-932430f7fd2c" + "c3a9ec23-d9c2-4368-9d18-7f82b034d427")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "490dbaab-727f-48a0-98b2-c764d050c5c3"
		var v string = "434b7bbb-78e0-4656-afce-f006fd30b8d5"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("27b150dd-8d8f-4214-92da-5e8181f4247f" + "37e60ee5-f99e-43a6-bb40-c332129e9190")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "003848a8-09ea-4578-8a16-252e1078101f"
		var v string = "5c5c2084-7a9c-4ded-860d-e3b4fcfd1963"
		hits := 0

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestStringMap_MarshalYAML(t *testing.T) {
	Convey("TestStringMap.MarshalYAML", t, func() {
		var k string = "e8ba214d-174a-4ead-a939-71e6c2324af7"
		var v string = "3f3b1981-bb7d-4546-b64a-7853ea7c6e8d"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestStringMap_ToYAML(t *testing.T) {
	Convey("TestStringMap.ToYAML", t, func() {
		var k string = "fcf781b3-ca72-48dd-9569-12a8c20e2ef0"
		var v string = "3f23d736-035a-4eb2-ad57-92613dd81856"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestStringMap_PutIfNotNil(t *testing.T) {
	Convey("TestStringMap.PutIfNotNil", t, func() {
		var k string = "0d8f9180-bb30-436c-9017-7a7f6ed74cc2"
		var v string = "19de4d04-4f0b-4bdc-b43c-af8e3aeb4ad4"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9a96a706-9bf5-4836-b5bc-3f84060f6016", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "d1bff5dc-1b94-418e-b375-c885b6418924"
		So(test.PutIfNotNil("b6e89470-7afb-4e74-bc46-90260bc9914f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "e2945eab-3ef7-4e45-bf59-2ce54b8320fc"
		var v string = "9f9883d7-3602-49a2-8915-3f3ef54bad44"
		var x string = "9dbc2275-d433-477e-bf2e-72e30c644fcb"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e8912ef1-f888-4749-b61f-37f0155014ae", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "5bf93774-c146-44a0-9561-bd1d2ac7fe8d"
		var v string = "740dce1a-2e51-4626-aa3d-fdda988ff739"
		var x string = "a31439bb-05d5-4ec2-a8c1-46f12a007126"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ac019441-0b5c-40b6-a741-4f89062af124", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_MarshalJSON(t *testing.T) {
	Convey("TestStringMap.MarshalJSON", t, func() {
		var k string = "45dec363-24ff-4570-bfe0-196a5c9a5210"
		var v string = "0c0cc061-8021-4f95-bc19-0051b16b0122"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"45dec363-24ff-4570-bfe0-196a5c9a5210","value":"0c0cc061-8021-4f95-bc19-0051b16b0122"}]`)
	})
}

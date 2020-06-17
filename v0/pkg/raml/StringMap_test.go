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
		var k string = "b75e49dc-34ba-4b35-b0bc-d1466df03852"
		var v string = "6a411539-2653-4c29-b7ee-ffefc9c82b05"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "26008166-7d6f-4215-a070-970b59804e0d"
		var v string = "b5d7aaf1-7d5d-47fe-8054-a343de6f7d27"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "12e1849a-94a4-49a7-bb9f-f78069c8351f"
		var v string = "c8ea4fc4-5669-4815-bb78-5286f17bd37f"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("236518eb-cfad-47b5-b323-1d7883626d2b"+"7e235f0d-3aaa-4325-96e6-26e110a1050e"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "ab00e6be-1df0-4872-867a-f8e126e70869"
		var v string = "f49fcdf2-17ba-4e6e-9d98-2c19b0c6df6c"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b0815f27-b8ab-4bb2-82bc-e200c2656ef5" + "5586921b-3016-4559-a184-8450c46215e4")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "5ce40514-641e-47e7-ac08-7187c251ee96"
		var v string = "65f2a68f-b76c-4137-83ea-e834b9e3f066"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("12ed8d62-19a8-4b2f-b496-42aea4dc0af4" + "f4409275-c413-4f2c-95a6-30a21172034e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "c56bbb43-cfb2-44cf-8008-1ee247748646"
		var v string = "0ba007d5-27a0-4281-b9f3-247954edac15"
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
		var k string = "f1edab6c-dfa4-4601-bcd2-05f12710cdce"
		var v string = "dc4f8640-567c-4e14-9b9b-cbf74f37ebd0"

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
		var k string = "28b90d4c-4fce-43a7-b46f-5361727fbb6b"
		var v string = "b88e652b-8599-4905-bd15-cc494afa615e"

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
		var k string = "db96b2c8-2d30-4b13-a8d5-f4c9822f6122"
		var v string = "ef810a17-852c-4efa-bb90-2602e3642237"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5e3de7fa-b3d6-4dcf-b82d-f5d075f12a91", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "af0f96c6-87f9-4e50-b59e-f2d4988cc6ad"
		So(test.PutIfNotNil("09946465-79f0-4e4e-90ba-8ef1929314f7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "f14b83e4-00d2-435f-84e4-79df0b8df64b"
		var v string = "e179c393-8c7f-40d2-9e18-3686b50184e7"
		var x string = "12a5bcdb-b4a2-42f1-8d7e-9ca56730b6a2"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("fcaa2a50-2cd8-4276-a3e6-7ac0740cf132", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "36b53333-4f16-4e98-8aa3-f55bc70e36c9"
		var v string = "325f7da0-6185-4ee4-918e-12e4f766dde1"
		var x string = "458761ea-506f-42dc-b081-d39c1dbac5e8"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("71680655-4e89-4114-9dd8-63a4c9373a17", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_MarshalJSON(t *testing.T) {
	Convey("TestStringMap.MarshalJSON", t, func() {
		var k string = "f8e616cc-39a0-4e33-9224-521d87741f7a"
		var v string = "9ba1d337-2d5e-453e-85fb-87cfc6ec8394"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f8e616cc-39a0-4e33-9224-521d87741f7a","value":"9ba1d337-2d5e-453e-85fb-87cfc6ec8394"}]`)
	})
}

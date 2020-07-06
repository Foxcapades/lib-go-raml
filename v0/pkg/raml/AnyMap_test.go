package raml_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
)

func TestAnyMap_Put(t *testing.T) {
	Convey("TestAnyMap.Put", t, func() {
		var k interface{} = "3f887f96-3f54-46a6-ba6e-fef7750cc25c"
		var v interface{} = "133ba455-6770-4ff4-b73e-ae8cd242ed5e"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "cfa6ba2f-9985-40a8-9836-fab6a1fa8b50"
		var v interface{} = "e419fe6a-0538-41bc-ac68-48a003f5c632"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "85c67a36-a21a-457e-8ea4-97e1896db31c"
		var v interface{} = "6224fcc6-0f0f-4347-ba4e-c122908dba64"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("dbf60dac-ab5b-4769-aa97-26863a33ebdc"+"4e9865d4-5c8a-46ef-a434-de58eec72105"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "6702a47e-322b-4a50-8482-bc1ea36c1290"
		var v interface{} = "be7b65ad-3ab4-4ada-8318-3a1b0539d23d"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("43f86337-027a-4aa2-96df-98741351c438" + "0244b1be-57e2-4a22-bc00-b8e9731d31a7")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "29f1efce-3ec9-4df4-80a5-8639ad9e4cf9"
		var v interface{} = "d988176b-ea6e-4e5f-a010-2b4c8383b997"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("68543fa5-aeb7-437f-ab6b-8aafd351161a" + "02c8d25d-ac0a-4bf3-a515-8130da369d4a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "d24a8eee-2676-4a9f-b7f4-a4bcd5dde63c"
		var v interface{} = "b8ba4dae-9eac-42bb-ab4a-72fd61c15809"
		hits := 0

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestAnyMap_MarshalYAML(t *testing.T) {
	Convey("TestAnyMap.MarshalYAML", t, func() {
		var k interface{} = "23dae8e3-441a-4d3d-919b-d33507170646"
		var v interface{} = "3dbfcb2f-4fe4-45e4-8407-1dc0c97ec82e"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestAnyMap_ToYAML(t *testing.T) {
	Convey("TestAnyMap.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "bfeaa7da-1b4c-4b67-8b12-201038eb2a50"
			var v interface{} = "4042160e-8bb6-475b-84ff-a750ff5e3999"

			test := raml.NewAnyMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k interface{} = "f2c05401-3f67-450b-8eeb-b01ede739f0a"
			var v interface{} = "6135b9e9-dbbc-45e3-b041-2f0590b9f70f"

			test := raml.NewAnyMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestAnyMap_PutIfNotNil(t *testing.T) {
	Convey("TestAnyMap.PutIfNotNil", t, func() {
		var k interface{} = "044fa5d2-58de-4ff3-a999-5d33e6477571"
		var v interface{} = "029d5d4e-ea7f-480f-ac4b-3e36afb114e0"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("fe7a5c03-f061-4438-906b-6d72afd1bce7", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "a8148f3b-59ff-45b0-b789-2bf69b1830d6"
		So(test.PutIfNotNil("04fe2ec5-cbaf-4951-b00d-6a1f52621c2b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "ca4345fd-8f71-4ec3-a3b6-6646ee3eade1"
		var v interface{} = "e1b5610c-14be-4777-bdab-cd9e2abae041"
		var x interface{} = "5700cbac-0206-460d-9a5e-81d1a57a531d"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a87adef2-daf4-4391-9519-31161652d0a0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "922fc1b2-e14f-4a63-b12a-07c2d1fb7b6f"
		var v interface{} = "e4390c0f-cd92-469c-9e57-5b1c0b62960a"
		var x interface{} = "39a2bdbb-3814-4fc2-bb5c-f182aa29992d"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("947eb900-b961-4bb7-8237-26373b159227", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "4e6511b8-c19b-46b8-b588-21aa8ade4d65"
			var v interface{} = "b67607dc-5709-4a06-a1dd-1d69f1382375"

			test := raml.NewAnyMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"4e6511b8-c19b-46b8-b588-21aa8ade4d65","value":"b67607dc-5709-4a06-a1dd-1d69f1382375"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "4e6511b8-c19b-46b8-b588-21aa8ade4d65"
			var v interface{} = "b67607dc-5709-4a06-a1dd-1d69f1382375"

			test := raml.NewAnyMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"4e6511b8-c19b-46b8-b588-21aa8ade4d65":"b67607dc-5709-4a06-a1dd-1d69f1382375"}`)
		})

	})
}

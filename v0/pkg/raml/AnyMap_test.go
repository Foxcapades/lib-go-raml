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
		var k interface{} = "1123610f-2ff0-4f55-bb33-8c1c13291db4"
		var v interface{} = "045b72c5-34c7-49c5-8d91-4119eb552d2d"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "d4b803db-a0d5-4cd4-ab05-0812a0d45334"
		var v interface{} = "c535b351-73f6-40b7-8faf-798e0e425eab"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "98ba83ef-4604-4374-9533-2222208713b1"
		var v interface{} = "cec09ec0-c7e9-459a-8b96-c74872fad7db"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b65e0319-4574-4d2c-a35d-ecc578990790"+"7a408d7c-0b48-43c8-8cfc-1417f479a0a3"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "15c77164-b6c8-4b3d-891a-6d0fdbe331a6"
		var v interface{} = "d9291357-3526-4572-9b31-e3f883d4ce39"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("787ddfc7-3704-40e0-8a0a-ea946e2add0e" + "64a46c45-ef12-4e26-9f37-e058f25bc9ad")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "df942fae-96ab-4a4b-b557-c8dea07a6961"
		var v interface{} = "6492abff-668d-418b-8cc3-198e6f2c016f"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("faf4c784-fa3a-4ca3-a07f-ec0e0fc855cc" + "bad63e4e-c95a-46a1-a58e-2818728aa01d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "34039cdb-1adf-4517-8022-46523f00d053"
		var v interface{} = "042d570d-78fe-48cc-bf71-1693bfe53843"
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
		var k interface{} = "0d45bda0-6a72-4969-a859-c742b6a66501"
		var v interface{} = "656fc93a-5538-455a-b643-08ad4f876e3e"

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
			var k interface{} = "a36feb57-2377-4b82-b51e-9676f446244f"
			var v interface{} = "feaa4f4d-2374-4229-8341-ed5f21f9a6b3"

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
			var k interface{} = "4bf88849-bd76-45ff-b2f2-5030764611bd"
			var v interface{} = "4b913a3b-f183-4352-8441-ed9e2b6b5c38"

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
		var k interface{} = "4307279a-f46f-4b0f-ac58-cd33de090ab7"
		var v interface{} = "206e0f27-bafd-48c6-bdbc-eda3435f309e"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("20a7b281-5aef-4618-ae44-b432abbe12c5", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "5daace89-8309-4257-bccc-5658f80da23a"
		So(test.PutIfNotNil("8e43103d-d2a3-432b-a438-ce5fa591945c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "d693ff3d-726a-4c7c-b314-551290dd2cb4"
		var v interface{} = "6ebf2d4c-cc82-45fd-896a-b3a803128f24"
		var x interface{} = "80416915-0303-41f9-aac8-5bb049962eba"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6e2aaf5f-7ea0-406e-9f94-562ff57665a1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "e7be5f85-41fa-45bf-9ea9-88247948ceac"
		var v interface{} = "59633c63-c4a7-4bc5-8555-3a4dc088f5df"
		var x interface{} = "9c9a0382-e9e1-4060-a2dc-6d2e2aec4483"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("6aa006db-e7dd-4018-9af3-8bae0ad0bd2c", x), ShouldPointTo, test)
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
			var k interface{} = "4b4f11a6-0dc1-4d5e-95f1-f46f07225ba7"
			var v interface{} = "dee7e0ff-8232-453e-9fd3-648e55ae07e9"

			test := raml.NewAnyMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"4b4f11a6-0dc1-4d5e-95f1-f46f07225ba7","value":"dee7e0ff-8232-453e-9fd3-648e55ae07e9"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "4b4f11a6-0dc1-4d5e-95f1-f46f07225ba7"
			var v interface{} = "dee7e0ff-8232-453e-9fd3-648e55ae07e9"

			test := raml.NewAnyMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"4b4f11a6-0dc1-4d5e-95f1-f46f07225ba7":"dee7e0ff-8232-453e-9fd3-648e55ae07e9"}`)
		})

	})
}

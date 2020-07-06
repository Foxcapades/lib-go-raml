package raml_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
)

func TestAnyExampleMap_Put(t *testing.T) {
	Convey("TestAnyExampleMap.Put", t, func() {
		var k string = "429cfffb-c7da-4492-b95c-2bb60e5a3f88"
		var v interface{} = "dc49dbbc-f410-45b8-b3f3-362ddf6c5c30"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyExampleMap_Delete(t *testing.T) {
	Convey("TestAnyExampleMap.Delete", t, func() {
		var k string = "801701a7-a189-482f-856b-1ee7b7139d8b"
		var v interface{} = "cd622982-919d-47c3-96cf-80c3bbe4437c"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyExampleMap_Has(t *testing.T) {
	Convey("TestAnyExampleMap.Has", t, func() {
		var k string = "7d7a5560-7501-4b40-953f-3d1c7f9c6869"
		var v interface{} = "4e9ebc0e-42e2-40e8-a8d7-b2d5d08269a8"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("187be80a-f901-48db-a52e-51d9026db0e2"+"f5495fea-4bf6-48ec-92c7-6b99cfd33f48"), ShouldBeFalse)
	})
}

func TestAnyExampleMap_Get(t *testing.T) {
	Convey("TestAnyExampleMap.Get", t, func() {
		var k string = "02ed5fde-98ff-4fd6-88bb-43b2661fc254"
		var v interface{} = "96e8bc4e-cc1b-4b31-89d6-fed9cedfee4c"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5b3e56a7-7cd0-43c1-8d87-6840eb218dd6" + "74500c24-a06c-454a-8a97-9ccaea0cf5dc")
		So(b, ShouldBeFalse)
	})
}

func TestAnyExampleMap_GetOpt(t *testing.T) {
	Convey("TestAnyExampleMap.GetOpt", t, func() {
		var k string = "99a7ff3f-be18-4723-be48-619bf35f2750"
		var v interface{} = "a6515a80-97b2-4794-b0b8-bc203208e197"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0058aeec-8161-452c-a78a-47a9b63b76d3" + "5dba048a-1940-4e65-ab23-bb2fab54116c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyExampleMap_ForEach(t *testing.T) {
	Convey("TestAnyExampleMap.ForEach", t, func() {
		var k string = "f5e1325c-360c-4009-9e68-444febec9671"
		var v interface{} = "c2fd3ef5-9763-46af-956d-53dfd65a8725"
		hits := 0

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestAnyExampleMap_MarshalYAML(t *testing.T) {
	Convey("TestAnyExampleMap.MarshalYAML", t, func() {
		var k string = "df61a4fa-45cf-4607-aedb-67450e2869c0"
		var v interface{} = "1047ff26-59be-494d-aab7-3ddc8fdbd4d3"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestAnyExampleMap_ToYAML(t *testing.T) {
	Convey("TestAnyExampleMap.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "130a9f48-7fb3-4ddd-97a0-c44487eadc82"
			var v interface{} = "08a4b664-4d38-4092-9df0-4de58b1dbd9d"

			test := raml.NewAnyExampleMap(1)

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
			var k string = "d2ada575-5262-4179-8a34-e74d38d8865f"
			var v interface{} = "6c0746ec-bcc3-444a-88b3-efd076327176"

			test := raml.NewAnyExampleMap(1)
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

func TestAnyExampleMap_PutIfNotNil(t *testing.T) {
	Convey("TestAnyExampleMap.PutIfNotNil", t, func() {
		var k string = "869e34be-72ce-412a-bab7-b53c61642c7f"
		var v interface{} = "8919565a-1c6b-48b4-b0c8-a3629b032f11"

		test := raml.NewAnyExampleMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c4c37a47-90b6-4005-8138-91b5449d5803", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "08153e01-abe3-47dd-8e82-02544a2b09b8"
		So(test.PutIfNotNil("d303635e-bff4-4ca1-b52b-4b2971c55a55", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceIfExists", t, func() {
		var k string = "063f5e73-e037-4e65-80f5-d00197192c1a"
		var v interface{} = "b3c787a1-3f9d-4148-897b-c36d12815a51"
		var x interface{} = "30a1cfa8-2258-432b-a497-c28c8a4da901"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e7c8e7c5-cb96-405b-817f-cf7494c53b10", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceOrPut", t, func() {
		var k string = "8771abd5-9885-4eba-9441-417009d1f728"
		var v interface{} = "aa927520-4a40-4732-aeb3-6b23fadd2215"
		var x interface{} = "f2d21945-b3e2-4886-b563-73fca933053a"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3f9fcec3-1234-4e4d-921c-599deaa40830", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyExampleMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "d25bf5a7-a7ee-4098-92fa-5d1080c87cde"
			var v interface{} = "286360c1-8818-4ad7-b2b8-10f5fdf08b03"

			test := raml.NewAnyExampleMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"d25bf5a7-a7ee-4098-92fa-5d1080c87cde","value":"286360c1-8818-4ad7-b2b8-10f5fdf08b03"}]`)
		})

		Convey("Unordered", func() {
			var k string = "d25bf5a7-a7ee-4098-92fa-5d1080c87cde"
			var v interface{} = "286360c1-8818-4ad7-b2b8-10f5fdf08b03"

			test := raml.NewAnyExampleMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"d25bf5a7-a7ee-4098-92fa-5d1080c87cde":"286360c1-8818-4ad7-b2b8-10f5fdf08b03"}`)
		})

	})
}

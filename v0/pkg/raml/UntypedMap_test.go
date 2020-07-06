package raml_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
)

func TestUntypedMap_Put(t *testing.T) {
	Convey("TestUntypedMap.Put", t, func() {
		var k string = "42a7b8ff-b445-43da-aedf-56fc6cd32054"
		var v interface{} = "49cf3a93-e0f6-4939-88d1-caf1f04a6820"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "41839e16-3ca2-4545-8f73-056fd248beed"
		var v interface{} = "0d8f1eda-ae2c-45d2-849f-2ec42da30e0a"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "a6ca2be7-dd40-4355-8dcd-c960eaf04dd7"
		var v interface{} = "dc5530ad-91e2-4d90-8145-66247cc35ea3"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("575ae8b2-71bc-4a65-82d8-0d6609836917"+"ffb44ef1-a911-4d3b-bfa5-050625a1af7e"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "e97953de-1ac0-4a9b-8181-3f7a52a6b281"
		var v interface{} = "0523b154-4554-455f-a674-abf1bbd78974"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("369cda83-67e4-4d51-bf3b-53e2328c5f4b" + "1dac760d-9ec1-4277-b146-7e1454b70349")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "fd609df5-3708-4781-8b28-104c52fbacb4"
		var v interface{} = "547d1ed4-858f-439f-b759-4cfee1daa688"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f40de22f-65e9-4873-9c08-64ab9c569de6" + "06b34797-e9bc-4f06-801d-baaf27129b08")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "adfb42b6-6c76-48d8-a063-429436e34f5c"
		var v interface{} = "2578af20-35d3-43b8-8cb4-40c1f01f6b36"
		hits := 0

		test := raml.NewUntypedMap(1)

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

func TestUntypedMap_MarshalYAML(t *testing.T) {
	Convey("TestUntypedMap.MarshalYAML", t, func() {
		var k string = "6d2b81c6-3cd4-4c14-9bba-e54bd4f2dab5"
		var v interface{} = "e6ca433f-e47e-4121-a519-216292076694"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestUntypedMap_ToYAML(t *testing.T) {
	Convey("TestUntypedMap.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "f7e2b825-b225-42d2-b9ed-e9a6f84f1eab"
			var v interface{} = "b9e542bd-286d-4960-a97b-bf20a6e7faec"

			test := raml.NewUntypedMap(1)

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
			var k string = "0538ff15-92ba-4c49-b39d-d6abfede10e4"
			var v interface{} = "bcc29934-0ea3-4594-a74f-2edbafeb3187"

			test := raml.NewUntypedMap(1)
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

func TestUntypedMap_PutIfNotNil(t *testing.T) {
	Convey("TestUntypedMap.PutIfNotNil", t, func() {
		var k string = "64774918-def1-4d49-958a-4722db0f83fc"
		var v interface{} = "9f66f179-8c77-440c-a2a2-5c4eeb1924a0"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1242ce04-88dd-474d-b115-16e9d07f1fc0", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "13fd93e0-1897-4f1d-8fcc-6b4a2848b6ec"
		So(test.PutIfNotNil("dc911361-e8d0-4b16-812e-282219017c9f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "c47856c8-975d-4ee6-8f28-8089ebb53929"
		var v interface{} = "8a986d48-294c-4d51-bf22-b082323edd41"
		var x interface{} = "f7f5b71e-1d7a-4b9b-9629-977e6a94baf2"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("641603bb-ffbd-4cad-bb0b-bcf658fd47bb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "ab1e5a8b-8f64-4927-950d-f3d4f17f7d7f"
		var v interface{} = "1c2e5df0-7b92-4c66-9053-0b4efacb3f18"
		var x interface{} = "3f746b95-13cb-4fbf-a362-a668abd29ac7"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("29fc2cd9-7ab7-41ca-be87-03d0837e5df8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_MarshalJSON(t *testing.T) {
	Convey("TestUntypedMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "644ab16b-2267-4663-ba9f-c1c9623f3f75"
			var v interface{} = "d8a03412-bda2-4a90-ac7c-644657f226a5"

			test := raml.NewUntypedMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"644ab16b-2267-4663-ba9f-c1c9623f3f75","value":"d8a03412-bda2-4a90-ac7c-644657f226a5"}]`)
		})

		Convey("Unordered", func() {
			var k string = "644ab16b-2267-4663-ba9f-c1c9623f3f75"
			var v interface{} = "d8a03412-bda2-4a90-ac7c-644657f226a5"

			test := raml.NewUntypedMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"644ab16b-2267-4663-ba9f-c1c9623f3f75":"d8a03412-bda2-4a90-ac7c-644657f226a5"}`)
		})

	})
}

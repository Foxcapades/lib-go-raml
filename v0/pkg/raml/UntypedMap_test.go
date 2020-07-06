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
		var k string = "983f30aa-39aa-46be-b425-9aa6fde6a699"
		var v interface{} = "bc3620ff-aca8-48e1-99dc-52e4e389a684"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "629cbe8f-cb91-4083-96a9-d88f2e53e395"
		var v interface{} = "871d76a9-8eea-470d-9d2e-7ecc85107f48"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "4a621feb-c7e7-416e-a9ea-5aa1eb231881"
		var v interface{} = "626020f7-9053-4fd6-a572-a3eec027e2a5"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4632ea87-7fc6-4793-92cc-b02fb3caaed6"+"877d9464-2c60-4f15-8813-5a5a9f1d6698"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "07d3b788-69e3-44d5-89f7-b04294a2b1e8"
		var v interface{} = "504ee68f-8f5d-4bcb-8603-8685b15a5acb"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("17d86faf-704d-4488-8297-3b75b7e92d85" + "2d6e7c13-3526-4908-a234-63ebe935c05c")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "bf07b8ee-39b8-4693-9cec-76a8c9d3b090"
		var v interface{} = "39a65e35-1871-4c39-9fc5-3da837a21b29"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a196dd6d-a2cb-4d0c-9632-6741a2c7d0ea" + "815b7bc8-6321-4542-ba4c-f13ee18cf200")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "e8de4772-07d1-4302-9ad2-c7666761da5a"
		var v interface{} = "c6be034c-ea5e-4f7c-bf09-6e7d4b0ee977"
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
		var k string = "107ed53e-4f00-4f1b-9990-7e765fbf73a3"
		var v interface{} = "3f3d367a-cc35-4816-b772-5982f7e83b3f"

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
			var k string = "71fa87d9-9721-46e9-bc5d-eaec857ead48"
			var v interface{} = "55db224f-1ddf-4aae-a5f3-6ee6a9ee12cf"

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
			var k string = "7e9826ac-65de-483b-ab24-f20fa56469fc"
			var v interface{} = "5d4656a9-d9f1-43fb-808f-b92fe015ee6e"

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
		var k string = "bf6b8e0d-9cf2-40b4-97bd-440e45a4448d"
		var v interface{} = "c2ff4206-70fa-462b-aaf4-5299ce8280d3"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a23df5ca-c3c8-4129-80f1-4692131b4ff7", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "5aced314-43d3-4e99-b778-15d53b0a8ebf"
		So(test.PutIfNotNil("119c4a5a-baea-43aa-ae34-ae0dcccc5aa1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "9f1ef931-9360-4f0d-9ba5-5433fa244b75"
		var v interface{} = "0fba9857-28ec-437c-9b19-4cd729e5d61c"
		var x interface{} = "357adebb-b1c9-4145-b024-f0d03882a915"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c27c5d5b-2e67-4684-b06c-bbdc3e722254", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "5df1dd6d-d2ea-4a9c-98b0-781a7da92760"
		var v interface{} = "caf1d7c3-f9b2-4187-98e9-57ec17fd6be2"
		var x interface{} = "215c1b92-1608-47ce-9bb2-3cfcb7e44996"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("84ecdaa7-5564-457c-ac81-2dad13e340c8", x), ShouldPointTo, test)
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
			var k string = "a4df88b7-b621-44ff-a223-e9cac97289fc"
			var v interface{} = "53f14465-3748-4f48-abae-a61d8bafa2ac"

			test := raml.NewUntypedMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"a4df88b7-b621-44ff-a223-e9cac97289fc","value":"53f14465-3748-4f48-abae-a61d8bafa2ac"}]`)
		})

		Convey("Unordered", func() {
			var k string = "a4df88b7-b621-44ff-a223-e9cac97289fc"
			var v interface{} = "53f14465-3748-4f48-abae-a61d8bafa2ac"

			test := raml.NewUntypedMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"a4df88b7-b621-44ff-a223-e9cac97289fc":"53f14465-3748-4f48-abae-a61d8bafa2ac"}`)
		})

	})
}

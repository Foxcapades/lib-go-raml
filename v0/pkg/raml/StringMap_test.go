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
		var k string = "e444af9c-0854-4a83-9503-0883280c4075"
		var v string = "47947634-9f0e-4275-9017-deffd3b4f853"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "7ca364a8-622d-4330-81dd-f647e843fed7"
		var v string = "9891d7a8-aaba-45fe-871b-2b8b7379f1b2"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "808f413d-cd22-406a-840b-071d862a1813"
		var v string = "dcb00652-d953-4cc4-b74b-0072283f4494"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7e5c7195-27c6-4373-8c67-dda6a0e45f2f"+"a09fea6a-6bc4-4614-b606-fc16f27d67db"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "dd24ed3d-5ea9-4ec9-9e70-83b19745f26d"
		var v string = "826f005b-8e8c-4899-ba8b-fd3ef2e76c6f"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("17612954-e2c1-43b0-9d64-de2bbb1bd88c" + "d4871f25-801d-4eeb-83b0-54b769f22c05")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "434ae8b7-07b0-4835-b348-82bcde91e8b3"
		var v string = "d47ecadf-a2d7-42be-8baf-8472ac4c097f"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ae0ea688-960b-4773-9651-2a1e11748011" + "880f4e6f-3c9a-4d87-80b0-808f7a180d2f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "a29a907f-51aa-47a2-9c8e-cd96e650705c"
		var v string = "ca8f7d01-5111-448a-9460-2a7195f40fd3"
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
		var k string = "feb93fd2-51b0-4fca-b224-d0232c80d853"
		var v string = "5df57acb-cdf4-4215-91a0-8202e9ac1754"

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
		Convey("Ordered", func() {
			var k string = "b777905a-bfa7-4e82-8941-ee289eb1a8ea"
			var v string = "882ac1cd-cddf-429f-b338-59aff54f65f9"

			test := raml.NewStringMap(1)

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
			var k string = "ddf5aa27-7889-4fdf-b22c-684a73a8ea5f"
			var v string = "3b702aa9-ac8b-4e00-8c5f-424245b6286f"

			test := raml.NewStringMap(1)
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

func TestStringMap_PutIfNotNil(t *testing.T) {
	Convey("TestStringMap.PutIfNotNil", t, func() {
		var k string = "df01f7ce-4909-4f61-b63d-697fab9583bc"
		var v string = "cf057da5-33ee-4f5b-990a-a20b26d2e65c"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e4057882-6cea-4218-b422-962a6b3a75a1", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "71909de1-f9c1-4c55-9233-3dbe10752332"
		So(test.PutIfNotNil("b7e715d5-14de-4f67-8df9-a901cbf34ec7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "47218abd-a897-4c2a-8648-3de743f409bc"
		var v string = "8e312660-dff1-4b1e-9509-3b3b9e81091d"
		var x string = "b6bdf158-3d33-4505-ad12-65ba5cebfe82"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("257b52b3-042a-439f-a30c-bf92d7ab4832", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "5a847a5d-7882-4112-8c66-af1f091b8427"
		var v string = "c5a97ff8-1543-4ae1-8012-5c516f34e19a"
		var x string = "9b917565-e609-4c06-b3b4-1062bf88dc9b"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("45ecd6c3-9cbb-4b62-8cd8-9e8de90a304e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_MarshalJSON(t *testing.T) {
	Convey("TestStringMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "824d4556-9604-44c2-b2f6-b24cf6a59d9c"
			var v string = "ed8a486d-77db-4983-824f-917a42bd234d"

			test := raml.NewStringMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"824d4556-9604-44c2-b2f6-b24cf6a59d9c","value":"ed8a486d-77db-4983-824f-917a42bd234d"}]`)
		})

		Convey("Unordered", func() {
			var k string = "824d4556-9604-44c2-b2f6-b24cf6a59d9c"
			var v string = "ed8a486d-77db-4983-824f-917a42bd234d"

			test := raml.NewStringMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"824d4556-9604-44c2-b2f6-b24cf6a59d9c":"ed8a486d-77db-4983-824f-917a42bd234d"}`)
		})

	})
}

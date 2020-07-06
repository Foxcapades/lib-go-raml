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
		var k string = "83dcfcd1-1f67-46d6-a809-b62ebc79b27f"
		var v string = "bdb146e9-52b1-4c03-8c8a-1aa11c8d1062"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "93ba6349-936d-44e9-bacc-739ea79e9c6b"
		var v string = "fc867c62-2640-4506-93c8-1e852da24748"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "a5800be5-407c-4bc3-8994-55680d8ad8d5"
		var v string = "dddab907-136e-430d-8dbf-894c2da22fe6"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("444db6ab-9c5f-4e00-878a-b6391f8016dd"+"0e024692-caa9-4527-b8c3-06bd1307daa3"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "92ebff95-40a6-4d7d-89da-36b65a8bc70d"
		var v string = "5be2a6ca-60fd-49d3-8e0a-b2e172395f02"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("6c27fce7-a70f-403f-93e3-24246525728e" + "52e8b8f9-ff1c-45dc-a589-ef46d6ff8d8b")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "085f070d-52c9-4847-87b7-e936bb1bc010"
		var v string = "f578bcb6-b453-4dab-96f4-6ca3f707830c"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("aafd94e7-4e4b-472b-91d8-e2a58cf70eb0" + "efe03e7a-8dde-46e4-a2de-8289ef4bcdba")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "c09cfc15-ccb3-46ed-9e63-19752617f919"
		var v string = "032d6998-089e-465e-9b8e-012666c494f8"
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
		var k string = "af53038e-df4d-4b7e-bc63-eddfbf0a03cb"
		var v string = "470aa940-faf2-46c6-aca9-5e9836fff317"

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
			var k string = "4e4b7a97-2db4-411c-b086-b8d5dd437c1b"
			var v string = "fa00aea2-aad7-4b09-8d5f-158d12c4edd2"

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
			var k string = "8d9d1a0e-957a-46bc-9ac9-24a9f6f1f080"
			var v string = "43a59877-433a-4547-bb7b-b4580f9dc340"

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
		var k string = "df8fc036-0bbc-4cc4-8b73-5ea6257c142a"
		var v string = "be1c4831-9b51-473e-ba43-0d655f0b4250"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("72733652-ec78-4a98-a35f-62f181c48d11", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "2d44d778-488e-4cbb-b847-4c8621eebbf9"
		So(test.PutIfNotNil("52628fbf-b891-4e2a-a81f-5f68ad7ff865", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "407cd497-76b8-4541-a1d8-a2f2440a91dc"
		var v string = "cc187158-5a19-4224-bcf3-e3232cdc73ad"
		var x string = "bbd1d0f7-384d-4cd9-b325-e248a768f692"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("48bfd02c-8419-43dd-a375-f320efcc6e40", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "f9088b43-d068-464d-9042-21c68a24b870"
		var v string = "69cece6e-7683-4666-9a79-ab56ac253bd6"
		var x string = "d9816240-e52d-47fc-9939-1a1229d2c81f"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("99dfa5f9-a199-48d1-b9fe-9385ef83efe1", x), ShouldPointTo, test)
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
			var k string = "1bbac07f-0a53-4be1-aac8-ff459658f5ac"
			var v string = "0ec14d1f-a296-4151-8c6c-2de6c0c0ff10"

			test := raml.NewStringMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"1bbac07f-0a53-4be1-aac8-ff459658f5ac","value":"0ec14d1f-a296-4151-8c6c-2de6c0c0ff10"}]`)
		})

		Convey("Unordered", func() {
			var k string = "1bbac07f-0a53-4be1-aac8-ff459658f5ac"
			var v string = "0ec14d1f-a296-4151-8c6c-2de6c0c0ff10"

			test := raml.NewStringMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1bbac07f-0a53-4be1-aac8-ff459658f5ac":"0ec14d1f-a296-4151-8c6c-2de6c0c0ff10"}`)
		})

	})
}

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
		var k string = "25879baa-6493-4668-812b-19eee3a6829c"
		var v string = "4f092f42-27dc-4b0d-a850-e79479e7aa3b"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "498b7356-6f5f-4911-945a-a436dba87a91"
		var v string = "8f33f94f-198a-42d8-ab4e-788951b1a676"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "0f298b54-9587-4413-a673-4abfcbf993fc"
		var v string = "4767ad0d-5ccf-4694-8e21-c2b6bbab798a"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5552b153-75cd-4e3d-ae1a-e7001a418e20"+"bebf682a-15df-467e-a767-70cc4dfb8f0f"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "f530d51c-1311-4ccb-86be-a43a70ab6b1f"
		var v string = "5b04a19c-ca19-41b0-af53-83dd3ee93eb5"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("51934d4a-dbec-4f87-acad-0e6c263c082c" + "3d265c56-6781-447a-885a-45df68d319ac")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "9814de78-692b-4223-bb0b-7dd64719f2a6"
		var v string = "347c3cf2-cd24-41b3-b28c-76fb6bd98468"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("60a898ce-fa85-482f-a5a3-e2270476b3a5" + "15f76f5c-46e2-4e8b-ae0c-3798b038d0a9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "d6f7e300-d118-4da7-9495-6789b234f382"
		var v string = "ce5b1c12-b138-4c9c-aa2f-24eb8eceaed5"
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
		var k string = "7deaad1b-ef81-48e2-b9a9-0a4596366bb1"
		var v string = "406c7c0a-8150-4dcd-ab11-bd05e8900471"

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
			var k string = "802e70fd-6360-49f6-ac23-fd7284dc9944"
			var v string = "00f80cc0-acf7-43aa-8bac-ecd11634e778"

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
			var k string = "269b30ad-bb11-45c1-8338-285e9de1bbbc"
			var v string = "80ab0f66-2595-4908-acd6-45a68817d41e"

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
		var k string = "c78e3771-5b3a-4863-9a0a-e69260e3f521"
		var v string = "afd7e14c-c2c1-465d-9fb2-e1f5a36f858e"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("88b5e5b7-3df0-4900-ad33-b0cff51d9736", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "59c0b37e-40fe-4c83-8dba-2f83c42111c0"
		So(test.PutIfNotNil("eb1046a5-8c93-4f0d-9101-8996057f6abc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "04b47b25-b45d-470e-bf8a-6e2090d35dcd"
		var v string = "14a05088-9af0-4fbe-bfda-695a6994efa8"
		var x string = "cbc0eb04-8212-43e6-b4de-d62964e2b2d5"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f574dd8c-e541-4a49-bb24-6008301c0e47", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "d7dd373e-00b3-459f-ad04-d6afa39e4998"
		var v string = "e7e370d8-a566-45a0-9624-c211d0c65790"
		var x string = "518bd755-2571-4003-905d-0090577b096f"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d173d174-396c-4e9d-bc5b-59aa5d52a9d4", x), ShouldPointTo, test)
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
			var k string = "7f9efd42-938a-47c0-988a-725163324e38"
			var v string = "770ee045-bb8f-4d69-83b7-3a9e35214672"

			test := raml.NewStringMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"7f9efd42-938a-47c0-988a-725163324e38","value":"770ee045-bb8f-4d69-83b7-3a9e35214672"}]`)
		})

		Convey("Unordered", func() {
			var k string = "7f9efd42-938a-47c0-988a-725163324e38"
			var v string = "770ee045-bb8f-4d69-83b7-3a9e35214672"

			test := raml.NewStringMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"7f9efd42-938a-47c0-988a-725163324e38":"770ee045-bb8f-4d69-83b7-3a9e35214672"}`)
		})

	})
}

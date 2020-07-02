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
		var k string = "07342681-f46f-4f27-875b-777ac3b1c140"
		var v interface{} = "19ae1297-0c96-4f33-8c46-6f36bbf54de2"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "dd8e6113-be1b-4b0f-9e80-4f33d63c14fd"
		var v interface{} = "f4a7137b-4cc8-4d7f-ae55-945d56144d5f"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "4cb5b928-d9a4-4c46-96ac-07ee17016e4f"
		var v interface{} = "cb2fc1c0-b0d7-4e15-8df7-c6c8fd656139"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("83e106ee-0b2d-4c09-8d7c-d6aab95d1e8b"+"bc68105d-c633-41bb-9dbe-484bceb49500"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "abae380f-02e1-427e-aa60-4429c94d0aa2"
		var v interface{} = "ad5e4137-31f0-42ac-8370-96a45597525d"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("d381169f-97e1-473c-a381-52e885d05f89" + "e742e6f3-8643-46df-97e6-d922ac40aceb")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "9afd41eb-67a5-45ef-a02d-7cc142b18d21"
		var v interface{} = "10b6c463-659f-4e6e-acb2-0972e796f14f"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("9d3644bb-4d1b-4c71-986e-437615e52377" + "417f8e3c-b837-494c-9183-71ce05c4a499")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "33d55c09-8d56-486c-8f72-776e7fecb96d"
		var v interface{} = "563f2e5d-e566-4c96-bb93-94f99514c28e"
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
		var k string = "900485a8-fda2-4818-b728-65215c601e41"
		var v interface{} = "262738ea-f754-49fb-809d-2894a4328f7a"

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
		var k string = "bef43054-89e8-40c0-9e62-96c4a836531b"
		var v interface{} = "c2dc72cc-2297-4f4f-8365-d2e39a70d8e1"

		test := raml.NewUntypedMap(1)

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

func TestUntypedMap_PutIfNotNil(t *testing.T) {
	Convey("TestUntypedMap.PutIfNotNil", t, func() {
		var k string = "dda9e790-2aa2-4173-8c3b-4a0443840952"
		var v interface{} = "87f36279-c1f5-4ae1-bd02-d6eb705a64ee"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1f57b3f7-054d-4e5a-8e47-25861db1ae7a", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "aaecdefc-fc4e-427f-9cd8-fe796ae6c6bb"
		So(test.PutIfNotNil("1cff545a-9526-4293-9120-e5482907fc99", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "5d722205-36b9-493f-b228-3890d7a136d1"
		var v interface{} = "17d405f1-7f42-47b2-a504-344a7d2aa69e"
		var x interface{} = "94ee2e41-3155-40f5-ba7d-7ad2330b4842"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4de7b479-a84d-4d45-8f3e-88eae111ea20", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "d4e0da5d-146d-4d1e-908d-396a000fcddc"
		var v interface{} = "1334bd75-84ca-404f-97d8-bcd0e7c0e77a"
		var x interface{} = "498520d1-f165-4e7b-8bc1-4e270c8c521e"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9b3ac7f4-6a90-4d07-9ceb-a0b472250481", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_MarshalJSON(t *testing.T) {
	Convey("TestUntypedMap.MarshalJSON", t, func() {
		var k string = "4688086d-d250-42b7-a9af-7bf3cd62f583"
		var v interface{} = "c54388be-050b-4a5d-a22d-47007fe3dfb2"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4688086d-d250-42b7-a9af-7bf3cd62f583","value":"c54388be-050b-4a5d-a22d-47007fe3dfb2"}]`)
	})
}

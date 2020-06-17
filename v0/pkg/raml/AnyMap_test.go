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
		var k interface{} = "c7eeebdd-eb26-49fe-bd14-d1b9ca6f0a2a"
		var v interface{} = "ce90221d-26a8-412f-8694-4990b6d9b6d6"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "a884131c-95c0-46da-8c90-57b748717099"
		var v interface{} = "238edf5c-b0e9-476b-ab09-0cb4407c69f7"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "19b6562f-d646-4b75-a5b6-bed2349fe13e"
		var v interface{} = "ca27eb83-ce92-4cdf-bd72-725d1e76d895"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("aa752397-5eaa-4510-a15e-3dc84f763fa3"+"bb97f862-75e5-4efc-b28d-1f564688ea24"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "5535a7f6-ce87-43c0-b22b-eb6eaab1d7b6"
		var v interface{} = "75f8786e-a6f6-4a9c-8bea-b8abfedef5dc"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("28538d11-e5f1-4034-ba62-cc2de4007de7" + "d00e3362-6cff-4c4b-b79d-d072c9a5c88c")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "e6dbd463-2528-4c36-9d37-2b916f79c68d"
		var v interface{} = "1556daf9-4b9e-4b1b-83f3-4b4e5271da4f"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("49f7b2d3-b65e-472f-b265-edd175242656" + "c7545447-8927-41fb-bdae-68a6770732c1")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "5e017e1e-d3d8-48e1-abe1-54e4350a61ce"
		var v interface{} = "a2e417e3-6f2e-438f-bbf5-8b71a69a6700"
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
		var k interface{} = "15ff9386-8042-4d8d-a251-ac2828b645fe"
		var v interface{} = "3cf82788-aac2-4e9e-8919-63abf49838bb"

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
		var k interface{} = "c7b2a2fb-de89-4b03-9a15-b4793167011c"
		var v interface{} = "7db96a71-282e-4c93-aee5-a3ab0138f100"

		test := raml.NewAnyMap(1)

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

func TestAnyMap_PutIfNotNil(t *testing.T) {
	Convey("TestAnyMap.PutIfNotNil", t, func() {
		var k interface{} = "14d50211-52f7-40a5-8990-c2a915bc0b5d"
		var v interface{} = "1934e593-8376-4a95-95d2-b3499a1f171b"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("af754a54-d59f-45d6-8724-c4b7f7f4d105", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "0b6e6e22-fc9f-4aa1-92a2-edbd5a073052"
		So(test.PutIfNotNil("7b6e2e64-ea91-4a64-8c35-abfa16f4be6b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "4fd5ea11-d6a4-44c6-9c0c-79bbbb54997c"
		var v interface{} = "244bb1e8-87bc-4b99-a365-097ae48f0a24"
		var x interface{} = "af2c0497-b442-44ef-98d6-912068f254e0"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d1e7e663-c7dc-487c-a7b9-13ae63f71a6c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "2225cbc9-6709-429e-bac9-9420a3c75a2b"
		var v interface{} = "5dab2b40-b1e2-4d4b-904b-701a5729d916"
		var x interface{} = "bda52531-a319-4bbc-84ae-f2ca82e1ec06"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1972acfe-04ff-4771-a918-9c9b4e80831d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyMap.MarshalJSON", t, func() {
		var k interface{} = "90582223-8eef-46fd-a7d4-e102aeca9a37"
		var v interface{} = "06657f76-bd5f-4e8d-8cde-0ede58f5a88a"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"90582223-8eef-46fd-a7d4-e102aeca9a37","value":"06657f76-bd5f-4e8d-8cde-0ede58f5a88a"}]`)
	})
}

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
		var k string = "86697954-6a91-477f-a0c2-efa4205d5e12"
		var v interface{} = "e483bc07-e872-4fb0-aed4-2fde2677d0ea"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyExampleMap_Delete(t *testing.T) {
	Convey("TestAnyExampleMap.Delete", t, func() {
		var k string = "18accc9b-ffbf-4bac-8917-05b16e58e5be"
		var v interface{} = "7348d1be-84ee-49a8-9aaa-ee4a68b839fe"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyExampleMap_Has(t *testing.T) {
	Convey("TestAnyExampleMap.Has", t, func() {
		var k string = "4a71d222-f0bd-46b4-beef-3c928829a1e7"
		var v interface{} = "b95ff7fc-1642-43ff-bac7-a94b85474be3"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5861bbaa-1b25-4eba-aede-0baa1989f1b0"+"9fbc6e9e-403f-4f0d-80f4-2ddad571bc48"), ShouldBeFalse)
	})
}

func TestAnyExampleMap_Get(t *testing.T) {
	Convey("TestAnyExampleMap.Get", t, func() {
		var k string = "73ffb4ef-cca8-4c9a-b047-abf6505c9b43"
		var v interface{} = "4ee98132-79e9-4251-a9b0-527632219dc9"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("34f60ac4-027b-4085-a21f-d79e20a37b82" + "8d6aee8c-a4ca-4f43-a9e5-bb8f8e44217c")
		So(b, ShouldBeFalse)
	})
}

func TestAnyExampleMap_GetOpt(t *testing.T) {
	Convey("TestAnyExampleMap.GetOpt", t, func() {
		var k string = "8cfd4254-3cbb-49d7-bce9-ad279e73b623"
		var v interface{} = "3003904c-0920-4e7f-aa99-e52a616dce5e"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("99819e96-332f-45ea-8b38-1f1012a8030d" + "affa6c34-a98f-4043-a418-deac9fce587d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyExampleMap_ForEach(t *testing.T) {
	Convey("TestAnyExampleMap.ForEach", t, func() {
		var k string = "3c3b21bf-9334-4a9e-b0f3-52416452ee60"
		var v interface{} = "985cf297-dbae-424e-9543-2993d1d748f7"
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
		var k string = "3cecc648-0604-406d-9d7e-20b7e1128568"
		var v interface{} = "bb0043b7-d37e-44c3-b5aa-ff7373e0e84d"

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
		var k string = "2006a546-6578-4fb4-ba32-047d00ef0fe8"
		var v interface{} = "2fb7f232-3128-4dd5-90f5-9ca80ec86eb9"

		test := raml.NewAnyExampleMap(1)

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

func TestAnyExampleMap_PutIfNotNil(t *testing.T) {
	Convey("TestAnyExampleMap.PutIfNotNil", t, func() {
		var k string = "3e08f9e3-235d-4ca6-9de2-b15d04186e14"
		var v interface{} = "188e52a0-d230-416a-8a00-e2745a6987fb"

		test := raml.NewAnyExampleMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9dad866f-081f-49ac-9164-503b20f47497", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "3f0b4d71-232b-41ff-8b8a-7d0e950ede8e"
		So(test.PutIfNotNil("16e9833e-20c4-4deb-963e-d9d1b49b6113", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceIfExists", t, func() {
		var k string = "4919458e-53fb-4ea6-b339-bbe1714085ce"
		var v interface{} = "a53eccb0-bf5c-49c3-8a95-ee23c35b610a"
		var x interface{} = "6b05b572-312d-40c9-99cf-ca3e7fcc020e"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cca7b59c-6ada-4bf9-b868-08295e816b3a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceOrPut", t, func() {
		var k string = "775203a3-65c3-4ef1-a61d-11befaafd5d1"
		var v interface{} = "4ef8b875-8b17-4478-ba08-10d04f6761df"
		var x interface{} = "edf07833-b28e-41ed-8a76-07f568ed031f"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e7b5b845-4f13-45bd-b8f4-75565c65b418", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyExampleMap.MarshalJSON", t, func() {
		var k string = "5276052e-03c1-421c-928f-529aff332432"
		var v interface{} = "1e719819-95fe-49cb-b13e-b3df821f063b"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5276052e-03c1-421c-928f-529aff332432","value":"1e719819-95fe-49cb-b13e-b3df821f063b"}]`)
	})
}

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
		var k interface{} = "54ec1ee8-b766-45ee-87a2-107db113c5b4"
		var v interface{} = "d6044ab7-e9ca-4997-b4e4-61e272df9891"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "ed0fc193-3dba-4e83-b1ab-c2463b82140f"
		var v interface{} = "ec5a21a9-848d-49a4-b199-419bf2915dbc"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "ccae58f8-acf9-454c-b49d-7603186d2a4b"
		var v interface{} = "1de0f415-ec4e-440f-89dd-e34c46994790"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d5842e96-c038-4160-9eb9-5a9ba3dce7c0"+"e95f8b8d-cf02-433c-affc-d77f7ce1e868"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "e266fd4f-c21e-467d-b349-d25d5665aefe"
		var v interface{} = "2f355fea-f75a-4174-ad8b-35fb15510243"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7076bf7a-cd7c-4bf1-a4c8-f5a7ef9ba981" + "5298b6e3-e12d-435c-8bf0-a54783020192")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "bfc2a325-a869-4516-9ebd-0ab75e2ed692"
		var v interface{} = "0127ad1b-00c8-4afd-a0e7-eba025a3f620"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("2500bb12-8e3b-4cfd-a083-5aa51f7347d8" + "1f4a158b-a661-43cd-95ad-7bda0e55608e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "342fd6e1-5d6d-454b-84a3-92824ea054c1"
		var v interface{} = "f678a903-7ded-45da-8302-7bde3d9c89f3"
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
		var k interface{} = "65a043a5-f187-499d-823f-97b8e5eb3a05"
		var v interface{} = "807139c1-b8df-4017-9c2a-0ad25eb92743"

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
		var k interface{} = "e8f45fa4-4c72-4718-902e-2af4687cd67e"
		var v interface{} = "386ec4df-276b-4790-acd4-c17521503815"

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
		var k interface{} = "472f71cf-df87-4f2f-b1f0-c3100c2f6265"
		var v interface{} = "075dd4ea-2041-4a6e-9cca-0bbe01e649ba"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("aa2ca183-ad77-462f-bcef-d57750978722", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "20a53753-20b6-4385-a74e-aea0d1f96b4e"
		So(test.PutIfNotNil("a7434f15-d4a2-43e9-a62a-3ad458cc1394", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "0324247b-c820-4be9-b21d-4c160e9635df"
		var v interface{} = "0b3cb06f-5894-4749-bd66-79d7ff7b02d6"
		var x interface{} = "5449e864-6ba1-4928-80c0-af1634b287e3"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6d592ae5-5072-4d48-9794-d2c9d3bbcd5e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "a636b6be-1ef5-4ec7-a851-4a2b0dd5085e"
		var v interface{} = "c5204d4a-32c2-4bf4-b3a6-2b576ed97054"
		var x interface{} = "687fbfd9-4dce-43de-951c-e7e2e55590b0"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e283b3d9-64f0-4fa4-b1fb-ae43dea5de5a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyMap.MarshalJSON", t, func() {
		var k interface{} = "bf1b4fff-6416-457d-b91b-b5bd14c49c6b"
		var v interface{} = "cc6bb9fe-acf4-4b0e-9c05-612f00f728bc"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"bf1b4fff-6416-457d-b91b-b5bd14c49c6b","value":"cc6bb9fe-acf4-4b0e-9c05-612f00f728bc"}]`)
	})
}

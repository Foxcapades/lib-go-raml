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
		var k string = "1eb7261e-5e03-456a-8c5e-63aedb95132d"
		var v interface{} = "06ade346-d2f3-4be1-99fe-1ade49b2b61c"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyExampleMap_Delete(t *testing.T) {
	Convey("TestAnyExampleMap.Delete", t, func() {
		var k string = "c1946012-3558-45b2-8c29-2a4ead361d74"
		var v interface{} = "0f4b51d8-b117-451a-a15b-fd484c3637a3"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyExampleMap_Has(t *testing.T) {
	Convey("TestAnyExampleMap.Has", t, func() {
		var k string = "2cf2804e-5544-46b7-adfc-399c50426f77"
		var v interface{} = "ed089e1b-3e48-4dbf-a50d-c0272bd3f65e"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("0995f490-f979-4a84-8243-45b47de8fa88"+"9bed2c72-28d6-4111-a27c-74b8c9a1f93d"), ShouldBeFalse)
	})
}

func TestAnyExampleMap_Get(t *testing.T) {
	Convey("TestAnyExampleMap.Get", t, func() {
		var k string = "fc9ddd5c-4428-4284-8916-2d835861bcbe"
		var v interface{} = "2aa75d31-1295-4544-b893-f04e540b08f0"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5fe83ce0-99e3-4ff2-9502-7dd29592c594" + "15bd6566-ddb0-4b94-b533-457d060bebd2")
		So(b, ShouldBeFalse)
	})
}

func TestAnyExampleMap_GetOpt(t *testing.T) {
	Convey("TestAnyExampleMap.GetOpt", t, func() {
		var k string = "5cdca5e2-92c6-41ca-aeb2-03f5960e2f61"
		var v interface{} = "3374040e-8539-448d-8867-ffd90fc57006"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5786728f-cd6a-475e-8c5b-29b54f879457" + "e605187e-4bbb-4d0b-9dab-e370159864c6")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyExampleMap_ForEach(t *testing.T) {
	Convey("TestAnyExampleMap.ForEach", t, func() {
		var k string = "8be8181b-de83-4c1f-8827-e0138688dd8f"
		var v interface{} = "5d35b537-b6c2-484c-864c-b71132a44862"
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
		var k string = "d99fc728-c0d4-47f2-a5c2-174493fb052c"
		var v interface{} = "dd55f781-62e2-4edc-968d-d3116fbaf4ca"

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
			var k string = "acac2ee2-f38e-4179-874c-ba37f1abeb6f"
			var v interface{} = "6b653dc8-9368-431f-ad34-ff6c2407459b"

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
			var k string = "3d95318f-c80b-498f-afba-2a0075213900"
			var v interface{} = "5a40887f-8323-4e30-a354-c9f27cf2cdf4"

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
		var k string = "e2eb26aa-bee4-4d1b-b1da-d2d44b9da8bb"
		var v interface{} = "ba0298a6-199a-4084-b8e9-e1a22c367f7e"

		test := raml.NewAnyExampleMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("ac963117-39df-485d-a9ac-52623e4f186e", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "1da20d2e-136e-43c4-8529-33c82f4f0e02"
		So(test.PutIfNotNil("d85d29ad-01bf-45c0-85a8-fe4be74e39fc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceIfExists", t, func() {
		var k string = "a2730c99-3608-4ad5-be75-12f24e53db0b"
		var v interface{} = "2fba8f3f-1be2-4d1f-8475-7c5ea4ce5794"
		var x interface{} = "a799ff23-5840-40df-b2f9-c5ae695b3ec4"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("86478908-a734-4b35-8d1b-5ee93bf1e4dd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceOrPut", t, func() {
		var k string = "8b25138c-1b46-47e1-86ee-d724ba4b2d8f"
		var v interface{} = "eac5dba6-f2e4-47df-bf9c-4120ff2ac466"
		var x interface{} = "baca153a-6dab-4ca1-9a43-b08623704caa"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8cd9e1d7-fe0b-43b1-ad0d-4a0398a572e6", x), ShouldPointTo, test)
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
			var k string = "f66e7f74-70ae-47e1-92a8-e61ee70aad79"
			var v interface{} = "3b376dc3-ecfa-4e60-9069-0c4c4ef638cc"

			test := raml.NewAnyExampleMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"f66e7f74-70ae-47e1-92a8-e61ee70aad79","value":"3b376dc3-ecfa-4e60-9069-0c4c4ef638cc"}]`)
		})

		Convey("Unordered", func() {
			var k string = "f66e7f74-70ae-47e1-92a8-e61ee70aad79"
			var v interface{} = "3b376dc3-ecfa-4e60-9069-0c4c4ef638cc"

			test := raml.NewAnyExampleMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"f66e7f74-70ae-47e1-92a8-e61ee70aad79":"3b376dc3-ecfa-4e60-9069-0c4c4ef638cc"}`)
		})

	})
}

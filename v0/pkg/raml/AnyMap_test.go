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
		var k interface{} = "9f6003fc-70b6-4f9b-8eee-416ddf4aaae9"
		var v interface{} = "8f44fcd7-a03c-43fe-ba45-8c72e69bdfaf"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "c32d4efd-8a8d-4906-a73e-b7e504b2d84f"
		var v interface{} = "380c5589-65c1-4c5a-bf25-f08a54fa5242"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "78688f76-a5b1-4802-bce1-75804b0cf4bc"
		var v interface{} = "f69c0e5e-6bfa-422f-b475-4fdba54c13bb"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7ec93458-3739-4f12-8423-5d0caa12a5f5"+"bdbc9edc-2c2b-421a-82ff-c0965c65acbc"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "b2c1597f-98b2-49ed-9e7d-a397ceef7c2d"
		var v interface{} = "08390e7d-8e87-4f93-9388-9fa24ea912fc"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("3036cd45-9ad1-4525-8b89-684d5f004071" + "5c24cc05-d68d-4377-af91-cf1f8b6c9e7d")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "f99b17b2-bf2a-4990-9d58-26045438c5b9"
		var v interface{} = "ef418387-4060-4e4d-aef6-f498d3eac863"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7c2b36da-a17b-4c68-b28e-e4fe7dc0bcf9" + "2908394a-4c52-42e3-955a-a8f5f64fe83a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "aed1e8b8-1359-464f-a0e4-29a576f05b47"
		var v interface{} = "9a347c22-ccc1-44b4-8278-6584db5fd571"
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
		var k interface{} = "0912c09e-79b1-45f3-8e93-af044834b230"
		var v interface{} = "af548683-2081-4cd4-9bda-c68cbec49c2a"

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
		Convey("Ordered", func() {
			var k interface{} = "18b3fae0-edb5-4ad4-ac2a-7c7fe333f3d5"
			var v interface{} = "e1e82c5e-acfd-4a99-a185-dc8d863e698c"

			test := raml.NewAnyMap(1)

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
			var k interface{} = "cee27e7c-4243-49b7-a28d-172da7fa2c42"
			var v interface{} = "b4bfb230-100a-48fb-8950-84dd7d3c7693"

			test := raml.NewAnyMap(1)
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

func TestAnyMap_PutIfNotNil(t *testing.T) {
	Convey("TestAnyMap.PutIfNotNil", t, func() {
		var k interface{} = "94053ddd-ba1a-42ab-8e72-7343a11baf34"
		var v interface{} = "212220ec-9f9d-46c1-8dbf-b6ae15fd6158"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6aa8662d-c493-4c42-86f0-8c8a4c96b5c4", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "15ee0961-7257-4df6-bd07-3d9f902ea163"
		So(test.PutIfNotNil("b3cccc7a-c263-4d77-86fd-b8b1ee5f59ff", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "471f5e11-6fed-4cb9-b541-bd2ff514630c"
		var v interface{} = "d4ef175e-21bf-4c9d-9f2f-8121deb1b87e"
		var x interface{} = "f49a3a7c-2d60-4f9f-8f6f-0609425a98da"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c98a8d5e-d61f-450c-95c7-af72959c9a7c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "7a6ca32e-438a-47d3-9990-af755e3851d1"
		var v interface{} = "c4dcdc13-5b39-4325-98cd-2f9d29f669a0"
		var x interface{} = "a7a0cdf2-2184-4ff7-aef5-b5cd677519f5"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d6b72b9f-e2e9-4e44-96e1-6be95d2e5198", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_MarshalJSON(t *testing.T) {
	Convey("TestAnyMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "e89a7b2f-bf2d-4dfa-a3cf-4339da10f196"
			var v interface{} = "0d0eede3-2b13-47ae-8735-2399f0dbfa35"

			test := raml.NewAnyMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"e89a7b2f-bf2d-4dfa-a3cf-4339da10f196","value":"0d0eede3-2b13-47ae-8735-2399f0dbfa35"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "e89a7b2f-bf2d-4dfa-a3cf-4339da10f196"
			var v interface{} = "0d0eede3-2b13-47ae-8735-2399f0dbfa35"

			test := raml.NewAnyMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"e89a7b2f-bf2d-4dfa-a3cf-4339da10f196":"0d0eede3-2b13-47ae-8735-2399f0dbfa35"}`)
		})

	})
}

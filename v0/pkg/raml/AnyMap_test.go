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
		var k interface{} = "ab8753a1-bf49-471b-89fd-a2077118af1b"
		var v interface{} = "705ca88e-d17d-4ff5-bb68-2c727e9c85f1"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyMap_Delete(t *testing.T) {
	Convey("TestAnyMap.Delete", t, func() {
		var k interface{} = "02ab74b9-bed3-4db9-9bfe-7635a1ef2f6b"
		var v interface{} = "cd9cc223-70b4-4e0d-9d98-89316446db5a"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyMap_Has(t *testing.T) {
	Convey("TestAnyMap.Has", t, func() {
		var k interface{} = "90ad18ac-8a2c-4f3f-8c4d-18c59333ae0f"
		var v interface{} = "3038ca9b-1c76-47e0-a862-bcd480fc70f3"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("91679003-7209-4cdd-a156-cb0dd06c7cba"+"4b952984-cd7d-4e9a-94bf-633a6185cd90"), ShouldBeFalse)
	})
}

func TestAnyMap_Get(t *testing.T) {
	Convey("TestAnyMap.Get", t, func() {
		var k interface{} = "7b10e88f-ec5e-4244-960a-4b6814ec61eb"
		var v interface{} = "299de777-8346-40cd-a52b-35277dd36eb6"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("556332fb-2acd-4cb4-a161-c04735a438ce" + "6f446a1c-cd8a-471b-93e5-934201d544ef")
		So(b, ShouldBeFalse)
	})
}

func TestAnyMap_GetOpt(t *testing.T) {
	Convey("TestAnyMap.GetOpt", t, func() {
		var k interface{} = "8b7db0be-b164-4e08-b9a5-66302747ae31"
		var v interface{} = "753dac0e-348f-489c-9f05-eee7cc837742"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b3f123d9-b185-4d09-9fc3-1415b86d279d" + "3be6bd1b-9d4e-4698-a147-0f54948ab66c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyMap_ForEach(t *testing.T) {
	Convey("TestAnyMap.ForEach", t, func() {
		var k interface{} = "67695db4-3a41-4d6c-acae-7ff8a947fd6d"
		var v interface{} = "b560423d-36d6-40dc-9fc0-b3ff7e6b9043"
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
		var k interface{} = "6d94edc5-8ad6-4024-baf3-5ecf66f1f2c6"
		var v interface{} = "afd543b8-76b0-4cb6-b20b-66c934336b9d"

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
			var k interface{} = "e10250cf-9d97-4661-806d-320b22c7577e"
			var v interface{} = "687f77d0-5c23-457d-b6ed-5aa2dfb3a872"

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
			var k interface{} = "2e21e1c2-930a-4798-a94d-3a9969137382"
			var v interface{} = "1f0043c1-7f87-408f-9213-001fe706b51d"

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
		var k interface{} = "849b94d6-10a3-4e47-8eb5-b8f7a4a1c0cd"
		var v interface{} = "3e56962a-14aa-4f7c-83a2-95b49e9109f4"

		test := raml.NewAnyMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("715051f9-89f7-48e6-bf96-d8763e429a69", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "0903e98c-109f-4b64-9eed-1f0996596650"
		So(test.PutIfNotNil("0e0c5636-3196-42f2-8676-3ce675827ac2", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyMap.ReplaceIfExists", t, func() {
		var k interface{} = "a06a265f-137e-4bb2-848b-41ef4d8e6294"
		var v interface{} = "7c19a1ea-5125-49cf-abeb-d13fa6997aa1"
		var x interface{} = "aff9b8bd-2e2b-48d2-8d4d-9f15813cc50b"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("72ea90d5-b33c-4992-9d31-7420b64b44d5", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyMap.ReplaceOrPut", t, func() {
		var k interface{} = "6d569ae3-3f47-4c8d-8397-82a660b54f10"
		var v interface{} = "d515965d-8fb6-4998-b3a8-fd9ac5068d31"
		var x interface{} = "dd7c945f-e0d3-4d4b-bfdf-31342a52aea6"

		test := raml.NewAnyMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9c12169f-6ece-47f4-a3ee-aee652bce6b9", x), ShouldPointTo, test)
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
			var k interface{} = "57a63799-8fcf-4ad7-b626-95f24a32b7bc"
			var v interface{} = "99eabece-75a3-476f-8024-01565c2d4725"

			test := raml.NewAnyMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"57a63799-8fcf-4ad7-b626-95f24a32b7bc","value":"99eabece-75a3-476f-8024-01565c2d4725"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "57a63799-8fcf-4ad7-b626-95f24a32b7bc"
			var v interface{} = "99eabece-75a3-476f-8024-01565c2d4725"

			test := raml.NewAnyMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"57a63799-8fcf-4ad7-b626-95f24a32b7bc":"99eabece-75a3-476f-8024-01565c2d4725"}`)
		})

	})
}

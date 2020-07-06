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
		var k string = "67ef42bb-0958-413f-9825-13fbd3ac9bdd"
		var v string = "13d46b8e-eeaa-4130-bc76-72df89a162c9"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestStringMap_Delete(t *testing.T) {
	Convey("TestStringMap.Delete", t, func() {
		var k string = "78060eac-5a2e-4f68-8b08-f89ed9c53766"
		var v string = "6b41e73f-b407-448c-81a4-a32a57819e46"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestStringMap_Has(t *testing.T) {
	Convey("TestStringMap.Has", t, func() {
		var k string = "ef626acd-6818-43d7-a743-2644346bd441"
		var v string = "824a313d-e1ac-42e2-bc21-d66d3cdd3495"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3e0bd88a-764d-4c52-983b-573bba0cda41"+"0373f297-cc79-48a7-9f68-96b53125a869"), ShouldBeFalse)
	})
}

func TestStringMap_Get(t *testing.T) {
	Convey("TestStringMap.Get", t, func() {
		var k string = "b944958b-d40a-4bab-a987-9e09063adff5"
		var v string = "f4aa9075-1bb7-40aa-8dd4-0d390d0b3ea1"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("3b5e2698-b557-44e7-a1d4-207fc6e0a6e7" + "612a2483-54f4-4fa0-b9d5-83b31b2ca407")
		So(b, ShouldBeFalse)
	})
}

func TestStringMap_GetOpt(t *testing.T) {
	Convey("TestStringMap.GetOpt", t, func() {
		var k string = "983c6944-3331-48cc-8b62-92d33ec73fbb"
		var v string = "40a20fba-4d7b-4822-9ecd-ff4519dec514"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b010e02a-6fa7-4961-ae2f-bd31f7cee346" + "448a2caa-844e-4543-9412-57c66ab3627a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestStringMap_ForEach(t *testing.T) {
	Convey("TestStringMap.ForEach", t, func() {
		var k string = "98cc4786-7a4b-4a08-82b3-7eb90a780486"
		var v string = "a2616089-b548-4358-841e-e8109477499e"
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
		var k string = "7f22dc5c-0696-4c90-873e-cee4749b2b81"
		var v string = "b61c38ad-9e56-4227-96a5-4ceac750aede"

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
			var k string = "5ebaf02f-2378-4eb2-b1f0-1dad67ea4379"
			var v string = "6eff5455-93b5-48a2-af75-712a214fa91c"

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
			var k string = "fccfac61-7cd5-4885-b418-18c4f950b247"
			var v string = "22a85ca1-feda-477d-8126-b89cf418bda5"

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
		var k string = "61103dcc-2a0a-4667-9755-fffd28cb613b"
		var v string = "6ef320e2-11d1-4fbf-8a88-e3c978d17300"

		test := raml.NewStringMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0c0a91e8-9616-4c68-aeb1-f324478ee808", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "0107167f-aeb8-4383-b4b6-b63bae9b7507"
		So(test.PutIfNotNil("eee04ec9-f25c-447c-84d0-8ff934713ff3", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceIfExists(t *testing.T) {
	Convey("TestStringMap.ReplaceIfExists", t, func() {
		var k string = "a81fcffd-6c4a-4970-a3fe-87e6f4ed63fa"
		var v string = "0151c370-a9f5-4622-a6fe-7f8b23316691"
		var x string = "fb35e48b-7070-44d6-a3da-e2b8d95e357c"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e4ad96c5-22d1-43df-8e60-c916b9032632", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestStringMap_ReplaceOrPut(t *testing.T) {
	Convey("TestStringMap.ReplaceOrPut", t, func() {
		var k string = "638f553f-51b0-4010-bddc-00ef0b071afa"
		var v string = "7e18d59a-d31e-46ef-bd9b-166748e761f9"
		var x string = "0482a6d0-1dae-47a0-b262-d14903831d1b"

		test := raml.NewStringMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c829fb73-4d49-488a-ad6f-eebeae91d38e", x), ShouldPointTo, test)
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
			var k string = "dd4c914f-e17a-4427-a534-321c1f24aefa"
			var v string = "59ff3a36-e341-4f13-9c46-2f3ff0cdd5b7"

			test := raml.NewStringMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"dd4c914f-e17a-4427-a534-321c1f24aefa","value":"59ff3a36-e341-4f13-9c46-2f3ff0cdd5b7"}]`)
		})

		Convey("Unordered", func() {
			var k string = "dd4c914f-e17a-4427-a534-321c1f24aefa"
			var v string = "59ff3a36-e341-4f13-9c46-2f3ff0cdd5b7"

			test := raml.NewStringMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"dd4c914f-e17a-4427-a534-321c1f24aefa":"59ff3a36-e341-4f13-9c46-2f3ff0cdd5b7"}`)
		})

	})
}

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
		var k string = "45f6c27a-88bb-4f67-9181-0fc1aebaa964"
		var v interface{} = "17443e4b-eab2-42f1-a476-56e526477798"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "d6da6c65-0c11-483e-b908-a6e53eae1cda"
		var v interface{} = "a6ce7cac-e2d6-4c00-b2a7-399a218ec137"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "3ec1d4db-9264-4f3f-b4ab-e1fe6e0780f8"
		var v interface{} = "241690ce-72b1-40e9-a1d0-ec002ba08e16"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("63052f84-a3ef-426e-b40d-b4a4cb6ff925"+"a7f456ef-f548-4ac6-abf4-ba332456da65"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "00d34017-5ff7-4e01-ac1b-41cc7c18e969"
		var v interface{} = "51d07ecd-e3e3-4094-891b-99d8cd02ac8e"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b04a7c56-25e0-49e7-96dc-7cf6c56ce154" + "365f77ac-1a67-44e1-991c-b41da59206a9")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "8ba8919e-da1d-47d2-9234-b4430d2b0b4a"
		var v interface{} = "a5d21466-b4f1-4344-92de-141e5aed3954"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0c37d0d7-b529-4cf1-8d37-b0df9a61b0b7" + "e5957691-ff5b-4ac3-b94f-274d29aa4c60")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "d34b5885-d181-467c-aac6-ea986def1a45"
		var v interface{} = "e4054f74-a421-4f60-8530-e944f25b694c"
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
		var k string = "276071b1-4e6f-4513-bd7e-f781e1392eeb"
		var v interface{} = "af6c317c-4e9e-4a43-950b-7ac291a4f78b"

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
		Convey("Ordered", func() {
			var k string = "134086e7-5469-49bc-84dc-c21147394ee3"
			var v interface{} = "cc78f6f7-6c16-41ed-86dc-e344559884fc"

			test := raml.NewUntypedMap(1)

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
			var k string = "9c09b324-5a75-4528-810b-ea3b1495c628"
			var v interface{} = "47dd7560-e042-47ab-8c29-00cf1d9c9737"

			test := raml.NewUntypedMap(1)
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

func TestUntypedMap_PutIfNotNil(t *testing.T) {
	Convey("TestUntypedMap.PutIfNotNil", t, func() {
		var k string = "52480b76-decf-4cdf-82e1-8c1a47e959d6"
		var v interface{} = "72ce3db9-9519-470e-bf98-7e178428c315"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6c0fca51-4823-4a65-b743-734324bc6c71", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "0e2b4c79-d775-4e4c-b685-b083db65ce56"
		So(test.PutIfNotNil("de59dcd6-7c33-4ecf-964a-19b55045ff44", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "cf3b4792-6b9f-483c-a5cf-3ad1dfaf30af"
		var v interface{} = "39dda4f9-e728-4003-a0c2-dae5f1324edc"
		var x interface{} = "88be149c-d4d9-4e71-a452-59cb3f7fe547"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d27396b0-6a3b-4d03-92f8-a4c6c27eb8cc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "b68ba2cb-0352-4d4c-8e69-1097f5fe4721"
		var v interface{} = "28d2d298-85d7-4450-b69e-b709f1a7e20b"
		var x interface{} = "e908a9ea-d962-47f3-9866-6b60fab8400d"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("cfd06081-3ddf-4950-b216-8d8dc075a2aa", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_MarshalJSON(t *testing.T) {
	Convey("TestUntypedMap.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "24ea23aa-6179-4dfd-b2ee-f2049fbaaaf0"
			var v interface{} = "41eb3d72-b1e3-49dd-9c1d-577488e66ed6"

			test := raml.NewUntypedMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"24ea23aa-6179-4dfd-b2ee-f2049fbaaaf0","value":"41eb3d72-b1e3-49dd-9c1d-577488e66ed6"}]`)
		})

		Convey("Unordered", func() {
			var k string = "24ea23aa-6179-4dfd-b2ee-f2049fbaaaf0"
			var v interface{} = "41eb3d72-b1e3-49dd-9c1d-577488e66ed6"

			test := raml.NewUntypedMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"24ea23aa-6179-4dfd-b2ee-f2049fbaaaf0":"41eb3d72-b1e3-49dd-9c1d-577488e66ed6"}`)
		})

	})
}

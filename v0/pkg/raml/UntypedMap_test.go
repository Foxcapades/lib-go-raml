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
		var k string = "4e21ad01-a3d8-4ac2-a80a-73a0fc524d8f"
		var v interface{} = "a1be5af1-8447-41e7-8346-5ae108f4face"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "23bc20b8-fc6a-40f1-a74b-9ed17bd9df95"
		var v interface{} = "e18eff3d-9079-4733-b4d4-551c75607a5f"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "fdb9526a-ef4f-4ca9-9dcc-891db29070ee"
		var v interface{} = "31eec84c-4af0-43d3-809c-1b91a857fb4f"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f7c0258a-ff35-4f2b-b592-ad17c5fbbe6b"+"9a4a7fa7-b9e9-4952-be70-04c9b2439e09"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "d7a24d8d-9242-4903-8e83-9accd0f74983"
		var v interface{} = "25da7127-9485-419b-bd77-b67de5f8cbbe"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("17c1da43-fd69-4338-a814-5f789e7de031" + "86cba70f-9220-4d2c-a617-0fae461f05d7")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "47aa4b11-d093-48da-a9b1-c514c9c75a9b"
		var v interface{} = "227be31e-5d5e-4acb-acc7-e7fe8cd3ed05"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fdd86fef-435d-4ea5-ad99-ef1a71b4e2fd" + "0b8d2c81-eb62-4b7c-9464-39d61b8bcb56")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "00449344-2c1b-48c6-9a80-d2a562bcfc12"
		var v interface{} = "d5d61054-907a-4603-a5a0-6ef599a7ff3a"
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
		var k string = "7ed4d18e-21b2-415c-8ed9-78c251d71a5f"
		var v interface{} = "7bfab149-c762-46ea-b799-4a0eb734f248"

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
			var k string = "a68dc90b-9105-4317-b861-d9994af0dffd"
			var v interface{} = "cb15a039-9bd5-49f1-82ee-165d6510a52b"

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
			var k string = "9b7750eb-7f10-4b11-b79d-38f539c85b71"
			var v interface{} = "182a9703-9231-451f-a082-c76abc33a71e"

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
		var k string = "0d68c6b4-1d97-46c6-98d4-f5cb0f7d428b"
		var v interface{} = "d3190f28-5c85-476b-862f-33de730c5221"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a2bf27a6-8426-4c05-a8fb-931767d3af62", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "d01fa123-c312-423a-9991-2281dce46bef"
		So(test.PutIfNotNil("0dbdf513-3eba-49fc-a7a1-7f69901dad59", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "61fd8e87-f7a7-42dd-a87d-789bd956f916"
		var v interface{} = "c2df94e9-aaad-4a27-bd9e-8a3a0546daab"
		var x interface{} = "d2f4e752-2a84-48dc-8616-b0e6207fe564"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6ad46b66-c2bd-4e8c-81cb-9789c7489738", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "e5c9f12b-299c-43d8-9869-c184f843c254"
		var v interface{} = "85b5969e-1db0-464f-8925-3072cdbc24c8"
		var x interface{} = "c4537e9f-83eb-4e99-bb81-da2ce807316c"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("63cf6eb8-8c8b-4d66-9d07-9dd5cc6de026", x), ShouldPointTo, test)
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
			var k string = "f507f0c9-8e39-4451-a592-8925e9efe01a"
			var v interface{} = "0deb80e4-cee7-4bad-88f0-a6d03345284e"

			test := raml.NewUntypedMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"f507f0c9-8e39-4451-a592-8925e9efe01a","value":"0deb80e4-cee7-4bad-88f0-a6d03345284e"}]`)
		})

		Convey("Unordered", func() {
			var k string = "f507f0c9-8e39-4451-a592-8925e9efe01a"
			var v interface{} = "0deb80e4-cee7-4bad-88f0-a6d03345284e"

			test := raml.NewUntypedMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"f507f0c9-8e39-4451-a592-8925e9efe01a":"0deb80e4-cee7-4bad-88f0-a6d03345284e"}`)
		})

	})
}

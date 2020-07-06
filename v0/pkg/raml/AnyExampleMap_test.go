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
		var k string = "d1eb314b-df5b-43ab-8f76-7fb37b280538"
		var v interface{} = "4dd0497e-c877-482a-b211-c8050541391a"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyExampleMap_Delete(t *testing.T) {
	Convey("TestAnyExampleMap.Delete", t, func() {
		var k string = "7c29e5c7-712d-435e-8bfd-48a45d6eb7f3"
		var v interface{} = "dd70d1ad-053a-4a94-9e5b-dac9226b83d5"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyExampleMap_Has(t *testing.T) {
	Convey("TestAnyExampleMap.Has", t, func() {
		var k string = "dbde568f-c476-449c-b15a-16345fe88c25"
		var v interface{} = "eb46b11b-c423-4c71-b9d0-cb938ff8f0c1"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("0346ec5a-e621-4484-af88-bf182b89367f"+"66cc88ad-3817-4277-afcd-774bc633e50c"), ShouldBeFalse)
	})
}

func TestAnyExampleMap_Get(t *testing.T) {
	Convey("TestAnyExampleMap.Get", t, func() {
		var k string = "319ab564-6d1c-42ff-ae0e-4b4832fd79fb"
		var v interface{} = "5ceae8db-ffdd-426c-95e9-1545bcb46ffc"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("0cffccba-0907-48b8-bf42-3de35aa68e1b" + "448405aa-f766-421c-95bb-30e88298f7d1")
		So(b, ShouldBeFalse)
	})
}

func TestAnyExampleMap_GetOpt(t *testing.T) {
	Convey("TestAnyExampleMap.GetOpt", t, func() {
		var k string = "69f1413e-f855-476e-afff-836daceaef4a"
		var v interface{} = "ea083136-2d0f-4b07-ab32-0b794af52eb4"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f52a984a-d2aa-441b-88bf-c18f28175464" + "a1cafd06-f8da-40ac-8838-a30972eaab85")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyExampleMap_ForEach(t *testing.T) {
	Convey("TestAnyExampleMap.ForEach", t, func() {
		var k string = "e9864b42-6d2e-4be6-aca7-8c2b52f8d950"
		var v interface{} = "42225666-942e-4664-bd98-9cef76ff2f7e"
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
		var k string = "a6877c88-6637-42cd-a141-fc0a792bf140"
		var v interface{} = "b63aed49-2854-4921-85c3-06e23eec11e8"

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
			var k string = "d43846df-9df3-4a71-a465-e8fea4c5cd3a"
			var v interface{} = "52341384-a188-4dd8-95e7-5e978e388896"

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
			var k string = "ec3377db-9a6f-4d68-8019-a16b0ccfec24"
			var v interface{} = "ecf2a77a-d205-4926-9d97-4d928dca05c5"

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
		var k string = "e5196aa4-3c29-4507-9aa7-6ade4970c86d"
		var v interface{} = "ee04a11a-caf2-4d52-b042-0263922bd9cf"

		test := raml.NewAnyExampleMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("494c5001-6e04-4fae-ad77-7f2fa3ce79e6", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "f3f1a4aa-2be1-4eb2-a44a-a544c8cd8f16"
		So(test.PutIfNotNil("83186c3e-af79-48ef-8f3a-56a0b20f16b9", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceIfExists", t, func() {
		var k string = "7ec8de6c-c79b-49d2-8a9b-516cba57b676"
		var v interface{} = "c11b0e88-6a57-4522-9818-2de1d2d110c5"
		var x interface{} = "e782538f-4cae-43ad-97cd-2ef6e3665dcf"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b7edf537-b30c-40c9-8f8c-53105e4dec55", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceOrPut", t, func() {
		var k string = "96e8e83b-8c37-4a44-a6dc-52741210db52"
		var v interface{} = "a7f44580-b2d9-490c-a40b-32bd1b700e90"
		var x interface{} = "006039a9-ccb8-4334-9164-dce97222360e"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9a45dec9-ba6c-4422-b558-3f1e580c1f83", x), ShouldPointTo, test)
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
			var k string = "8765d403-f04d-42a1-8230-11a4f2a87788"
			var v interface{} = "80ba8d9f-58c8-4b98-bb33-c433e184abed"

			test := raml.NewAnyExampleMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"8765d403-f04d-42a1-8230-11a4f2a87788","value":"80ba8d9f-58c8-4b98-bb33-c433e184abed"}]`)
		})

		Convey("Unordered", func() {
			var k string = "8765d403-f04d-42a1-8230-11a4f2a87788"
			var v interface{} = "80ba8d9f-58c8-4b98-bb33-c433e184abed"

			test := raml.NewAnyExampleMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"8765d403-f04d-42a1-8230-11a4f2a87788":"80ba8d9f-58c8-4b98-bb33-c433e184abed"}`)
		})

	})
}

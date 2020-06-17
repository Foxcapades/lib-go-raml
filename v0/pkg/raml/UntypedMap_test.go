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
		var k string = "24e5397d-e4c2-452e-9ea0-334b7766a187"
		var v interface{} = "ba4272f8-55f5-4cb2-9890-58689b1c94a1"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestUntypedMap_Delete(t *testing.T) {
	Convey("TestUntypedMap.Delete", t, func() {
		var k string = "d745a27e-2035-4e79-91e0-90017e5502c4"
		var v interface{} = "7b397342-ffde-4a21-8b08-69cdf4a57d26"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestUntypedMap_Has(t *testing.T) {
	Convey("TestUntypedMap.Has", t, func() {
		var k string = "393fedc8-93e3-46c7-832d-e90cc22cbded"
		var v interface{} = "d3d4a523-ce73-47c5-b1c2-cf684c276a06"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5b3388bf-93ac-4e90-8def-d1ef5c329caf"+"91727943-7861-49c5-b83c-06aa6b93a0a5"), ShouldBeFalse)
	})
}

func TestUntypedMap_Get(t *testing.T) {
	Convey("TestUntypedMap.Get", t, func() {
		var k string = "62150afc-c341-40fb-9b1b-4bd32210111d"
		var v interface{} = "01889e83-abdb-41bb-a404-8158af7c6e38"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("a4b8f7f2-d291-4767-a3a8-58ee4400f74e" + "95fb5b48-5436-4770-9319-7ddf3a4e09f8")
		So(b, ShouldBeFalse)
	})
}

func TestUntypedMap_GetOpt(t *testing.T) {
	Convey("TestUntypedMap.GetOpt", t, func() {
		var k string = "7f22a72e-2d9d-4a1b-9f77-2fb7c78fef10"
		var v interface{} = "17d44740-86c0-4aef-ad82-48588292a2e5"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b7f32fdd-6f27-4ac8-bce2-606fd9b004a9" + "dd5c38b6-8eca-4a2e-b884-d3b3f2ad3b4b")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestUntypedMap_ForEach(t *testing.T) {
	Convey("TestUntypedMap.ForEach", t, func() {
		var k string = "c2b875d6-65ba-4657-9d79-aa0893ee4951"
		var v interface{} = "c2968c30-aff3-4f5a-81a1-1663e023fbfc"
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
		var k string = "972aebf2-3373-4900-8742-87e691c4f5a1"
		var v interface{} = "0f49365c-98b8-454b-aeaa-1aad0e0201c3"

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
		var k string = "e599c0b3-9416-4fdb-bed8-88cae36e08d4"
		var v interface{} = "7e17d985-960d-4fe1-8600-fddd8f7ed6c9"

		test := raml.NewUntypedMap(1)

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

func TestUntypedMap_PutIfNotNil(t *testing.T) {
	Convey("TestUntypedMap.PutIfNotNil", t, func() {
		var k string = "002adc12-f6d2-4c12-b514-6c9a477d49d2"
		var v interface{} = "cee5acb6-6019-427e-bd9a-f0df75cbd3bf"

		test := raml.NewUntypedMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("779c7486-9e82-4db6-8e6c-08a1a883f72f", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "7b4889be-f4fd-46c0-aa7a-d9f559abebda"
		So(test.PutIfNotNil("90c9b320-984d-4674-bfe4-58032837a34a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceIfExists(t *testing.T) {
	Convey("TestUntypedMap.ReplaceIfExists", t, func() {
		var k string = "ca8af0f2-bb5a-463c-9d49-a6db657cdac7"
		var v interface{} = "658659d1-980a-4595-af6e-2912ff7cf691"
		var x interface{} = "70199ec2-df18-4434-9fc9-1241867f56c3"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1e4bbe6f-cace-462f-960d-a09022f418ce", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_ReplaceOrPut(t *testing.T) {
	Convey("TestUntypedMap.ReplaceOrPut", t, func() {
		var k string = "ec48a8cd-b3db-4275-88bb-0bd6dd635071"
		var v interface{} = "30dca002-8fdb-4ba7-bc11-7c8d5a8cf3d5"
		var x interface{} = "d9ea15ee-6375-4d87-935e-3c9bf79d91de"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4b73b636-c68f-4e09-b869-ad5a6c196ace", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestUntypedMap_MarshalJSON(t *testing.T) {
	Convey("TestUntypedMap.MarshalJSON", t, func() {
		var k string = "21b70afc-c6f7-4599-af82-0bafc9a7d515"
		var v interface{} = "ef9ffddd-200e-4eee-82ab-58f3ee27dece"

		test := raml.NewUntypedMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"21b70afc-c6f7-4599-af82-0bafc9a7d515","value":"ef9ffddd-200e-4eee-82ab-58f3ee27dece"}]`)
	})
}

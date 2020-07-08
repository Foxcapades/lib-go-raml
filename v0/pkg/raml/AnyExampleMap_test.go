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
		var k string = "2ef5acdc-998a-470e-8ad4-fd34a1770a02"
		var v interface{} = "56a90f83-3e07-40ff-93cf-85f2f59ca2cc"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestAnyExampleMap_Delete(t *testing.T) {
	Convey("TestAnyExampleMap.Delete", t, func() {
		var k string = "06efeb53-ce3d-4369-b69a-03ca4b8e3dcb"
		var v interface{} = "8093c103-d262-4930-955e-ce582177eea5"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestAnyExampleMap_Has(t *testing.T) {
	Convey("TestAnyExampleMap.Has", t, func() {
		var k string = "065184c7-4269-43d9-82fc-94972be637b5"
		var v interface{} = "444fc017-465d-4b5e-8c53-321ac0af37e5"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c6b70999-7af6-4e0a-9914-80fb4fd54240"+"10f4c9d3-5cee-46e4-a5f4-8a2c320db5d8"), ShouldBeFalse)
	})
}

func TestAnyExampleMap_Get(t *testing.T) {
	Convey("TestAnyExampleMap.Get", t, func() {
		var k string = "fd6341e8-6405-4dd9-9abf-5f12906b481e"
		var v interface{} = "f39363ea-9d9f-4aff-a3db-65ef86b44571"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("a6865a26-ef6f-4af6-9359-2d782c971c77" + "ee9bbbcd-3523-4203-8b21-0a4a8542e731")
		So(b, ShouldBeFalse)
	})
}

func TestAnyExampleMap_GetOpt(t *testing.T) {
	Convey("TestAnyExampleMap.GetOpt", t, func() {
		var k string = "af793c8d-7f03-4a33-8d30-aa473b2bb428"
		var v interface{} = "fb6aa142-c390-4af2-b705-6cf601e5cf91"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ff18fcec-c843-414d-aaa9-270e4098d2b0" + "490c921d-1dd6-4730-9bd2-410e313a9890")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestAnyExampleMap_ForEach(t *testing.T) {
	Convey("TestAnyExampleMap.ForEach", t, func() {
		var k string = "2987eb55-419b-4c69-b905-e00198ccb27e"
		var v interface{} = "25e102c0-150f-4a8d-ad3f-2c567a5a522c"
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
		var k string = "a69f9294-0219-4518-aeb9-626033116c54"
		var v interface{} = "b5c39e83-af24-4342-bd0d-bdced8d6ad36"

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
			var k string = "878e5a55-c3ae-4857-b23f-75158f397568"
			var v interface{} = "aca8bcf6-5a22-4a3d-9143-eaac6cd297c6"

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
			var k string = "8862146a-c022-4cb7-9fb8-8be2e74d75dc"
			var v interface{} = "b7d74534-ed11-4c56-aee4-3d8260a02408"

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
		var k string = "b5e00d82-ed59-431e-89fd-a1a93d1f8cf8"
		var v interface{} = "ca69c05d-cbe2-44b3-8ca3-75033cd11034"

		test := raml.NewAnyExampleMap(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("613d19f5-f263-4db3-9fb2-76b9905d5dbc", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "58368a54-8037-438d-9020-404d92aca144"
		So(test.PutIfNotNil("db569e59-023d-4340-8148-eefbf7ede9c7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceIfExists(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceIfExists", t, func() {
		var k string = "74f442a0-5f59-4608-96a9-1fa4997cb7fb"
		var v interface{} = "07bf306e-e2e5-44f4-9c6d-974153811b1d"
		var x interface{} = "9352cd43-32a5-4eec-8af9-c00b215d4d20"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("2a406074-2119-4b98-95e6-9dce30b3d591", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestAnyExampleMap_ReplaceOrPut(t *testing.T) {
	Convey("TestAnyExampleMap.ReplaceOrPut", t, func() {
		var k string = "4913813e-ce0d-484a-8099-cbe99bc4dc72"
		var v interface{} = "b5d23328-03cf-4900-989f-5347f728a6f7"
		var x interface{} = "f0db180c-b5a3-49cd-b8c2-ab9133bbf65b"

		test := raml.NewAnyExampleMap(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("01a8f4fe-9305-4c65-a23f-5cc148a5496d", x), ShouldPointTo, test)
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
			var k string = "592bb290-ce7f-4ad4-aba0-008b74c55197"
			var v interface{} = "efac9a23-11cf-42c3-b85b-7dcd70cb085c"

			test := raml.NewAnyExampleMap(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"592bb290-ce7f-4ad4-aba0-008b74c55197","value":"efac9a23-11cf-42c3-b85b-7dcd70cb085c"}]`)
		})

		Convey("Unordered", func() {
			var k string = "592bb290-ce7f-4ad4-aba0-008b74c55197"
			var v interface{} = "efac9a23-11cf-42c3-b85b-7dcd70cb085c"

			test := raml.NewAnyExampleMap(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"592bb290-ce7f-4ad4-aba0-008b74c55197":"efac9a23-11cf-42c3-b85b-7dcd70cb085c"}`)
		})

	})
}

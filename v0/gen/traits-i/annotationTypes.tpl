package raml

type HasAnnotationTypes interface {
	{{define "hasAnnotationTypes"}}
	AnnotationTypes() UntypedMap
	{{- end}}
}

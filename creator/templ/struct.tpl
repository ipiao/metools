package {{.Pkg}}

{{range $i,$e:=$.Imports}}import "{{$e}}"
{{end}}
// {{$.Name}} {{.Comment}}
type {{$.Name}} struct{
    {{range $i,$e:=$.Fields}} {{if not $e.Annoy}}{{$e.Name}}{{end}}  {{$e.Type}}  {{if gt (len $e.Tags) 0}}`{{range $k,$v:=$e.Tags}}{{if ne $k 0}} {{end}}{{$v.Name}}:"{{$v.Value}}"{{end}}` {{end}}
    {{end}}
}
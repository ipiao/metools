package {{.Pkg}}

{{range $e:=$.Imports}}import "{{$e}}"
{{end}}

// {{$.Name}} {{.Comment}}
type {{$.Name}} struct{
    {{range $e:=$.Fields}} {{if not $e.Annoy}}{{$e.Name}}{{end}}  {{$e.Type}}  {{if gt (len $e.Tags) 0}}`{{range $k,$v:=$e.Tags}}{{if ne $k 0}} {{end}}{{$v.Name}}:"{{$v.Value}}"{{end}}` {{end}}
    {{end}}
}

{{ range $f:=$.Funcs }}
func {{if $f.HasReceiver}}({{$f.ReceiverAlias}} {{$f.Receiver}}){{end}}{{$f.Name}}({{range $i,$v:=$f.ArgsIn}}{{$v.Alias}} {{$v.Type}},{{end}}) ({{range $i,$v:=.ArgsOut}}{{if not $v.AliasAnnoy}}{{$v.Alias}} {{end}}{{$v.Type}}, {{end}}){
    {{$f.FuncBody}}
}
{{end}}
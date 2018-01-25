{{define "Func"}}
func {{if .HasReceiver}}({{.ReceiverAlias}} {{.Receiver}}){{end}}{{.Name}}({{range $i,$v:=.ArgsIn}}{{$v.Alias}} {{$v.Type}},{{end}}) ({{range $i,$v:=.ArgsOut}}{{if not $v.AliasAnnoy}}{{$v.Alias}} {{end}}{{$v.Type}}, {{end}}){
    {{.FuncBody}}
}
{{end}}
{%
{{ if not .Large -}} \renewcommand{\arraystretch}{\myNumArrayStretch}% {{- end}}
\setlength{\tabcolsep}{\myLenTabColSep}%
%
{{ .Month.DefineTable .TableType .Large }}
  {{ .Month.MaybeName .Large }}
  {{ if $.Large -}} \hline {{- end }}
  {{ .Month.WeekHeader .Large }} \\ {{ if .Large -}} \noalign{\hrule height \myLenLineThicknessThick} {{- else -}} \hline {{- end}}
  {{- range $i, $week := .Month.Weeks }}
  {{$week.WeekNumber $.Large}} &
    {{- range $j, $day := $week.Days -}}
      {{- $day.Day $.Today $.Large -}}
      {{- if eq $j 6 -}}
        \\ {{ if $.Large -}} \hline {{- end -}}
      {{- else -}} & {{- end -}}
    {{- end -}}
  {{ end }}
  {{- if not .Large -}}
  \multicolumn{8}{l}{ {{- .Month.GetHolidayText .Large -}} }
  {{- end }}
  {{ .Month.EndTable .TableType -}}
  {{ if .Large -}}
  \vspace{2mm}
  {{ .Month.GetHolidayText .Large }}
  {{- end -}}
}
package chart

import _ "embed"

//goland:noinspection GoSnakeCaseUsage
var (
	//go:embed template6/hs.json
	template6_hs string
	//go:embed template6/rks.json
	template6_rks string

	//go:embed template6/qs.json
	template6_qs string
	//go:embed template6/yy.json
	template6_yy string
	//go:embed template6/sx.json
	template6_sx string

	//go:embed template6/map-bubble.json
	template6_map_bubble string
	//go:embed template6/map-heat.json
	template6_map_heat string

	//go:embed template6/jt.json
	template6_jt string
	//go:embed template6/bf.json
	template6_bf string
)

var JsonDataMap = map[string]string{
	"template6_hs":         template6_hs,
	"template6_rks":        template6_rks,
	"template6_qs":         template6_qs,
	"template6_yy":         template6_yy,
	"template6_sx":         template6_sx,
	"template6_map_bubble": template6_map_bubble,
	"template6_map_heat":   template6_map_heat,
	"template6_jt":         template6_jt,
	"template6_bf":         template6_bf,
}

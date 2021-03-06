module httpclout

go 1.16

//replace github.com/andrewarrow/cloutcli => ../cloutcli
replace github.com/btcsuite/btcutil => ../btcutil

require (
	github.com/andrewarrow/cloutcli v0.0.12
	github.com/andrewarrow/mini v0.0.6
	github.com/gin-gonic/gin v1.7.2
	github.com/justincampbell/bigduration v0.0.0-20160531141349-e45bf03c0666 // indirect
	github.com/justincampbell/timeago v0.0.0-20160528003754-027f40306f1d
)

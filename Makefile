parser.go:

%.go: %.rl
	ragel -Z -G2 $*.rl
	sed '/\/\/line/d' -i $@
	gofmt -s -w $@

%.gv: %.rl
	ragel -V -p -o $@ $*.rl

%.png: %.gv
	dot -Tpng -o $@ $*.gv

%.pdf: %.gv
	dot -Tpdf -o $@ $*.gv

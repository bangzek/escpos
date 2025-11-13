%%{
	machine escpos;

	NUL = "NUL" %{ out = append(out, NUL) };
	SOH = "SOH" %{ out = append(out, SOH) };
	STX = "STX" %{ out = append(out, STX) };
	ETX = "ETX" %{ out = append(out, ETX) };
	EOT = "EOT" %{ out = append(out, EOT) };
	ENQ = "ENQ" %{ out = append(out, ENQ) };
	ACK = "ACK" %{ out = append(out, ACK) };
	BEL = "BEL" %{ out = append(out, BEL) };
	BS  = "BS"  %{ out = append(out, BS) };
	HT  = "HT"  %{ out = append(out, HT) };
	LF  = "LF"  %{ out = append(out, LF) };
	VT  = "VT"  %{ out = append(out, VT) };
	FF  = "FF"  %{ out = append(out, FF) };
	CR  = "CR"  %{ out = append(out, CR) };
	SO  = "SO"  %{ out = append(out, SO) };
	SI  = "SI"  %{ out = append(out, SI) };

	DLE = "DLE" %{ out = append(out, DLE) };
	DC1 = "DC1" %{ out = append(out, DC1) };
	DC2 = "DC2" %{ out = append(out, DC2) };
	DC3 = "DC3" %{ out = append(out, DC3) };
	DC4 = "DC4" %{ out = append(out, DC4) };
	NAK = "NAK" %{ out = append(out, NAK) };
	SYN = "SYN" %{ out = append(out, SYN) };
	ETB = "ETB" %{ out = append(out, ETB) };
	CAN = "CAN" %{ out = append(out, CAN) };
	EM  = "EM"  %{ out = append(out, EM) };
	SUB = "SUB" %{ out = append(out, SUB) };
	ESC = "ESC" %{ out = append(out, ESC) };
	FS  = "FS"  %{ out = append(out, FS) };
	GS  = "GS"  %{ out = append(out, GS) };
	RS  = "RS"  %{ out = append(out, RS) };
	US  = "US"  %{ out = append(out, US) };

	SP  = "SP"  %{ out = append(out, SP) };
	DQ  = "DQ"  %{ out = append(out, DQ) };
	SQ  = "SQ"  %{ out = append(out, SQ) };
	DEL = "DEL" %{ out = append(out, DEL) };

	action in_quote {
		out = append(out, fc)
	}
	action assign_hex {
		if fc < 'A' {
			num = fc - '0'
		} else {
			num = fc - 'A' + 10
		}
	}
	action update_hex {
        num <<= 4
		if fc < 'A' {
			num += fc - '0'
		} else {
			num += fc - 'A' + 10
		}
	}
	action assign_num {
		num = fc - '0'
	}
	action update_num {
		num *= 10
		num += fc - '0'
	}
	action append_num {
		out = append(out, num)
	}
	action save_pline {
		pline = p
	}
	action inc_lineno {
		lineno++
	}
	action error {
        for end = p; end < len(data); end++ {
			if data[end] == '\n' {
				break
			}
		}
		return nil, &ParseError{
			LineNo: lineno,
			Column: p - pline + 1,
			State:  cs,
			Line:   string(data[pline:end]),
		}
	}


	comment = /'[^\r\n]*/;
	string  = '"' [ !#-~]+ $in_quote '"';
    hspace  = [ \t]+;

    hex = '0x'
		[0-9A-F] $assign_hex
		[0-9A-F] $update_hex %append_num;
    num = (
		[0-9]     $assign_num

		| [1-9]   $assign_num
			[0-9] $update_num

		| '1'     $assign_num
			[0-9] $update_num
			[0-9] $update_num

		| '2'     $assign_num
			[0-4] $update_num
			[0-9] $update_num

		| '2'     $assign_num
			'5'   $update_num
			[0-5] $update_num
		) %append_num;

	word = NUL | SOH | STX | ETX | EOT | ENQ | ACK | BEL | BS
		| HT | LF | VT | FF | CR | SO | SI
		| DLE | DC1 | DC2 | DC3 | DC4 | NAK | SYN | ETB
		| CAN | EM | SUB | ESC | FS | GS | RS | US
		| SP | DQ | SQ | DEL
        | string
        | hex
        | num
        ;

    line = (
			hspace? (word hspace)* word? hspace?
			|  hspace? (word hspace)* word? comment
		) >save_pline %inc_lineno
		;

    main := (
			line
			| line '\r'? '\n'
			| (line '\r'? '\n')+ line '\r'? '\n'?
		) $!error;
}%%

package escpos

import (
	"fmt"
	"strings"
)

%% write data;

type ParseError struct {
	LineNo int
    Column int
    State  int
	Line   string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parsing error on line %d:%d state %d\n%s\n%s",
		e.LineNo, e.Column, e.State,
		e.Line,
		strings.Repeat(" ", e.Column-1) + "^")
}

func Parse(data []byte) (out []byte, err error) {
	var cs, p int
	%% write init;
    pe := len(data)
	eof := pe
    var num byte
    var lineno, pline, end int

	lineno = 1
	%% write exec;
    return
}

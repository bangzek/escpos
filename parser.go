package escpos

import (
	"fmt"
	"strings"
)

const escpos_start int = 68
const escpos_first_final int = 68
const escpos_error int = 0

const escpos_en_main int = 68

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
		strings.Repeat(" ", e.Column-1)+"^")
}

func Parse(data []byte) (out []byte, err error) {
	var cs, p int

	{
		cs = escpos_start
	}

	pe := len(data)
	eof := pe
	var num byte
	var lineno, pline, end int

	lineno = 1

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 68:
			goto st_case_68
		case 0:
			goto st_case_0
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 81:
			goto st_case_81
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 13:
			goto st_case_13
		case 90:
			goto st_case_90
		case 14:
			goto st_case_14
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 15:
			goto st_case_15
		case 93:
			goto st_case_93
		case 16:
			goto st_case_16
		case 94:
			goto st_case_94
		case 17:
			goto st_case_17
		case 95:
			goto st_case_95
		case 18:
			goto st_case_18
		case 96:
			goto st_case_96
		case 19:
			goto st_case_19
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 20:
			goto st_case_20
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 21:
			goto st_case_21
		case 101:
			goto st_case_101
		case 22:
			goto st_case_22
		case 102:
			goto st_case_102
		case 23:
			goto st_case_23
		case 103:
			goto st_case_103
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 104:
			goto st_case_104
		case 26:
			goto st_case_26
		case 105:
			goto st_case_105
		case 27:
			goto st_case_27
		case 106:
			goto st_case_106
		case 28:
			goto st_case_28
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 29:
			goto st_case_29
		case 112:
			goto st_case_112
		case 30:
			goto st_case_30
		case 113:
			goto st_case_113
		case 31:
			goto st_case_31
		case 114:
			goto st_case_114
		case 32:
			goto st_case_32
		case 115:
			goto st_case_115
		case 33:
			goto st_case_33
		case 116:
			goto st_case_116
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 125:
			goto st_case_125
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 47:
			goto st_case_47
		case 134:
			goto st_case_134
		case 48:
			goto st_case_48
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 49:
			goto st_case_49
		case 137:
			goto st_case_137
		case 50:
			goto st_case_50
		case 138:
			goto st_case_138
		case 51:
			goto st_case_51
		case 139:
			goto st_case_139
		case 52:
			goto st_case_52
		case 140:
			goto st_case_140
		case 53:
			goto st_case_53
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 54:
			goto st_case_54
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 55:
			goto st_case_55
		case 145:
			goto st_case_145
		case 56:
			goto st_case_56
		case 146:
			goto st_case_146
		case 57:
			goto st_case_57
		case 147:
			goto st_case_147
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 148:
			goto st_case_148
		case 60:
			goto st_case_60
		case 149:
			goto st_case_149
		case 61:
			goto st_case_61
		case 150:
			goto st_case_150
		case 62:
			goto st_case_62
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 63:
			goto st_case_63
		case 156:
			goto st_case_156
		case 64:
			goto st_case_64
		case 157:
			goto st_case_157
		case 65:
			goto st_case_65
		case 158:
			goto st_case_158
		case 66:
			goto st_case_66
		case 159:
			goto st_case_159
		case 67:
			goto st_case_67
		case 160:
			goto st_case_160
		}
		goto st_out
	st_case_68:
		switch data[p] {
		case 9:
			goto tr110
		case 10:
			goto tr111
		case 13:
			goto tr112
		case 32:
			goto tr110
		case 34:
			goto tr113
		case 39:
			goto tr114
		case 48:
			goto tr115
		case 49:
			goto tr116
		case 50:
			goto tr117
		case 65:
			goto tr119
		case 66:
			goto tr120
		case 67:
			goto tr121
		case 68:
			goto tr122
		case 69:
			goto tr123
		case 70:
			goto tr124
		case 71:
			goto tr125
		case 72:
			goto tr126
		case 76:
			goto tr127
		case 78:
			goto tr128
		case 82:
			goto tr129
		case 83:
			goto tr130
		case 85:
			goto tr131
		case 86:
			goto tr132
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto tr118
		}
		goto tr0
	tr0:

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

		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr110:

		pline = p

		goto st69
	tr353:

		out = append(out, num)

		goto st69
	tr360:
		out = append(out, ACK)
		goto st69
	tr363:
		out = append(out, BEL)
		goto st69
	tr366:
		out = append(out, BS)
		goto st69
	tr369:
		out = append(out, CAN)
		goto st69
	tr372:
		out = append(out, CR)
		goto st69
	tr375:
		out = append(out, DC1)
		goto st69
	tr378:
		out = append(out, DC2)
		goto st69
	tr381:
		out = append(out, DC3)
		goto st69
	tr384:
		out = append(out, DC4)
		goto st69
	tr387:
		out = append(out, DEL)
		goto st69
	tr390:
		out = append(out, DLE)
		goto st69
	tr393:
		out = append(out, DQ)
		goto st69
	tr396:
		out = append(out, EM)
		goto st69
	tr399:
		out = append(out, ENQ)
		goto st69
	tr402:
		out = append(out, EOT)
		goto st69
	tr405:
		out = append(out, ESC)
		goto st69
	tr408:
		out = append(out, ETB)
		goto st69
	tr411:
		out = append(out, ETX)
		goto st69
	tr414:
		out = append(out, FF)
		goto st69
	tr417:
		out = append(out, FS)
		goto st69
	tr420:
		out = append(out, GS)
		goto st69
	tr423:
		out = append(out, HT)
		goto st69
	tr426:
		out = append(out, LF)
		goto st69
	tr429:
		out = append(out, NAK)
		goto st69
	tr432:
		out = append(out, NUL)
		goto st69
	tr435:
		out = append(out, RS)
		goto st69
	tr438:
		out = append(out, SI)
		goto st69
	tr441:
		out = append(out, SO)
		goto st69
	tr445:
		out = append(out, SOH)
		goto st69
	tr448:
		out = append(out, SP)
		goto st69
	tr451:
		out = append(out, SQ)
		goto st69
	tr454:
		out = append(out, STX)
		goto st69
	tr457:
		out = append(out, SUB)
		goto st69
	tr460:
		out = append(out, SYN)
		goto st69
	tr463:
		out = append(out, US)
		goto st69
	tr466:
		out = append(out, VT)
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 9:
			goto st69
		case 10:
			goto tr134
		case 13:
			goto tr135
		case 32:
			goto st69
		case 34:
			goto st35
		case 39:
			goto st118
		case 48:
			goto tr138
		case 49:
			goto tr139
		case 50:
			goto tr140
		case 65:
			goto st39
		case 66:
			goto st41
		case 67:
			goto st43
		case 68:
			goto st45
		case 69:
			goto st49
		case 70:
			goto st54
		case 71:
			goto st55
		case 72:
			goto st56
		case 76:
			goto st57
		case 78:
			goto st58
		case 82:
			goto st61
		case 83:
			goto st62
		case 85:
			goto st66
		case 86:
			goto st67
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto tr141
		}
		goto tr0
	tr111:

		pline = p

		lineno++

		goto st70
	tr134:

		lineno++

		goto st70
	tr201:

		out = append(out, num)

		lineno++

		goto st70
	tr209:
		out = append(out, ACK)

		lineno++

		goto st70
	tr213:
		out = append(out, BEL)

		lineno++

		goto st70
	tr217:
		out = append(out, BS)

		lineno++

		goto st70
	tr221:
		out = append(out, CAN)

		lineno++

		goto st70
	tr225:
		out = append(out, CR)

		lineno++

		goto st70
	tr229:
		out = append(out, DC1)

		lineno++

		goto st70
	tr233:
		out = append(out, DC2)

		lineno++

		goto st70
	tr237:
		out = append(out, DC3)

		lineno++

		goto st70
	tr241:
		out = append(out, DC4)

		lineno++

		goto st70
	tr245:
		out = append(out, DEL)

		lineno++

		goto st70
	tr249:
		out = append(out, DLE)

		lineno++

		goto st70
	tr253:
		out = append(out, DQ)

		lineno++

		goto st70
	tr257:
		out = append(out, EM)

		lineno++

		goto st70
	tr261:
		out = append(out, ENQ)

		lineno++

		goto st70
	tr265:
		out = append(out, EOT)

		lineno++

		goto st70
	tr269:
		out = append(out, ESC)

		lineno++

		goto st70
	tr273:
		out = append(out, ETB)

		lineno++

		goto st70
	tr277:
		out = append(out, ETX)

		lineno++

		goto st70
	tr281:
		out = append(out, FF)

		lineno++

		goto st70
	tr285:
		out = append(out, FS)

		lineno++

		goto st70
	tr289:
		out = append(out, GS)

		lineno++

		goto st70
	tr293:
		out = append(out, HT)

		lineno++

		goto st70
	tr297:
		out = append(out, LF)

		lineno++

		goto st70
	tr301:
		out = append(out, NAK)

		lineno++

		goto st70
	tr305:
		out = append(out, NUL)

		lineno++

		goto st70
	tr309:
		out = append(out, RS)

		lineno++

		goto st70
	tr313:
		out = append(out, SI)

		lineno++

		goto st70
	tr317:
		out = append(out, SO)

		lineno++

		goto st70
	tr322:
		out = append(out, SOH)

		lineno++

		goto st70
	tr326:
		out = append(out, SP)

		lineno++

		goto st70
	tr330:
		out = append(out, SQ)

		lineno++

		goto st70
	tr334:
		out = append(out, STX)

		lineno++

		goto st70
	tr338:
		out = append(out, SUB)

		lineno++

		goto st70
	tr342:
		out = append(out, SYN)

		lineno++

		goto st70
	tr346:
		out = append(out, US)

		lineno++

		goto st70
	tr350:
		out = append(out, VT)

		lineno++

		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 9:
			goto tr156
		case 10:
			goto tr111
		case 13:
			goto tr157
		case 32:
			goto tr156
		case 34:
			goto tr158
		case 39:
			goto tr159
		case 48:
			goto tr160
		case 49:
			goto tr161
		case 50:
			goto tr162
		case 65:
			goto tr164
		case 66:
			goto tr165
		case 67:
			goto tr166
		case 68:
			goto tr167
		case 69:
			goto tr168
		case 70:
			goto tr169
		case 71:
			goto tr170
		case 72:
			goto tr171
		case 76:
			goto tr172
		case 78:
			goto tr173
		case 82:
			goto tr174
		case 83:
			goto tr175
		case 85:
			goto tr176
		case 86:
			goto tr177
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto tr163
		}
		goto tr0
	tr156:

		pline = p

		goto st71
	tr200:

		out = append(out, num)

		goto st71
	tr208:
		out = append(out, ACK)
		goto st71
	tr212:
		out = append(out, BEL)
		goto st71
	tr216:
		out = append(out, BS)
		goto st71
	tr220:
		out = append(out, CAN)
		goto st71
	tr224:
		out = append(out, CR)
		goto st71
	tr228:
		out = append(out, DC1)
		goto st71
	tr232:
		out = append(out, DC2)
		goto st71
	tr236:
		out = append(out, DC3)
		goto st71
	tr240:
		out = append(out, DC4)
		goto st71
	tr244:
		out = append(out, DEL)
		goto st71
	tr248:
		out = append(out, DLE)
		goto st71
	tr252:
		out = append(out, DQ)
		goto st71
	tr256:
		out = append(out, EM)
		goto st71
	tr260:
		out = append(out, ENQ)
		goto st71
	tr264:
		out = append(out, EOT)
		goto st71
	tr268:
		out = append(out, ESC)
		goto st71
	tr272:
		out = append(out, ETB)
		goto st71
	tr276:
		out = append(out, ETX)
		goto st71
	tr280:
		out = append(out, FF)
		goto st71
	tr284:
		out = append(out, FS)
		goto st71
	tr288:
		out = append(out, GS)
		goto st71
	tr292:
		out = append(out, HT)
		goto st71
	tr296:
		out = append(out, LF)
		goto st71
	tr300:
		out = append(out, NAK)
		goto st71
	tr304:
		out = append(out, NUL)
		goto st71
	tr308:
		out = append(out, RS)
		goto st71
	tr312:
		out = append(out, SI)
		goto st71
	tr316:
		out = append(out, SO)
		goto st71
	tr321:
		out = append(out, SOH)
		goto st71
	tr325:
		out = append(out, SP)
		goto st71
	tr329:
		out = append(out, SQ)
		goto st71
	tr333:
		out = append(out, STX)
		goto st71
	tr337:
		out = append(out, SUB)
		goto st71
	tr341:
		out = append(out, SYN)
		goto st71
	tr345:
		out = append(out, US)
		goto st71
	tr349:
		out = append(out, VT)
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 9:
			goto st71
		case 10:
			goto tr134
		case 13:
			goto tr179
		case 32:
			goto st71
		case 34:
			goto st1
		case 39:
			goto st74
		case 48:
			goto tr182
		case 49:
			goto tr183
		case 50:
			goto tr184
		case 65:
			goto st5
		case 66:
			goto st7
		case 67:
			goto st9
		case 68:
			goto st11
		case 69:
			goto st15
		case 70:
			goto st20
		case 71:
			goto st21
		case 72:
			goto st22
		case 76:
			goto st23
		case 78:
			goto st24
		case 82:
			goto st27
		case 83:
			goto st28
		case 85:
			goto st32
		case 86:
			goto st33
		}
		if 51 <= data[p] && data[p] <= 57 {
			goto tr185
		}
		goto tr0
	tr157:

		pline = p

		lineno++

		goto st72
	tr179:

		lineno++

		goto st72
	tr202:

		out = append(out, num)

		lineno++

		goto st72
	tr210:
		out = append(out, ACK)

		lineno++

		goto st72
	tr214:
		out = append(out, BEL)

		lineno++

		goto st72
	tr218:
		out = append(out, BS)

		lineno++

		goto st72
	tr222:
		out = append(out, CAN)

		lineno++

		goto st72
	tr226:
		out = append(out, CR)

		lineno++

		goto st72
	tr230:
		out = append(out, DC1)

		lineno++

		goto st72
	tr234:
		out = append(out, DC2)

		lineno++

		goto st72
	tr238:
		out = append(out, DC3)

		lineno++

		goto st72
	tr242:
		out = append(out, DC4)

		lineno++

		goto st72
	tr246:
		out = append(out, DEL)

		lineno++

		goto st72
	tr250:
		out = append(out, DLE)

		lineno++

		goto st72
	tr254:
		out = append(out, DQ)

		lineno++

		goto st72
	tr258:
		out = append(out, EM)

		lineno++

		goto st72
	tr262:
		out = append(out, ENQ)

		lineno++

		goto st72
	tr266:
		out = append(out, EOT)

		lineno++

		goto st72
	tr270:
		out = append(out, ESC)

		lineno++

		goto st72
	tr274:
		out = append(out, ETB)

		lineno++

		goto st72
	tr278:
		out = append(out, ETX)

		lineno++

		goto st72
	tr282:
		out = append(out, FF)

		lineno++

		goto st72
	tr286:
		out = append(out, FS)

		lineno++

		goto st72
	tr290:
		out = append(out, GS)

		lineno++

		goto st72
	tr294:
		out = append(out, HT)

		lineno++

		goto st72
	tr298:
		out = append(out, LF)

		lineno++

		goto st72
	tr302:
		out = append(out, NAK)

		lineno++

		goto st72
	tr306:
		out = append(out, NUL)

		lineno++

		goto st72
	tr310:
		out = append(out, RS)

		lineno++

		goto st72
	tr314:
		out = append(out, SI)

		lineno++

		goto st72
	tr318:
		out = append(out, SO)

		lineno++

		goto st72
	tr323:
		out = append(out, SOH)

		lineno++

		goto st72
	tr327:
		out = append(out, SP)

		lineno++

		goto st72
	tr331:
		out = append(out, SQ)

		lineno++

		goto st72
	tr335:
		out = append(out, STX)

		lineno++

		goto st72
	tr339:
		out = append(out, SUB)

		lineno++

		goto st72
	tr343:
		out = append(out, SYN)

		lineno++

		goto st72
	tr347:
		out = append(out, US)

		lineno++

		goto st72
	tr351:
		out = append(out, VT)

		lineno++

		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 10 {
			goto st70
		}
		goto tr0
	tr158:

		pline = p

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 126 {
				goto tr1
			}
		case data[p] >= 32:
			goto tr1
		}
		goto tr0
	tr1:

		out = append(out, data[p])

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 34 {
			goto st73
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto tr1
		}
		goto tr0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 9:
			goto st71
		case 10:
			goto tr134
		case 13:
			goto tr179
		case 32:
			goto st71
		case 39:
			goto st74
		}
		goto tr0
	tr159:

		pline = p

		goto st74
	tr203:

		out = append(out, num)

		goto st74
	tr211:
		out = append(out, ACK)
		goto st74
	tr215:
		out = append(out, BEL)
		goto st74
	tr219:
		out = append(out, BS)
		goto st74
	tr223:
		out = append(out, CAN)
		goto st74
	tr227:
		out = append(out, CR)
		goto st74
	tr231:
		out = append(out, DC1)
		goto st74
	tr235:
		out = append(out, DC2)
		goto st74
	tr239:
		out = append(out, DC3)
		goto st74
	tr243:
		out = append(out, DC4)
		goto st74
	tr247:
		out = append(out, DEL)
		goto st74
	tr251:
		out = append(out, DLE)
		goto st74
	tr255:
		out = append(out, DQ)
		goto st74
	tr259:
		out = append(out, EM)
		goto st74
	tr263:
		out = append(out, ENQ)
		goto st74
	tr267:
		out = append(out, EOT)
		goto st74
	tr271:
		out = append(out, ESC)
		goto st74
	tr275:
		out = append(out, ETB)
		goto st74
	tr279:
		out = append(out, ETX)
		goto st74
	tr283:
		out = append(out, FF)
		goto st74
	tr287:
		out = append(out, FS)
		goto st74
	tr291:
		out = append(out, GS)
		goto st74
	tr295:
		out = append(out, HT)
		goto st74
	tr299:
		out = append(out, LF)
		goto st74
	tr303:
		out = append(out, NAK)
		goto st74
	tr307:
		out = append(out, NUL)
		goto st74
	tr311:
		out = append(out, RS)
		goto st74
	tr315:
		out = append(out, SI)
		goto st74
	tr319:
		out = append(out, SO)
		goto st74
	tr324:
		out = append(out, SOH)
		goto st74
	tr328:
		out = append(out, SP)
		goto st74
	tr332:
		out = append(out, SQ)
		goto st74
	tr336:
		out = append(out, STX)
		goto st74
	tr340:
		out = append(out, SUB)
		goto st74
	tr344:
		out = append(out, SYN)
		goto st74
	tr348:
		out = append(out, US)
		goto st74
	tr352:
		out = append(out, VT)
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 10:
			goto tr134
		case 13:
			goto tr179
		}
		goto st74
	tr160:

		pline = p

		num = data[p] - '0'

		goto st75
	tr182:

		num = data[p] - '0'

		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		case 120:
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto tr3
			}
		case data[p] >= 48:
			goto tr3
		}
		goto tr0
	tr3:

		if data[p] < 'A' {
			num = data[p] - '0'
		} else {
			num = data[p] - 'A' + 10
		}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto tr4
			}
		case data[p] >= 48:
			goto tr4
		}
		goto tr0
	tr4:

		num <<= 4
		if data[p] < 'A' {
			num += data[p] - '0'
		} else {
			num += data[p] - 'A' + 10
		}

		goto st76
	tr206:

		num *= 10
		num += data[p] - '0'

		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		}
		goto tr0
	tr161:

		pline = p

		num = data[p] - '0'

		goto st77
	tr183:

		num = data[p] - '0'

		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr205
		}
		goto tr0
	tr163:

		pline = p

		num = data[p] - '0'

		goto st78
	tr185:

		num = data[p] - '0'

		goto st78
	tr205:

		num *= 10
		num += data[p] - '0'

		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr206
		}
		goto tr0
	tr162:

		pline = p

		num = data[p] - '0'

		goto st79
	tr184:

		num = data[p] - '0'

		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		case 53:
			goto tr207
		}
		switch {
		case data[p] > 52:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr206
			}
		case data[p] >= 48:
			goto tr205
		}
		goto tr0
	tr207:

		num *= 10
		num += data[p] - '0'

		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 9:
			goto tr200
		case 10:
			goto tr201
		case 13:
			goto tr202
		case 32:
			goto tr200
		case 39:
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto tr206
		}
		goto tr0
	tr164:

		pline = p

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 67 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 75 {
			goto st81
		}
		goto tr0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 9:
			goto tr208
		case 10:
			goto tr209
		case 13:
			goto tr210
		case 32:
			goto tr208
		case 39:
			goto tr211
		}
		goto tr0
	tr165:

		pline = p

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 69:
			goto st8
		case 83:
			goto st83
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 76 {
			goto st82
		}
		goto tr0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 9:
			goto tr212
		case 10:
			goto tr213
		case 13:
			goto tr214
		case 32:
			goto tr212
		case 39:
			goto tr215
		}
		goto tr0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 9:
			goto tr216
		case 10:
			goto tr217
		case 13:
			goto tr218
		case 32:
			goto tr216
		case 39:
			goto tr219
		}
		goto tr0
	tr166:

		pline = p

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 65:
			goto st10
		case 82:
			goto st85
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 78 {
			goto st84
		}
		goto tr0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 9:
			goto tr220
		case 10:
			goto tr221
		case 13:
			goto tr222
		case 32:
			goto tr220
		case 39:
			goto tr223
		}
		goto tr0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 9:
			goto tr224
		case 10:
			goto tr225
		case 13:
			goto tr226
		case 32:
			goto tr224
		case 39:
			goto tr227
		}
		goto tr0
	tr167:

		pline = p

		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 67:
			goto st12
		case 69:
			goto st13
		case 76:
			goto st14
		case 81:
			goto st92
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 49:
			goto st86
		case 50:
			goto st87
		case 51:
			goto st88
		case 52:
			goto st89
		}
		goto tr0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 9:
			goto tr228
		case 10:
			goto tr229
		case 13:
			goto tr230
		case 32:
			goto tr228
		case 39:
			goto tr231
		}
		goto tr0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 9:
			goto tr232
		case 10:
			goto tr233
		case 13:
			goto tr234
		case 32:
			goto tr232
		case 39:
			goto tr235
		}
		goto tr0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 9:
			goto tr236
		case 10:
			goto tr237
		case 13:
			goto tr238
		case 32:
			goto tr236
		case 39:
			goto tr239
		}
		goto tr0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 9:
			goto tr240
		case 10:
			goto tr241
		case 13:
			goto tr242
		case 32:
			goto tr240
		case 39:
			goto tr243
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 76 {
			goto st90
		}
		goto tr0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 9:
			goto tr244
		case 10:
			goto tr245
		case 13:
			goto tr246
		case 32:
			goto tr244
		case 39:
			goto tr247
		}
		goto tr0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 69 {
			goto st91
		}
		goto tr0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 9:
			goto tr248
		case 10:
			goto tr249
		case 13:
			goto tr250
		case 32:
			goto tr248
		case 39:
			goto tr251
		}
		goto tr0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 9:
			goto tr252
		case 10:
			goto tr253
		case 13:
			goto tr254
		case 32:
			goto tr252
		case 39:
			goto tr255
		}
		goto tr0
	tr168:

		pline = p

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 77:
			goto st93
		case 78:
			goto st16
		case 79:
			goto st17
		case 83:
			goto st18
		case 84:
			goto st19
		}
		goto tr0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 9:
			goto tr256
		case 10:
			goto tr257
		case 13:
			goto tr258
		case 32:
			goto tr256
		case 39:
			goto tr259
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 81 {
			goto st94
		}
		goto tr0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 9:
			goto tr260
		case 10:
			goto tr261
		case 13:
			goto tr262
		case 32:
			goto tr260
		case 39:
			goto tr263
		}
		goto tr0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 84 {
			goto st95
		}
		goto tr0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 9:
			goto tr264
		case 10:
			goto tr265
		case 13:
			goto tr266
		case 32:
			goto tr264
		case 39:
			goto tr267
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 67 {
			goto st96
		}
		goto tr0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 9:
			goto tr268
		case 10:
			goto tr269
		case 13:
			goto tr270
		case 32:
			goto tr268
		case 39:
			goto tr271
		}
		goto tr0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 66:
			goto st97
		case 88:
			goto st98
		}
		goto tr0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 9:
			goto tr272
		case 10:
			goto tr273
		case 13:
			goto tr274
		case 32:
			goto tr272
		case 39:
			goto tr275
		}
		goto tr0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 9:
			goto tr276
		case 10:
			goto tr277
		case 13:
			goto tr278
		case 32:
			goto tr276
		case 39:
			goto tr279
		}
		goto tr0
	tr169:

		pline = p

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 70:
			goto st99
		case 83:
			goto st100
		}
		goto tr0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 9:
			goto tr280
		case 10:
			goto tr281
		case 13:
			goto tr282
		case 32:
			goto tr280
		case 39:
			goto tr283
		}
		goto tr0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 9:
			goto tr284
		case 10:
			goto tr285
		case 13:
			goto tr286
		case 32:
			goto tr284
		case 39:
			goto tr287
		}
		goto tr0
	tr170:

		pline = p

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 83 {
			goto st101
		}
		goto tr0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 9:
			goto tr288
		case 10:
			goto tr289
		case 13:
			goto tr290
		case 32:
			goto tr288
		case 39:
			goto tr291
		}
		goto tr0
	tr171:

		pline = p

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 84 {
			goto st102
		}
		goto tr0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 9:
			goto tr292
		case 10:
			goto tr293
		case 13:
			goto tr294
		case 32:
			goto tr292
		case 39:
			goto tr295
		}
		goto tr0
	tr172:

		pline = p

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 70 {
			goto st103
		}
		goto tr0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 9:
			goto tr296
		case 10:
			goto tr297
		case 13:
			goto tr298
		case 32:
			goto tr296
		case 39:
			goto tr299
		}
		goto tr0
	tr173:

		pline = p

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 65:
			goto st25
		case 85:
			goto st26
		}
		goto tr0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 75 {
			goto st104
		}
		goto tr0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 9:
			goto tr300
		case 10:
			goto tr301
		case 13:
			goto tr302
		case 32:
			goto tr300
		case 39:
			goto tr303
		}
		goto tr0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 76 {
			goto st105
		}
		goto tr0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 9:
			goto tr304
		case 10:
			goto tr305
		case 13:
			goto tr306
		case 32:
			goto tr304
		case 39:
			goto tr307
		}
		goto tr0
	tr174:

		pline = p

		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 83 {
			goto st106
		}
		goto tr0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 9:
			goto tr308
		case 10:
			goto tr309
		case 13:
			goto tr310
		case 32:
			goto tr308
		case 39:
			goto tr311
		}
		goto tr0
	tr175:

		pline = p

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 73:
			goto st107
		case 79:
			goto st108
		case 80:
			goto st110
		case 81:
			goto st111
		case 84:
			goto st29
		case 85:
			goto st30
		case 89:
			goto st31
		}
		goto tr0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 9:
			goto tr312
		case 10:
			goto tr313
		case 13:
			goto tr314
		case 32:
			goto tr312
		case 39:
			goto tr315
		}
		goto tr0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 9:
			goto tr316
		case 10:
			goto tr317
		case 13:
			goto tr318
		case 32:
			goto tr316
		case 39:
			goto tr319
		case 72:
			goto st109
		}
		goto tr0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 9:
			goto tr321
		case 10:
			goto tr322
		case 13:
			goto tr323
		case 32:
			goto tr321
		case 39:
			goto tr324
		}
		goto tr0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 9:
			goto tr325
		case 10:
			goto tr326
		case 13:
			goto tr327
		case 32:
			goto tr325
		case 39:
			goto tr328
		}
		goto tr0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 9:
			goto tr329
		case 10:
			goto tr330
		case 13:
			goto tr331
		case 32:
			goto tr329
		case 39:
			goto tr332
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 88 {
			goto st112
		}
		goto tr0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 9:
			goto tr333
		case 10:
			goto tr334
		case 13:
			goto tr335
		case 32:
			goto tr333
		case 39:
			goto tr336
		}
		goto tr0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 66 {
			goto st113
		}
		goto tr0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 9:
			goto tr337
		case 10:
			goto tr338
		case 13:
			goto tr339
		case 32:
			goto tr337
		case 39:
			goto tr340
		}
		goto tr0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 78 {
			goto st114
		}
		goto tr0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 9:
			goto tr341
		case 10:
			goto tr342
		case 13:
			goto tr343
		case 32:
			goto tr341
		case 39:
			goto tr344
		}
		goto tr0
	tr176:

		pline = p

		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 83 {
			goto st115
		}
		goto tr0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 9:
			goto tr345
		case 10:
			goto tr346
		case 13:
			goto tr347
		case 32:
			goto tr345
		case 39:
			goto tr348
		}
		goto tr0
	tr177:

		pline = p

		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 84 {
			goto st116
		}
		goto tr0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 9:
			goto tr349
		case 10:
			goto tr350
		case 13:
			goto tr351
		case 32:
			goto tr349
		case 39:
			goto tr352
		}
		goto tr0
	tr112:

		pline = p

		lineno++

		goto st34
	tr135:

		lineno++

		goto st34
	tr354:

		out = append(out, num)

		lineno++

		goto st34
	tr361:
		out = append(out, ACK)

		lineno++

		goto st34
	tr364:
		out = append(out, BEL)

		lineno++

		goto st34
	tr367:
		out = append(out, BS)

		lineno++

		goto st34
	tr370:
		out = append(out, CAN)

		lineno++

		goto st34
	tr373:
		out = append(out, CR)

		lineno++

		goto st34
	tr376:
		out = append(out, DC1)

		lineno++

		goto st34
	tr379:
		out = append(out, DC2)

		lineno++

		goto st34
	tr382:
		out = append(out, DC3)

		lineno++

		goto st34
	tr385:
		out = append(out, DC4)

		lineno++

		goto st34
	tr388:
		out = append(out, DEL)

		lineno++

		goto st34
	tr391:
		out = append(out, DLE)

		lineno++

		goto st34
	tr394:
		out = append(out, DQ)

		lineno++

		goto st34
	tr397:
		out = append(out, EM)

		lineno++

		goto st34
	tr400:
		out = append(out, ENQ)

		lineno++

		goto st34
	tr403:
		out = append(out, EOT)

		lineno++

		goto st34
	tr406:
		out = append(out, ESC)

		lineno++

		goto st34
	tr409:
		out = append(out, ETB)

		lineno++

		goto st34
	tr412:
		out = append(out, ETX)

		lineno++

		goto st34
	tr415:
		out = append(out, FF)

		lineno++

		goto st34
	tr418:
		out = append(out, FS)

		lineno++

		goto st34
	tr421:
		out = append(out, GS)

		lineno++

		goto st34
	tr424:
		out = append(out, HT)

		lineno++

		goto st34
	tr427:
		out = append(out, LF)

		lineno++

		goto st34
	tr430:
		out = append(out, NAK)

		lineno++

		goto st34
	tr433:
		out = append(out, NUL)

		lineno++

		goto st34
	tr436:
		out = append(out, RS)

		lineno++

		goto st34
	tr439:
		out = append(out, SI)

		lineno++

		goto st34
	tr442:
		out = append(out, SO)

		lineno++

		goto st34
	tr446:
		out = append(out, SOH)

		lineno++

		goto st34
	tr449:
		out = append(out, SP)

		lineno++

		goto st34
	tr452:
		out = append(out, SQ)

		lineno++

		goto st34
	tr455:
		out = append(out, STX)

		lineno++

		goto st34
	tr458:
		out = append(out, SUB)

		lineno++

		goto st34
	tr461:
		out = append(out, SYN)

		lineno++

		goto st34
	tr464:
		out = append(out, US)

		lineno++

		goto st34
	tr467:
		out = append(out, VT)

		lineno++

		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 10 {
			goto st70
		}
		goto tr0
	tr113:

		pline = p

		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 126 {
				goto tr56
			}
		case data[p] >= 32:
			goto tr56
		}
		goto tr0
	tr56:

		out = append(out, data[p])

		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 34 {
			goto st117
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto tr56
		}
		goto tr0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 9:
			goto st69
		case 10:
			goto tr134
		case 13:
			goto tr135
		case 32:
			goto st69
		case 39:
			goto st118
		}
		goto tr0
	tr114:

		pline = p

		goto st118
	tr355:

		out = append(out, num)

		goto st118
	tr362:
		out = append(out, ACK)
		goto st118
	tr365:
		out = append(out, BEL)
		goto st118
	tr368:
		out = append(out, BS)
		goto st118
	tr371:
		out = append(out, CAN)
		goto st118
	tr374:
		out = append(out, CR)
		goto st118
	tr377:
		out = append(out, DC1)
		goto st118
	tr380:
		out = append(out, DC2)
		goto st118
	tr383:
		out = append(out, DC3)
		goto st118
	tr386:
		out = append(out, DC4)
		goto st118
	tr389:
		out = append(out, DEL)
		goto st118
	tr392:
		out = append(out, DLE)
		goto st118
	tr395:
		out = append(out, DQ)
		goto st118
	tr398:
		out = append(out, EM)
		goto st118
	tr401:
		out = append(out, ENQ)
		goto st118
	tr404:
		out = append(out, EOT)
		goto st118
	tr407:
		out = append(out, ESC)
		goto st118
	tr410:
		out = append(out, ETB)
		goto st118
	tr413:
		out = append(out, ETX)
		goto st118
	tr416:
		out = append(out, FF)
		goto st118
	tr419:
		out = append(out, FS)
		goto st118
	tr422:
		out = append(out, GS)
		goto st118
	tr425:
		out = append(out, HT)
		goto st118
	tr428:
		out = append(out, LF)
		goto st118
	tr431:
		out = append(out, NAK)
		goto st118
	tr434:
		out = append(out, NUL)
		goto st118
	tr437:
		out = append(out, RS)
		goto st118
	tr440:
		out = append(out, SI)
		goto st118
	tr443:
		out = append(out, SO)
		goto st118
	tr447:
		out = append(out, SOH)
		goto st118
	tr450:
		out = append(out, SP)
		goto st118
	tr453:
		out = append(out, SQ)
		goto st118
	tr456:
		out = append(out, STX)
		goto st118
	tr459:
		out = append(out, SUB)
		goto st118
	tr462:
		out = append(out, SYN)
		goto st118
	tr465:
		out = append(out, US)
		goto st118
	tr468:
		out = append(out, VT)
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 10:
			goto tr134
		case 13:
			goto tr135
		}
		goto st118
	tr115:

		pline = p

		num = data[p] - '0'

		goto st119
	tr138:

		num = data[p] - '0'

		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		case 120:
			goto st37
		}
		goto tr0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto tr58
			}
		case data[p] >= 48:
			goto tr58
		}
		goto tr0
	tr58:

		if data[p] < 'A' {
			num = data[p] - '0'
		} else {
			num = data[p] - 'A' + 10
		}

		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto tr59
			}
		case data[p] >= 48:
			goto tr59
		}
		goto tr0
	tr59:

		num <<= 4
		if data[p] < 'A' {
			num += data[p] - '0'
		} else {
			num += data[p] - 'A' + 10
		}

		goto st120
	tr358:

		num *= 10
		num += data[p] - '0'

		goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		}
		goto tr0
	tr116:

		pline = p

		num = data[p] - '0'

		goto st121
	tr139:

		num = data[p] - '0'

		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr357
		}
		goto tr0
	tr118:

		pline = p

		num = data[p] - '0'

		goto st122
	tr141:

		num = data[p] - '0'

		goto st122
	tr357:

		num *= 10
		num += data[p] - '0'

		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr358
		}
		goto tr0
	tr117:

		pline = p

		num = data[p] - '0'

		goto st123
	tr140:

		num = data[p] - '0'

		goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		case 53:
			goto tr359
		}
		switch {
		case data[p] > 52:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr358
			}
		case data[p] >= 48:
			goto tr357
		}
		goto tr0
	tr359:

		num *= 10
		num += data[p] - '0'

		goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 9:
			goto tr353
		case 10:
			goto tr201
		case 13:
			goto tr354
		case 32:
			goto tr353
		case 39:
			goto tr355
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto tr358
		}
		goto tr0
	tr119:

		pline = p

		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 67 {
			goto st40
		}
		goto tr0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 75 {
			goto st125
		}
		goto tr0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 9:
			goto tr360
		case 10:
			goto tr209
		case 13:
			goto tr361
		case 32:
			goto tr360
		case 39:
			goto tr362
		}
		goto tr0
	tr120:

		pline = p

		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 69:
			goto st42
		case 83:
			goto st127
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 76 {
			goto st126
		}
		goto tr0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 9:
			goto tr363
		case 10:
			goto tr213
		case 13:
			goto tr364
		case 32:
			goto tr363
		case 39:
			goto tr365
		}
		goto tr0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 9:
			goto tr366
		case 10:
			goto tr217
		case 13:
			goto tr367
		case 32:
			goto tr366
		case 39:
			goto tr368
		}
		goto tr0
	tr121:

		pline = p

		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 65:
			goto st44
		case 82:
			goto st129
		}
		goto tr0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 78 {
			goto st128
		}
		goto tr0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 9:
			goto tr369
		case 10:
			goto tr221
		case 13:
			goto tr370
		case 32:
			goto tr369
		case 39:
			goto tr371
		}
		goto tr0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 9:
			goto tr372
		case 10:
			goto tr225
		case 13:
			goto tr373
		case 32:
			goto tr372
		case 39:
			goto tr374
		}
		goto tr0
	tr122:

		pline = p

		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 67:
			goto st46
		case 69:
			goto st47
		case 76:
			goto st48
		case 81:
			goto st136
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 49:
			goto st130
		case 50:
			goto st131
		case 51:
			goto st132
		case 52:
			goto st133
		}
		goto tr0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 9:
			goto tr375
		case 10:
			goto tr229
		case 13:
			goto tr376
		case 32:
			goto tr375
		case 39:
			goto tr377
		}
		goto tr0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 9:
			goto tr378
		case 10:
			goto tr233
		case 13:
			goto tr379
		case 32:
			goto tr378
		case 39:
			goto tr380
		}
		goto tr0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 9:
			goto tr381
		case 10:
			goto tr237
		case 13:
			goto tr382
		case 32:
			goto tr381
		case 39:
			goto tr383
		}
		goto tr0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 9:
			goto tr384
		case 10:
			goto tr241
		case 13:
			goto tr385
		case 32:
			goto tr384
		case 39:
			goto tr386
		}
		goto tr0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 76 {
			goto st134
		}
		goto tr0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 9:
			goto tr387
		case 10:
			goto tr245
		case 13:
			goto tr388
		case 32:
			goto tr387
		case 39:
			goto tr389
		}
		goto tr0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 69 {
			goto st135
		}
		goto tr0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 9:
			goto tr390
		case 10:
			goto tr249
		case 13:
			goto tr391
		case 32:
			goto tr390
		case 39:
			goto tr392
		}
		goto tr0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 9:
			goto tr393
		case 10:
			goto tr253
		case 13:
			goto tr394
		case 32:
			goto tr393
		case 39:
			goto tr395
		}
		goto tr0
	tr123:

		pline = p

		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 77:
			goto st137
		case 78:
			goto st50
		case 79:
			goto st51
		case 83:
			goto st52
		case 84:
			goto st53
		}
		goto tr0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 9:
			goto tr396
		case 10:
			goto tr257
		case 13:
			goto tr397
		case 32:
			goto tr396
		case 39:
			goto tr398
		}
		goto tr0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 81 {
			goto st138
		}
		goto tr0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 9:
			goto tr399
		case 10:
			goto tr261
		case 13:
			goto tr400
		case 32:
			goto tr399
		case 39:
			goto tr401
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 84 {
			goto st139
		}
		goto tr0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 9:
			goto tr402
		case 10:
			goto tr265
		case 13:
			goto tr403
		case 32:
			goto tr402
		case 39:
			goto tr404
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 67 {
			goto st140
		}
		goto tr0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 9:
			goto tr405
		case 10:
			goto tr269
		case 13:
			goto tr406
		case 32:
			goto tr405
		case 39:
			goto tr407
		}
		goto tr0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 66:
			goto st141
		case 88:
			goto st142
		}
		goto tr0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 9:
			goto tr408
		case 10:
			goto tr273
		case 13:
			goto tr409
		case 32:
			goto tr408
		case 39:
			goto tr410
		}
		goto tr0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 9:
			goto tr411
		case 10:
			goto tr277
		case 13:
			goto tr412
		case 32:
			goto tr411
		case 39:
			goto tr413
		}
		goto tr0
	tr124:

		pline = p

		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 70:
			goto st143
		case 83:
			goto st144
		}
		goto tr0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 9:
			goto tr414
		case 10:
			goto tr281
		case 13:
			goto tr415
		case 32:
			goto tr414
		case 39:
			goto tr416
		}
		goto tr0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 9:
			goto tr417
		case 10:
			goto tr285
		case 13:
			goto tr418
		case 32:
			goto tr417
		case 39:
			goto tr419
		}
		goto tr0
	tr125:

		pline = p

		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 83 {
			goto st145
		}
		goto tr0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 9:
			goto tr420
		case 10:
			goto tr289
		case 13:
			goto tr421
		case 32:
			goto tr420
		case 39:
			goto tr422
		}
		goto tr0
	tr126:

		pline = p

		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 84 {
			goto st146
		}
		goto tr0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 9:
			goto tr423
		case 10:
			goto tr293
		case 13:
			goto tr424
		case 32:
			goto tr423
		case 39:
			goto tr425
		}
		goto tr0
	tr127:

		pline = p

		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 70 {
			goto st147
		}
		goto tr0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 9:
			goto tr426
		case 10:
			goto tr297
		case 13:
			goto tr427
		case 32:
			goto tr426
		case 39:
			goto tr428
		}
		goto tr0
	tr128:

		pline = p

		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 65:
			goto st59
		case 85:
			goto st60
		}
		goto tr0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 75 {
			goto st148
		}
		goto tr0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 9:
			goto tr429
		case 10:
			goto tr301
		case 13:
			goto tr430
		case 32:
			goto tr429
		case 39:
			goto tr431
		}
		goto tr0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 76 {
			goto st149
		}
		goto tr0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 9:
			goto tr432
		case 10:
			goto tr305
		case 13:
			goto tr433
		case 32:
			goto tr432
		case 39:
			goto tr434
		}
		goto tr0
	tr129:

		pline = p

		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 83 {
			goto st150
		}
		goto tr0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 9:
			goto tr435
		case 10:
			goto tr309
		case 13:
			goto tr436
		case 32:
			goto tr435
		case 39:
			goto tr437
		}
		goto tr0
	tr130:

		pline = p

		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 73:
			goto st151
		case 79:
			goto st152
		case 80:
			goto st154
		case 81:
			goto st155
		case 84:
			goto st63
		case 85:
			goto st64
		case 89:
			goto st65
		}
		goto tr0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 9:
			goto tr438
		case 10:
			goto tr313
		case 13:
			goto tr439
		case 32:
			goto tr438
		case 39:
			goto tr440
		}
		goto tr0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 9:
			goto tr441
		case 10:
			goto tr317
		case 13:
			goto tr442
		case 32:
			goto tr441
		case 39:
			goto tr443
		case 72:
			goto st153
		}
		goto tr0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 9:
			goto tr445
		case 10:
			goto tr322
		case 13:
			goto tr446
		case 32:
			goto tr445
		case 39:
			goto tr447
		}
		goto tr0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 9:
			goto tr448
		case 10:
			goto tr326
		case 13:
			goto tr449
		case 32:
			goto tr448
		case 39:
			goto tr450
		}
		goto tr0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 9:
			goto tr451
		case 10:
			goto tr330
		case 13:
			goto tr452
		case 32:
			goto tr451
		case 39:
			goto tr453
		}
		goto tr0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 88 {
			goto st156
		}
		goto tr0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 9:
			goto tr454
		case 10:
			goto tr334
		case 13:
			goto tr455
		case 32:
			goto tr454
		case 39:
			goto tr456
		}
		goto tr0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 66 {
			goto st157
		}
		goto tr0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 9:
			goto tr457
		case 10:
			goto tr338
		case 13:
			goto tr458
		case 32:
			goto tr457
		case 39:
			goto tr459
		}
		goto tr0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 78 {
			goto st158
		}
		goto tr0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 9:
			goto tr460
		case 10:
			goto tr342
		case 13:
			goto tr461
		case 32:
			goto tr460
		case 39:
			goto tr462
		}
		goto tr0
	tr131:

		pline = p

		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 83 {
			goto st159
		}
		goto tr0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 9:
			goto tr463
		case 10:
			goto tr346
		case 13:
			goto tr464
		case 32:
			goto tr463
		case 39:
			goto tr465
		}
		goto tr0
	tr132:

		pline = p

		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 84 {
			goto st160
		}
		goto tr0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 9:
			goto tr466
		case 10:
			goto tr350
		case 13:
			goto tr467
		case 32:
			goto tr466
		case 39:
			goto tr468
		}
		goto tr0
	st_out:
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof156:
		cs = 156
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof159:
		cs = 159
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof160:
		cs = 160
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 69, 71, 73, 74, 117, 118:

				lineno++

			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67:

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

			case 105, 149:
				out = append(out, NUL)

				lineno++

			case 109, 153:
				out = append(out, SOH)

				lineno++

			case 112, 156:
				out = append(out, STX)

				lineno++

			case 98, 142:
				out = append(out, ETX)

				lineno++

			case 95, 139:
				out = append(out, EOT)

				lineno++

			case 94, 138:
				out = append(out, ENQ)

				lineno++

			case 81, 125:
				out = append(out, ACK)

				lineno++

			case 82, 126:
				out = append(out, BEL)

				lineno++

			case 83, 127:
				out = append(out, BS)

				lineno++

			case 102, 146:
				out = append(out, HT)

				lineno++

			case 103, 147:
				out = append(out, LF)

				lineno++

			case 116, 160:
				out = append(out, VT)

				lineno++

			case 99, 143:
				out = append(out, FF)

				lineno++

			case 85, 129:
				out = append(out, CR)

				lineno++

			case 108, 152:
				out = append(out, SO)

				lineno++

			case 107, 151:
				out = append(out, SI)

				lineno++

			case 91, 135:
				out = append(out, DLE)

				lineno++

			case 86, 130:
				out = append(out, DC1)

				lineno++

			case 87, 131:
				out = append(out, DC2)

				lineno++

			case 88, 132:
				out = append(out, DC3)

				lineno++

			case 89, 133:
				out = append(out, DC4)

				lineno++

			case 104, 148:
				out = append(out, NAK)

				lineno++

			case 114, 158:
				out = append(out, SYN)

				lineno++

			case 97, 141:
				out = append(out, ETB)

				lineno++

			case 84, 128:
				out = append(out, CAN)

				lineno++

			case 93, 137:
				out = append(out, EM)

				lineno++

			case 113, 157:
				out = append(out, SUB)

				lineno++

			case 96, 140:
				out = append(out, ESC)

				lineno++

			case 100, 144:
				out = append(out, FS)

				lineno++

			case 101, 145:
				out = append(out, GS)

				lineno++

			case 106, 150:
				out = append(out, RS)

				lineno++

			case 115, 159:
				out = append(out, US)

				lineno++

			case 110, 154:
				out = append(out, SP)

				lineno++

			case 92, 136:
				out = append(out, DQ)

				lineno++

			case 111, 155:
				out = append(out, SQ)

				lineno++

			case 90, 134:
				out = append(out, DEL)

				lineno++

			case 75, 76, 77, 78, 79, 80, 119, 120, 121, 122, 123, 124:

				out = append(out, num)

				lineno++

			case 68, 70:

				pline = p

				lineno++

			}
		}

	_out:
		{
		}
	}

	return
}

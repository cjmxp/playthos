package ui

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Font struct {
	charmap map[string]mgl32.Vec4
}

func (f *Font) GetVec4(char string) mgl32.Vec4 {
	return f.charmap[char]
}

func (f *Font) NewFont() {
	f.charmap = make(map[string]mgl32.Vec4)
	f.charmap["\n"] = mgl32.Vec4{0x000000, 0x000000, 0x000000, 1}
	f.charmap[" "] = mgl32.Vec4{0x000000, 0x000000, 0x000000, 0x000000}
	f.charmap["0"] = mgl32.Vec4{0x007CC6, 0xD6D6D6, 0xD6D6C6, 0x7C0000}
	f.charmap["1"] = mgl32.Vec4{0x001030, 0xF03030, 0x303030, 0xFC0000}
	f.charmap["2"] = mgl32.Vec4{0x0078CC, 0xCC0C18, 0x3060CC, 0xFC0000}
	f.charmap["3"] = mgl32.Vec4{0x0078CC, 0x0C0C38, 0x0C0CCC, 0x780000}
	f.charmap["4"] = mgl32.Vec4{0x000C1C, 0x3C6CCC, 0xFE0C0C, 0x1E0000}
	f.charmap["5"] = mgl32.Vec4{0x00FCC0, 0xC0C0F8, 0x0C0CCC, 0x780000}
	f.charmap["6"] = mgl32.Vec4{0x003860, 0xC0C0F8, 0xCCCCCC, 0x780000}
	f.charmap["7"] = mgl32.Vec4{0x00FEC6, 0xC6060C, 0x183030, 0x300000}
	f.charmap["8"] = mgl32.Vec4{0x0078CC, 0xCCEC78, 0xDCCCCC, 0x780000}
	f.charmap["9"] = mgl32.Vec4{0x0078CC, 0xCCCC7C, 0x181830, 0x700000}
	f.charmap["A"] = mgl32.Vec4{0x003078, 0xCCCCCC, 0xFCCCCC, 0xCC0000}
	f.charmap["B"] = mgl32.Vec4{0x00FC66, 0x66667C, 0x666666, 0xFC0000}
	f.charmap["C"] = mgl32.Vec4{0x003C66, 0xC6C0C0, 0xC0C666, 0x3C0000}
	f.charmap["D"] = mgl32.Vec4{0x00F86C, 0x666666, 0x66666C, 0xF80000}
	f.charmap["E"] = mgl32.Vec4{0x00FE62, 0x60647C, 0x646062, 0xFE0000}
	f.charmap["F"] = mgl32.Vec4{0x00FE66, 0x62647C, 0x646060, 0xF00000}
	f.charmap["G"] = mgl32.Vec4{0x003C66, 0xC6C0C0, 0xCEC666, 0x3E0000}
	f.charmap["H"] = mgl32.Vec4{0x00CCCC, 0xCCCCFC, 0xCCCCCC, 0xCC0000}
	f.charmap["I"] = mgl32.Vec4{0x007830, 0x303030, 0x303030, 0x780000}
	f.charmap["J"] = mgl32.Vec4{0x001E0C, 0x0C0C0C, 0xCCCCCC, 0x780000}
	f.charmap["K"] = mgl32.Vec4{0x00E666, 0x6C6C78, 0x6C6C66, 0xE60000}
	f.charmap["L"] = mgl32.Vec4{0x00F060, 0x606060, 0x626666, 0xFE0000}
	f.charmap["M"] = mgl32.Vec4{0x00C6EE, 0xFEFED6, 0xC6C6C6, 0xC60000}
	f.charmap["N"] = mgl32.Vec4{0x00C6C6, 0xE6F6FE, 0xDECEC6, 0xC60000}
	f.charmap["O"] = mgl32.Vec4{0x00386C, 0xC6C6C6, 0xC6C66C, 0x380000}
	f.charmap["P"] = mgl32.Vec4{0x00FC66, 0x66667C, 0x606060, 0xF00000}
	f.charmap["Q"] = mgl32.Vec4{0x00386C, 0xC6C6C6, 0xCEDE7C, 0x0C1E00}
	f.charmap["R"] = mgl32.Vec4{0x00FC66, 0x66667C, 0x6C6666, 0xE60000}
	f.charmap["S"] = mgl32.Vec4{0x0078CC, 0xCCC070, 0x18CCCC, 0x780000}
	f.charmap["T"] = mgl32.Vec4{0x00FCB4, 0x303030, 0x303030, 0x780000}
	f.charmap["U"] = mgl32.Vec4{0x00CCCC, 0xCCCCCC, 0xCCCCCC, 0x780000}
	f.charmap["V"] = mgl32.Vec4{0x00CCCC, 0xCCCCCC, 0xCCCC78, 0x300000}
	f.charmap["W"] = mgl32.Vec4{0x00C6C6, 0xC6C6D6, 0xD66C6C, 0x6C0000}
	f.charmap["X"] = mgl32.Vec4{0x00CCCC, 0xCC7830, 0x78CCCC, 0xCC0000}
	f.charmap["Y"] = mgl32.Vec4{0x00CCCC, 0xCCCC78, 0x303030, 0x780000}
	f.charmap["Z"] = mgl32.Vec4{0x00FECE, 0x981830, 0x6062C6, 0xFE0000}
	f.charmap["a"] = mgl32.Vec4{0x000000, 0x00780C, 0x7CCCCC, 0x760000}
	f.charmap["b"] = mgl32.Vec4{0x00E060, 0x607C66, 0x666666, 0xDC0000}
	f.charmap["c"] = mgl32.Vec4{0x000000, 0x0078CC, 0xC0C0CC, 0x780000}
	f.charmap["d"] = mgl32.Vec4{0x001C0C, 0x0C7CCC, 0xCCCCCC, 0x760000}
	f.charmap["e"] = mgl32.Vec4{0x000000, 0x0078CC, 0xFCC0CC, 0x780000}
	f.charmap["f"] = mgl32.Vec4{0x00386C, 0x6060F8, 0x606060, 0xF00000}
	f.charmap["g"] = mgl32.Vec4{0x000000, 0x0076CC, 0xCCCC7C, 0x0CCC78}
	f.charmap["h"] = mgl32.Vec4{0x00E060, 0x606C76, 0x666666, 0xE60000}
	f.charmap["i"] = mgl32.Vec4{0x001818, 0x007818, 0x181818, 0x7E0000}
	f.charmap["j"] = mgl32.Vec4{0x000C0C, 0x003C0C, 0x0C0C0C, 0xCCCC78}
	f.charmap["k"] = mgl32.Vec4{0x00E060, 0x60666C, 0x786C66, 0xE60000}
	f.charmap["l"] = mgl32.Vec4{0x007818, 0x181818, 0x181818, 0x7E0000}
	f.charmap["m"] = mgl32.Vec4{0x000000, 0x00FCD6, 0xD6D6D6, 0xC60000}
	f.charmap["n"] = mgl32.Vec4{0x000000, 0x00F8CC, 0xCCCCCC, 0xCC0000}
	f.charmap["o"] = mgl32.Vec4{0x000000, 0x0078CC, 0xCCCCCC, 0x780000}
	f.charmap["p"] = mgl32.Vec4{0x000000, 0x00DC66, 0x666666, 0x7C60F0}
	f.charmap["q"] = mgl32.Vec4{0x000000, 0x0076CC, 0xCCCCCC, 0x7C0C1E}
	f.charmap["r"] = mgl32.Vec4{0x000000, 0x00EC6E, 0x766060, 0xF00000}
	f.charmap["s"] = mgl32.Vec4{0x000000, 0x0078CC, 0x6018CC, 0x780000}
	f.charmap["t"] = mgl32.Vec4{0x000020, 0x60FC60, 0x60606C, 0x380000}
	f.charmap["u"] = mgl32.Vec4{0x000000, 0x00CCCC, 0xCCCCCC, 0x760000}
	f.charmap["v"] = mgl32.Vec4{0x000000, 0x00CCCC, 0xCCCC78, 0x300000}
	f.charmap["w"] = mgl32.Vec4{0x000000, 0x00C6C6, 0xD6D66C, 0x6C0000}
	f.charmap["x"] = mgl32.Vec4{0x000000, 0x00C66C, 0x38386C, 0xC60000}
	f.charmap["y"] = mgl32.Vec4{0x000000, 0x006666, 0x66663C, 0x0C18F0}
	f.charmap["z"] = mgl32.Vec4{0x000000, 0x00FC8C, 0x1860C4, 0xFC0000}
	f.charmap["!"] = mgl32.Vec4{0x003078, 0x787830, 0x300030, 0x300000}

	f.charmap["\""] = mgl32.Vec4{0x006666, 0x662400, 0x000000, 0x000000}
	f.charmap["#"] = mgl32.Vec4{0x006C6C, 0xFE6C6C, 0x6CFE6C, 0x6C0000}
	f.charmap["$"] = mgl32.Vec4{0x30307C, 0xC0C078, 0x0C0CF8, 0x303000}
	f.charmap["%"] = mgl32.Vec4{0x000000, 0xC4CC18, 0x3060CC, 0x8C0000}
	f.charmap["&"] = mgl32.Vec4{0x0070D8, 0xD870FA, 0xDECCDC, 0x760000}
	f.charmap["'"] = mgl32.Vec4{0x003030, 0x306000, 0x000000, 0x000000}
	f.charmap["("] = mgl32.Vec4{0x000C18, 0x306060, 0x603018, 0x0C0000}
	f.charmap[")"] = mgl32.Vec4{0x006030, 0x180C0C, 0x0C1830, 0x600000}
	f.charmap["*"] = mgl32.Vec4{0x000000, 0x663CFF, 0x3C6600, 0x000000}
	f.charmap["+"] = mgl32.Vec4{0x000000, 0x18187E, 0x181800, 0x000000}
	f.charmap[","] = mgl32.Vec4{0x000000, 0x000000, 0x000038, 0x386000}
	f.charmap["-"] = mgl32.Vec4{0x000000, 0x0000FE, 0x000000, 0x000000}
	f.charmap["."] = mgl32.Vec4{0x000000, 0x000000, 0x000038, 0x380000}
	f.charmap["/"] = mgl32.Vec4{0x000002, 0x060C18, 0x3060C0, 0x800000}
	f.charmap[":"] = mgl32.Vec4{0x000000, 0x383800, 0x003838, 0x000000}
	f.charmap[";"] = mgl32.Vec4{0x000000, 0x383800, 0x003838, 0x183000}
	f.charmap["<"] = mgl32.Vec4{0x000C18, 0x3060C0, 0x603018, 0x0C0000}
	f.charmap["="] = mgl32.Vec4{0x000000, 0x007E00, 0x7E0000, 0x000000}
	f.charmap[">"] = mgl32.Vec4{0x006030, 0x180C06, 0x0C1830, 0x600000}
	f.charmap["?"] = mgl32.Vec4{0x0078CC, 0x0C1830, 0x300030, 0x300000}
	f.charmap["@"] = mgl32.Vec4{0x007CC6, 0xC6DEDE, 0xDEC0C0, 0x7C0000}
	f.charmap["["] = mgl32.Vec4{0x003C30, 0x303030, 0x303030, 0x3C0000}
	f.charmap["\\"] = mgl32.Vec4{0x000080, 0xC06030, 0x180C06, 0x020000}
	f.charmap["]"] = mgl32.Vec4{0x003C0C, 0x0C0C0C, 0x0C0C0C, 0x3C0000}
	f.charmap["^"] = mgl32.Vec4{0x10386C, 0xC60000, 0x000000, 0x000000}
	f.charmap["_"] = mgl32.Vec4{0x000000, 0x000000, 0x000000, 0x00FF00}
	f.charmap["{"] = mgl32.Vec4{0x001C30, 0x3060C0, 0x603030, 0x1C0000}
	f.charmap["|"] = mgl32.Vec4{0x001818, 0x181800, 0x181818, 0x180000}
	f.charmap["}"] = mgl32.Vec4{0x00E030, 0x30180C, 0x183030, 0xE00000}
	f.charmap["~"] = mgl32.Vec4{0x0073DA, 0xCE0000, 0x000000, 0x000000}
	f.charmap["ch_lar"] = mgl32.Vec4{0x000000, 0x10386C, 0xC6C6FE, 0x000000}
}

// f.charmap[" "] = mgl32.Vec2{0.0, 0.0}
// f.charmap["spc"] = mgl32.Vec2{0.0, 0.0}
// f.charmap["!"] = mgl32.Vec2{276705.0, 32776.0}
// f.charmap["\""] = mgl32.Vec2{1797408.0, 0.0}
// f.charmap["#"] = mgl32.Vec2{10738.0, 1134484.0}
// f.charmap["dol"] = mgl32.Vec2{538883.0, 19976.0}
// f.charmap["%"] = mgl32.Vec2{1664033.0, 68006.0}
// f.charmap["&"] = mgl32.Vec2{545090.0, 174362.0}
// f.charmap["'"] = mgl32.Vec2{798848.0, 0.0}
// f.charmap["("] = mgl32.Vec2{270466.0, 66568.0}
// f.charmap[")"] = mgl32.Vec2{528449.0, 33296.0}
// f.charmap["*"] = mgl32.Vec2{10471.0, 1688832.0}
// f.charmap["crs"] = mgl32.Vec2{4167.0, 1606144.0}
// f.charmap["per"] = mgl32.Vec2{0.0, 1560.0}
// f.charmap["-"] = mgl32.Vec2{7.0, 1572864.0}
// f.charmap[","] = mgl32.Vec2{0.0, 1544.0}
// f.charmap["lsl"] = mgl32.Vec2{1057.0, 67584.0}
// f.charmap["0"] = mgl32.Vec2{935221.0, 731292.0}
// f.charmap["1"] = mgl32.Vec2{274497.0, 33308.0}
// f.charmap["2"] = mgl32.Vec2{934929.0, 1116222.0}
// f.charmap["3"] = mgl32.Vec2{934931.0, 1058972.0}
// f.charmap["4"] = mgl32.Vec2{137380.0, 1302788.0}
// f.charmap["5"] = mgl32.Vec2{2048263.0, 1058972.0}
// f.charmap["6"] = mgl32.Vec2{401671.0, 1190044.0}
// f.charmap["7"] = mgl32.Vec2{2032673.0, 66576.0}
// f.charmap["8"] = mgl32.Vec2{935187.0, 1190044.0}
// f.charmap["9"] = mgl32.Vec2{935187.0, 1581336.0}
// f.charmap[":"] = mgl32.Vec2{195.0, 1560.0}
// f.charmap[";"] = mgl32.Vec2{195.0, 1544.0}
// f.charmap["<"] = mgl32.Vec2{135300.0, 66052.0}
// f.charmap["="] = mgl32.Vec2{496.0, 3968.0}
// f.charmap[">"] = mgl32.Vec2{528416.0, 541200.0}
// f.charmap["?"] = mgl32.Vec2{934929.0, 1081352.0}
// f.charmap["ats"] = mgl32.Vec2{935285.0, 714780.0}
// f.charmap["A"] = mgl32.Vec2{935188.0, 780450.0}
// f.charmap["B"] = mgl32.Vec2{1983767.0, 1190076.0}
// f.charmap["C"] = mgl32.Vec2{935172.0, 133276.0}
// f.charmap["D"] = mgl32.Vec2{1983764.0, 665788.0}
// f.charmap["E"] = mgl32.Vec2{2048263.0, 1181758.0}
// f.charmap["F"] = mgl32.Vec2{2048263.0, 1181728.0}
// f.charmap["G"] = mgl32.Vec2{935173.0, 1714334.0}
// f.charmap["H"] = mgl32.Vec2{1131799.0, 1714338.0}
// f.charmap["I"] = mgl32.Vec2{921665.0, 33308.0}
// f.charmap["J"] = mgl32.Vec2{66576.0, 665756.0}
// f.charmap["K"] = mgl32.Vec2{1132870.0, 166178.0}
// f.charmap["L"] = mgl32.Vec2{1065220.0, 133182.0}
// f.charmap["M"] = mgl32.Vec2{1142100.0, 665762.0}
// f.charmap["N"] = mgl32.Vec2{1140052.0, 1714338.0}
// f.charmap["O"] = mgl32.Vec2{935188.0, 665756.0}
// f.charmap["P"] = mgl32.Vec2{1983767.0, 1181728.0}
// f.charmap["Q"] = mgl32.Vec2{935188.0, 698650.0}
// f.charmap["R"] = mgl32.Vec2{1983767.0, 1198242.0}
// f.charmap["S"] = mgl32.Vec2{935171.0, 1058972.0}
// f.charmap["T"] = mgl32.Vec2{2035777.0, 33288.0}
// f.charmap["U"] = mgl32.Vec2{1131796.0, 665756.0}
// f.charmap["V"] = mgl32.Vec2{1131796.0, 664840.0}
// f.charmap["W"] = mgl32.Vec2{1131861.0, 699028.0}
// f.charmap["X"] = mgl32.Vec2{1131681.0, 84130.0}
// f.charmap["Y"] = mgl32.Vec2{1131794.0, 1081864.0}
// f.charmap["Z"] = mgl32.Vec2{1968194.0, 133180.0}
// f.charmap["lsb"] = mgl32.Vec2{925826.0, 66588.0}
// f.charmap["rsl"] = mgl32.Vec2{16513.0, 16512.0}
// f.charmap["rsb"] = mgl32.Vec2{919584.0, 1065244.0}
// f.charmap["pow"] = mgl32.Vec2{272656.0, 0.0}
// f.charmap["_"] = mgl32.Vec2{0.0, 62.0}
// f.charmap["a"] = mgl32.Vec2{224.0, 649374.0}
// f.charmap["b"] = mgl32.Vec2{1065444.0, 665788.0}
// f.charmap["c"] = mgl32.Vec2{228.0, 657564.0}
// f.charmap["d"] = mgl32.Vec2{66804.0, 665758.0}
// f.charmap["e"] = mgl32.Vec2{228.0, 772124.0}
// f.charmap["f"] = mgl32.Vec2{401543.0, 1115152.0}
// f.charmap["g"] = mgl32.Vec2{244.0, 665474.0}
// f.charmap["h"] = mgl32.Vec2{1065444.0, 665762.0}
// f.charmap["i"] = mgl32.Vec2{262209.0, 33292.0}
// f.charmap["j"] = mgl32.Vec2{131168.0, 1066252.0}
// f.charmap["k"] = mgl32.Vec2{1065253.0, 199204.0}
// f.charmap["l"] = mgl32.Vec2{266305.0, 33292.0}
// f.charmap["m"] = mgl32.Vec2{421.0, 698530.0}
// f.charmap["n"] = mgl32.Vec2{452.0, 1198372.0}
// f.charmap["o"] = mgl32.Vec2{228.0, 665756.0}
// f.charmap["p"] = mgl32.Vec2{484.0, 667424.0}
// f.charmap["q"] = mgl32.Vec2{244.0, 665474.0}
// f.charmap["r"] = mgl32.Vec2{354.0, 590904.0}
// f.charmap["s"] = mgl32.Vec2{228.0, 114844.0}
// f.charmap["t"] = mgl32.Vec2{8674.0, 66824.0}
// f.charmap["u"] = mgl32.Vec2{292.0, 1198868.0}
// f.charmap["v"] = mgl32.Vec2{276.0, 664840.0}
// f.charmap["w"] = mgl32.Vec2{276.0, 700308.0}
// f.charmap["x"] = mgl32.Vec2{292.0, 1149220.0}
// f.charmap["y"] = mgl32.Vec2{292.0, 1163824.0}
// f.charmap["z"] = mgl32.Vec2{480.0, 1148988.0}
// f.charmap["lpa"] = mgl32.Vec2{401542.0, 66572.0}
// f.charmap["|"] = mgl32.Vec2{266304.0, 33288.0}
// f.charmap["rpa"] = mgl32.Vec2{788512.0, 1589528.0}
// f.charmap["~"] = mgl32.Vec2{675840.0, 0.0}
// f.charmap["lar"] = mgl32.Vec2{8387.0, 1147904.0}

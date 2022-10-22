package harvester

import m "github.com/drabadan/gostealthclient/pkg/model"

const DEBUG = false
const UNLOAD_BAG_BANK = 0x4260ACE3
const SEARCH_TREES_RANGE = 16
const CAPTCHA_GUMP_ID = 0x21C7F58

const FORGE = 0x4047E107
const FORGE_MINOC = 0x4000CBE3

var SMELT_POINT = m.Point2D{X: 3719, Y: 2465}
var SMELT_POINT_MINOC = m.Point2D{X: 2562, Y: 502}
var UNLOAD_POINT = m.Point2D{X: 3704, Y: 2511}
var UNLOAD_POINT_MINOC_BANK = m.Point2D{X: 2505, Y: 542}

var MINABLE_ROCKS_OCCLO = []m.Point2D{
	{
		X: 3672,
		Y: 2455,
	},
	{
		X: 3667,
		Y: 2454,
	},
	{
		X: 3667,
		Y: 2455,
	},
	{
		X: 3667,
		Y: 2456,
	},
	{
		X: 3668,
		Y: 2456,
	},
	{
		X: 3722,
		Y: 2465,
	},
	{
		X: 3722,
		Y: 2466,
	},
	{
		X: 3723,
		Y: 2472,
	},
	{
		X: 3723,
		Y: 2473,
	},
	{
		X: 3672,
		Y: 2456,
	},
	{
		X: 3673,
		Y: 2456,
	},
}

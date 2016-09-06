package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"clound_video_api_project/video_logic"
)

type GPSFrameController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @router /*.* [get]
func (this *GPSFrameController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	path := this.Ctx.Input.Param(":path")
	ext := this.Ctx.Input.Param(":ext")

	t1 := this.GetString("t1")
	t2 := this.GetString("t2")
	t3 := this.GetString("t3")

	t4 := this.GetString("t4")
	t5 := this.GetString("t5")
	t6 := this.GetString("t6")
	t7 := this.GetString("t7")
	t8 := this.GetString("t8")
	t9 := this.GetString("t9")
	t10 := this.GetString("t10")

	if len(t1) == 0 {
		this.Data["json"] = map[string]interface{}{"state": -1001, "msg": "时间1不能为空"}
		this.ServeJSON()
		return
	}

	t1_st, err := strconv.ParseFloat(t1, 32)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"state": -1002, "msg": "时间1格式错误"}
		this.ServeJSON()
		return
	}
	var t2_st float64 = -1
	if len(t2) != 0 {
		_tmp_t2_st, err := strconv.ParseFloat(t2, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1003, "msg": "时间2格式错误"}
			this.ServeJSON()
			return
		}
		t2_st = _tmp_t2_st

		//时间必须由小到大
		if t2_st <= t1_st {
			this.Data["json"] = map[string]interface{}{"state": -1005, "msg": "时间2必须大于时间1"}
			this.ServeJSON()
			return
		}
	}

	//3
	var t3_st float64 = -1
	if len(t3) != 0 {
		_tmp_t3_st, err := strconv.ParseFloat(t3, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1004, "msg": "时间3格式错误"}
			this.ServeJSON()
			return
		}
		t3_st = _tmp_t3_st

		if t3_st <= t2_st {
			this.Data["json"] = map[string]interface{}{"state": -1006, "msg": "时间3必须大于时间2"}
			this.ServeJSON()
			return
		}
	}

	//4
	var t4_st float64 = -1
	if len(t4) != 0 {
		_tmp_t4_st, err := strconv.ParseFloat(t4, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1007, "msg": "时间4格式错误"}
			this.ServeJSON()
			return
		}
		t4_st = _tmp_t4_st

		if t4_st <= t3_st {
			this.Data["json"] = map[string]interface{}{"state": -1008, "msg": "时间4必须大于时间3"}
			this.ServeJSON()
			return
		}
	}
	//5
	var t5_st float64 = -1
	if len(t5) != 0 {
		_tmp_t5_st, err := strconv.ParseFloat(t5, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1009, "msg": "时间5格式错误"}
			this.ServeJSON()
			return
		}
		t5_st = _tmp_t5_st

		if t5_st <= t4_st {
			this.Data["json"] = map[string]interface{}{"state": -1010, "msg": "时间5必须大于时间4"}
			this.ServeJSON()
			return
		}
	}
	//6
	var t6_st float64 = -1
	if len(t6) != 0 {
		_tmp_t6_st, err := strconv.ParseFloat(t6, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1011, "msg": "时间6格式错误"}
			this.ServeJSON()
			return
		}
		t6_st = _tmp_t6_st

		if t6_st <= t5_st {
			this.Data["json"] = map[string]interface{}{"state": -1012, "msg": "时间6必须大于时间5"}
			this.ServeJSON()
			return
		}
	}
	//7
	var t7_st float64 = -1
	if len(t7) != 0 {
		_tmp_t7_st, err := strconv.ParseFloat(t7, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1013, "msg": "时间7格式错误"}
			this.ServeJSON()
			return
		}
		t7_st = _tmp_t7_st

		if t7_st <= t6_st {
			this.Data["json"] = map[string]interface{}{"state": -1014, "msg": "时间7必须大于时间6"}
			this.ServeJSON()
			return
		}
	}
	//8
	var t8_st float64 = -1
	if len(t8) != 0 {
		_tmp_t8_st, err := strconv.ParseFloat(t8, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1015, "msg": "时间8格式错误"}
			this.ServeJSON()
			return
		}
		t8_st = _tmp_t8_st

		if t8_st <= t7_st {
			this.Data["json"] = map[string]interface{}{"state": -1016, "msg": "时间8必须大于时7"}
			this.ServeJSON()
			return
		}
	}
	//9
	var t9_st float64 = -1
	if len(t9) != 0 {
		_tmp_t9_st, err := strconv.ParseFloat(t9, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1017, "msg": "时间9格式错误"}
			this.ServeJSON()
			return
		}
		t9_st = _tmp_t9_st

		if t9_st <= t8_st {
			this.Data["json"] = map[string]interface{}{"state": -1018, "msg": "时间9必须大于时间8"}
			this.ServeJSON()
			return
		}
	}
	//10
	var t10_st float64 = -1
	if len(t10) != 0 {
		_tmp_t10_st, err := strconv.ParseFloat(t10, 32)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"state": -1019, "msg": "时间10格式错误"}
			this.ServeJSON()
			return
		}
		t10_st = _tmp_t10_st

		if t10_st <= t9_st {
			this.Data["json"] = map[string]interface{}{"state": -1020, "msg": "时间10必须大于时间9"}
			this.ServeJSON()
			return
		}
	}

	result, result_json := video_logic.Video_thum_gps(path+"."+ext, float32(t1_st), float32(t2_st), float32(t3_st), float32(t4_st), float32(t5_st), float32(t6_st), float32(t7_st), float32(t8_st), float32(t9_st), float32(t10_st))
	if result == 0 {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Ctx.WriteString(result_json)
	} else {
		this.Data["json"] = map[string]interface{}{"state": result}
		this.ServeJSON()
	}
}

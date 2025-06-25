package model

type Response struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type ClassInfos struct {
	Code int        `json:"code"`
	Info string     `json:"info"`
	Data []MetaData `json:"data"`
}

type MetaData struct {
	Xnxq    string `json:"xnxq"`
	Jxb     string `json:"jxb"`
	Kcbh    string `json:"kcbh"`
	Kcmc    string `json:"kcmc"`
	Xf      string `json:"xf"`
	TeaName string `json:"teaName"`
	RsLimit int    `json:"rsLimit"`
	RwType  int    `json:"rwType"`
	Kclb    string `json:"kclb"`
	KchType string `json:"kchType"`
	Memo    string `json:"memo"`

	// 以下是小学期独有属性
	KcInfo string `json:"kcInfo"` // 课程信息，小学期独有
	Tea    string `json:"tea"`    // 这部分是小学期的老师信息
	SkInfo string `json:"skInfo"` // 小学期特有属性，上课信息
}

type SmallRequest struct {
	Action string `json:"action"`
	Jxb    string `json:"jxb"`
	Kclb   string `json:"kclb"`
	Kcbh   string `json:"kcbh"`
}

// 以下是正式选课的结构体部分：

// CourseData 表示选课系统中一门课程的完整信息
type SecCourseData struct {
	IsOver  int      `json:"isOver"`  // 是否已结束（0 = 否）
	IsSksy  int      `json:"isSksy"`  // 是否为实验课或试用课（0 = 否）
	IsTyfx  int      `json:"isTyfx"`  // 是否为体育方向（0 = 否）
	Jxb     string   `json:"jxb"`     // 教学班编号
	Kcbh    string   `json:"kcbh"`    // 课程编号
	KchType string   `json:"kchType"` // 课程号类型（如“全校选修”）
	Kclb    string   `json:"kclb"`    // 课程类别（如“选修”、“必修”）
	Kcmc    string   `json:"kcmc"`    // 课程名称
	Nj      string   `json:"nj"`      // 年级或面向对象（如“全校”）
	Sd      []string `json:"sd"`      // 上课节次代码（如 ["134"]）
	SdPrint string   `json:"sdPrint"` // 上课时间（人类可读格式，如 “星期1第3-4节1-16周”）
	Syxs    int      `json:"syxs"`    // 剩余可选人数
	Teacher string   `json:"teacher"` // 教师姓名
	Xf      string   `json:"xf"`      // 学分
	Xkzt    int      `json:"xkzt"`    // 选课状态（1 = 可选）
	Zc      []string `json:"zc"`      // 周次信息编码（如 ["111111111111111100000"]）
	Zym     string   `json:"zym"`     // 专业信息（为空时为通选课）
}

type SecResponse struct {
	Code int                      `json:"code"`
	Info string                   `json:"info"`
	Data map[string]SecCourseData `json:"data"`
}

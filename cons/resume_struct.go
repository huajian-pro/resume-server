package cons

// ResumeData 简历内容结构体
type ResumeData struct {
	ID    string `json:"ID"`    // 模版ID
	NAME  string `json:"NAME"`  // 模版名称
	TITLE string `json:"TITLE"` // 模版标题
	LIST  []struct {
		Id       string `json:"id"`
		Iconfont string `json:"iconfont"`
		Model    string `json:"model"`
		Show     bool   `json:"show"`
		Title    string `json:"title"`
		Style    struct {
			TextColor       string `json:"textColor"`
			TextFontSize    string `json:"textFontSize"`
			TextFontWeight  int    `json:"textFontWeight"`
			BackgroundColor string `json:"backgroundColor,omitempty"`
			MBottom         string `json:"mBottom"`
			MTop            string `json:"mTop"`
			TitleColor      string `json:"titleColor,omitempty"`
			TitleFontSize   string `json:"titleFontSize,omitempty"`
			TitleFontWeight int    `json:"titleFontWeight,omitempty"`
		} `json:"style"`
		Name        string `json:"name,omitempty"`
		Age         int    `json:"age,omitempty"`
		Address     string `json:"address,omitempty"`
		Avatar      string `json:"avatar,omitempty"`
		WorkService int    `json:"workService,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
		Email       string `json:"email,omitempty"`
		Abstract    string `json:"abstract,omitempty"`
		Degree      string `json:"degree,omitempty"`
		IsShow      struct {
			Age               bool `json:"age,omitempty"`
			Address           bool `json:"address,omitempty"`
			Avatar            bool `json:"avatar,omitempty"`
			WorkService       bool `json:"workService,omitempty"`
			PhoneNumber       bool `json:"phoneNumber,omitempty"`
			Email             bool `json:"email,omitempty"`
			Abstract          bool `json:"abstract,omitempty"`
			Degree            bool `json:"degree,omitempty"`
			IntendedPositions bool `json:"intendedPositions,omitempty"`
			IntendedCity      bool `json:"intendedCity,omitempty"`
			ExpectSalary      bool `json:"expectSalary,omitempty"`
			JobStatus         bool `json:"jobStatus,omitempty"`
			JobSearchType     bool `json:"jobSearchType,omitempty"`
			Date              bool `json:"date,omitempty"`
			SchoolName        bool `json:"schoolName,omitempty"`
			Specialized       bool `json:"specialized,omitempty"`
			MajorCourse       bool `json:"majorCourse,omitempty"`
			CampusBriefly     bool `json:"campusBriefly,omitempty"`
			CampusDuty        bool `json:"campusDuty,omitempty"`
			CampusContent     bool `json:"campusContent,omitempty"`
			CompanyName       bool `json:"companyName,omitempty"`
			Posts             bool `json:"posts,omitempty"`
			JobContent        bool `json:"jobContent,omitempty"`
			ProjectName       bool `json:"projectName,omitempty"`
			AwardsName        bool `json:"awardsName,omitempty"`
			AwardsGrade       bool `json:"awardsGrade,omitempty"`
		} `json:"isShow,omitempty"`
		IntendedPositions string `json:"intendedPositions,omitempty"`
		IntendedCity      string `json:"intendedCity,omitempty"`
		ExpectSalary      string `json:"expectSalary,omitempty"`
		JobStatus         string `json:"jobStatus,omitempty"`
		JobSearchType     string `json:"jobSearchType,omitempty"`
		LIST              []struct {
			Date          interface{} `json:"date,omitempty"`
			SchoolName    string      `json:"schoolName,omitempty"`
			Specialized   string      `json:"specialized,omitempty"`
			Degree        string      `json:"degree,omitempty"`
			MajorCourse   string      `json:"majorCourse,omitempty"`
			SkillName     string      `json:"skillName,omitempty"`
			Proficiency   string      `json:"proficiency,omitempty"`
			Introduce     string      `json:"introduce,omitempty"`
			CampusBriefly string      `json:"campusBriefly,omitempty"`
			CampusDuty    string      `json:"campusDuty,omitempty"`
			CampusContent string      `json:"campusContent,omitempty"`
			CompanyName   string      `json:"companyName,omitempty"`
			Posts         string      `json:"posts,omitempty"`
			JobContent    []struct {
				Content string `json:"content"`
			} `json:"jobContent,omitempty"`
			ProjectName    string `json:"projectName,omitempty"`
			ProjectContent []struct {
				Content string `json:"content"`
			} `json:"projectContent,omitempty"`
			AwardsName     string `json:"awardsName,omitempty"`
			AwardsGrade    string `json:"awardsGrade,omitempty"`
			WorksName      string `json:"worksName,omitempty"`
			WorksLink      string `json:"worksLink,omitempty"`
			WorksIntroduce string `json:"worksIntroduce,omitempty"`
		} `json:"LIST,omitempty"`
		Content string `json:"content,omitempty"`
	} `json:"LIST"` // 数据列表
	GlobalStyle struct {
		ThemeColor          string `json:"themeColor"`          // 主题颜色
		FirstTitleFontSize  string `json:"firstTitleFontSize"`  // 导航栏标题字体大小
		SecondTitleFontSize string `json:"secondTitleFontSize"` // 字体大小
		TextFontSize        string `json:"textFontSize"`        // 字体大小
		SecondTitleColor    string `json:"secondTitleColor"`    // 字体颜色
		TextFontColor       string `json:"textFontColor"`       // 文字颜色
		SecondTitleWeight   int    `json:"secondTitleWeight"`   // 字体粗细
		TextFontWeight      int    `json:"textFontWeight"`      // 文字字体粗细
		ModelMarginTop      string `json:"modelMarginTop"`      // 	模块距离顶部距离
		ModelMarginBottom   string `json:"modelMarginBottom"`   // 模块边距
	} `json:"GLOBAL_STYLE"` // 全局样式
}

package cons

// ResumeData 简历内容结构体
type ResumeData struct {
	ID    string `json:"ID"`
	NAME  string `json:"NAME"`
	TITLE string `json:"TITLE"`
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
	} `json:"LIST"`
	GlobalStyle struct {
		ThemeColor          string `json:"themeColor"`
		FirstTitleFontSize  string `json:"firstTitleFontSize"`
		SecondTitleFontSize string `json:"secondTitleFontSize"`
		TextFontSize        string `json:"textFontSize"`
		SecondTitleColor    string `json:"secondTitleColor"`
		TextFontColor       string `json:"textFontColor"`
		SecondTitleWeight   int    `json:"secondTitleWeight"`
		TextFontWeight      int    `json:"textFontWeight"`
		ModelMarginTop      string `json:"modelMarginTop"`
		ModelMarginBottom   string `json:"modelMarginBottom"`
	} `json:"GLOBAL_STYLE"`
}

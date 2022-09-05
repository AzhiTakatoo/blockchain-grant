package lib

//系统用户
type WyuUser struct {
	WyuUserId   string `json:"wyuUserId"`   //用户ID
	WyuUserName string `json:"wyuUserName"` //用户名
	WyuPasswd   string `json:"wyuPasswd"`   //用户密码
	Power       string `json:"power"`
}

type RegisterCertify struct {
	RegisterId   string `json:"wyuUserId"`   //用户ID
	RegisterName string `json:"wyuUserName"` //用户名
	RegisterTime string `json:"registerTime"`
}

type ProofCertify struct {
	StipendId    string `json:"stipendId"` //申请助学金学生ID
	RegisterTime string `json:"registerTime"`
}

//助学金评定
type ProofMaterial struct {
	StipendId               string  `json:"stipendId"`               //申请助学金学生ID
	AnnualHouseholdIncome   float64 `json:"annualHouseholdIncome"`   //家庭人均年收入
	ComprehensiveEvaluation float64 `json:"comprehensiveEvaluation"` //综合测评
	VolunteerTime           int64   `json:"volunteerTime"`           //义工时长
	StipendScore            float64 `json:"stipendScore"`            //助学金评定得分
	//PhotoMaterial           string  `json:"photoMaterial"`           //照片认证材料
	//Ranking int64 `json:"ranking"` //排名
}

type SortStipend struct {
	StipendId    string  `json:"stipendId"`    //申请助学金学生ID
	StipendScore float64 `json:"stipendScore"` //助学金评定得分
	Ranking      int64   `json:"ranking"`      //排名
	Grade        string  `json:"grade"`        //等级
	Money        string  `json:"money"`
}

type PhotoMaterial struct {
	WyuUserId string `json:"wyuUserId"` //ID
	Photo     string `json:"photo"`     //照片
}

type Vote struct {
	StipendId   string  `json:"stipendId"`   //被评分者
	VoteId      string  `json:"voteId"`      //评分者
	Vote        float64 `json:"vote"`        //评分
	AverageVote float64 `json:"averageVote"` //平均分
	StuNum      float64 `json:"stuNum"`      //评分人数
}

// Account 账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"grantUserId"` //账号ID
	UserName  string  `json:"userName"`    //账号名
	Balance   float64 `json:"balance"`     //余额
}

const (
	WyuUserKey         = "wyuuser-key"
	RegisterCertifyKey = "register-certify-key"
	ProofMaterialKey   = "proof-material-key"
	PhotoMaterialKey   = "photo-material-key"
	ProofCertifyKey    = "proof-certify-key"
	VoteKey            = "vote-key"
	StipendRankingKey  = "stipend-ranking-key"
)

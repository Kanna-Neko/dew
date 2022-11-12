package link

type UsersStatusApi struct {
	Status string   `json:"status"`
	Result []Status `json:"result"`
}
type ContestStandingInterface struct {
	Status  string                `json:"status"`
	Result  ContestStandingResult `json:"result"`
	Comment string                `json:"comment"`
}
type ContestStandingResult struct {
	Problems []Problem `json:"problems"`
}
type Status struct {
	Problem Problem `json:"problem"`
	// Id                  int     `json:"id"`
	// ContestId           int     `json:"contestId"`
	// CreationTimeSeconds int     `json:"creationTimeSeconds"`
	// RelativeTimeSeconds int     `json:"relativeTimeSeconds"`
	// Author              any     `json:"author"`
	// ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict string `json:"verdict"`
	// Testset             string  `json:"testset"`
	// PassedTestCount     int     `json:"passedTestCount"`
	// TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	// MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}
type Problem struct {
	ContestId int    `json:"contestId"`
	Index     string `json:"index"`
	// Name      string   `json:"name"`
	// Type      string   `json:"type"`
	// Points    float64  `json:"points"`
	// Rating    int      `json:"rating"`
	// Tags      []string `json:"tags"`
}

type ProblemJson struct {
	Id    string `json:"id"`
	Index string `json:"index"`
}

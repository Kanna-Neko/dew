package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Kanna-Neko/dew/link"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var div3Diffculty = []ProblemCondition{{Difficult: []string{"800"}}, {Difficult: []string{"800", "900"}}, {Difficult: []string{"900", "1000", "1100"}}, {Difficult: []string{"1100", "1200", "1300", "1400"}}, {Difficult: []string{"1400", "1500", "1600", "1700"}}, {Difficult: []string{"1700", "1800", "1900"}}, {Difficult: []string{"1900", "2000", "2100"}}}
var div2Diffculty = []ProblemCondition{{Difficult: []string{"800", "900", "1000"}, Bad: []string{"interactive"}}, {Difficult: []string{"1000", "1100", "1200"}, Bad: []string{"interactive"}}, {Difficult: []string{"1200", "1300", "1400", "1500", "1600"}}, {Difficult: []string{"1600", "1700", "1800", "1900"}}, {Difficult: []string{"2000", "2100", "2200", "2300", "2400"}}, {Difficult: []string{"2500", "2600", "2700", "2800"}}}
var div1Diffculty = []ProblemCondition{{Difficult: []string{"1500", "1600", "1700"}}, {Difficult: []string{"1800", "1900", "2000", "2100", "2200", "2300"}}, {Difficult: []string{"2400", "2500", "2600", "2700", "2800"}}, {Difficult: []string{"2900", "3000", "3100", "3200", "3300"}}, {Difficult: []string{"3400", "3500"}}}

func init() {
	rootCmd.AddCommand(NewCmd)
	NewCmd.AddCommand(div1)
	NewCmd.AddCommand(div2)
	NewCmd.AddCommand(div3)
	NewCmd.AddCommand(randomOne)
	NewCmd.AddCommand(customCmd)
}

const (
	title    = "miaonei"
	duration = "120"
)

var NewCmd = &cobra.Command{
	Use:   "generate",
	Short: "create a contest",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var div3 = &cobra.Command{
	Use:   "div3",
	Short: "create a contest, whose difficulty like div3",
	Run: func(cmd *cobra.Command, args []string) {
		newContest(ContestInfo{
			Duration:          "120",
			ContestTitle:      "miaonei",
			ProblemConditions: div3Diffculty,
			BanProblems:       nil,
		})
	},
}
var div2 = &cobra.Command{
	Use:   "div2",
	Short: "create a contest, whose difficulty like div2",
	Run: func(cmd *cobra.Command, args []string) {
		newContest(ContestInfo{
			Duration:          "120",
			ContestTitle:      "miaonei",
			ProblemConditions: div2Diffculty,
			BanProblems:       nil,
		})
	},
}
var div1 = &cobra.Command{
	Use:   "div1",
	Short: "create a contest, whose difficulty like div1",
	Run: func(cmd *cobra.Command, args []string) {
		newContest(ContestInfo{
			Duration:          "120",
			ContestTitle:      "miaonei",
			ProblemConditions: div1Diffculty,
			BanProblems:       nil,
		})
	},
}

func newContest(contest ContestInfo) {
	if contest.ContestTitle == "" {
		contest.ContestTitle = title
	}
	if contest.Duration == "" {
		contest.Duration = "120"
	}
	if len(contest.ProblemConditions) == 0 {
		contest.ProblemConditions = div2Diffculty
	}
	link.Login()
	pro := PickSomeProblem(contest.ProblemConditions, contest.BanProblems)
	link.CreateContest(contest.ContestTitle, contest.Duration, pro)
	OpenWebsite(codeforcesDomain + "/mashups")
}

var randomOne = &cobra.Command{
	Use:   "random",
	Short: "random select one problem",
	Run: func(cmd *cobra.Command, args []string) {
		Random()
	},
}

func Random() {
	isExist := checkConfigFile()
	if !isExist {
		log.Fatal("config file is not exist, please use cf init command")
	}
	ReadConfig()
	if !viper.IsSet("rating") {
		log.Fatal("we notice the info of rating is not exist, please use init config command first, or modify rating in ./codeforces/config.yaml (you can add a line and write 'rating: 1234').")
	}
	var rating = viper.GetInt("rating")
	if rating < 800 {
		rating = 800
	}
	rating = (rating / 100) * 100
	lowRating := rating + 200
	if lowRating > 3500 {
		lowRating = 3500
	}
	highRating := rating + 300
	if highRating > 3500 {
		highRating = 3500
	}
	var pro []string
	for i := lowRating; i <= highRating; i += 100 {
		pro = append(pro, strconv.Itoa(i))
	}
	var thisOne = PickOneProblem(ProblemCondition{
		Difficult: pro,
	}, nil)
	viper.Set("problem", strconv.Itoa(thisOne.ContestId)+thisOne.Index)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
	OpenRandomFunc(thisOne)
	GetTestcases(strconv.Itoa(thisOne.ContestId) + thisOne.Index)
}

func PickSomeProblem(in []ProblemCondition, banProblems []string) []string {
	cj := uispinner.New()
	cj.Start()
	login := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("picking problems").SetComplete("pick problem complete")
	var pro []string
	var mp map[string]bool = make(map[string]bool)
	for i := 0; i < len(in); i++ {
		var one = PickOneProblem(in[i], banProblems)
		var goal = strconv.Itoa(one.ContestId) + one.Index
		if mp[goal] {
			i--
			continue
		}
		pro = append(pro, goal)
		mp[goal] = true
	}
	login.Done()
	cj.Stop()
	return pro
}

func PickOneProblem(r ProblemCondition, banProblems []string) problemInfo {
	data := PickProblems(r.Difficult)
	data = Deduplication(data, link.GetStatus())
	data = filterGood(data, r.Good)
	data = filterBad(data, r.Bad)
	if banProblems != nil {
		data = filterBanProblems(data, banProblems)
	}
	if len(data) == 0 {
		log.Fatal("you are so good, you have solve all problems of the range", r)
	}
	rand.Seed(time.Now().Unix())
	var pos = rand.Int() % len(data)
	return data[pos]
}

func PickProblems(in []string) []problemInfo {
	var res []problemInfo
	for i := 0; i < len(in); i++ {
		data, err := ioutil.ReadFile("./codeforces/" + in[i] + ".json")
		if err != nil {
			log.Fatal(err.Error() + "\nyou should use update command before generate")
		}
		var tmp []problemInfo
		json.Unmarshal(data, &tmp)
		res = append(res, tmp...)
	}
	return res
}

func Deduplication(data []problemInfo, s map[string]bool) []problemInfo {
	var res []problemInfo
	for i := 0; i < len(data); i++ {
		if _, exist := s[strconv.Itoa(data[i].ContestId)+data[i].Index]; !exist {
			res = append(res, data[i])
		}
	}
	return res
}

func filterGood(data []problemInfo, good []string) []problemInfo {
	if len(good) == 0 {
		return data
	}
	var mp = make(map[string]bool)
	for _, v := range good {
		mp[v] = true
	}
	var result []problemInfo
	for _, v := range data {
		for _, vv := range v.Tags {
			if mp[vv] {
				result = append(result, v)
				break
			}
		}
	}
	return result
}
func filterBad(data []problemInfo, bad []string) []problemInfo {
	if len(bad) == 0 {
		return data
	}
	var mp = make(map[string]bool)
	for _, v := range bad {
		mp[v] = true
	}
	var result []problemInfo
	for _, v := range data {
		for _, vv := range v.Tags {
			if mp[vv] {
				break
			}
		}
		result = append(result, v)
	}
	return result
}

func filterBanProblems(data []problemInfo, banProblems []string) []problemInfo {
	var mp = make(map[string]bool)
	for _, v := range banProblems {
		mp[v] = true
	}
	var result []problemInfo
	for _, v := range data {
		if !mp[strconv.Itoa(v.ContestId)+v.Index] {
			result = append(result, v)
		}
	}
	return result
}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "custom a virtual contest",
	Run: func(cmd *cobra.Command, args []string) {
		custom()
	},
}

type ContestInfo struct {
	Duration          string             `json:"duration"`
	ContestTitle      string             `json:"contestTitle"`
	Name              string             `json:"name"`
	ProblemConditions []ProblemCondition `json:"problemConditions"`
	BanProblems       []string           `json:"banProblems"`
}
type ProblemCondition struct {
	Difficult []string `json:"difficult"`
	Good      []string `json:"good"`
	Bad       []string `json:"bad"`
}
type ContestInfos struct {
	Templates []ContestInfo `json:"templates"`
}

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item ContestInfo

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   ContestInfo
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

var temp ContestInfos

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = ContestInfo(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice.Name != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to you.", m.choice.Name))
	}
	if m.quitting {
		return quitTextStyle.Render("No want? That's cool.")
	}
	return "\n" + m.list.View()
}

func custom() {
	data, err := ioutil.ReadFile("./codeforces/contestTemplate.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &temp)
	if err != nil {
		log.Fatal(err)
	}
	items := []list.Item{}
	for _, v := range temp.Templates {
		items = append(items, item(v))
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "choose one contest template"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}
	mod, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if m, ok := mod.(model); ok && m.choice.Name != "" {
		newContest(ContestInfo{
			Duration:          m.choice.Duration,
			ContestTitle:      m.choice.ContestTitle,
			BanProblems:       m.choice.BanProblems,
			ProblemConditions: m.choice.ProblemConditions,
		})
	}
}

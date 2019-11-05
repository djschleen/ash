// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cve, err := UnmarshalCve(bytes)
//    bytes, err = cve.Marshal()

package calc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UnmarshalCve(data []byte) (Cve, error) {
	var r Cve
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Cve) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (c Cve) ToShortCve() ShortCve {
	var s = ShortCve{Cvss: c.Cvss, ID: c.ID, CvssTime: c.CvssTime}
	return s
}

func Info(cveID string) Cve {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://cve.circl.lu/api/cve/%s", cveID), nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	cve, err := UnmarshalCve(body)
	if err != nil {
		log.Fatal(err)
	}

	return cve
}

func (c ShortCve) SeverityWeight() float64 {
	if (c.Cvss >= 0.0) && (c.Cvss <= 3.9) {
		return 0.1
	} else if (c.Cvss >= 4.0) && (c.Cvss <= 6.9) {
		return 0.2
	} else if (c.Cvss >= 7.0) && (c.Cvss <= 9.9) {
		return 0.3
	}
	return 0.4
}

type ShortCve struct {
	Cvss     float64 `json:"cvss"`
	ID       string  `json:"id"`
	CvssTime string  `json:"cvss-time"`
}

type Cve struct {
	Modified                      string                    `json:"Modified"`
	Published                     string                    `json:"Published"`
	Access                        Access                    `json:"access"`
	Capec                         []Capec                   `json:"capec"`
	Cvss                          float64                  `json:"cvss"`
	CvssTime                      string                    `json:"cvss-time"`
	Cwe                           string                    `json:"cwe"`
	ExploitDB                     []ExploitDB               `json:"exploit-db"`
	ID                            string                    `json:"id"`
	Impact                        Impact                    `json:"impact"`
	LastModified                  string                    `json:"last-modified"`
	Metasploit                    []ExploitDB               `json:"metasploit"`
	Msbulletin                    []Msbulletin              `json:"msbulletin"`
	Nessus                        []Nessus                  `json:"nessus"`
	Oval                          []Oval                    `json:"oval"`
	Packetstorm                   []ExploitDB               `json:"packetstorm"`
	Ranking                       [][]Ranking               `json:"ranking"`
	References                    []string                  `json:"references"`
	Refmap                        Refmap                    `json:"refmap"`
	Saint                         []Saint                   `json:"saint"`
	Summary                       string                    `json:"summary"`
	TheHackerNews                 []ExploitDB               `json:"the hacker news"`
	VulnerableConfiguration       []VulnerableConfiguration `json:"vulnerable_configuration"`
	VulnerableConfigurationCpe2_2 []string                  `json:"vulnerable_configuration_cpe_2_2"`
}

type Access struct {
	Authentication string `json:"authentication"`
	Complexity     string `json:"complexity"`
	Vector         string `json:"vector"`
}

type Capec struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Prerequisites   string   `json:"prerequisites"`
	RelatedWeakness []string `json:"related_weakness"`
	Solutions       string   `json:"solutions"`
	Summary         string   `json:"summary"`
}

type ExploitDB struct {
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	LastSeen    string  `json:"last seen"`
	Modified    *string `json:"modified,omitempty"`
	Published   string  `json:"published"`
	Reporter    string  `json:"reporter"`
	Source      string  `json:"source"`
	Title       string  `json:"title"`
	Reliability *string `json:"reliability,omitempty"`
	DataSource  *string `json:"data source,omitempty"`
}

type Impact struct {
	Availability    string `json:"availability"`
	Confidentiality string `json:"confidentiality"`
	Integrity       string `json:"integrity"`
}

type Msbulletin struct {
	BulletinID       string      `json:"bulletin_id"`
	BulletinURL      interface{} `json:"bulletin_url"`
	Date             string      `json:"date"`
	Impact           string      `json:"impact"`
	KnowledgebaseID  string      `json:"knowledgebase_id"`
	KnowledgebaseURL interface{} `json:"knowledgebase_url"`
	Severity         string      `json:"severity"`
	Title            string      `json:"title"`
}

type Nessus struct {
	NASLFamily  string `json:"NASL family"`
	NASLID      string `json:"NASL id"`
	Description string `json:"description"`
	LastSeen    string `json:"last seen"`
	Modified    string `json:"modified"`
	PluginID    string `json:"plugin id"`
	Published   string `json:"published"`
	Reporter    string `json:"reporter"`
	Source      string `json:"source"`
	Title       string `json:"title"`
}

type Oval struct {
	Accepted             string                `json:"accepted"`
	Class                string                `json:"class"`
	Contributors         []Contributor         `json:"contributors"`
	DefinitionExtensions []DefinitionExtension `json:"definition_extensions"`
	Description          string                `json:"description"`
	Family               string                `json:"family"`
	ID                   string                `json:"id"`
	Status               string                `json:"status"`
	Submitted            string                `json:"submitted"`
	Title                string                `json:"title"`
	Version              string                `json:"version"`
}

type Contributor struct {
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type DefinitionExtension struct {
	Comment string `json:"comment"`
	Oval    string `json:"oval"`
}

type Ranking struct {
	Circl int64 `json:"circl"`
}

type Refmap struct {
	Bid      []string `json:"bid"`
	CERT     []string `json:"cert"`
	Idefense []string `json:"idefense"`
	MS       []string `json:"ms"`
	Sectrack []string `json:"sectrack"`
	Secunia  []string `json:"secunia"`
	Sreason  []string `json:"sreason"`
	Vupen    []string `json:"vupen"`
}

type Saint struct {
	Bid         string `json:"bid"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Osvdb       string `json:"osvdb"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type VulnerableConfiguration struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

package gpm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/net/context"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	ret := &Client{
		baseURL:    "https://api.github.com",
		httpClient: &http.Client{},
		token:      token,
	}
	return ret
}
func (c *Client) basicHeader() map[string]string {
	ret := make(map[string]string)
	ret["Accept"] = "application/vnd.github.inertia-preview+json"
	ret["Authorization"] = "token " + c.token
	return ret
}

func (c *Client) ListRepoProjects(ctx context.Context, repoName string) ([]Project, error) {
	endpoint := "/repos/" + repoName + "/projects"
	ret := []Project{}
	resp, err := c.get(endpoint, c.basicHeader())
	if err != nil {
		return ret, errors.Wrap(err, "Something wrong occurred on http request")
	}
	err = decodeBody(resp, &ret)
	if err != nil {
		return ret, errors.Wrap(err, "Something wrong occurred on decoding response")
	}
	return ret, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) get(endpoint string, headers map[string]string) (*http.Response, error) {
	url := c.baseURL + endpoint
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return c.httpClient.Do(req)
}

type Project struct {
	OwnerURL string `json:"owner_url"`
	URL      string `json:"url"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Body     string `json:"body"`
	Number   int    `json:"number"`
	Creator  struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

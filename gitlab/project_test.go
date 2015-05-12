package gitlab
import (
	"testing"
)

type mockconfig struct {
	host string
}

func (this *mockconfig) Host() (string, error) {
	return this.host, nil
}

func (this *mockconfig) ApiPath() (string, error) {
	return "api/v3", nil
}

func (this *mockconfig) Token() (string, error) {
	return "token", nil
}

func (this *mockconfig) Scheme() (string, error) {
	return "https", nil
}

func TestProject_NewProjectFromURL(t *testing.T) {
	remote := "git@gitlab.com:numa08/cookbook-gitlab.git"
	config := &mockconfig{host: "gitlab.com",}
	project, e := NewProject(config, remote)

	if e != nil {
		t.Errorf("create project error %s", e)
		return
	}
	if project  == nil {
		t.Error("create nil project")
		return
	}

	if project.Host != "gitlab.com" {
		t.Error("actual project host is " + project.Host + " but expected gitlab.com")
	}
	if project.Name != "cookbook-gitlab" {
		t.Error("actual project name is " + project.Name + " but expected cookbook-gitlab")
	}
	if project.NameSpace != "numa08" {
		t.Error("actual project name space is " + project.NameSpace + " but expected numa08")
	}
	if project.Protocol != "https" {
		t.Error("actual project protocol is " + project.Protocol + " but expected https")
	}
}
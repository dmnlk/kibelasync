package kibela

import (
	"encoding/json"

	"github.com/Songmu/kibelasync/client"
	"golang.org/x/xerrors"
)

type Group struct {
	ID   `json:"id"`
	Name string `json:"name"`
}

func (ki *Kibela) getGroupCount() (int, error) {
	data, err := ki.cli.Do(&client.Payload{Query: totalGroupCountQuery})
	if err != nil {
		return 0, xerrors.Errorf("failed to ki.getGroupCount: %w", err)
	}
	var res struct {
		Groups struct {
			TotalCount int `json:"totalCount"`
		} `json:"groups"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return 0, xerrors.Errorf("failed to ki.getNotesCount: %w", err)
	}
	return res.Groups.TotalCount, nil
}

func (ki *Kibela) getGroups() ([]*Group, error) {
	num, err := ki.getGroupCount()
	if err != nil {
		return nil, xerrors.Errorf("failed to getGroups: %w", err)
	}
	data, err := ki.cli.Do(&client.Payload{Query: listGroupQuery(num)})
	if err != nil {
		return nil, xerrors.Errorf("failed to ki.getGroups: %w", err)
	}
	var res struct {
		Groups struct {
			Nodes []*Group `json:"nodes"`
		} `json:"groups"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, xerrors.Errorf("failed to ki.getNotesCount: %w", err)
	}
	return res.Groups.Nodes, nil
}
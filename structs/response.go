package structs

import (
  "errors"
)

type Response struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Push        *Push       `json:"push"`
  Fork        *Repository `json:"fork"`
  Issue       *Issue      `json:"issue"`
  Comment     *Comment    `json:"comment"`
  PullRequest *PullRequest `json:"pull_request"`
  Commit      *Hstore     `json:"commit"`
  Approval    *Approval   `json:"approval"`
  Changes     *JSON       `json:"changes"`
}

func (r *Response) Type() (string, error) {
  if r.Actor == nil || r.Repository == nil {
    return "", errors.New("Empty response")
  }

  if r.Push != nil {
    return "Push", nil
  }

  if r.Fork != nil {
    return "Fork", nil
  }

  if r.Commit != nil && r.Comment != nil {
    return "Commit Comment", nil
  }

  if r.Issue != nil && r.Comment == nil && r.Changes == nil {
    return "Issue Created", nil
  }

  if r.Issue != nil && r.Comment != nil && r.Changes != nil {
    return "Issue Updated", nil
  }

  if r.Issue != nil && r.Comment != nil && r.Changes == nil {
    return "Issue Comment Created", nil
  }

  if r.PullRequest != nil && r.Approval == nil && r.Comment == nil {
    switch r.PullRequest.Title {
    case "OPEN":
      return "Pull Request Open", nil
    case "MERGED":
      return "Pull Request Merged", nil
    case "DECLINED":
      return "Pull Request Declined", nil
    }
  }

  if r.PullRequest != nil && r.Approval != nil {
    return "Pull Request Approval", nil
  }

  if r.PullRequest != nil && r.Comment != nil {
    return "Pull Request Comment", nil
  }

  return "", errors.New("Unrecognizable response")
}

package structs

import (
  "time"
)

type Links map[string]map[string]string

type JSON map[string]interface{}

type User struct {
  Username    string      `json:"username"`
  DisplayName string      `json:"display_name"`
  Uuid        string      `json:"uuid"`
  Links       Links       `json:"links"`
}

type Repository struct {
  Name        string      `json:"name"`
  FullName    string      `json:"full_name"`
  Uuid        string      `json:"uuid"`
  Links       Links       `json:"links"`
}

type Parent struct {
  Type        string      `json:"type"`
  Hash        string      `json:"hash"`
  Links       Links       `json:"links"`
}

type Target struct {
  Type        string      `json:"type"`
  Hash        string      `json:"hash"`
  Author      User        `json:"author"`
  Message     string      `json:"message"`
  Date        time.Time   `json:"date"`
  Parents     []Parent    `json:"parents"`
  Links       Links       `json:"links"`
}

type Obj struct {
  Type        string      `json:"string"`
  Name        string      `json:"name"`
  Target      Target      `json:"target"`
  Links       Links       `json:"links"`
}

type Change struct {
  New         Obj         `json:"new"`
  Old         Obj         `json:"old"`
  Links       Links       `json:"links"`
  Created     bool        `json:"created"`
  Forced      bool        `json:"forced"`
  Closed      bool        `json:"closed"`
}

type Push struct {
  Changes     []Change    `json:"changes"`
}

type BaseResponse struct {
  User        User        `json:"user"`
  Repository  Repository  `json:"repository"`
}

type Hstore map[string]string

type Issue struct {
  Id          int         `json:"id"`
  Component   string      `json:"component"`
  Title       string      `json:"title"`
  Content     Hstore      `json:"content"`
  Priority    string      `json:"priority"`
  State       string      `json:"state"`
  Type        string      `json:"type"`
  Milestone   Hstore      `json:"milestone"`
  Version     Hstore      `json:"version"`
  CreatedOn   time.Time   `json:"created_on"`
  UpdatedOn   time.Time   `json:"updated_on"`
  Links       Links       `json:"links"`
}

type Inline struct {
  Path        string      `json:"path"`
  From        int         `json:"from"`
  To          int         `json:"to"`
}

type Comment struct {
  Id          int         `json:"id"`
  Parent   map[string]int `json:"parent"`
  Content     Hstore      `json:"content"`
  Inline      Inline      `json:"inline"`
  CreatedOn   time.Time   `json:"created_on"`
  UpdatedOn   time.Time   `json:"updated_on"`
  Links       Links       `json:"links"`
}

type Source struct {
  Branch      Hstore      `json:"branch"`
  Commit      Hstore      `json:"commit"`
  Repository  Repository  `json:"repository"`
}

type PullRequest struct {
  Id          int         `json:"id"`
  Title       string      `json:"title"`
  Description string      `json:"description"`
  State       string      `json:"state"`
  Author      User        `json:"author"`
  Source      Source      `json:"source"`
  Destination Source      `json:"destination"`
  MergeCommit Hstore      `json:"merge_commit"`
  Participants []User     `json:"participants"`
  Reviewers   []User      `json:"reviewers"`
  CloseSourceBranch bool  `json:"close_source_branch"`
  ClosedBy    User        `json:"closed_by"`
  Reason      string      `json:"reason"`
  CreatedOn   time.Time   `json:"created_on"`
  UpdatedOn   time.Time   `json:"updated_on"`
  Links       Links       `json:"links"`
}

type PushResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Push        *Push       `json:"push"`
}

type ForkResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Fork        *Repository `json:"fork"`
}

type RepositoryCommentResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Comment     *Comment    `json:"comment"`
  Commit      *Hstore     `json:"commit"`
}

type IssueResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Issue       *Issue      `json:"issue"`
}

type IssueCommentResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Issue       *Issue      `json:"issue"`
  Commit      *Hstore     `json:"commit"`
}

type PullRequestResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  PullRequest *PullRequest `json:"pull_request"`
}

type PullRequestCommentResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  PullRequest *PullRequest `json:"pull_request"`
  Comment     *Comment    `json:"comment"`
}

type Approval struct {
  Date        time.Time   `json:"date"`
  User        User        `json:"user"`
}

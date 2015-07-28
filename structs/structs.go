package structs

import (
  "time"
)

type Links map[string]map[string]string

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

type PushResponse struct {
  Actor       *User       `json:"actor"`
  Repository  *Repository `json:"repository"`
  Push        *Push       `json:"push"`
}

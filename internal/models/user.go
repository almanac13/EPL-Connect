package models

type Role string

const (
  RoleFan      Role = "fan"
  RoleVIP      Role = "vip"
  RoleAdmin    Role = "admin"
  RoleClubAdmin Role = "club_admin"
)

type User struct {
  ID    string `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
  Role  Role   `json:"role"`
}
import "time"

type Club struct {
  ID          string    `json:"id"`
  OwnerUserID string    `json:"ownerUserId"`
  Name        string    `json:"name"`
  Description string    `json:"description"`
  CreatedAt   time.Time `json:"createdAt"`
}

package workflows

import (
	"fmt"
	"time"
	"os"
	"text/tabwriter"
  "encoding/json"

	client "github.com/semaphoreci/cli/api/client"
	"github.com/semaphoreci/cli/api/models"
	"github.com/semaphoreci/cli/cmd/utils"
)

func Describe(projectID, wfID string) {
  c := client.NewPipelinesV1AlphaApi()
  body, err := c.ListPplByWfID(projectID, wfID)
  utils.Check(err)

  prettyPrintPipelineList(body)
}

func prettyPrintPipelineList(jsonList []byte) {
  j := models.PipelinesListV1Alpha{}
	err := json.Unmarshal(jsonList, &j)
  utils.Check(err)

  const padding = 3
  w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

  fmt.Fprintln(w, "PIPELINE ID\tPIPELINE NAME\tCREATION TIME\tSTATE")

  for _, p := range j {
    createdAt := time.Unix(p.CreatedAt.Seconds, 0).Format("2006-01-02 15:04:05")
    fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", p.Id, p.Name, createdAt, p.State)
  }

  w.Flush()
}

package cmd

import (
	"GitActivity/internal"
	"GitActivity/model"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:     "github-activity",
	Short:   "show all activity of a user",
	Example: "github-activity" + " <username>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.Get("https://api.github.com/users/" + args[0] + "/events")

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("User not found")
			return
		} else if resp.StatusCode != http.StatusOK {
			s := fmt.Sprintf("error fetching data: %d\"", resp.StatusCode)
			fmt.Println(s)
			return
		} else if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var activity []model.Activity
		err = json.Unmarshal(body, &activity)

		if len(activity) == 0 {
			fmt.Println("The user has no activity")
			return
		}

		for _, v := range activity {

			var action string

			switch v.Type {
			case internal.PushActivity:
				commitCount := len(v.Payload.Commits)
				action = fmt.Sprintf("Pushed %d commit(s) to %s %s", commitCount, v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			case internal.PullActivity:
				action = fmt.Sprintf("Pulling from %s %s", v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			case internal.IssuesActivity:
				action = fmt.Sprintf("%s an issue in %s %s", v.Payload.Action, v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			case internal.WatchActivity:
				action = fmt.Sprintf("Starred %s %s", v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			case internal.ForkActivity:
				action = fmt.Sprintf("Forked %s %s", v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			case internal.CreateActivity:
				action = fmt.Sprintf("Created %s in %s %s", v.Payload.RefType, v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			default:
				action = fmt.Sprintf("%s in %s %s", v.Type, v.Repo.Name, v.CreatedAt.Format(time.DateTime))
			}
			fmt.Println(action)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

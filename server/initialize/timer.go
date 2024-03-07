package initialize

import (
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/task"

	"github.com/robfig/cron/v3"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// Clean DB scheduled task
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // Scheduled task method located in the task package
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "Scheduled cleanup of database content [logs, blacklist]", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// Other scheduled tasks go here. Refer to the usage above.

		//_, err := global.GVA_Timer.AddTaskByFunc("ScheduledTaskIdentifier", "cronExpression", func() {
		//	Specific execution content...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}

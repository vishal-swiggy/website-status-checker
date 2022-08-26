package Checker

import (
	"context"
	"log"
	"net/http"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type WebsiteChecker struct {
}

func (checker WebsiteChecker) Check(ctx context.Context, name string) (status bool, err error) {
	routineNo := ctx.Value("routineNo")
	if response, err := http.Get("https://" + name); err == nil {
		if response.StatusCode == 200 {
			log.Printf("Website: `%s` UP  Status code: `%d` RoutineNo: `%d`\n", name, response.StatusCode, routineNo)
			return true, nil
		} else {
			log.Printf("Website: `%s` DOWN  Status code: `%d` RoutineNo: `%d`\n", name, response.StatusCode, routineNo)
			return false, nil
		}
	} else {
		log.Printf("Error: `%s` RoutineNo: `%d`\n", err.Error(), routineNo)
		return false, err
	}
}

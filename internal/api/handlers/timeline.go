package handlers

import (
	"fmt"
	"time"

	"github.com/4lxprime/Goldmp/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type TimelineEvent struct {
	EventType   string    `json:"eventType"`
	ActiveUntil time.Time `json:"activeUntil"`
	ActiveSince time.Time `json:"activeSince"`
}

// must be in utils but i cannot
func timelineNewEvent(event string) TimelineEvent {
	return TimelineEvent{
		EventType:   fmt.Sprintf(`EventFlag.%s`, event),
		ActiveUntil: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		ActiveSince: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

// ["/fortnite/api/calendar/v1/timeline"]
// work actually on season 3-5
func HandleTimeline(c *fiber.Ctx) error {
	version, err := utils.GetVersion(&c.Request().Header)
	if err != nil {
		return err
	}

	var activeEvents []TimelineEvent

	activeEvents = append(activeEvents, timelineNewEvent(fmt.Sprintf(`Season%d`, version.Season)))
	activeEvents = append(activeEvents, timelineNewEvent(version.Lobby))

	if version.Season == 3 {
		activeEvents = append(activeEvents, timelineNewEvent("Spring2018Phase1"))

		if version.Build >= 3.1 {
			activeEvents = append(activeEvents, timelineNewEvent("Spring2018Phase2"))
		}

		if version.Build >= 3.3 {
			activeEvents = append(activeEvents, timelineNewEvent("Spring2018Phase3"))
		}

		if version.Build >= 3.4 {
			activeEvents = append(activeEvents, timelineNewEvent("Spring2018Phase4"))
		}

	} else if version.Season == 4 {
		activeEvents = append(activeEvents, timelineNewEvent("Blockbuster2018"))
		activeEvents = append(activeEvents, timelineNewEvent("Blockbuster2018Phase1"))

		if version.Build >= 4.3 {
			activeEvents = append(activeEvents, timelineNewEvent("Blockbuster2018Phase2"))
		}

		if version.Build >= 4.4 {
			activeEvents = append(activeEvents, timelineNewEvent("Blockbuster2018Phase3"))
		}

		if version.Build >= 4.5 {
			activeEvents = append(activeEvents, timelineNewEvent("Blockbuster2018Phase4"))
		}

	} else if version.Season == 5 {
		activeEvents = append(activeEvents, timelineNewEvent("RoadTrip2018"))
		activeEvents = append(activeEvents, timelineNewEvent("Horde"))
		activeEvents = append(activeEvents, timelineNewEvent("Anniversary2018_BR"))
		activeEvents = append(activeEvents, timelineNewEvent("LTM_Heist"))

		if version.Build >= 5.1 {
			activeEvents = append(activeEvents, timelineNewEvent("BirthdayBattleBus"))
		}
	}

	return c.JSON(activeEvents)
}

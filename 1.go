package main

import (
    "fmt"
    "time"
)

type Event struct {
    EventName   string
    Date  time.Time
    Description string
}

type Userr struct {
    Name   string
    Events []Event
}


func (u *Userr) AddEvent(event Event) {

    u.Events = append(u.Events, event)
}

func (u *Userr) RemoveEvent(eventName string) {
    filteredEvents := []Event{}
    for _, event := range u.Events {
        if event.EventName != eventName {
            filteredEvents = append(filteredEvents, event)
        }
    }
    u.Events = filteredEvents
}
func (u *Userr) UpdateEvent(eventName string, newEvent Event) {
    for i, event := range u.Events {
        if event.EventName == eventName {
            u.Events[i] = newEvent
            break
        }
    }
}

func (u *Userr) GetEventsByDate(Date time.Time) []Event {
    result := []Event{}
    for _, event := range u.Events {
        if event.Date.Format("2006-01-02") == Date.Format("2006-01-02") {
            result = append(result, event)
        }
    }
    return result
}

func main() {
   
    me := Userr{Name: "Xadicha"} 
	
    today := time.Now()
    tomorrow := today.AddDate(0, 0, 1)

	var chosen int
	for {
		fmt.Printf("1.AddEvent\n2.RemoveEvent\n3.UpdateEvent\n4.My events\n5.Exit\nChoose: ")
		fmt.Scan(&chosen)
		switch chosen{
		case 1:
			me.AddEvent(Event{"Meeting", today, "Team meeting"})
			me.AddEvent(Event{"Lunch", today.Add(time.Hour * 2), "Lunch"})
			me.AddEvent(Event{"Class", tomorrow, "Najot ta'lim"})
		case 2:  
			me.RemoveEvent("Lunch")
		case 3:
			newEvent := Event{"Updated Meeting", today.Add(time.Hour * 1), "Meeting"}
			me.UpdateEvent("Meeting", newEvent)
		case 4:
			fmt.Println("All Events:")
			for _, event := range me.Events {
				fmt.Println(event)
			}
		case 5:
			break

		}
	
	}
}

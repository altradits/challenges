package main

import (
	"fmt"
	"sort"
	"time"
)

func cmdDue(root string, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: drill due")
	}
	cat, err := loadCatalog(root)
	if err != nil {
		return err
	}
	prog, err := loadProgress(root)
	if err != nil {
		return err
	}
	at := now()
	type item struct {
		c   Challenge
		due time.Time
		new bool
	}
	var overdue, upcoming []item
	var fresh []Challenge
	for _, c := range cat {
		r := prog.record(c.Num)
		due, ok := dueAt(r)
		if !ok {
			continue // mastered
		}
		if r.Passes == 0 {
			fresh = append(fresh, c)
			continue
		}
		if !due.After(at) {
			overdue = append(overdue, item{c, due, false})
		} else {
			upcoming = append(upcoming, item{c, due, false})
		}
	}
	sort.Slice(overdue, func(i, j int) bool { return overdue[i].due.Before(overdue[j].due) })
	sort.Slice(upcoming, func(i, j int) bool { return upcoming[i].due.Before(upcoming[j].due) })

	if len(overdue) == 0 && len(fresh) == 0 {
		fmt.Println("Nothing due. Rest — or drill ahead with `drill start NN`.")
		return nil
	}
	for _, it := range overdue {
		r := prog.record(it.c.Num)
		fmt.Printf("OVERDUE  %-22s stage %d/%d  target %s  (was due %s)\n",
			it.c.Dir, r.Stage, len(intervals),
			formatDur(targetTime(it.c.Limit, r.Stage)), it.due.Format("Mon 2006-01-02"))
	}
	// Suggest the next unseen challenge in number order.
	if len(fresh) > 0 {
		c := fresh[0]
		fmt.Printf("NEXT     %-22s first attempt  target %s\n", c.Dir, formatDur(targetTime(c.Limit, 0)))
		if len(fresh) > 1 {
			fmt.Printf("         ... and %d more unseen challenges after it.\n", len(fresh)-1)
		}
	}
	for _, it := range upcoming {
		r := prog.record(it.c.Num)
		fmt.Printf("waiting  %-22s stage %d/%d  due %s\n",
			it.c.Dir, r.Stage, len(intervals), it.due.Format("Mon 2006-01-02"))
	}
	return nil
}

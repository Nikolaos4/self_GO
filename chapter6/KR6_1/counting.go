package main

func check(people map[string]int) int {
    keys := make([]string, 0, len(people))
    for k := range people {
        keys = append(keys, k)
    }
    count := 0
    for i := 0; i < len(keys); i++ {
        for j := i + 1; j < len(keys); j++ {
            if keys[i][0] == keys[j][0] {
                diff := people[keys[i]] - people[keys[j]]
                if diff < 0 {
                    diff = -diff
                }
                if diff <= 5 {
                    count++
                }
            }
        }
    }
    return count
}
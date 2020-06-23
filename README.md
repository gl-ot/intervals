# intervals

### install
```
go get -u github.com/gl-ot/intervals
```

### API
```
intervals.FirstDay
intervals.LastDay
intervals.InSameInterval
```

### examples
```
const yourLayout = "2006-01-02"

day1 := time.Parse(yourLayout, "2020-12-12") // sat
day2 := time.Parse(yourLayout, "2020-12-13") // sun
intervals.InSameInterval(day1, day2, intervals.Week) // true

day3 := time.Parse(yourLayout, "2020-12-14") // mon
intervals.InSameInterval(day1, day3, intervals.Week) // false
```

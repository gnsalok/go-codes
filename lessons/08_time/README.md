# Lesson 08: Time, Instants, and Locations

## Learning goals

- Distinguish an instant from its display in a time zone.
- Parse and format with Go's reference-time layouts.
- Explain Unix timestamps, monotonic clock readings, and duration arithmetic.
- Make time-dependent behavior deterministic in tests.

## Mental model

A `time.Time` represents an instant plus location and optional monotonic-clock data.
Unix seconds are one serialization choice, not automatically the best database
representation. Preserve the precision and zone semantics required by the domain;
convert for presentation at an explicit boundary.

```bash
go run ./lessons/08_time/examples/epoch
go run ./lessons/08_time/examples/time_zones
go test ./lessons/08_time/examples/...
```

## Production tradeoffs and common mistakes

- Never parse a zone abbreviation and assume it uniquely identifies a location.
- Calendar arithmetic is not equivalent to adding a fixed duration across daylight
  saving transitions.
- `time.Since` can use monotonic data carried by an in-process `time.Time`; serialized
  values do not retain that component.
- Inject a clock or pass the current time into domain logic rather than asserting
  against the wall clock in tests.

## Interview prompts

1. What information is lost when a timestamp is stored as Unix seconds?
2. When should you use `AddDate` instead of adding `24*time.Hour`?
3. How would you test expiration logic at an exact boundary?

## Practice

Parse an RFC 3339 timestamp, display it in two named locations, and test a deadline
function using a fixed instant rather than `time.Now` inside the function.

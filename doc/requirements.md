# Requirements

## v1.0.0

1. Be provided as a Go-package.
2. Organize computation-stages into pipes.
3. Concurrently process many items at each stage.
4. Give each item to the next stage as soon as the current one ends.
5. Process items through the stages in the order the stages were given.
6. Hide concurrency.
7. Make all types should be usable.
8. Define a helper that inserts all the items into a pipe and returns the
   processed items once they're all delivered.

## v1.1.0

1. Allow consumers to receives items.
2. Define a helper that inserts all the items into a pipe and consumes the
   processed items as they're delivered.

## v1.2.0

1. Allow producers to create items and signal they are done creating.
2. Define a helper that inserts all the items from a producer into a pipe and
   returns the processed items once they're all delivered.
3. Define a helper that inserts all the items from a producer into a pipe and
   consumes the processed items as they're delivered.

# Requirements

## v1.0.0

1. Stages of computations should be organized into pipes.
2. Each stage should be able to concurrently process many items.
3. Each item should be give to the next stage as soon as the current one ends.
4. Elements should proceed through the stages in order.
5. Concurrency should be hidden.
6. All types should be usable.

## v1.1.0

1. A wrapper should be provided to process items and pass them to a consumer as
   they're done.

## v1.2.0

1. A wrapper should be provided to produce items and pass them into all pipe
   interfaces.

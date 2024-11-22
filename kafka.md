## Kafka

Brokers
Producer
Consumer
Consumer Groups
Topic
Partitions 
Replications


Message delivery semantics
-------------------------
At most once : Message is delivered to the consumer mostly once. There is a chance of loss of data.

At least once : Message is delivered to the consumer at least once.In case multiple times you need to retrieve the message, it is possible

Only Once: Message is delivered only once.But for sure it is delivered, cannot retrieve it again

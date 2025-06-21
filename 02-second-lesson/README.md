# Second Lesson - Error Handling with Go

Functions in Go can return two values:
- The value/response;
- And the error;

Thus, every function should return its expected type and an error.
Errors can panic the system, breaking its flow, log the error and etc.  

However, blank identifiers "_" can replace the error, but it won't describe what happened with precision.
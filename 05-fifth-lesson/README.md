# Fifth Lesson - Structs and Compositions

Compositions is a way to embed one struct inside another, allowing for code reuse and method sharing, similar to inheritance in OOP.

* You can convert structs to JSON using the `json.Marshal` function, which serializes the struct into JSON format.
* Similarly, you can parse JSON back into a struct using json.Unmarshal, making it easy to work with external data sources.

Finally, you can add tags to control how field names appear in the JSON output, you can add struct tags, such as this:

```go
    type Employee struct {
    Company string `json:"company_name"`
}
```
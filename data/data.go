package data

type Question struct{ Text string }
type Entity  struct{ Name string }

var Questions = []Question{
    {Text: "Is it a real person?"},
    {Text: "Is it male?"},
    {Text: "Is he famous for science?"},
}

// For the moment we only have one possible entity
var Secret = Entity{Name: "Albert Einstein"}
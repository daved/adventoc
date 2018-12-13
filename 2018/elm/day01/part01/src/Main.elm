module Main exposing (result)

import Input exposing (data)


lines : List String
lines =
    String.split "\n" data


result : Int
result =
    List.filterMap String.toInt lines |> List.sum

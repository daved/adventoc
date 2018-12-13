module Main exposing (result)

import Input exposing (data)


result : Int
result =
    String.split "\n" data
        |> List.filterMap String.toInt
        |> List.sum

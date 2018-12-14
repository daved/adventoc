module Main exposing (result)

import Input exposing (data)


result : Int
result =
    data
        |> List.filterMap String.toInt
        |> List.sum

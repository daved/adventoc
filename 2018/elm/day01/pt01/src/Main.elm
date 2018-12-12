module Main exposing (result)

import Input exposing (data)


lines =
    String.split "\n" data


result =
    List.filterMap String.toInt lines |> List.sum

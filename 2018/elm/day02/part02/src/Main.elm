module Main exposing (result)

import Input exposing (data)
import List.Extra as List


type alias Comparison =
    { sames : List Char
    , diffs : List ( Char, Char )
    }


comparison : String -> String -> Comparison
comparison a b =
    let
        ( sames, diffs ) =
            List.zip (String.toList a) (String.toList b)
                |> List.partition (\( x, y ) -> x == y)
    in
    { sames = List.map (\( x, _ ) -> x) sames
    , diffs = diffs
    }


comparisonByDiffs : Int -> ( String, String ) -> Maybe Comparison
comparisonByDiffs diffs ( a, b ) =
    let
        c =
            comparison a b
    in
    if List.length c.diffs == diffs then
        Just c

    else
        Nothing


pairs : List ( String, String ) -> String -> List String -> List ( String, String )
pairs tups string list =
    list
        |> List.zip (List.repeat (List.length list) string)
        |> List.append tups


listPairs : List ( String, String ) -> List String -> List ( String, String )
listPairs tups list =
    case list of
        x :: xs ->
            listPairs (pairs tups x xs) xs

        _ ->
            tups


result =
    data
        |> listPairs []
        |> List.filterMap (comparisonByDiffs 1)
        |> List.map .sames
        |> List.map String.fromList

module Main exposing (result)

import Dict
import Dict.Extra as Dict
import Input exposing (data)
import List.Extra as List


type alias DubsTripsTally =
    { dubs : Int
    , trps : Int
    }


neatDubsTripsTally : DubsTripsTally
neatDubsTripsTally =
    { dubs = 0
    , trps = 0
    }


oneOrNone : Bool -> Int
oneOrNone bool =
    case bool of
        True ->
            1

        False ->
            0


dubsTripsCombiner : String -> DubsTripsTally -> DubsTripsTally
dubsTripsCombiner string tally =
    let
        uniqCounts =
            string
                |> String.toList
                |> Dict.frequencies
                |> Dict.toList
                |> List.map (\( _, b ) -> b)
                |> List.unique
    in
    { tally
        | dubs = tally.dubs + (oneOrNone <| List.member 2 uniqCounts)
        , trps = tally.trps + (oneOrNone <| List.member 3 uniqCounts)
    }


result =
    let
        tal =
            List.foldl dubsTripsCombiner neatDubsTripsTally data
    in
    tal.dubs * tal.trps

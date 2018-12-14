module Main exposing (result)

import Dict
import Dict.Extra as Dict
import Input exposing (data)


type alias DubsTripsTally =
    { dubs : Int
    , trps : Int
    }


neatDubsTripsTally : DubsTripsTally
neatDubsTripsTally =
    { dubs = 0
    , trps = 0
    }


countOfFrequencies : List ( Char, Int ) -> Int -> Int
countOfFrequencies charFreqs frequency =
    List.foldl
        (\charFreq count ->
            let
                ( char, freq ) =
                    charFreq
            in
            if freq == frequency then
                count + 1

            else
                count
        )
        0
        charFreqs


oneOrNone : Int -> Int
oneOrNone val =
    if val > 0 then
        1

    else
        0


dubsTripsCombiner : String -> DubsTripsTally -> DubsTripsTally
dubsTripsCombiner string tally =
    let
        countOfFrequenciesOf =
            countOfFrequencies
                (string
                    |> String.toList
                    |> Dict.frequencies
                    |> Dict.toList
                )
    in
    { tally
        | dubs = tally.dubs + (oneOrNone <| countOfFrequenciesOf 2)
        , trps = tally.trps + (oneOrNone <| countOfFrequenciesOf 3)
    }


result =
    let
        tal =
            List.foldl dubsTripsCombiner neatDubsTripsTally data
    in
    tal.dubs * tal.trps

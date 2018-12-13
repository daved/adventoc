module Main exposing (result)

import Input exposing (data)


lines : List String
lines =
    String.split "\n" data


ints : List Int
ints =
    List.filterMap String.toInt lines


headInt : List Int -> Int
headInt list =
    case List.head list of
        Just n ->
            n

        Nothing ->
            0


tailInts : List Int -> List Int
tailInts list =
    case List.tail list of
        Just l ->
            l

        Nothing ->
            []


type alias Accumulation =
    { total : Int
    , finds : List Int
    , found : Bool
    , rmndr : List Int
    }


baseAccumulation : Accumulation
baseAccumulation =
    { total = 0
    , finds = []
    , found = False
    , rmndr = ints
    }


accumulateByTail : List Int -> Accumulation -> Accumulation
accumulateByTail tail curr =
    case List.isEmpty tail of
        True ->
            curr

        False ->
            accumulate curr


accumulate : Accumulation -> Accumulation
accumulate curr =
    let
        total =
            curr.total + headInt curr.rmndr

        tail =
            tailInts curr.rmndr
    in
    case List.member total curr.finds of
        True ->
            { curr | total = total, found = True, rmndr = tail }

        False ->
            accumulateByTail tail
                { curr
                    | total = total
                    , finds = total :: curr.finds
                    , rmndr = tail
                }


accumulation : Accumulation -> Accumulation
accumulation init =
    let
        acc =
            accumulate init
    in
    case acc.found of
        True ->
            acc

        False ->
            accumulation { acc | rmndr = ints }


result =
    (accumulation baseAccumulation).total

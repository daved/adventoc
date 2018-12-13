module Main exposing (result)

import Array exposing (Array)
import Input exposing (data)
import Set exposing (Set)


lines : List String
lines =
    String.split "\n" data


ints : List Int
ints =
    List.filterMap String.toInt lines


type alias Accumulation =
    { total : Int
    , finds : Set Int
    , index : Int
    , adnds : Array Int
    }


initial : Accumulation
initial =
    { total = 0
    , finds = Set.empty
    , index = 0
    , adnds = Array.fromList ints
    }


intByIndex : Int -> Array Int -> Int
intByIndex index array =
    case Array.get index array of
        Just n ->
            n

        Nothing ->
            0


nextIndex : Int -> Array a -> Int
nextIndex index array =
    if index + 1 >= Array.length array then
        0

    else
        index + 1


accumulate : Accumulation -> Accumulation
accumulate curr =
    let
        total =
            curr.total + intByIndex curr.index curr.adnds

        next =
            nextIndex curr.index curr.adnds
    in
    case Set.member total curr.finds of
        True ->
            { curr | total = total, index = next }

        False ->
            accumulate
                { curr
                    | total = total
                    , finds = Set.insert total curr.finds
                    , index = next
                }


result : Int
result =
    (accumulate initial).total

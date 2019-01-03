module Main exposing (result)

import Array exposing (Array)
import Input exposing (data)
import Parser exposing ((|.), (|=), spaces, succeed, symbol)


type alias Coords =
    { x : Int
    , y : Int
    }


parseCoords : Parser.Parser Coords
parseCoords =
    succeed Coords
        |. spaces
        |= Parser.int
        |. spaces
        |. symbol ","
        |. spaces
        |= Parser.int
        |. spaces


type alias Size =
    { width : Int
    , height : Int
    }


parseSize : Parser.Parser Size
parseSize =
    succeed Size
        |. spaces
        |= Parser.int
        |. spaces
        |. symbol "x"
        |. spaces
        |= Parser.int
        |. spaces


type alias Claim =
    { id : Int
    , start : Coords
    , size : Size
    }


parseClaimId : Parser.Parser Int
parseClaimId =
    succeed identity
        |. spaces
        |. symbol "#"
        |= Parser.int
        |. spaces


parseClaim : Parser.Parser Claim
parseClaim =
    succeed Claim
        |= parseClaimId
        |. symbol "@"
        |= parseCoords
        |. symbol ":"
        |= parseSize


runMaybeParseClaim : String -> Maybe Claim
runMaybeParseClaim string =
    case Parser.run parseClaim string of
        Ok claim ->
            Just claim

        Err _ ->
            Nothing


claimXAZ : Claim -> ( Int, Int )
claimXAZ claim =
    ( claim.start.x, claim.start.x + claim.size.width - 1 )


claimYAZ : Claim -> ( Int, Int )
claimYAZ claim =
    ( claim.start.y, claim.start.y + claim.size.height - 1 )


type alias Claims =
    List Claim


allClaims : Claims
allClaims =
    data
        |> List.filterMap runMaybeParseClaim


lowestFirst : ( Int, Int ) -> ( Int, Int )
lowestFirst ( first, second ) =
    if first < second then
        ( first, second )

    else
        ( second, first )


intArraySafeGet : Int -> Array Int -> Int
intArraySafeGet index array =
    case Array.get index array of
        Just n ->
            n

        Nothing ->
            0


type alias Row =
    Array Int


neatRow : Row
neatRow =
    Array.initialize 1000 (always 0)


markedRow : ( Int, Int ) -> Row -> Row
markedRow ( xa, xz ) curr =
    let
        ( a, z ) =
            lowestFirst ( xa, xz )
    in
    if a < z then
        curr
            |> Array.set a (intArraySafeGet a curr + 1)
            |> markedRow ( a + 1, z )

    else if a == z then
        curr
            |> Array.set a (intArraySafeGet a curr + 1)

    else
        curr


type alias Fabric =
    Array Row


rowArraySafeGet : Int -> Array Row -> Row
rowArraySafeGet index array =
    case Array.get index array of
        Just row ->
            row

        Nothing ->
            Array.empty


neatFabric : Fabric
neatFabric =
    Array.initialize 1000 (always neatRow)


columnsMarkedFabric : ( Int, Int ) -> ( Int, Int ) -> Fabric -> Fabric
columnsMarkedFabric ( ya, yz ) xaz curr =
    let
        ( a, z ) =
            lowestFirst ( ya, yz )
    in
    if a < z then
        curr
            |> Array.set a (markedRow xaz (rowArraySafeGet a curr))
            |> columnsMarkedFabric
                ( a + 1, z )
                xaz

    else if a == z then
        curr
            |> Array.set a (markedRow xaz (rowArraySafeGet a curr))

    else
        curr


markedFabric : Claims -> Fabric -> Fabric
markedFabric claims curr =
    case claims of
        [] ->
            curr

        claim :: rest ->
            let
                yaz =
                    claimYAZ claim

                xaz =
                    claimXAZ claim

                next =
                    columnsMarkedFabric yaz xaz curr
            in
            markedFabric rest next


areIsolated : ( Int, Int ) -> Row -> Bool
areIsolated ( xa, xz ) row =
    let
        ( a, z ) =
            lowestFirst ( xa, xz )
    in
    if a < z && intArraySafeGet a row == 1 then
        areIsolated ( a + 1, z ) row

    else if a == z then
        intArraySafeGet a row == 1

    else
        False


isIsolated : ( Int, Int ) -> ( Int, Int ) -> Fabric -> Bool
isIsolated ( ya, yz ) xaz fabric =
    let
        ( a, z ) =
            lowestFirst ( ya, yz )
    in
    if a < z && areIsolated xaz (rowArraySafeGet a fabric) then
        isIsolated ( a + 1, z ) xaz fabric

    else if a == z then
        areIsolated xaz (rowArraySafeGet a fabric)

    else
        False


firstIsolatedClaimId : Claims -> Fabric -> Int
firstIsolatedClaimId claims fabric =
    case claims of
        [] ->
            -1

        claim :: rest ->
            let
                yaz =
                    claimYAZ claim

                xaz =
                    claimXAZ claim
            in
            if isIsolated yaz xaz fabric then
                claim.id

            else
                firstIsolatedClaimId rest fabric


result =
    firstIsolatedClaimId allClaims (markedFabric allClaims neatFabric)

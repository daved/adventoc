module Input exposing (data)


data : List String
data =
    """
"""
        |> String.lines
        |> List.filter (not << String.isEmpty)

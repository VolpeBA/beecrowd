a = IO.gets("") |> String.trim() |> String.to_float
b = IO.gets("") |> String.trim() |> String.to_float
c = IO.gets("") |> String.trim() |> String.to_float

media = ((a * 2) + (b * 3) + (c * 5)) / 10

IO.puts "MEDIA = #{:erlang.float_to_binary(media, decimals: 1)}"

n1 = IO.gets("") |> String.trim() |> String.to_float
n2 = IO.gets("") |> String.trim() |> String.to_float

media = ((n1 * 3.5) + (n2 * 7.5)) / 11

IO.puts "MEDIA = #{:erlang.float_to_binary(media, decimals: 5)}"

# 1014
km = IO.gets("") |> String.trim() |> String.to_integer
litros = IO.gets("") |> String.trim() |> String.to_float

consumo = km / litros

IO.puts "#{:erlang.float_to_binary(consumo, decimals: 3)} km/l"

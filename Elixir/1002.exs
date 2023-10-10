pi = 3.14159
raio = IO.gets("") |> String.trim() |> String.to_float

area = pi * (:math.pow(raio,2))
IO.puts "A=#{:erlang.float_to_binary(area, decimals: 4)}"
